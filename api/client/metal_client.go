// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/client/firewall"
	"github.com/metal-stack/metal-go/api/client/health"
	"github.com/metal-stack/metal-go/api/client/image"
	"github.com/metal-stack/metal-go/api/client/ip"
	"github.com/metal-stack/metal-go/api/client/machine"
	"github.com/metal-stack/metal-go/api/client/network"
	"github.com/metal-stack/metal-go/api/client/partition"
	"github.com/metal-stack/metal-go/api/client/project"
	"github.com/metal-stack/metal-go/api/client/size"
	"github.com/metal-stack/metal-go/api/client/switch_operations"
	"github.com/metal-stack/metal-go/api/client/version"
)

// Default metal HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new metal HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Metal {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new metal HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Metal {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new metal client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Metal {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Metal)
	cli.Transport = transport

	cli.Firewall = firewall.New(transport, formats)

	cli.Health = health.New(transport, formats)

	cli.Image = image.New(transport, formats)

	cli.IP = ip.New(transport, formats)

	cli.Machine = machine.New(transport, formats)

	cli.Network = network.New(transport, formats)

	cli.Partition = partition.New(transport, formats)

	cli.Project = project.New(transport, formats)

	cli.Size = size.New(transport, formats)

	cli.SwitchOperations = switch_operations.New(transport, formats)

	cli.Version = version.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Metal is a client for metal
type Metal struct {
	Firewall *firewall.Client

	Health *health.Client

	Image *image.Client

	IP *ip.Client

	Machine *machine.Client

	Network *network.Client

	Partition *partition.Client

	Project *project.Client

	Size *size.Client

	SwitchOperations *switch_operations.Client

	Version *version.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Metal) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Firewall.SetTransport(transport)

	c.Health.SetTransport(transport)

	c.Image.SetTransport(transport)

	c.IP.SetTransport(transport)

	c.Machine.SetTransport(transport)

	c.Network.SetTransport(transport)

	c.Partition.SetTransport(transport)

	c.Project.SetTransport(transport)

	c.Size.SetTransport(transport)

	c.SwitchOperations.SetTransport(transport)

	c.Version.SetTransport(transport)

}
