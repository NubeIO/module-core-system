package pkg

import (
	"github.com/NubeIO/flow-framework/module/shared"
	"github.com/NubeIO/flow-framework/utils/nstring"
	"github.com/NubeIO/system-module/marshal"
)

type Module struct {
	dbHelper       shared.DBHelper
	moduleName     string
	grpcMarshaller marshal.Marshaller
	config         *Config
	pluginUUID     string
	networkUUID    string
	interruptChan  chan struct{}
}

var module *Module

func (m *Module) Init(dbHelper shared.DBHelper, moduleName string) error {
	grpcMarshaller := marshal.GrpcMarshaller{DbHelper: dbHelper}
	m.dbHelper = dbHelper
	m.moduleName = moduleName
	m.grpcMarshaller = &grpcMarshaller
	module = m
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
