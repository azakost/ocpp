package main

import (
	"log"
	"time"

	ocpp16 "github.com/lorenzodonini/ocpp-go/ocpp1.6"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/types"
)

const defaultHeartbeatInterval = 600

type CentralSystemHandler struct {
}

var centralSystem ocpp16.CentralSystem

func main() {
	centralSystem = ocpp16.NewCentralSystem(nil, nil)
	handler := &CentralSystemHandler{}
	centralSystem.SetNewChargePointHandler(func(chargePoint ocpp16.ChargePointConnection) {
		log.Printf("new charge point connected: %v", chargePoint.ID())
	})
	centralSystem.SetChargePointDisconnectedHandler(func(chargePoint ocpp16.ChargePointConnection) {
		log.Printf("charge point disconnected: %v", chargePoint.ID())
	})
	centralSystem.SetCoreHandler(handler)
	listenPort := 8887
	log.Printf("starting central system")
	centralSystem.Start(listenPort, "/{ws}")
	log.Println("stopped central system")
}

func (*CentralSystemHandler) OnDataTransfer(chargePointId string, request *core.DataTransferRequest) (*core.DataTransferConfirmation, error) {
	log.Printf("data transfer from %v: %v\n", chargePointId, request)
	return core.NewDataTransferConfirmation(core.DataTransferStatusAccepted), nil
}

func (*CentralSystemHandler) OnHeartbeat(chargePointId string, request *core.HeartbeatRequest) (*core.HeartbeatConfirmation, error) {
	log.Printf("heartbeat from %v\n", chargePointId)
	return core.NewHeartbeatConfirmation(types.NewDateTime(time.Now())), nil
}

func (*CentralSystemHandler) OnMeterValues(chargePointId string, request *core.MeterValuesRequest) (*core.MeterValuesConfirmation, error) {
	log.Printf("meter values from %v: %v\n", chargePointId, request)
	return core.NewMeterValuesConfirmation(), nil
}

func (*CentralSystemHandler) OnStartTransaction(chargePointId string, request *core.StartTransactionRequest) (*core.StartTransactionConfirmation, error) {
	log.Printf("start transaction from %v: %v\n", chargePointId, request)
	return core.NewStartTransactionConfirmation(types.NewIdTagInfo(types.AuthorizationStatusAccepted), 1), nil
}

func (*CentralSystemHandler) OnStatusNotification(chargePointId string, request *core.StatusNotificationRequest) (confirmation *core.StatusNotificationConfirmation, err error) {
	log.Printf("status notification from %v: %v\n", chargePointId, request)
	return core.NewStatusNotificationConfirmation(), nil
}

func (*CentralSystemHandler) OnStopTransaction(chargePointId string, request *core.StopTransactionRequest) (confirmation *core.StopTransactionConfirmation, err error) {
	log.Printf("stop transaction from %v: %v\n", chargePointId, request)
	return core.NewStopTransactionConfirmation(), nil
}

func (handler *CentralSystemHandler) OnAuthorize(chargePointId string, request *core.AuthorizeRequest) (confirmation *core.AuthorizeConfirmation, err error) {
	log.Printf("authorize from %v: %v\n", chargePointId, request)
	return core.NewAuthorizationConfirmation(types.NewIdTagInfo(types.AuthorizationStatusAccepted)), nil
}

func (handler *CentralSystemHandler) OnBootNotification(chargePointId string, request *core.BootNotificationRequest) (confirmation *core.BootNotificationConfirmation, err error) {
	log.Printf("boot notification from %v: %v\n", chargePointId, request)
	return core.NewBootNotificationConfirmation(types.NewDateTime(time.Now()), defaultHeartbeatInterval, core.RegistrationStatusAccepted), nil
}
