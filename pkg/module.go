package pkg

import (
	"github.com/NubeIO/rubix-os/module/shared"
	"github.com/patrickmn/go-cache"
	"time"
)

type Module struct {
	dbHelper       shared.DBHelper
	moduleName     string
	grpcMarshaller shared.Marshaller
	config         *Config
	store          *cache.Cache
}

func (m *Module) Init(dbHelper shared.DBHelper, moduleName string) error {
	grpcMarshaller := shared.GRPCMarshaller{DbHelper: dbHelper}
	m.dbHelper = dbHelper
	m.moduleName = moduleName
	m.grpcMarshaller = &grpcMarshaller
	m.store = cache.New(5*time.Minute, 10*time.Minute)
	return nil
}

func (m *Module) GetInfo() (*shared.Info, error) {
	return &shared.Info{
		Name:       m.moduleName,
		Author:     "Nube iO",
		Website:    "https://nube-io.com",
		License:    "N/A",
		HasNetwork: true,
	}, nil
}
