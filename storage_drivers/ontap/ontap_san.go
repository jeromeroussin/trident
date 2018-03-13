// Copyright 2018 NetApp, Inc. All Rights Reserved.

package ontap

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	trident "github.com/netapp/trident/config"
	"github.com/netapp/trident/storage"
	sa "github.com/netapp/trident/storage_attribute"
	drivers "github.com/netapp/trident/storage_drivers"
	"github.com/netapp/trident/storage_drivers/ontap/api"
	"github.com/netapp/trident/storage_drivers/ontap/api/azgo"
	"github.com/netapp/trident/utils"
)

const LUNAttributeFSType = "com.netapp.ndvp.fstype"

func lunPath(name string) string {
	return fmt.Sprintf("/vol/%v/lun0", name)
}

// SANStorageDriver is for iSCSI storage provisioning
type SANStorageDriver struct {
	initialized bool
	Config      drivers.OntapStorageDriverConfig
	API         *api.Client
	Telemetry   *Telemetry
}

func (d *SANStorageDriver) GetConfig() *drivers.OntapStorageDriverConfig {
	return &d.Config
}

func (d *SANStorageDriver) GetAPI() *api.Client {
	return d.API
}

func (d *SANStorageDriver) GetTelemetry() *Telemetry {
	return d.Telemetry
}

// Name is for returning the name of this driver
func (d SANStorageDriver) Name() string {
	return drivers.OntapSANStorageDriverName
}

// Initialize from the provided config
func (d *SANStorageDriver) Initialize(
	context trident.DriverContext, configJSON string, commonConfig *drivers.CommonStorageDriverConfig,
) error {

	if commonConfig.DebugTraceFlags["method"] {
		fields := log.Fields{"Method": "Initialize", "Type": "SANStorageDriver"}
		log.WithFields(fields).Debug(">>>> Initialize")
		defer log.WithFields(fields).Debug("<<<< Initialize")
	}

	// Parse the config
	config, err := InitializeOntapConfig(context, configJSON, commonConfig)
	if err != nil {
		return fmt.Errorf("error initializing %s driver: %v", d.Name(), err)
	}

	if config.IgroupName == "" {
		config.IgroupName = drivers.GetDefaultIgroupName(context)
	}

	d.API, err = InitializeOntapDriver(config)
	if err != nil {
		return fmt.Errorf("error initializing %s driver: %v", d.Name(), err)
	}
	d.Config = *config

	err = d.validate()
	if err != nil {
		return fmt.Errorf("error validating %s driver: %v", d.Name(), err)
	}

	// Set up the autosupport heartbeat
	d.Telemetry = InitializeOntapTelemetry(d)
	StartEmsHeartbeat(d)

	d.initialized = true
	return nil
}

func (d *SANStorageDriver) Initialized() bool {
	return d.initialized
}

func (d *SANStorageDriver) Terminate() {

	if d.Config.DebugTraceFlags["method"] {
		fields := log.Fields{"Method": "Terminate", "Type": "SANStorageDriver"}
		log.WithFields(fields).Debug(">>>> Terminate")
		defer log.WithFields(fields).Debug("<<<< Terminate")
	}

	d.initialized = false
}

// Validate the driver configuration and execution environment
func (d *SANStorageDriver) validate() error {

	if d.Config.DebugTraceFlags["method"] {
		fields := log.Fields{"Method": "validate", "Type": "SANStorageDriver"}
		log.WithFields(fields).Debug(">>>> validate")
		defer log.WithFields(fields).Debug("<<<< validate")
	}

	dataLIFs, err := d.API.NetInterfaceGetDataLIFs("iscsi")
	if err != nil {
		return err
	}

	if len(dataLIFs) == 0 {
		return fmt.Errorf("no iSCSI data LIFs found on SVM %s", d.Config.SVM)
	} else {
		log.WithField("dataLIFs", dataLIFs).Debug("Found iSCSI LIFs.")
	}

	// If they didn't set a LIF to use in the config, we'll set it to the first iSCSI LIF we happen to find
	if d.Config.DataLIF == "" {
		d.Config.DataLIF = dataLIFs[0]
	} else {
		err := ValidateDataLIFs(&d.Config, dataLIFs)
		if err != nil {
			return fmt.Errorf("data LIF validation failed: %v", err)
		}

		d.Config.DataLIF = dataLIFs[0]
	}

	if d.Config.DriverContext == trident.ContextDocker {
		// Make sure this host is logged into the ONTAP iSCSI target
		err := utils.EnsureISCSISession(d.Config.DataLIF)
		if err != nil {
			return fmt.Errorf("error establishing iSCSI session: %v", err)
		}

		// Make sure the configured aggregate is available
		err = ValidateAggregate(d.API, &d.Config)
		if err != nil {
			return err
		}
	}

	return nil
}

