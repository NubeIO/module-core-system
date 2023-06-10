package pkg

import (
	"github.com/NubeIO/rubix-os/module/shared"
	"github.com/NubeIO/rubix-os/src/cachestore"
	"github.com/NubeIO/rubix-os/utils/nstring"
)

type Module struct {
	dbHelper       shared.DBHelper
	moduleName     string
	grpcMarshaller shared.Marshaller
	config         *Config
	pluginUUID     string
	networkUUID    string
	interruptChan  chan struct{}
	store          cachestore.Handler
}

func (m *Module) Init(dbHelper shared.DBHelper, moduleName string) error {
	grpcMarshaller := shared.GRPCMarshaller{DbHelper: dbHelper}
	m.dbHelper = dbHelper
	m.moduleName = moduleName
	m.grpcMarshaller = &grpcMarshaller
	return nil
}

func (m *Module) GetInfo() (*shared.Info, error) {
	return &shared.Info{
		Name:       name,
		Author:     "i8e4",
		Website:    "https://nube-io.com",
		License:    "N/A",
		HasNetwork: true,
	}, nil
}

func (m *Module) GetUrlPrefix() (*string, error) {
	return nstring.New(urlPrefix), nil
}
