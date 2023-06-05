package pkg

import (
	"encoding/json"
	"errors"
	"github.com/NubeIO/flow-framework/module/common"
	"github.com/NubeIO/flow-framework/plugin/nube/system/smodel" // May need to refactor as smodel is imported from plugin
	"github.com/NubeIO/lib-schema/systemschema"
)

const (
	schemaNetwork     = "/schema/network"
	schemaDevice      = "/schema/device"
	schemaPoint       = "/schema/point"
	jsonSchemaNetwork = "/schema/json/network"
	jsonSchemaDevice  = "/schema/json/device"
	jsonSchemaPoint   = "/schema/json/point"
)

func (m *Module) Get(path string) ([]byte, error) {
	if path == schemaNetwork {
		fns, err := m.grpcMarshaller.GetFlowNetworks("")
		if err != nil {
			return nil, err
		}

		fnsNames := make([]string, 0)
		for _, fn := range fns {
			fnsNames = append(fnsNames, fn.Name)
		}

		networkSchema := smodel.GetNetworkSchema() // Not sure if func GetNetworkSchema of smodel can be used
		networkSchema.AutoMappingFlowNetworkName.Options = fnsNames
		return json.Marshal(networkSchema)
	} else if path == schemaDevice {
		return json.Marshal(smodel.GetDeviceSchema())
	} else if path == schemaPoint {
		return json.Marshal(smodel.GetPointSchema())
	} else if path == jsonSchemaNetwork {
		fns, err := m.grpcMarshaller.GetFlowNetworks("")
		if err != nil {
			return nil, err
		}

		networkSchema := systemschema.GetNetworkSchema()
		networkSchema.AutoMappingFlowNetworkName.Options = common.GetFlowNetworkNames(fns)
		return json.Marshal(networkSchema)
	} else if path == jsonSchemaDevice {
		return json.Marshal(systemschema.GetDeviceSchema())
	} else if path == jsonSchemaPoint {
		return json.Marshal(systemschema.GetPointSchema())
	}

	return nil, errors.New("not found")
}

func (m *Module) Post(path string, body []byte) ([]byte, error) {
	return nil, errors.New("not found")
}

func (m *Module) Put(path, uuid string, body []byte) ([]byte, error) {
	return nil, errors.New("not found")
}

func (m *Module) Patch(path, uuid string, body []byte) ([]byte, error) {
	return nil, errors.New("not found")
}

func (m *Module) Delete(path, uuid string) ([]byte, error) {
	return nil, errors.New("not found")
}