// Create a volume+LUN with the specified options
func (d *SANStorageDriver) Create(name string, sizeBytes uint64, opts map[string]string) error {

	if d.Config.DebugTraceFlags["method"] {
		fields := log.Fields{
			"Method":    "Create",
			"Type":      "SANStorageDriver",
			"name":      name,
			"sizeBytes": sizeBytes,
			"opts":      opts,
		}
		log.WithFields(fields).Debug(">>>> Create")
		defer log.WithFields(fields).Debug("<<<< Create")
	}

	// If the volume already exists, bail out
	volExists, err := d.API.VolumeExists(name)
	if err != nil {
		return fmt.Errorf("error checking for existing volume: %v", err)
	}
	if volExists {
		return fmt.Errorf("volume %s already exists", name)
	}

	if sizeBytes < MinimumVolumeSizeBytes {
		return fmt.Errorf("requested volume size (%d bytes) is too small; the minimum volume size is %d bytes",
			sizeBytes, MinimumVolumeSizeBytes)
	}

	// Get options with default fallback values
	// see also: ontap_common.go#PopulateConfigurationDefaults
	size := strconv.FormatUint(sizeBytes, 10)
	spaceReserve := utils.GetV(opts, "spaceReserve", d.Config.SpaceReserve)
	snapshotPolicy := utils.GetV(opts, "snapshotPolicy", d.Config.SnapshotPolicy)
	unixPermissions := utils.GetV(opts, "unixPermissions", d.Config.UnixPermissions)
	snapshotDir := utils.GetV(opts, "snapshotDir", d.Config.SnapshotDir)
	exportPolicy := utils.GetV(opts, "exportPolicy", d.Config.ExportPolicy)
	aggregate := utils.GetV(opts, "aggregate", d.Config.Aggregate)
	securityStyle := utils.GetV(opts, "securityStyle", d.Config.SecurityStyle)
	encryption := utils.GetV(opts, "encryption", d.Config.Encryption)

	encrypt, err := ValidateEncryptionAttribute(encryption, d.API)
	if err != nil {
		return err
	}

	// Check for a supported file system type
	fstype := strings.ToLower(utils.GetV(opts, "fstype|fileSystemType", d.Config.FileSystemType))
	switch fstype {
	case "xfs", "ext3", "ext4":
		log.WithFields(log.Fields{"fileSystemType": fstype, "name": name}).Debug("Filesystem format.")
	default:
		return fmt.Errorf("unsupported fileSystemType option: %s", fstype)
	}

	log.WithFields(log.Fields{
		"name":            name,
		"size":            size,
		"spaceReserve":    spaceReserve,
		"snapshotPolicy":  snapshotPolicy,
		"unixPermissions": unixPermissions,
		"snapshotDir":     snapshotDir,
		"exportPolicy":    exportPolicy,
		"aggregate":       aggregate,
		"securityStyle":   securityStyle,
		"encryption":      encryption,
	}).Debug("Creating Flexvol.")

	// Create the volume
	volCreateResponse, err := d.API.VolumeCreate(
		name, aggregate, size, spaceReserve, snapshotPolicy,
		unixPermissions, exportPolicy, securityStyle, encrypt)

	if err = api.GetError(volCreateResponse, err); err != nil {
		if zerr, ok := err.(api.ZapiError); ok {
			// Handle case where the Create is passed to every Docker Swarm node
			if zerr.Code() == azgo.EAPIERROR && strings.HasSuffix(strings.TrimSpace(zerr.Reason()), "Job exists") {
				log.WithField("volume", name).Warn("Volume create job already exists, " +
					"skipping volume create on this node.")
				return nil
			}
		}
		return fmt.Errorf("error creating volume: %v", err)
	}

	lunPath := lunPath(name)
	osType := "linux"

	// Create the LUN
	lunCreateResponse, err := d.API.LunCreate(lunPath, int(sizeBytes), osType, false)
	if err = api.GetError(lunCreateResponse, err); err != nil {
		return fmt.Errorf("error creating LUN: %v", err)
	}

	// Save the fstype in a LUN attribute so we know what to do in Attach
	attrResponse, err := d.API.LunSetAttribute(lunPath, LUNAttributeFSType, fstype)
	if err = api.GetError(attrResponse, err); err != nil {
		defer d.API.LunDestroy(lunPath)
		return fmt.Errorf("error saving file system type for LUN: %v", err)
	}
	// Save the context
	attrResponse, err = d.API.LunSetAttribute(lunPath, "context", string(d.Config.DriverContext))
	if err = api.GetError(attrResponse, err); err != nil {
		log.WithField("name", name).Warning("Failed to save the driver context attribute for new volume.")
	}

	return nil
}

