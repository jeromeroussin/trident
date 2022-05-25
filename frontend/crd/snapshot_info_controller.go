// Copyright 2021 NetApp, Inc. All Rights Reserved.

package crd

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"

	"github.com/netapp/trident/frontend/csi"
	. "github.com/netapp/trident/logger"
	netappv1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	"github.com/netapp/trident/storage"
	"github.com/netapp/trident/utils"
)

// addSnapshotInfo is the add handler for the TridentSnapshotInfo watcher.
func (c *TridentCrdController) addSnapshotInfo(obj interface{}) {
	ctx := GenerateRequestContext(context.Background(), "", ContextSourceCRD)
	ctx = context.WithValue(ctx, CRDControllerEvent, string(EventAdd))

	Logx(ctx).Trace("TridentCrdController#addSnapshotInfo")

	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		Logx(ctx).Error(err)
		return
	}

	keyItem := KeyItem{
		key:        key,
		event:      EventAdd,
		ctx:        ctx,
		objectType: ObjectTypeTridentSnapshotInfo,
	}

	// Only add it if it's new
	newSnapInfo := obj.(*netappv1.TridentSnapshotInfo)
	if newSnapInfo.Status.SnapshotHandle == "" {
		c.workqueue.Add(keyItem)
	}
}

// updateSnapshotInfo is the update handler for the TridentSnapshotInfo watcher.
func (c *TridentCrdController) updateSnapshotInfo(old, new interface{}) {
	ctx := GenerateRequestContext(context.Background(), "", ContextSourceCRD)
	ctx = context.WithValue(ctx, CRDControllerEvent, string(EventUpdate))

	Logx(ctx).Trace("TridentCrdController#updateSnapshotInfo")

	newSnapInfo := new.(*netappv1.TridentSnapshotInfo)
	if newSnapInfo == nil {
		Logx(ctx).Warn("No updated snapshot info provided, skipping update")
		return
	}
	oldSnapInfo := old.(*netappv1.TridentSnapshotInfo)

	// If the snapshot info has not changed, do nothing
	needsUpdate := false
	if oldSnapInfo != nil {
		if newSnapInfo.GetGeneration() != oldSnapInfo.GetGeneration() && newSnapInfo.GetGeneration() != 0 {
			Logx(ctx).WithField("TridentSnapshotInfo", newSnapInfo.Name).Debugf("Spec has been updated for TSI")
			needsUpdate = true
		}
	}

	if !needsUpdate {
		Logx(ctx).WithField("TridentSnapshotInfo", newSnapInfo.Name).Debugf("No required update for TSI")
		return
	}

	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(new); err != nil {
		Logx(ctx).Error(err)
		return
	}

	keyItem := KeyItem{
		key:        key,
		event:      EventUpdate,
		ctx:        ctx,
		objectType: ObjectTypeTridentSnapshotInfo,
	}

	c.workqueue.Add(keyItem)
}

// deleteSnapshotInfo is the delete handler for the TridentSnapshotInfo watcher.
func (c *TridentCrdController) deleteSnapshotInfo(obj interface{}) {
	ctx := GenerateRequestContext(context.Background(), "", ContextSourceCRD)
	ctx = context.WithValue(ctx, CRDControllerEvent, string(EventDelete))

	Logx(ctx).Debug("TridentCrdController#deleteSnapshotInfo")

	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		Logx(ctx).Error(err)
		return
	}

	keyItem := KeyItem{
		key:        key,
		event:      EventDelete,
		ctx:        ctx,
		objectType: ObjectTypeTridentSnapshotInfo,
	}

	c.workqueue.Add(keyItem)
}

// updateTSIStatus updates the TridentSnapshotInfo.status fields on the specified TridentSnapshotInfo resource
// using the kubernetes api
func (c *TridentCrdController) updateTSIStatus(
	ctx context.Context, snapshotInfo *netappv1.TridentSnapshotInfo,
	status *netappv1.TridentSnapshotInfoStatus,
) (*netappv1.TridentSnapshotInfo, error) {
	// Create new status
	infoStatusCopy := snapshotInfo.DeepCopy()
	status.ObservedGeneration = int(infoStatusCopy.Generation)

	currentTime := time.Now()
	status.LastTransitionTime = currentTime.Format(time.RFC3339)
	infoStatusCopy.Status = *status

	return c.crdClientset.TridentV1().TridentSnapshotInfos(infoStatusCopy.Namespace).UpdateStatus(ctx,
		infoStatusCopy, updateOpts)
}

