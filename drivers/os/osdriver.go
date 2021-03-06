package osdriver

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/pkg/mount"
	"github.com/emccode/rexray/util"
)

var (
	driverInitFuncs map[string]InitFunc
	drivers         map[string]Driver
	debug           string
)

var (
	ErrDriverInstanceDiscovery = errors.New("Driver Instance discovery failed")
)

var Adapters map[string]Driver

type Driver interface {
	// Shows the existing mount points
	GetMounts(string, string) ([]*mount.Info, error)

	// Check whether path is mounted or not
	Mounted(string) (bool, error)

	// Unmount based on a path
	Unmount(string) error

	// Mount based on a device, target, options, label
	Mount(string, string, string, string) error

	// Format a device with a FS type
	Format(string, string, bool) error
}

type InitFunc func() (Driver, error)

func Register(name string, initFunc InitFunc) error {
	driverInitFuncs[name] = initFunc
	return nil
}

func init() {
	driverInitFuncs = make(map[string]InitFunc)
	drivers = make(map[string]Driver)
	debug = strings.ToUpper(os.Getenv("REXRAY_DEBUG"))
}

func GetDrivers(osDrivers string) (map[string]Driver, error) {
	var err error
	var osDriversArr []string
	if osDrivers != "" {
		osDriversArr = strings.Split(osDrivers, ",")
	}

	if debug == "TRUE" {
		fmt.Println(driverInitFuncs)
	}

	for name, initFunc := range driverInitFuncs {
		if len(osDriversArr) > 0 && !util.StringInSlice(name, osDriversArr) {
			continue
		}
		drivers[name], err = initFunc()
		if err != nil {
			if debug == "TRUE" {
				fmt.Println(fmt.Sprintf("Info (%s): %s", name, err))
			}
			delete(drivers, name)
		}
	}

	return drivers, nil
}