// Create a volume clone
func (d *SANStorageDriver) CreateClone(name, source, snapshot string, opts map[string]string) error {

	if d.Config.DebugTraceFlags["method"] {
		fields := log.Fields{
			"Method":   "CreateClone",
			"Type":     "SANStorageDriver",
			"name":     name,
			"source":   source,
			"snapshot": snapshot,
			"opts":     opts,
		}
		log.WithFields(fields).Debug(">>>> CreateClone")
		defer log.WithFields(fields).Debug("<<<< CreateClone")
	}

	split, err := strconv.ParseBool(utils.GetV(opts, "splitOnClone", d.Config.SplitOnClone))
	if err != nil {
		return fmt.Errorf("invalid boolean value for splitOnClone: %v", err)
	}

	log.WithField("splitOnClone", split).Debug("Creating volume clone.")
	return CreateOntapClone(name, source, snapshot, split, &d.Config, d.API)
}

// Destroy the requested (volume,lun) storage tuple
func (d *SANStorageDriver) Destroy(name string) error {

	if d.Config.DebugTraceFlags["method"] {
		fields := log.Fields{
			"Method": "Destroy",
			"Type":   "SANStorageDriver",
			"name":   name,
		}
		log.WithFields(fields).Debug(">>>> Destroy")
		defer log.WithFields(fields).Debug("<<<< Destroy")
	}

	var (
		err             error
		iSCSINodeName   string
		iSCSIInterfaces []string
		lunID           int
	)

	// Validate Flexvol exists before trying to destroy
	volExists, err := d.API.VolumeExists(name)
	if err != nil {
		return fmt.Errorf("error checking for existing volume: %v", err)
	}
	if !volExists {
		log.WithField("volume", name).Debug("Volume already deleted, skipping destroy.")
		return nil
	}

	if d.Config.DriverContext == trident.ContextDocker {

		// Get target info
		iSCSINodeName, iSCSIInterfaces, err = d.getISCSITargetInfo()
		if err != nil {
			log.WithField("error", err).Error("Could not get target info.")
			return err
		}

		// Get the LUN ID
		lunPath := fmt.Sprintf("/vol/%s/lun0", name)
		lunMapResponse, err := d.API.LunMapListInfo(lunPath)
		if err != nil {
			return fmt.Errorf("error reading LUN maps for volume %s: %v", name, err)
		}
		lunID = -1
		for _, lunMapResponse := range lunMapResponse.Result.InitiatorGroups() {
			if lunMapResponse.InitiatorGroupName() == d.Config.IgroupName {
				lunID = lunMapResponse.LunId()
			}
		}
		if lunID >= 0 {
			// Inform the host about the device removal
			utils.PrepareDeviceForRemoval(lunID, iSCSINodeName, iSCSIInterfaces)
		}
	}

	// Delete the Flexvol & LUN
	volDestroyResponse, err := d.API.VolumeDestroy(name, true)
	if err != nil {
		return fmt.Errorf("error destroying volume %v: %v", name, err)
	}
	if zerr := api.NewZapiError(volDestroyResponse); !zerr.IsPassed() {
		// Handle case where the Destroy is passed to every Docker Swarm node
		if zerr.Code() == azgo.EVOLUMEDOESNOTEXIST {
			log.WithField("volume", name).Warn("Volume already deleted.")
		} else {
			return fmt.Errorf("error destroying volume %v: %v", name, zerr)
		}
	}

	return nil
}