// updateTSICR updates the TridentSnapshotInfo CR
func (c *TridentCrdController) updateTSICR(ctx context.Context, tsi *netappv1.TridentSnapshotInfo,
) (*netappv1.TridentSnapshotInfo, error) {
	logFields := log.Fields{"TridentSnapshotInfo": tsi.Name}

	// Update phase of the tsiCR
	Logx(ctx).WithFields(logFields).Debug("Updating the TridentSnapshotInfo CR")

	newTSI, err := c.crdClientset.TridentV1().TridentSnapshotInfos(tsi.Namespace).Update(ctx, tsi, updateOpts)
	if err != nil {
		Logx(ctx).WithFields(logFields).Errorf("could not update TridentSnapshotInfo CR; %v", err)
	}

	return newTSI, err
}

func (c *TridentCrdController) reconcileTSI(keyItem *KeyItem) error {
	Logx(keyItem.ctx).Trace("TridentCrdController#reconcileTSI")

	if err := c.handleTridentSnapshotInfo(keyItem); err != nil {
		c.workqueue.AddRateLimited(*keyItem)

		if utils.IsReconcileDeferredError(err) {
			errMessage := fmt.Sprintf("deferred syncing TSI '%v', requeuing; %v", keyItem.key, err.Error())
			Logx(keyItem.ctx).Info(errMessage)
			return utils.ReconcileDeferredError(fmt.Errorf(errMessage))
		} else {
			errMessage := fmt.Sprintf("error syncing TSI '%v', requeuing; %v", keyItem.key, err.Error())
			Logx(keyItem.ctx).Error(errMessage)

			return fmt.Errorf(errMessage)
		}
	}

	return nil
}

// handleTridentSnapshotInfo ensures we move to the desired state and the desired state is maintained
func (c *TridentCrdController) handleTridentSnapshotInfo(keyItem *KeyItem) error {
	key := keyItem.key
	ctx := keyItem.ctx

	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		Logx(ctx).WithField("key", key).Error("Invalid key.")
		return nil
	}

	// Get the resource with this namespace/name
	snapshotInfo, err := c.snapshotInfoLister.TridentSnapshotInfos(namespace).Get(name)
	if err != nil {
		// The resource may no longer exist, in which case we stop processing.
		if errors.IsNotFound(err) {
			Logx(ctx).WithField("key", key).Debug("Object in work queue no longer exists.")
			return nil
		}

		return err
	}

	snapshotInfoCopy := snapshotInfo.DeepCopy()
	// Ensure TSI is not deleting, then ensure it has a finalizer
	if snapshotInfoCopy.ObjectMeta.DeletionTimestamp.IsZero() {
		if !snapshotInfoCopy.HasTridentFinalizers() {
			Logx(ctx).WithField("TSI.Name", snapshotInfoCopy.Name).Debugf("Adding finalizer.")
			snapshotInfoCopy.AddTridentFinalizers()

			if snapshotInfoCopy, err = c.updateTSICR(ctx, snapshotInfoCopy); err != nil {
				return fmt.Errorf("error setting finalizer; %v", err)
			}
		}
	} else {
		Logx(ctx).WithFields(log.Fields{
			"TridentSnapshotInfo.Name":                         snapshotInfoCopy.Name,
			"TridentSnapshotInfo.ObjectMeta.DeletionTimestamp": snapshotInfoCopy.ObjectMeta.DeletionTimestamp,
		}).Trace("TridentCrdController#handleTridentSnapshotInfo CR is being deleted, not updating.")

		Logx(ctx).Debugf("Removing TridentSnapshotInfo '%v' finalizers.", snapshotInfoCopy.Name)
		return c.removeFinalizers(ctx, snapshotInfoCopy, false)
	}

	validTSI, reason := snapshotInfoCopy.IsValid()
	var status *netappv1.TridentSnapshotInfoStatus
	if validTSI {
		logFields := log.Fields{
			"snapshotInfoName": snapshotInfoCopy.Name,
		}
		Logx(ctx).WithFields(logFields).Debug("Valid TridentSnapshotInfo provided.")
		snapshotHandle, err := c.getSnapshotHandle(ctx, snapshotInfoCopy)
		if err != nil {
			return err
		}
		status = &netappv1.TridentSnapshotInfoStatus{SnapshotHandle: snapshotHandle}
	} else {
		logFields := log.Fields{
			"snapshotInfoName": snapshotInfoCopy.Name,
			"reason":           reason,
		}
		Logx(ctx).WithFields(logFields).Debug("Invalid TridentSnapshotInfo provided.")
		c.recorder.Eventf(snapshotInfoCopy, corev1.EventTypeWarning, netappv1.SnapshotInfoInvalid, reason)
		status = &netappv1.TridentSnapshotInfoStatus{}
	}

	_, err = c.updateTSIStatus(ctx, snapshotInfoCopy, status)
	if err != nil {
		err = fmt.Errorf("could not update TridentSnapshotInfo status; %v", err)
		Logx(ctx).Error(err)
		c.recorder.Eventf(snapshotInfo, corev1.EventTypeWarning, netappv1.SnapshotInfoUpdateFailed,
			"Could not update TridentSnapshotInfo")
	} else {
		c.recorder.Eventf(snapshotInfo, corev1.EventTypeNormal, netappv1.SnapshotInfoUpdated, "Snapshot info updated")
	}
	return err
}

