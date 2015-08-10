package isilon

import (
	"sync"

	//"github.com/bkeyoumarsi/isilon-docker-plugin/rest"
	"github.com/emccode/rexray/drivers/storage"
)

var (
	providerName string
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
