// Copyright 2022 NetApp, Inc. All Rights Reserved.

package ontap

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mockapi "github.com/netapp/trident/mocks/mock_storage_drivers/mock_ontap"
	"github.com/netapp/trident/storage_drivers/ontap/api"
)

const (
	localVolumeHandle  = "svm-2:volume-b"
	remoteVolumeHandle = "svm-1:volume-a"
	snapshotHandle     = ""
	replicationPolicy  = "async_mirror"
	localSVMName       = "svm-2"
	localFlexvolName   = "volume-b"
	remoteSVMName      = "svm-1"
	remoteFlexvolName  = "volume-a"
)

var errNotReady = api.NotReadyError("operation still in progress, fail")

func TestPromoteMirror_NoErrors(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()

	mockAPI.EXPECT().SnapmirrorGet(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(2).
		Return(&api.Snapmirror{State: api.SnapmirrorStateSnapmirrored}, nil)
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, replicationPolicy).Times(1).
		Return(&api.SnapmirrorPolicy{Type: api.SnapmirrorPolicyTypeAsync}, nil)
	firstCall := mockAPI.EXPECT().SnapmirrorQuiesce(ctx, localFlexvolName, localSVMName, remoteFlexvolName,
		remoteSVMName).Times(1)
	secondCall := mockAPI.EXPECT().SnapmirrorAbort(ctx, localFlexvolName, localSVMName, remoteFlexvolName,
		remoteSVMName).Times(1).After(firstCall)
	thirdCall := mockAPI.EXPECT().SnapmirrorBreak(ctx, localFlexvolName, localSVMName, remoteFlexvolName,
		remoteSVMName, snapshotHandle).Times(1).After(secondCall)
	mockAPI.EXPECT().SnapmirrorDelete(ctx, localFlexvolName, localSVMName, remoteFlexvolName,
		remoteSVMName).Times(1).After(thirdCall)

	wait, err := promoteMirror(ctx, localVolumeHandle, remoteVolumeHandle, snapshotHandle, replicationPolicy, mockAPI)

	assert.False(t, wait, "wait should be false")
	assert.NoError(t, err, "promote mirror should not return an error")
}

func TestPromoteMirror_NoRemoteHandle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()
	noRemoteVolumeHandle := ""

	wait, err := promoteMirror(ctx, localVolumeHandle, noRemoteVolumeHandle, snapshotHandle, replicationPolicy,
		mockAPI)

	assert.False(t, wait, "wait should be false")
	assert.NoError(t, err, "promote mirror should not return an error")
}

func TestPromoteMirror_QuiesceErrorNotReady(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()

	mockAPI.EXPECT().SnapmirrorGet(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(1).
		Return(&api.Snapmirror{State: api.SnapmirrorStateSnapmirrored}, nil)
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, replicationPolicy).Times(1).
		Return(&api.SnapmirrorPolicy{Type: api.SnapmirrorPolicyTypeAsync}, nil)
	mockAPI.EXPECT().SnapmirrorQuiesce(ctx, localFlexvolName, localSVMName, remoteFlexvolName,
		remoteSVMName).Times(1).Return(errNotReady)

	wait, err := promoteMirror(ctx, localVolumeHandle, remoteVolumeHandle, snapshotHandle, replicationPolicy, mockAPI)

	assert.False(t, wait, "wait should be false")
	assert.Error(t, err, "promote mirror should return a not ready error")
	assert.True(t, api.IsNotReadyError(err), "not NotReadyError")
}

func TestPromoteMirror_AbortErrorNotReady(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()

	mockAPI.EXPECT().SnapmirrorGet(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(1).
		Return(&api.Snapmirror{State: api.SnapmirrorStateSnapmirrored}, nil)
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, replicationPolicy).Times(1).
		Return(&api.SnapmirrorPolicy{Type: api.SnapmirrorPolicyTypeAsync}, nil)
	mockAPI.EXPECT().SnapmirrorQuiesce(ctx, localFlexvolName, localSVMName, remoteFlexvolName,
		remoteSVMName).Times(1)
	mockAPI.EXPECT().SnapmirrorAbort(ctx, localFlexvolName, localSVMName, remoteFlexvolName,
		remoteSVMName).Times(1).Return(errNotReady)

	wait, err := promoteMirror(ctx, localVolumeHandle, remoteVolumeHandle, snapshotHandle, replicationPolicy, mockAPI)

	assert.False(t, wait, "wait should be false")
	assert.Error(t, err, "promote mirror should return a not ready error")
	assert.True(t, api.IsNotReadyError(err), "not NotReadyError")
}

