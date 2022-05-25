// Copyright 2022 NetApp, Inc. All Rights Reserved.

package ontap

import (
	"context"
	"net"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	tridentconfig "github.com/netapp/trident/config"
	mockapi "github.com/netapp/trident/mocks/mock_storage_drivers/mock_ontap"
	"github.com/netapp/trident/storage"
	sa "github.com/netapp/trident/storage_attribute"
	drivers "github.com/netapp/trident/storage_drivers"
	"github.com/netapp/trident/storage_drivers/ontap/api"
	"github.com/netapp/trident/utils"
)

func TestOntapSanStorageDriverConfigString(t *testing.T) {
	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAdminPort := "0"
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME

	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)

	ontapSanDrivers := []SANStorageDriver{
		*newTestOntapSANDriver(vserverAdminHost, vserverAdminPort, vserverAggrName, true, mockAPI),
		*newTestOntapSANDriver(vserverAdminHost, vserverAdminPort, vserverAggrName, false, mockAPI),
	}

	sensitiveIncludeList := map[string]string{
		"username":        "ontap-san-user",
		"password":        "password1!",
		"client username": "client_username",
		"client password": "client_password",
	}

	externalIncludeList := map[string]string{
		"<REDACTED>":                   "<REDACTED>",
		"username":                     "Username:<REDACTED>",
		"password":                     "Password:<REDACTED>",
		"api":                          "API:<REDACTED>",
		"chap username":                "ChapUsername:<REDACTED>",
		"chap initiator secret":        "ChapInitiatorSecret:<REDACTED>",
		"chap target username":         "ChapTargetUsername:<REDACTED>",
		"chap target initiator secret": "ChapTargetInitiatorSecret:<REDACTED>",
		"client private key":           "ClientPrivateKey:<REDACTED>",
	}

	for _, ontapSanDriver := range ontapSanDrivers {
		for key, val := range externalIncludeList {
			assert.Contains(t, ontapSanDriver.String(), val, "ontap-san driver does not contain %v", key)
			assert.Contains(t, ontapSanDriver.GoString(), val, "ontap-san driver does not contain %v", key)
		}

		for key, val := range sensitiveIncludeList {
			assert.NotContains(t, ontapSanDriver.String(), val, "ontap-san driver contains %v", key)
			assert.NotContains(t, ontapSanDriver.GoString(), val, "ontap-san driver contains %v", key)
		}
	}
}

func newTestOntapSANDriver(
	vserverAdminHost, vserverAdminPort, vserverAggrName string, useREST bool, apiOverride api.OntapAPI,
) *SANStorageDriver {
	config := &drivers.OntapStorageDriverConfig{}
	sp := func(s string) *string { return &s }

	config.CommonStorageDriverConfig = &drivers.CommonStorageDriverConfig{}
	config.CommonStorageDriverConfig.DebugTraceFlags = make(map[string]bool)
	config.CommonStorageDriverConfig.DebugTraceFlags["method"] = true
	config.CommonStorageDriverConfig.DebugTraceFlags["api"] = true
	// config.Labels = map[string]string{"app": "wordpress"}
	config.ManagementLIF = vserverAdminHost + ":" + vserverAdminPort
	config.SVM = "SVM1"
	config.Aggregate = vserverAggrName
	config.Username = "ontap-san-user"
	config.Password = "password1!"
	config.StorageDriverName = "ontap-san"
	config.StoragePrefix = sp("test_")
	config.UseREST = useREST

	sanDriver := &SANStorageDriver{}
	sanDriver.Config = *config

	numRecords := api.DefaultZapiRecords
	if config.DriverContext == tridentconfig.ContextDocker {
		numRecords = api.MaxZapiRecords
	}

	var ontapAPI api.OntapAPI

	if apiOverride != nil {
		ontapAPI = apiOverride
	} else {
		if config.UseREST {
			ontapAPI, _ = api.NewRestClientFromOntapConfig(context.TODO(), config)
		} else {
			ontapAPI, _ = api.NewZAPIClientFromOntapConfig(context.TODO(), config, numRecords)
		}
	}

	sanDriver.API = ontapAPI
	sanDriver.telemetry = &TelemetryAbstraction{
		Plugin:        sanDriver.Name(),
		SVM:           sanDriver.GetConfig().SVM,
		StoragePrefix: *sanDriver.GetConfig().StoragePrefix,
		Driver:        sanDriver,
	}

	return sanDriver
}

