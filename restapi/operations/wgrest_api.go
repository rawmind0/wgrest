// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/suquant/wgrest/restapi/operations/wireguard"
)

// NewWgrestAPI creates a new Wgrest instance
func NewWgrestAPI(spec *loads.Document) *WgrestAPI {
	return &WgrestAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		WireguardDeviceCreateHandler: wireguard.DeviceCreateHandlerFunc(func(params wireguard.DeviceCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wireguard.DeviceCreate has not yet been implemented")
		}),
		WireguardDeviceDeleteHandler: wireguard.DeviceDeleteHandlerFunc(func(params wireguard.DeviceDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wireguard.DeviceDelete has not yet been implemented")
		}),
		WireguardDeviceGetHandler: wireguard.DeviceGetHandlerFunc(func(params wireguard.DeviceGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wireguard.DeviceGet has not yet been implemented")
		}),
		WireguardDeviceListHandler: wireguard.DeviceListHandlerFunc(func(params wireguard.DeviceListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wireguard.DeviceList has not yet been implemented")
		}),
		WireguardPeerCreateHandler: wireguard.PeerCreateHandlerFunc(func(params wireguard.PeerCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wireguard.PeerCreate has not yet been implemented")
		}),
		WireguardPeerDeleteHandler: wireguard.PeerDeleteHandlerFunc(func(params wireguard.PeerDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wireguard.PeerDelete has not yet been implemented")
		}),
		WireguardPeerGetHandler: wireguard.PeerGetHandlerFunc(func(params wireguard.PeerGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wireguard.PeerGet has not yet been implemented")
		}),
		WireguardPeerListHandler: wireguard.PeerListHandlerFunc(func(params wireguard.PeerListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation wireguard.PeerList has not yet been implemented")
		}),

		// Applies when the "Token" header is set
		KeyAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (key) Token from header param [Token] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*WgrestAPI Manage WireGuard VPN tunnels by RESTful manner */
type WgrestAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// KeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Token provided in the header
	KeyAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// WireguardDeviceCreateHandler sets the operation handler for the device create operation
	WireguardDeviceCreateHandler wireguard.DeviceCreateHandler
	// WireguardDeviceDeleteHandler sets the operation handler for the device delete operation
	WireguardDeviceDeleteHandler wireguard.DeviceDeleteHandler
	// WireguardDeviceGetHandler sets the operation handler for the device get operation
	WireguardDeviceGetHandler wireguard.DeviceGetHandler
	// WireguardDeviceListHandler sets the operation handler for the device list operation
	WireguardDeviceListHandler wireguard.DeviceListHandler
	// WireguardPeerCreateHandler sets the operation handler for the peer create operation
	WireguardPeerCreateHandler wireguard.PeerCreateHandler
	// WireguardPeerDeleteHandler sets the operation handler for the peer delete operation
	WireguardPeerDeleteHandler wireguard.PeerDeleteHandler
	// WireguardPeerGetHandler sets the operation handler for the peer get operation
	WireguardPeerGetHandler wireguard.PeerGetHandler
	// WireguardPeerListHandler sets the operation handler for the peer list operation
	WireguardPeerListHandler wireguard.PeerListHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *WgrestAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *WgrestAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *WgrestAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *WgrestAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *WgrestAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *WgrestAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *WgrestAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *WgrestAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *WgrestAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the WgrestAPI
func (o *WgrestAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.KeyAuth == nil {
		unregistered = append(unregistered, "TokenAuth")
	}

	if o.WireguardDeviceCreateHandler == nil {
		unregistered = append(unregistered, "wireguard.DeviceCreateHandler")
	}
	if o.WireguardDeviceDeleteHandler == nil {
		unregistered = append(unregistered, "wireguard.DeviceDeleteHandler")
	}
	if o.WireguardDeviceGetHandler == nil {
		unregistered = append(unregistered, "wireguard.DeviceGetHandler")
	}
	if o.WireguardDeviceListHandler == nil {
		unregistered = append(unregistered, "wireguard.DeviceListHandler")
	}
	if o.WireguardPeerCreateHandler == nil {
		unregistered = append(unregistered, "wireguard.PeerCreateHandler")
	}
	if o.WireguardPeerDeleteHandler == nil {
		unregistered = append(unregistered, "wireguard.PeerDeleteHandler")
	}
	if o.WireguardPeerGetHandler == nil {
		unregistered = append(unregistered, "wireguard.PeerGetHandler")
	}
	if o.WireguardPeerListHandler == nil {
		unregistered = append(unregistered, "wireguard.PeerListHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *WgrestAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *WgrestAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.KeyAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *WgrestAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *WgrestAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *WgrestAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *WgrestAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the wgrest API
func (o *WgrestAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *WgrestAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/devices"] = wireguard.NewDeviceCreate(o.context, o.WireguardDeviceCreateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/devices/{dev}"] = wireguard.NewDeviceDelete(o.context, o.WireguardDeviceDeleteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/devices/{dev}"] = wireguard.NewDeviceGet(o.context, o.WireguardDeviceGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/devices"] = wireguard.NewDeviceList(o.context, o.WireguardDeviceListHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/devices/{dev}/peers"] = wireguard.NewPeerCreate(o.context, o.WireguardPeerCreateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/devices/{dev}/peers/{peer_id}"] = wireguard.NewPeerDelete(o.context, o.WireguardPeerDeleteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/devices/{dev}/peers/{peer_id}"] = wireguard.NewPeerGet(o.context, o.WireguardPeerGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/devices/{dev}/peers"] = wireguard.NewPeerList(o.context, o.WireguardPeerListHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *WgrestAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *WgrestAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *WgrestAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *WgrestAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *WgrestAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