func TestPromoteMirror_BreakErrorNotReady(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()

	mockAPI.EXPECT().SnapmirrorGet(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(2).
		Return(&api.Snapmirror{State: api.SnapmirrorStateSnapmirrored}, nil)
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, replicationPolicy).Times(1).
		Return(&api.SnapmirrorPolicy{Type: api.SnapmirrorPolicyTypeAsync}, nil)
	mockAPI.EXPECT().SnapmirrorQuiesce(ctx, localFlexvolName, localSVMName, remoteFlexvolName,
		remoteSVMName).Times(1)
	mockAPI.EXPECT().SnapmirrorAbort(ctx, localFlexvolName, localSVMName, remoteFlexvolName,
		remoteSVMName).Times(1)
	mockAPI.EXPECT().SnapmirrorBreak(ctx, localFlexvolName, localSVMName, remoteFlexvolName,
		remoteSVMName, snapshotHandle).Times(1).Return(errNotReady)

	wait, err := promoteMirror(ctx, localVolumeHandle, remoteVolumeHandle, snapshotHandle, replicationPolicy, mockAPI)

	assert.False(t, wait, "wait should be false")
	assert.Error(t, err, "promote mirror should return a not ready error")
	assert.True(t, api.IsNotReadyError(err), "not NotReadyError")
}