// Attach the lun
func (d *SANStorageDriver) Attach(name, mountpoint string, opts map[string]string) error {

	if d.Config.DebugTraceFlags["method"] {
		fields := log.Fields{
			"Method":     "Attach",
			"Type":       "SANStorageDriver",
			"name":       name,
			"mountpoint": mountpoint,
			"opts":       opts,
		}
		log.WithFields(fields).Debug(">>>> Attach")
		defer log.WithFields(fields).Debug("<<<< Attach")
	}

	var err error

	// Error if no iSCSI session exists for the specified iscsi portal
	sessionExists, err := utils.ISCSISessionExists(d.Config.DataLIF)
	if err != nil {
		return fmt.Errorf("unexpected iSCSI session error: %v", err)
	}
	if !sessionExists {
		return fmt.Errorf("expected iSCSI session %v not found; please login to the iSCSI portal", d.Config.DataLIF)
	}

	// Get target info
	iSCSINodeName, iSCSIInterfaces, err := d.getISCSITargetInfo()
	if err != nil {
		return err
	}

	igroupName := d.Config.IgroupName
	lunPath := lunPath(name)

	// Get the fstype
	fstype := DefaultFileSystemType
	attrResponse, err := d.API.LunGetAttribute(lunPath, LUNAttributeFSType)
	if err = api.GetError(attrResponse, err); err != nil {
		log.WithFields(log.Fields{
			"LUN":    lunPath,
			"fstype": fstype,
		}).Warn("LUN attribute fstype not found, using default.")
	} else {
		fstype = attrResponse.Result.Value()
		log.WithFields(log.Fields{"LUN": lunPath, "fstype": fstype}).Debug("Found LUN attribute fstype.")
	}

	// Create igroup
	igroupResponse, err := d.API.IgroupCreate(igroupName, "iscsi", "linux")
	if err != nil {
		return fmt.Errorf("error creating igroup: %v", err)
	}
	if zerr := api.NewZapiError(igroupResponse); !zerr.IsPassed() {
		// Handle case where the igroup already exists
		if zerr.Code() != azgo.EVDISK_ERROR_INITGROUP_EXISTS {
			return fmt.Errorf("error creating igroup %v: %v", igroupName, zerr)
		}
	}

	// Lookup host IQNs
	iqns, err := utils.GetInitiatorIqns()
	if err != nil {
		return fmt.Errorf("error determining host initiator IQNs: %v", err)
	}

	// Add each IQN found to group
	for _, iqn := range iqns {
		igroupAddResponse, err := d.API.IgroupAdd(igroupName, iqn)
		if err := api.GetError(igroupAddResponse, err); err != nil {
			if zerr, ok := err.(api.ZapiError); ok {
				if zerr.Code() == azgo.EVDISK_ERROR_INITGROUP_HAS_NODE {
					continue
				}
			}
			return fmt.Errorf("error adding IQN %v to igroup %v: %v", iqn, igroupName, err)
		}
	}

	// Map LUN
	lunID, err := d.API.LunMapIfNotMapped(igroupName, lunPath)
	if err != nil {
		return err
	}

	// Rescan and wait for the device(s) to appear
	err = utils.RescanTargetAndWaitForDevice(lunID, iSCSINodeName, iSCSIInterfaces)
	if err != nil {
		return fmt.Errorf("could not find iSCSI device: %v", err)
	}

	// Lookup all the SCSI device information
	deviceInfo, err := utils.GetDeviceInfoForLUN(lunID, iSCSINodeName, iSCSIInterfaces)
	if err != nil {
		return fmt.Errorf("error getting iSCSI device information: %v", err)
	} else if deviceInfo == nil {
		return fmt.Errorf("could not get iSCSI device information for LUN %d", lunID)
	}

	log.WithFields(log.Fields{
		"scsiLun":          deviceInfo.LUN,
		"multipathDevFile": deviceInfo.MultipathDevice,
		"devFile":          deviceInfo.Device,
		"fsType":           deviceInfo.Filesystem,
		"iqn":              deviceInfo.IQN,
	}).Debug("Found device.")

	// Make sure we use the proper device (multipath if in use)
	deviceToUse := deviceInfo.Device
	if deviceInfo.MultipathDevice != "" {
		deviceToUse = deviceInfo.MultipathDevice
	}

	if deviceToUse == "" {
		return fmt.Errorf("could not determine device to use for %v", name)
	}

	// Put a filesystem on it if there isn't one already there
	if deviceInfo.Filesystem == "" {
		log.WithFields(log.Fields{"LUN": lunPath, "fstype": fstype}).Debug("Formatting LUN.")
		err := utils.FormatVolume(deviceToUse, fstype)
		if err != nil {
			return fmt.Errorf("error formatting LUN %v, device %v: %v", name, deviceToUse, err)
		}
	} else if deviceInfo.Filesystem != fstype {
		log.WithFields(log.Fields{
			"LUN":             lunPath,
			"existingFstype":  deviceInfo.Filesystem,
			"requestedFstype": fstype,
		}).Warn("LUN already formatted with a different file system type.")
	} else {
		log.WithFields(log.Fields{"LUN": lunPath, "fstype": deviceInfo.Filesystem}).Debug("LUN already formatted.")
	}

	// Mount it
	err = utils.Mount(deviceToUse, mountpoint)
	if err != nil {
		return fmt.Errorf("error mounting LUN %v, device %v, mountpoint %v: %v",
			name, deviceToUse, mountpoint, err)
	}

	return nil
}

