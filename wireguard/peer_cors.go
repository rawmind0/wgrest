package wireguard

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/suquant/wgrest/restapi/operations/wireguard"
)

// PeersCorsHandler wireguard devices cors
func PeersCorsHandler(
	params wireguard.PeersCorsParams,
) middleware.Responder {
	cors := wireguard.NewPeersCorsOK()
	return cors
}

// PeerCorsHandler wireguard device cors
func PeerCorsHandler(
	params wireguard.PeerCorsParams,
) middleware.Responder {
	cors := wireguard.NewPeerCorsOK()
	return cors
}
