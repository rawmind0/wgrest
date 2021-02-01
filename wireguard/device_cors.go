package wireguard

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/suquant/wgrest/restapi/operations/wireguard"
)

// DevicesCorsHandler wireguard devices cors
func DevicesCorsHandler(
	params wireguard.DevicesCorsParams,
) middleware.Responder {
	cors := wireguard.NewDevicesCorsOK()
	return cors
}

// DeviceCorsHandler wireguard device cors
func DeviceCorsHandler(
	params wireguard.DeviceCorsParams,
) middleware.Responder {
	cors := wireguard.NewDeviceCorsOK()
	return cors
}