func (d *SANStorageDriver) getISCSITargetInfo() (iSCSINodeName string, iSCSIInterfaces []string, returnError error) {

	// Get the SVM iSCSI IQN
	nodeNameResponse, err := d.API.IscsiNodeGetNameRequest()
	if err != nil {
		returnError = fmt.Errorf("could not get SVM iSCSI node name: %v", err)
		return
	}
	iSCSINodeName = nodeNameResponse.Result.NodeName()

	// Get the SVM iSCSI interfaces
	interfaceResponse, err := d.API.IscsiInterfaceGetIterRequest()
	if err != nil {
		returnError = fmt.Errorf("could not get SVM iSCSI interfaces: %v", err)
		return
	}
	for _, iscsiAttrs := range interfaceResponse.Result.AttributesList() {
		if !iscsiAttrs.IsInterfaceEnabled() {
			continue
		}
		iSCSIInterface := fmt.Sprintf("%s:%d", iscsiAttrs.IpAddress(), iscsiAttrs.IpPort())
		iSCSIInterfaces = append(iSCSIInterfaces, iSCSIInterface)
	}
	if len(iSCSIInterfaces) == 0 {
		returnError = fmt.Errorf("SVM %s has no active iSCSI interfaces", d.Config.SVM)
		return
	}

	return
}

// Detach the volume
func (d *SANStorageDriver) Detach(name, mountpoint string) error {

	if d.Config.DebugTraceFlags["method"] {
		fields := log.Fields{
			"Method":     "Detach",
			"Type":       "SANStorageDriver",
			"name":       name,
			"mountpoint": mountpoint,
		}
		log.WithFields(fields).Debug(">>>> Detach")
		defer log.WithFields(fields).Debug("<<<< Detach")
	}

	cmd := fmt.Sprintf("umount %s", mountpoint)
	log.WithField("command", cmd).Debug("Unmounting volume.")

	if out, err := exec.Command("sh", "-c", cmd).CombinedOutput(); err != nil {
		log.WithField("output", string(out)).Debug("Unmount failed.")
		return fmt.Errorf("error unmounting volume %v, mountpoint %v: %v", name, mountpoint, err)
	}

	return nil
}

// Return the list of snapshots associated with the named volume
func (d *SANStorageDriver) SnapshotList(name string) ([]storage.Snapshot, error) {

	if d.Config.DebugTraceFlags["method"] {
		fields := log.Fields{
			"Method": "SnapshotList",
			"Type":   "SANStorageDriver",
			"name":   name,
		}
		log.WithFields(fields).Debug(">>>> SnapshotList")
		defer log.WithFields(fields).Debug("<<<< SnapshotList")
	}

	return GetSnapshotList(name, &d.Config, d.API)
}

// Return the list of volumes associated with this tenant
func (d *SANStorageDriver) List() ([]string, error) {

	if d.Config.DebugTraceFlags["method"] {
		fields := log.Fields{"Method": "List", "Type": "SANStorageDriver"}
		log.WithFields(fields).Debug(">>>> List")
		defer log.WithFields(fields).Debug("<<<< List")
	}

	return GetVolumeList(d.API, &d.Config)
}

