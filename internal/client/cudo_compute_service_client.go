// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client/api_keys"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client/disks"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client/networks"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client/object_storage"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client/permissions"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client/projects"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client/search"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client/ssh_keys"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client/user"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client/virtual_machines"
)

// Default cudo compute service HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "rest.compute.cudo.org"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new cudo compute service HTTP client.
func NewHTTPClient(formats strfmt.Registry) *CudoComputeService {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cudo compute service HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *CudoComputeService {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cudo compute service client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *CudoComputeService {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(CudoComputeService)
	cli.Transport = transport
	cli.APIKeys = api_keys.New(transport, formats)
	cli.Disks = disks.New(transport, formats)
	cli.Networks = networks.New(transport, formats)
	cli.ObjectStorage = object_storage.New(transport, formats)
	cli.Permissions = permissions.New(transport, formats)
	cli.Projects = projects.New(transport, formats)
	cli.Search = search.New(transport, formats)
	cli.SSHKeys = ssh_keys.New(transport, formats)
	cli.User = user.New(transport, formats)
	cli.VirtualMachines = virtual_machines.New(transport, formats)
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

// CudoComputeService is a client for cudo compute service
type CudoComputeService struct {
	APIKeys api_keys.ClientService

	Disks disks.ClientService

	Networks networks.ClientService

	ObjectStorage object_storage.ClientService

	Permissions permissions.ClientService

	Projects projects.ClientService

	Search search.ClientService

	SSHKeys ssh_keys.ClientService

	User user.ClientService

	VirtualMachines virtual_machines.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *CudoComputeService) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.APIKeys.SetTransport(transport)
	c.Disks.SetTransport(transport)
	c.Networks.SetTransport(transport)
	c.ObjectStorage.SetTransport(transport)
	c.Permissions.SetTransport(transport)
	c.Projects.SetTransport(transport)
	c.Search.SetTransport(transport)
	c.SSHKeys.SetTransport(transport)
	c.User.SetTransport(transport)
	c.VirtualMachines.SetTransport(transport)
}