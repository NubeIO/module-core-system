package marshal

import (
	"encoding/json"
	"github.com/NubeIO/flow-framework/module/common"
	"github.com/NubeIO/flow-framework/module/shared"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
)

type Marshaller interface {
	GetFlowNetworks(args string) ([]*model.FlowNetwork, error)

	GetNetwork(uuid, args string) (*model.Network, error)
	GetDevice(uuid, args string) (*model.Device, error)
	GetPoint(uuid, args string) (*model.Point, error)

	GetNetworksByPluginName(pluginName, args string) ([]*model.Network, error)

	GetOneNetworkByArgs(args string) (*model.Network, error)
	GetOneDeviceByArgs(args string) (*model.Device, error)
	GetOnePointByArgs(args string) (*model.Point, error)

	CreateNetwork(body *model.Network) (*model.Network, error)
	CreateDevice(body *model.Device) (*model.Device, error)
	CreatePoint(body *model.Point) (*model.Point, error)

	UpdateNetwork(uuid string, body *model.Network) (*model.Network, error)
	UpdateDevice(uuid string, body *model.Device) (*model.Device, error)
	UpdatePoint(uuid string, body *model.Point) (*model.Point, error)

	UpdateNetworkErrors(uuid string, body *model.Network) error
	UpdateDeviceErrors(uuid string, body *model.Device) error
	UpdatePointErrors(uuid string, body *model.Point) error
	UpdatePointSuccess(uuid string, body *model.Point) error

	DeleteNetwork(uuid string) error
	DeleteDevice(uuid string) error
	DeletePoint(uuid string) error

	PointWrite(uuid string, pointWriter *model.PointWriter) (*common.PointWriteResponse, error)
}

type GrpcMarshaller struct {
	DbHelper shared.DBHelper
}

func (g *GrpcMarshaller) GetFlowNetworks(args string) ([]*model.FlowNetwork, error) {
	res, err := g.DbHelper.GetWithoutParam("flow_networks", args)
	if err != nil {
		return nil, err
	}

	var fns []*model.FlowNetwork
	if err = json.Unmarshal(res, &fns); err != nil {
		return nil, err
	}
	return fns, nil
}

func (g *GrpcMarshaller) GetNetwork(uuid, args string) (*model.Network, error) {
	res, err := g.DbHelper.Get("networks", uuid, args)
	if err != nil {
		return nil, err
	}

	var network *model.Network
	if err = json.Unmarshal(res, &network); err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GrpcMarshaller) GetDevice(uuid, args string) (*model.Device, error) {
	res, err := g.DbHelper.Get("devices", uuid, args)
	if err != nil {
		return nil, err
	}

	var device *model.Device
	if err = json.Unmarshal(res, &device); err != nil {
		return nil, err
	}
	return device, nil
}

func (g *GrpcMarshaller) GetPoint(uuid, args string) (*model.Point, error) {
	res, err := g.DbHelper.Get("points", uuid, args)
	if err != nil {
		return nil, err
	}

	var point *model.Point
	if err = json.Unmarshal(res, &point); err != nil {
		return nil, err
	}

	return point, nil
}

func (g *GrpcMarshaller) GetNetworksByPluginName(pluginName, args string) ([]*model.Network, error) {
	res, err := g.DbHelper.Get("networks_by_plugin_name", pluginName, args)
	if err != nil {
		return nil, err
	}

	var networks []*model.Network
	if err = json.Unmarshal(res, &networks); err != nil {
		return nil, err
	}
	return networks, nil
}

func (g *GrpcMarshaller) GetOneNetworkByArgs(args string) (*model.Network, error) {
	res, err := g.DbHelper.GetWithoutParam("one_network_by_args", args)
	if err != nil {
		return nil, err
	}

	var network *model.Network
	if err = json.Unmarshal(res, &network); err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GrpcMarshaller) GetOneDeviceByArgs(args string) (*model.Device, error) {
	res, err := g.DbHelper.GetWithoutParam("one_device_by_args", args)
	if err != nil {
		return nil, err
	}

	var device *model.Device
	if err = json.Unmarshal(res, &device); err != nil {
		return nil, err
	}
	return device, nil
}