func TestPromoteMirror_ReplicationPolicySync(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()
	NewReplicationPolicy := "sync_mirror"

	mockAPI.EXPECT().SnapmirrorGet(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(2).
		Return(&api.Snapmirror{State: api.SnapmirrorStateSnapmirrored}, nil)
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, NewReplicationPolicy).Times(1).
		Return(&api.SnapmirrorPolicy{Type: api.SnapmirrorPolicyTypeSync}, nil)
	mockAPI.EXPECT().SnapmirrorQuiesce(ctx, localFlexvolName, localSVMName, remoteFlexvolName,
		remoteSVMName).Times(1)
	mockAPI.EXPECT().SnapmirrorAbort(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(1)
	mockAPI.EXPECT().SnapmirrorBreak(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName,
		snapshotHandle).Times(1)
	mockAPI.EXPECT().SnapmirrorDelete(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(1)

	wait, err := promoteMirror(ctx, localVolumeHandle, remoteVolumeHandle, "snapHandle", NewReplicationPolicy, mockAPI)

	assert.False(t, wait, "wait should be false")
	assert.NoError(t, err, "promote mirror should not return an error")
}

func TestPromoteMirror_WaitForSnapshot(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()

	mockAPI.EXPECT().SnapmirrorGet(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(1).
		Return(&api.Snapmirror{State: api.SnapmirrorStateSnapmirrored}, nil)
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, replicationPolicy).Times(1).
		Return(&api.SnapmirrorPolicy{Type: api.SnapmirrorPolicyTypeAsync}, nil)
	mockAPI.EXPECT().VolumeSnapshotList(ctx, localFlexvolName).Times(1).Return(api.Snapshots{
		api.Snapshot{Name: "snapshot-1", CreateTime: "1"},
	}, nil)

	wait, err := promoteMirror(ctx, localVolumeHandle, remoteVolumeHandle, "volume-a/snapshot-a", replicationPolicy,
		mockAPI)

	assert.True(t, wait, "wait should be true")
	assert.NoError(t, err, "promote mirror should not return an error")
}

func TestPromoteMirror_FoundSnapshot(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()

	mockAPI.EXPECT().SnapmirrorGet(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(2).
		Return(&api.Snapmirror{State: api.SnapmirrorStateSnapmirrored}, nil)
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, replicationPolicy).Times(1).
		Return(&api.SnapmirrorPolicy{Type: api.SnapmirrorPolicyTypeAsync}, nil)
	mockAPI.EXPECT().VolumeSnapshotList(ctx, localFlexvolName).Times(1).Return(api.Snapshots{
		api.Snapshot{Name: "snapshot-a", CreateTime: "1"},
	}, nil)
	mockAPI.EXPECT().SnapmirrorQuiesce(ctx, localFlexvolName, localSVMName, remoteFlexvolName,
		remoteSVMName).Times(1)
	mockAPI.EXPECT().SnapmirrorAbort(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(1)
	mockAPI.EXPECT().SnapmirrorBreak(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName,
		"snapshot-a").Times(1)
	mockAPI.EXPECT().SnapmirrorDelete(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(1)

	wait, err := promoteMirror(ctx, localVolumeHandle, remoteVolumeHandle, "volume-a/snapshot-a", replicationPolicy,
		mockAPI)

	assert.False(t, wait, "wait should be false")
	assert.NoError(t, err, "promote mirror should not return an error")
}

func TestPromoteMirror_SnapmirrorGetError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()

	mockAPI.EXPECT().SnapmirrorGet(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(1).
		Return(nil, api.ApiError("snapmirror get error"))

	wait, err := promoteMirror(ctx, localVolumeHandle, remoteVolumeHandle, snapshotHandle, replicationPolicy, mockAPI)

	assert.False(t, wait, "wait should be false")
	assert.Error(t, err, "snapmirror get error")
}

func TestPromoteMirror_SnapmirrorPolicyGetError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()

	mockAPI.EXPECT().SnapmirrorGet(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(1).
		Return(&api.Snapmirror{State: api.SnapmirrorStateSnapmirrored}, nil)
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, replicationPolicy).Times(1).
		Return(&api.SnapmirrorPolicy{Type: api.SnapmirrorPolicyTypeAsync},
			api.ApiError("error on snapmirror policy get"))

	wait, err := promoteMirror(ctx, localVolumeHandle, remoteVolumeHandle, snapshotHandle, replicationPolicy, mockAPI)

	assert.False(t, wait, "wait should be false")
	assert.Error(t, err, "snapmirror policy get error")
}

func TestPromoteMirror_SnapshotPresentError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()

	mockAPI.EXPECT().SnapmirrorGet(ctx, localFlexvolName, localSVMName, remoteFlexvolName, remoteSVMName).Times(1).
		Return(&api.Snapmirror{State: api.SnapmirrorStateSnapmirrored}, nil)
	mockAPI.EXPECT().SnapmirrorPolicyGet(ctx, replicationPolicy).Times(1).
		Return(&api.SnapmirrorPolicy{Type: api.SnapmirrorPolicyTypeAsync}, nil)
	mockAPI.EXPECT().VolumeSnapshotList(ctx, localFlexvolName).Times(1).Return(nil,
		api.ApiError("snapshot present error"))

	wait, err := promoteMirror(ctx, localVolumeHandle, remoteVolumeHandle, "volume-a/snapshot-a",
		replicationPolicy, mockAPI)

	assert.False(t, wait, "wait should be false")
	assert.Error(t, err, "snapshot present error")
}

func TestPromoteMirror_InvalidLocalVolumeHandle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()
	wrongVolumeHandle := ""

	wait, err := promoteMirror(ctx, wrongVolumeHandle, remoteVolumeHandle, snapshotHandle, replicationPolicy, mockAPI)

	assert.False(t, wait, "wait should be false")
	assert.Error(t, err, "should return an error if cannot parse volume handle")
}

func TestPromoteMirror_InvalidRemoteVolumeHandle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()
	wrongVolumeHandle := "pvc-a"

	wait, err := promoteMirror(ctx, localVolumeHandle, wrongVolumeHandle, snapshotHandle, replicationPolicy, mockAPI)

	assert.False(t, wait, "wait should be false")
	assert.Error(t, err, "should return an error if cannot parse volume handle")
}

func TestReleaseMirror_NoErrors(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()

	mockAPI.EXPECT().SnapmirrorRelease(localFlexvolName, localSVMName)

	err := releaseMirror(ctx, localVolumeHandle, mockAPI)
	assert.NoError(t, err, "release mirror should not return an error")
}

func TestReleaseMirror_InvalidLocalVolumeHandle(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()
	wrongLocalVolumeHandle := ""

	err := releaseMirror(ctx, wrongLocalVolumeHandle, mockAPI)

	assert.Error(t, err, "should return an error if cannot parse volume handle")
}

func TestReleaseMirror_ReleaseError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	ctx := context.Background()

	mockAPI.EXPECT().SnapmirrorRelease(localFlexvolName,
		localSVMName).Return(api.ApiError("error releasing snapmirror info for volume"))
	err := releaseMirror(ctx, localVolumeHandle, mockAPI)

	assert.Error(t, err, "should return an error if release fails")
}