func TestOntapSanReconcileNodeAccess(t *testing.T) {
	ctx := context.Background()

	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME

	server := api.NewFakeUnstartedVserver(ctx, vserverAdminHost, vserverAggrName)
	server.StartTLS()

	_, port, err := net.SplitHostPort(server.Listener.Addr().String())
	assert.Nil(t, err, "Unable to get Web host port %s", port)

	defer func() {
		if r := recover(); r != nil {
			server.Close()
			t.Error("Panic in fake filer", r)
		}
	}()

	cases := [][]struct {
		igroupName         string
		igroupExistingIQNs []string
		nodes              []*utils.Node
		igroupFinalIQNs    []string
	}{
		// Add a backend
		{
			{
				igroupName:         "igroup1",
				igroupExistingIQNs: []string{},
				nodes: []*utils.Node{
					{
						Name: "node1",
						IQN:  "IQN1",
					},
					{
						Name: "node2",
						IQN:  "IQN2",
					},
				},
				igroupFinalIQNs: []string{"IQN1", "IQN2"},
			},
		},
		// 2 same cluster backends/ nodes unchanged - both current
		{
			{
				igroupName:         "igroup1",
				igroupExistingIQNs: []string{"IQN1", "IQN2"},
				nodes: []*utils.Node{
					{
						Name: "node1",
						IQN:  "IQN1",
					},
					{
						Name: "node2",
						IQN:  "IQN2",
					},
				},
				igroupFinalIQNs: []string{"IQN1", "IQN2"},
			},
			{
				igroupName:         "igroup2",
				igroupExistingIQNs: []string{"IQN3", "IQN4"},
				nodes: []*utils.Node{
					{
						Name: "node3",
						IQN:  "IQN3",
					},
					{
						Name: "node4",
						IQN:  "IQN4",
					},
				},
				igroupFinalIQNs: []string{"IQN3", "IQN4"},
			},
		},
		// 2 different cluster backends - add node
		{
			{
				igroupName:         "igroup1",
				igroupExistingIQNs: []string{"IQN1"},
				nodes: []*utils.Node{
					{
						Name: "node1",
						IQN:  "IQN1",
					},
					{
						Name: "node2",
						IQN:  "IQN2",
					},
				},
				igroupFinalIQNs: []string{"IQN1", "IQN2"},
			},
			{
				igroupName:         "igroup2",
				igroupExistingIQNs: []string{"IQN3", "IQN4"},
				nodes: []*utils.Node{
					{
						Name: "node3",
						IQN:  "IQN3",
					},
					{
						Name: "node4",
						IQN:  "IQN4",
					},
				},
				igroupFinalIQNs: []string{"IQN3", "IQN4"},
			},
		},
		// 2 different cluster backends - remove node
		{
			{
				igroupName:         "igroup1",
				igroupExistingIQNs: []string{"IQN1", "IQN2"},
				nodes: []*utils.Node{
					{
						Name: "node1",
						IQN:  "IQN1",
					},
				},
				igroupFinalIQNs: []string{"IQN1"},
			},
			{
				igroupName:         "igroup2",
				igroupExistingIQNs: []string{"IQN3", "IQN4"},
				nodes: []*utils.Node{
					{
						Name: "node3",
						IQN:  "IQN3",
					},
					{
						Name: "node4",
						IQN:  "IQN4",
					},
				},
				igroupFinalIQNs: []string{"IQN3", "IQN4"},
			},
		},
	}

	for _, testCase := range cases {

		api.FakeIgroups = map[string]map[string]struct{}{}

		var ontapSanDrivers []SANStorageDriver

		for _, driverInfo := range testCase {

			// simulate existing IQNs on the vserver
			igroupsIQNMap := map[string]struct{}{}
			for _, iqn := range driverInfo.igroupExistingIQNs {
				igroupsIQNMap[iqn] = struct{}{}
			}

			api.FakeIgroups[driverInfo.igroupName] = igroupsIQNMap

			sanStorageDriver := newTestOntapSANDriver(vserverAdminHost, port, vserverAggrName, false, nil)
			sanStorageDriver.Config.IgroupName = driverInfo.igroupName
			ontapSanDrivers = append(ontapSanDrivers, *sanStorageDriver)
		}

		for driverIndex, driverInfo := range testCase {
			ontapSanDrivers[driverIndex].ReconcileNodeAccess(ctx, driverInfo.nodes, uuid.New().String())
		}

		for _, driverInfo := range testCase {

			assert.Equal(t, len(driverInfo.igroupFinalIQNs), len(api.FakeIgroups[driverInfo.igroupName]))

			for _, iqn := range driverInfo.igroupFinalIQNs {
				assert.Contains(t, api.FakeIgroups[driverInfo.igroupName], iqn)
			}
		}
	}
}