// Test for the existence of a volume
func (d *SANStorageDriver) Get(name string) error {

	if d.Config.DebugTraceFlags["method"] {
		fields := log.Fields{"Method": "Get", "Type": "SANStorageDriver"}
		log.WithFields(fields).Debug(">>>> Get")
		defer log.WithFields(fields).Debug("<<<< Get")
	}

	return GetVolume(name, d.API, &d.Config)
}

// Retrieve storage backend capabilities
func (d *SANStorageDriver) GetStorageBackendSpecs(backend *storage.Backend) error {

	backend.Name = "ontapsan_" + d.Config.DataLIF
	poolAttrs := d.GetStoragePoolAttributes()
	return getStorageBackendSpecsCommon(d, backend, poolAttrs)
}

func (d *SANStorageDriver) GetStoragePoolAttributes() map[string]sa.Offer {

	return map[string]sa.Offer{
		sa.BackendType:      sa.NewStringOffer(d.Name()),
		sa.Snapshots:        sa.NewBoolOffer(true),
		sa.Clones:           sa.NewBoolOffer(true),
		sa.Encryption:       sa.NewBoolOffer(d.API.SupportsFeature(api.NetAppVolumeEncryption)),
		sa.ProvisioningType: sa.NewStringOffer("thick", "thin"),
	}
}

func (d *SANStorageDriver) GetVolumeOpts(
	volConfig *storage.VolumeConfig,
	pool *storage.Pool,
	requests map[string]sa.Request,
) (map[string]string, error) {
	return getVolumeOptsCommon(volConfig, pool, requests), nil
}

func (d *SANStorageDriver) GetInternalVolumeName(name string) string {
	return getInternalVolumeNameCommon(d.Config.CommonStorageDriverConfig, name)
}

func (d *SANStorageDriver) CreatePrepare(volConfig *storage.VolumeConfig) bool {
	return createPrepareCommon(d, volConfig)
}

func (d *SANStorageDriver) CreateFollowup(volConfig *storage.VolumeConfig) error {

	if d.Config.DebugTraceFlags["method"] {
		fields := log.Fields{
			"Method":       "CreateFollowup",
			"Type":         "SANStorageDriver",
			"name":         volConfig.Name,
			"internalName": volConfig.InternalName,
		}
		log.WithFields(fields).Debug(">>>> CreateFollowup")
		defer log.WithFields(fields).Debug("<<<< CreateFollowup")
	}

	if d.Config.DriverContext == trident.ContextDocker {
		log.Debug("No follow-up create actions for Docker.")
		return nil
	}

	return d.mapOntapSANLun(volConfig)
}

func (d *SANStorageDriver) mapOntapSANLun(volConfig *storage.VolumeConfig) error {
	var (
		targetIQN string
		lunID     int
	)

	response, err := d.API.IscsiServiceGetIterRequest()
	if response.Result.ResultStatusAttr != "passed" || err != nil {
		return fmt.Errorf("problem retrieving iSCSI services: %v, %v",
			err, response.Result.ResultErrnoAttr)
	}
	for _, serviceInfo := range response.Result.AttributesList() {
		if serviceInfo.Vserver() == d.Config.SVM {
			targetIQN = serviceInfo.NodeName()
			log.WithFields(log.Fields{
				"volume":    volConfig.Name,
				"targetIQN": targetIQN,
			}).Debug("Successfully discovered target IQN for the volume.")
			break
		}
	}

	// Map LUN
	lunPath := fmt.Sprintf("/vol/%v/lun0", volConfig.InternalName)
	lunID, err = d.API.LunMapIfNotMapped(d.Config.IgroupName, lunPath)
	if err != nil {
		return err
	}

	volConfig.AccessInfo.IscsiTargetPortal = d.Config.DataLIF
	volConfig.AccessInfo.IscsiTargetIQN = targetIQN
	volConfig.AccessInfo.IscsiLunNumber = int32(lunID)
	volConfig.AccessInfo.IscsiIgroup = d.Config.IgroupName
	log.WithFields(log.Fields{
		"volume":          volConfig.Name,
		"volume_internal": volConfig.InternalName,
		"targetIQN":       volConfig.AccessInfo.IscsiTargetIQN,
		"lunNumber":       volConfig.AccessInfo.IscsiLunNumber,
		"igroup":          volConfig.AccessInfo.IscsiIgroup,
	}).Debug("Successfully mapped ONTAP LUN.")

	return nil
}

func (d *SANStorageDriver) GetProtocol() trident.Protocol {
	return trident.Block
}