func (c *TridentCrdController) getSnapshotHandle(
	ctx context.Context, snapshotInfo *netappv1.TridentSnapshotInfo,
) (string, error) {
	// Check if k8s snapshot exists
	k8sSnapshot, err := c.snapshotClientSet.SnapshotV1().VolumeSnapshots(snapshotInfo.Namespace).Get(
		ctx,
		snapshotInfo.Spec.SnapshotName, metav1.GetOptions{},
	)
	if statusErr, ok := err.(*errors.StatusError); ok && statusErr.Status().Reason == metav1.StatusReasonNotFound {
		message := fmt.Sprintf(
			"VolumeSnapshot '%v' for TridentSnapshotInfo '%v' does not yet exist.",
			snapshotInfo.Spec.SnapshotName, snapshotInfo.Name,
		)
		Logx(ctx).Debug(message)
		c.recorder.Eventf(snapshotInfo, corev1.EventTypeWarning, netappv1.SnapshotInfoUpdateFailed,
			"VolumeSnapshot '%v' does not exist in namespace '%v'", snapshotInfo.Spec.SnapshotName,
			snapshotInfo.Namespace)
		// If PVC does not yet exist, do not update the TSI and retry later
		return "", utils.ReconcileDeferredError(fmt.Errorf(message))
	} else if err != nil {
		return "", utils.ReconcileDeferredError(err)
	}

	// Check if volumeSnapshot is bound to a volumeSnapshotContent
	snapContentName := ""
	if k8sSnapshot.Status != nil && k8sSnapshot.Status.BoundVolumeSnapshotContentName != nil {
		snapContentName = *k8sSnapshot.Status.BoundVolumeSnapshotContentName
	}
	if snapContentName == "" {
		message := fmt.Sprintf(
			"VolumeSnapshotContent for VolumeSnapshot '%v' for TridentSnapshotInfo '%v' does"+
				" not yet exist.",
			k8sSnapshot.Name, snapshotInfo.Name,
		)
		Logx(ctx).Debug(message)
		c.recorder.Eventf(snapshotInfo, corev1.EventTypeWarning, netappv1.SnapshotInfoUpdateFailed,
			"VolumeSnapshot '%v' is not bound to a VolumeSnapshotContent", snapshotInfo.Spec.SnapshotName)
		return "", utils.ReconcileDeferredError(fmt.Errorf(message))
	}
	snapContent, err := c.snapshotClientSet.SnapshotV1().VolumeSnapshotContents().Get(
		ctx, snapContentName,
		metav1.GetOptions{},
	)
	if statusErr, ok := err.(*errors.StatusError); ok && statusErr.Status().Reason == metav1.StatusReasonNotFound {
		message := fmt.Sprintf(
			"VolumeSnapshotContent '%v' for VolumeSnapshot '%v' does not yet exist.",
			snapContentName, k8sSnapshot.Name,
		)
		Logx(ctx).Debug(message)
		c.recorder.Eventf(snapshotInfo, corev1.EventTypeWarning, netappv1.SnapshotInfoUpdateFailed,
			"VolumeSnapshotContent '%v' does not exist", snapContentName)
		// If VSC does not yet exist, do not update the TSI and retry later
		return "", utils.ReconcileDeferredError(fmt.Errorf(message))
	} else if err != nil {
		return "", utils.ReconcileDeferredError(err)
	}

	// Check if VolumeSnapshotContent is a Trident snapshot
	if snapContent.Spec.Driver != csi.Provisioner {
		message := fmt.Sprintf("snapshot '%v' is not a Trident snapshot", k8sSnapshot.Name)
		Logx(ctx).WithField("snapshotDriver", snapContent.Spec.Driver).Debug(message)
		c.recorder.Eventf(snapshotInfo, corev1.EventTypeWarning, netappv1.SnapshotInfoInvalid,
			"VolumeSnapshot '%v' is not a Trident snapshot", k8sSnapshot.Name)
		return "", fmt.Errorf(message)
	}

	// Check if VolumeSnapshotContent has internal name set
	if snapContent.Status == nil || snapContent.Status.SnapshotHandle == nil || *snapContent.Status.SnapshotHandle == "" {
		message := fmt.Sprintf("SnapshotHandle for VolumeSnapshotContent '%v' is not yet set.", snapContent.Name)
		Logx(ctx).Debug(message)
		c.recorder.Eventf(snapshotInfo, corev1.EventTypeWarning, netappv1.SnapshotInfoUpdateFailed,
			"SnapshotHandle for VolumeSnapshotContent '%v' is not set", k8sSnapshot.Name)
		return "", utils.ReconcileDeferredError(fmt.Errorf(message))
	}

	// Verify the snapshot is ONTAP
	snapshotHandle := *snapContent.Status.SnapshotHandle
	volumeName, _, err := storage.ParseSnapshotID(snapshotHandle)
	if err != nil {
		c.recorder.Event(snapshotInfo, corev1.EventTypeWarning, netappv1.SnapshotInfoUpdateFailed,
			"unrecognized snapshot handle")
		return "", fmt.Errorf("unrecognized snapshot handle '%v'; %v", snapshotHandle, err)
	}
	tridentVolume, err := c.orchestrator.GetVolume(ctx, volumeName)
	if err != nil {
		c.recorder.Event(snapshotInfo, corev1.EventTypeWarning, netappv1.SnapshotInfoUpdateFailed,
			"could not find volume in Trident")
		return "", fmt.Errorf("could not find volume '%v' in Trident; %v", volumeName, err)
	}
	backend, err := c.orchestrator.GetBackendByBackendUUID(ctx, tridentVolume.BackendUUID)
	if err != nil {
		c.recorder.Event(snapshotInfo, corev1.EventTypeWarning, netappv1.SnapshotInfoUpdateFailed,
			"could not find Trident backend")
		return "", fmt.Errorf("could not find backend '%v' in Trident; %v", tridentVolume.Backend, err)
	}
	canMirror, err := c.orchestrator.CanBackendMirror(ctx, backend.BackendUUID)
	if err != nil {
		c.recorder.Event(snapshotInfo, corev1.EventTypeWarning, netappv1.SnapshotInfoUpdateFailed,
			"could not determine if Trident backend can mirror data")
		return "", fmt.Errorf("error checking if backend can mirror; %v", err)
	}
	if !canMirror {
		c.recorder.Event(snapshotInfo, corev1.EventTypeWarning, netappv1.SnapshotInfoInvalid,
			"Backend does not support TridentSnapshotInfo")
		return "", fmt.Errorf("backend does not support TridentSnapshotInfo")
	}

	return snapshotHandle, nil
}
