package mcp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	log "github.com/sanservices/apilogger/v2"
	"{{ cookiecutter.module_name }}/config"
	"{{ cookiecutter.module_name }}/internal/{{ cookiecutter.service_name }}"
)

// MCP exposes the service's operations through the Model Context Protocol so
// that LLM agents can call them as tools, alongside the gRPC and REST APIs.
// All three surfaces share the same {{ cookiecutter.service_name }}.Service layer.
type MCP struct {
	config  *config.Settings
	service {{ cookiecutter.service_name }}.Service
	server  *mcp.Server
}

// New builds the MCP server and registers the service tools on it.
func New(config *config.Settings, service {{ cookiecutter.service_name }}.Service) *MCP {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    config.Service.Name,
		Version: config.Service.Version,
	}, nil)

	m := &MCP{
		config:  config,
		service: service,
		server:  server,
	}

	m.registerTools()

	return m
}

// StartServer serves the MCP server over the Streamable HTTP transport on the
// configured MCP port. It is meant to be run in its own goroutine.
func (m *MCP) StartServer(ctx context.Context) {
	handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
		return m.server
	}, nil)

	address := fmt.Sprintf(":%s", m.config.MCP.Port)
	log.Infof(ctx, log.LogCatStartUp, "starting MCP server on port %s", m.config.MCP.Port)

	// http.ErrServerClosed is returned on a graceful shutdown and is expected.
	if err := http.ListenAndServe(address, handler); err != nil && err != http.ErrServerClosed {
		log.Errorf(ctx, log.LogCatUncategorized, "MCP server failed: %v", err)
	}
}