func (d *SANStorageDriver) StoreConfig(
	b *storage.PersistentStorageBackendConfig,
) {
	drivers.SanitizeCommonStorageDriverConfig(d.Config.CommonStorageDriverConfig)
	b.OntapConfig = &d.Config
}

func (d *SANStorageDriver) GetExternalConfig() interface{} {
	return getExternalConfig(d.Config)
}

// GetVolumeExternal queries the storage backend for all relevant info about
// a single container volume managed by this driver and returns a VolumeExternal
// representation of the volume.
func (d *SANStorageDriver) GetVolumeExternal(name string) (*storage.VolumeExternal, error) {

	volumeAttrs, err := d.API.VolumeGet(name)
	if err != nil {
		return nil, err
	}

	lunPath := fmt.Sprintf("/vol/%v/lun0", name)
	lunAttrs, err := d.API.LunGet(lunPath)
	if err != nil {
		return nil, err
	}

	return d.getVolumeExternal(&lunAttrs, &volumeAttrs), nil
}

// GetVolumeExternalWrappers queries the storage backend for all relevant info about
// container volumes managed by this driver.  It then writes a VolumeExternal
// representation of each volume to the supplied channel, closing the channel
// when finished.
func (d *SANStorageDriver) GetVolumeExternalWrappers(
	channel chan *storage.VolumeExternalWrapper) {

	// Let the caller know we're done by closing the channel
	defer close(channel)

	// Get all volumes matching the storage prefix
	volumesResponse, err := d.API.VolumeGetAll(*d.Config.StoragePrefix)
	if err = api.GetError(volumesResponse, err); err != nil {
		channel <- &storage.VolumeExternalWrapper{nil, err}
		return
	}

	// Get all LUNs named 'lun0' in volumes matching the storage prefix
	lunPathPattern := fmt.Sprintf("/vol/%v/lun0", *d.Config.StoragePrefix+"*")
	lunsResponse, err := d.API.LunGetAll(lunPathPattern)
	if err = api.GetError(lunsResponse, err); err != nil {
		channel <- &storage.VolumeExternalWrapper{nil, err}
		return
	}

	// Make a map of volumes for faster correlation with LUNs
	volumeMap := make(map[string]azgo.VolumeAttributesType)
	for _, volumeAttrs := range volumesResponse.Result.AttributesList() {
		internalName := string(volumeAttrs.VolumeIdAttributesPtr.Name())
		volumeMap[internalName] = volumeAttrs
	}

	// Convert all LUNs to VolumeExternal and write them to the channel
	for _, lun := range lunsResponse.Result.AttributesList() {

		volume, ok := volumeMap[lun.Volume()]
		if !ok {
			log.WithField("path", lun.Path()).Warning("Flexvol not found for LUN.")
			continue
		}

		channel <- &storage.VolumeExternalWrapper{d.getVolumeExternal(&lun, &volume), nil}
	}
}

// getExternalVolume is a private method that accepts info about a volume
// as returned by the storage backend and formats it as a VolumeExternal
// object.
func (d *SANStorageDriver) getVolumeExternal(
	lunAttrs *azgo.LunInfoType, volumeAttrs *azgo.VolumeAttributesType,
) *storage.VolumeExternal {

	volumeIDAttrs := volumeAttrs.VolumeIdAttributesPtr
	volumeSnapshotAttrs := volumeAttrs.VolumeSnapshotAttributesPtr

	internalName := string(volumeIDAttrs.Name())
	name := internalName[len(*d.Config.StoragePrefix):]

	volumeConfig := &storage.VolumeConfig{
		Version:         trident.OrchestratorAPIVersion,
		Name:            name,
		InternalName:    internalName,
		Size:            strconv.FormatInt(int64(lunAttrs.Size()), 10),
		Protocol:        trident.Block,
		SnapshotPolicy:  volumeSnapshotAttrs.SnapshotPolicy(),
		ExportPolicy:    "",
		SnapshotDir:     "false",
		UnixPermissions: "",
		StorageClass:    "",
		AccessMode:      trident.ReadWriteOnce,
		AccessInfo:      storage.VolumeAccessInfo{},
		BlockSize:       "",
		FileSystem:      "",
	}

	return &storage.VolumeExternal{
		Config: volumeConfig,
		Pool:   volumeIDAttrs.ContainingAggregateName(),
	}
}