func TestOntapSanTerminate(t *testing.T) {
	ctx := context.Background()

	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME

	server := api.NewFakeUnstartedVserver(ctx, vserverAdminHost, vserverAggrName)
	server.StartTLS()

	_, port, err := net.SplitHostPort(server.Listener.Addr().String())
	assert.Nil(t, err, "Unable to get Web host port %s", port)

	defer func() {
		if r := recover(); r != nil {
			server.Close()
			t.Error("Panic in fake filer", r)
		}
	}()

	cases := [][]struct {
		igroupName         string
		igroupExistingIQNs []string
	}{
		// 2 different cluster backends - remove backend
		{
			{
				igroupName:         "igroup1",
				igroupExistingIQNs: []string{"IQN1", "IQN2"},
			},
			{
				igroupName:         "igroup2",
				igroupExistingIQNs: []string{"IQN3", "IQN4"},
			},
		},
		{
			{
				igroupName:         "igroup1",
				igroupExistingIQNs: []string{},
			},
		},
	}

	for _, testCase := range cases {

		api.FakeIgroups = map[string]map[string]struct{}{}

		var ontapSanDrivers []SANStorageDriver

		for _, driverInfo := range testCase {

			// simulate existing IQNs on the vserver
			igroupsIQNMap := map[string]struct{}{}
			for _, iqn := range driverInfo.igroupExistingIQNs {
				igroupsIQNMap[iqn] = struct{}{}
			}

			api.FakeIgroups[driverInfo.igroupName] = igroupsIQNMap

			sanStorageDriver := newTestOntapSANDriver(vserverAdminHost, port, vserverAggrName, false, nil)
			sanStorageDriver.Config.IgroupName = driverInfo.igroupName
			sanStorageDriver.telemetry = nil
			ontapSanDrivers = append(ontapSanDrivers, *sanStorageDriver)
		}

		for driverIndex, driverInfo := range testCase {
			ontapSanDrivers[driverIndex].Terminate(ctx, "")
			assert.NotContains(t, api.FakeIgroups, api.FakeIgroups[driverInfo.igroupName])
		}

	}
}