func (g *GrpcMarshaller) GetOnePointByArgs(args string) (*model.Point, error) {
	res, err := g.DbHelper.GetWithoutParam("one_point_by_args", args)
	if err != nil {
		return nil, err
	}

	var point *model.Point
	if err = json.Unmarshal(res, &point); err != nil {
		return nil, err
	}
	return point, nil
}

func (g *GrpcMarshaller) CreateNetwork(body *model.Network) (*model.Network, error) {
	net, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := g.DbHelper.Post("networks", net)
	if err != nil {
		return nil, err
	}

	var network *model.Network
	if err = json.Unmarshal(res, &network); err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GrpcMarshaller) CreateDevice(body *model.Device) (*model.Device, error) {
	dev, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := g.DbHelper.Post("devices", dev)
	if err != nil {
		return nil, err
	}

	var device *model.Device
	if err = json.Unmarshal(res, &device); err != nil {
		return nil, err
	}
	return device, nil
}

func (g *GrpcMarshaller) CreatePoint(body *model.Point) (*model.Point, error) {
	pnt, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := g.DbHelper.Post("points", pnt)
	if err != nil {
		return nil, err
	}

	var point *model.Point
	if err = json.Unmarshal(res, &point); err != nil {
		return nil, err
	}
	return point, nil
}

func (g *GrpcMarshaller) UpdateNetwork(uuid string, body *model.Network) (*model.Network, error) {
	net, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := g.DbHelper.Patch("networks", uuid, net)
	if err != nil {
		return nil, err
	}

	var network *model.Network
	if err = json.Unmarshal(res, &network); err != nil {
		return nil, err
	}
	return network, nil
}

func (g *GrpcMarshaller) UpdateDevice(uuid string, body *model.Device) (*model.Device, error) {
	dev, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := g.DbHelper.Patch("devices", uuid, dev)
	if err != nil {
		return nil, err
	}

	var device *model.Device
	if err = json.Unmarshal(res, &device); err != nil {
		return nil, err
	}
	return device, nil
}

func (g *GrpcMarshaller) UpdatePoint(uuid string, body *model.Point) (*model.Point, error) {
	pnt, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := g.DbHelper.Patch("points", uuid, pnt)
	if err != nil {
		return nil, err
	}

	var point *model.Point
	if err = json.Unmarshal(res, &point); err != nil {
		return nil, err
	}
	return point, nil
}

func (g *GrpcMarshaller) UpdateNetworkErrors(uuid string, body *model.Network) error {
	dev, err := json.Marshal(body)
	if err != nil {
		return err
	}

	if _, err = g.DbHelper.Patch("network_errors", uuid, dev); err != nil {
		return err
	}
	return nil
}

func (g *GrpcMarshaller) UpdateDeviceErrors(uuid string, body *model.Device) error {
	dev, err := json.Marshal(body)
	if err != nil {
		return err
	}

	if _, err = g.DbHelper.Patch("device_errors", uuid, dev); err != nil {
		return err
	}
	return nil
}

func (g *GrpcMarshaller) UpdatePointErrors(uuid string, body *model.Point) error {
	point, err := json.Marshal(body)
	if err != nil {
		return err
	}

	if _, err = g.DbHelper.Patch("point_errors", uuid, point); err != nil {
		return err
	}
	return nil
}

func (g *GrpcMarshaller) UpdatePointSuccess(uuid string, body *model.Point) error {
	point, err := json.Marshal(body)
	if err != nil {
		return err
	}

	if _, err = g.DbHelper.Patch("point_success", uuid, point); err != nil {
		return err
	}
	return nil
}

func (g *GrpcMarshaller) DeleteNetwork(uuid string) error {
	_, err := g.DbHelper.Delete("networks", uuid)
	return err
}

func (g *GrpcMarshaller) DeleteDevice(uuid string) error {
	_, err := g.DbHelper.Delete("devices", uuid)
	return err
}

func (g *GrpcMarshaller) DeletePoint(uuid string) error {
	_, err := g.DbHelper.Delete("points", uuid)
	return err
}

func (g *GrpcMarshaller) PointWrite(uuid string, body *model.PointWriter) (*common.PointWriteResponse, error) {
	pw, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := g.DbHelper.Patch("point_write", uuid, pw)
	if err != nil {
		return nil, err
	}

	var pwr *common.PointWriteResponse
	if err = json.Unmarshal(res, &pwr); err != nil {
		return nil, err
	}
	return pwr, nil
}