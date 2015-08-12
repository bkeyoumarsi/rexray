package isilon

import (
	"errors"
	"strconv"
	"sync"

	//"github.com/bkeyoumarsi/isilon-docker-plugin/rest"
	"github.com/emccode/rexray/drivers/storage"
)

var (
	providerName string
)

var (
	ErrMissingVolumeID         = errors.New("Missing VolumeID")
	ErrMultipleVolumesReturned = errors.New("Multiple Volumes returned")
	ErrNoVolumesReturned       = errors.New("No Volumes returned")
	ErrLocalVolumeMaps         = errors.New("Getting local volume mounts")
)

type Driver struct {
	clusterPath string
	//volumes     map[string]*volume
	//restClient *rest.Client
	m *sync.Mutex
}

func init() {
	providerName = "Isilon"
	storagedriver.Register("Isilon", Init)
}

func Init() (storagedriver.Driver, error) {
	driver := &Driver{
		clusterPath: "TEMP",
		//volumes:     map[string]*volume{},
		//restClient:  rest.NewClient(addr, usr, pass),
		m: &sync.Mutex{},
	}
	return driver, nil
}

// GetInstance retrieves the local instance.
func (driver *Driver) GetInstance() (*Instance, error) {
	initiator, err := driver.getInitiator()
	if err != nil {
		return &storagedriver.Instance{}, err
	}

	instance := &storagedriver.Instance{
		ProviderName: providerName,
		InstanceID:   strconv.Itoa(initiator.Index),
		Region:       "",
		Name:         initiator.Name,
	}

	// log.Println("Got Instance: " + fmt.Sprintf("%+v", instance))
	return instance, nil
}

// GetVolumeMapping lists the block devices that are attached to the instance.
func (driver *Driver) GetVolumeMapping() ([]*BlockDevice, error) {
	return nil, nil
}

// GetVolume returns all volumes for the instance based on either volumeID or volumeName
// that are available to the instance.
func (driver *Driver) GetVolume(volumeID, volumeName string) ([]*Volume, error) {
	return nil, nil
}

// GetVolumeAttach returns the attachment details based on volumeID or volumeName
// where the volume is currently attached.
func (driver *Driver) GetVolumeAttach(volumeID, instanceID string) ([]*VolumeAttachment, error) {
	return nil, nil
}

// CreateSnapshot is a synch/async operation that returns snapshots that have been
// performed based on supplying a snapshotName, source volumeID, and optional description.
func (driver *Driver) CreateSnapshot(runAsync bool, snapshotName, volumeID, description string) ([]*Snapshot, error) {
	return nil, nil
}

// GetSnapshot returns a list of snapshots for a volume based on volumeID, snapshotID, or snapshotName.
func (driver *Driver) GetSnapshot(volumeID, snapshotID, snapshotName string) ([]*Snapshot, error) {
	return nil, nil
}

// RemoveSnapshot will remove a snapshot based on the snapshotID.
func (driver *Driver) RemoveSnapshot(snapshotID string) error {
	return nil
}

// CreateVolume is sync/async and will create an return a new/existing Volume based on volumeID/snapshotID with
// a name of volumeName and a size in GB.  Optionally based on the storage driver, a volumeType, IOPS, and availabilityZone
// could be defined.
func (driver *Driver) CreateVolume(runAsync bool, volumeName string, volumeID string, snapshotID string, volumeType string, IOPS int64, size int64, availabilityZone string) (*Volume, error) {
	return nil, nil
}

// RemoveVolume will remove a volume based on volumeID.
func (driver *Driver) RemoveVolume(volumeID string) error {
	return nil
}

// GetDeviceNextAvailable return a device path that will retrieve the next available disk device that can be used.
func (driver *Driver) GetDeviceNextAvailable() (string, error) {
	return "", nil
}

// AttachVolume returns a list of VolumeAttachments is sync/async that will attach a volume to an instance based on volumeID and instanceID.
func (driver *Driver) AttachVolume(runAsync bool, volumeID, instanceID string) ([]*VolumeAttachment, error) {
	return nil, nil
}

// DetachVolume is sync/async that will detach the volumeID from the local instance or the instanceID.
func (driver *Driver) DetachVolume(runAsync bool, volumeID string, instanceID string) error {
	return nil
}

// CopySnapshot is a sync/async and returns a snapshot that will copy a snapshot based on volumeID/snapshotID/snapshotName and
// create a new snapshot of desinationSnapshotName in the destinationRegion location.
func (driver *Driver) CopySnapshot(runAsync bool, volumeID, snapshotID, snapshotName, destinationSnapshotName, destinationRegion string) (*Snapshot, error) {
	return nil, nil
}