func expectLunAndVolumeCreateSequence(ctx context.Context, mockAPI *mockapi.MockOntapAPI) {
	// expected call sequenece is:
	//   check the volume exists (should return false)
	//   create the volume
	//   create the LUN
	//   set attributes on the LUN

	mockAPI.EXPECT().VolumeExists(ctx, gomock.Any()).DoAndReturn(
		func(ctx context.Context, volumeName string) (bool, error) {
			return false, nil
		},
	).MaxTimes(1)

	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).DoAndReturn(
		func(ctx context.Context, volume api.Volume) error {
			return nil
		},
	).MaxTimes(1)

	mockAPI.EXPECT().LunCreate(ctx, gomock.Any()).DoAndReturn(
		func(ctx context.Context, lun api.Lun) error {
			return nil
		},
	).MaxTimes(1)

	mockAPI.EXPECT().LunSetAttribute(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, lunPath, attribute, fstype, context string) error {
			return nil
		},
	).MaxTimes(1)
}

func TestOntapSanVolumeCreate(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	expectLunAndVolumeCreateSequence(ctx, mockAPI)

	d := newTestOntapSANDriver(ONTAPTEST_LOCALHOST, "0", ONTAPTEST_VSERVER_AGGR_NAME, true, mockAPI)
	d.API = mockAPI

	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
	})
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}

	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}

	err := d.Create(ctx, volConfig, pool1, volAttrs)
	assert.Nil(t, err, "Error is not nil")
}

func TestGetChapInfo(t *testing.T) {
	type fields struct {
		initialized   bool
		Config        drivers.OntapStorageDriverConfig
		ips           []string
		API           api.OntapAPI
		telemetry     *TelemetryAbstraction
		physicalPools map[string]storage.Pool
		virtualPools  map[string]storage.Pool
	}
	type args struct {
		in0 context.Context
		in1 string
		in2 string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *utils.IscsiChapInfo
	}{
		{
			name: "driverInitialized", fields: fields{
				initialized: true,
				Config: drivers.OntapStorageDriverConfig{
					UseCHAP:                   true,
					ChapUsername:              "foo",
					ChapInitiatorSecret:       "bar",
					ChapTargetUsername:        "baz",
					ChapTargetInitiatorSecret: "biz",
				},
				ips:           nil,
				API:           nil,
				telemetry:     nil,
				physicalPools: nil,
				virtualPools:  nil,
			}, args: args{
				in0: nil,
				in1: "volume",
				in2: "node",
			}, want: &utils.IscsiChapInfo{
				UseCHAP:              true,
				IscsiUsername:        "foo",
				IscsiInitiatorSecret: "bar",
				IscsiTargetUsername:  "baz",
				IscsiTargetSecret:    "biz",
			},
		},
		{
			name: "driverUninitialized", fields: fields{
				initialized: false,
				Config: drivers.OntapStorageDriverConfig{
					UseCHAP:                   true,
					ChapUsername:              "biz",
					ChapInitiatorSecret:       "baz",
					ChapTargetUsername:        "bar",
					ChapTargetInitiatorSecret: "foo",
				},
				ips:           nil,
				API:           nil,
				telemetry:     nil,
				physicalPools: nil,
				virtualPools:  nil,
			}, args: args{
				in0: nil,
				in1: "volume",
				in2: "node",
			}, want: &utils.IscsiChapInfo{
				UseCHAP:              true,
				IscsiUsername:        "biz",
				IscsiInitiatorSecret: "baz",
				IscsiTargetUsername:  "bar",
				IscsiTargetSecret:    "foo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &SANStorageDriver{
				initialized:   tt.fields.initialized,
				Config:        tt.fields.Config,
				ips:           tt.fields.ips,
				API:           tt.fields.API,
				telemetry:     tt.fields.telemetry,
				physicalPools: tt.fields.physicalPools,
				virtualPools:  tt.fields.virtualPools,
			}
			got, err := d.GetChapInfo(tt.args.in0, tt.args.in1, tt.args.in2)
			if err != nil {
				t.Errorf("GetChapInfo(%v, %v, %v)", tt.args.in0, tt.args.in1, tt.args.in2)
			}
			assert.Equalf(t, tt.want, got, "GetChapInfo(%v, %v, %v)", tt.args.in0, tt.args.in1, tt.args.in2)
		})
	}
}
