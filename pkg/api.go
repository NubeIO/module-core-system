package pkg

import (
	"encoding/json"
	"errors"
	"github.com/NubeIO/rubix-os/schema/systemschema"
)

const (
	jsonSchemaNetwork = "/schema/json/network"
	jsonSchemaDevice  = "/schema/json/device"
	jsonSchemaPoint   = "/schema/json/point"
)

func (m *Module) Get(path string) ([]byte, error) {
	if path == jsonSchemaNetwork {
		return json.Marshal(systemschema.GetNetworkSchema())
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
