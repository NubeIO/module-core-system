package pkg

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bugs"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	log "github.com/sirupsen/logrus"
	"time"
)

func (m *Module) networkUpdateSuccess(uuid string) error {
	var network model.Network
	network.InFault = false
	network.MessageLevel = model.MessageLevel.Info
	network.MessageCode = model.CommonFaultCode.Ok
	network.Message = model.CommonFaultMessage.NetworkMessage
	network.LastOk = time.Now().UTC()
	err := m.grpcMarshaller.UpdateNetworkErrors(uuid, &network)
	if err != nil {
		log.Error(bugs.DebugPrint(m.moduleName, m.networkUpdateSuccess, err))
	}
	return err
}

func (m *Module) networkUpdateErr(uuid, port string, e error) error {
	var network model.Network
	network.InFault = true
	network.MessageLevel = model.MessageLevel.Fail
	network.MessageCode = model.CommonFaultCode.NetworkError
	network.Message = fmt.Sprintf(" port: %s message: %s", port, e.Error())
	network.LastFail = time.Now().UTC()
	err := m.grpcMarshaller.UpdateNetworkErrors(uuid, &network)
	if err != nil {
		log.Error(bugs.DebugPrint(m.moduleName, m.networkUpdateErr, err))
	}
	return err
}

func (m *Module) handleSerialPayload(data string) {}
