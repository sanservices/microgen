package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	log "github.com/sanservices/apilogger/v2"
)

// GetUserArgs are the input arguments for the get_user tool. Struct tags drive
// the JSON Schema the MCP client uses to call the tool.
type GetUserArgs struct {
	UserID uint32 `json:"user_id" jsonschema:"the unique identifier of the user to fetch"`
}

// registerTools wires every exposed service operation as an MCP tool.
// Add new tools here as the service grows; each handler should delegate to the
// service layer so behaviour stays consistent with the gRPC and REST APIs.
func (m *MCP) registerTools() {
	mcp.AddTool(m.server, &mcp.Tool{
		Name:        "get_user",
		Description: "Fetch a user by its unique ID.",
	}, m.getUser)
}

// getUser is the MCP tool handler for "get_user". It reuses the same service
// layer that backs the gRPC and REST endpoints.
func (m *MCP) getUser(ctx context.Context, req *mcp.CallToolRequest, args GetUserArgs) (*mcp.CallToolResult, any, error) {
	user, err := m.service.GetUser(ctx, args.UserID)
	if err != nil {
		log.Error(ctx, log.LogCatUncategorized, err)
		return &mcp.CallToolResult{
			IsError: true,
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("failed to get user: %v", err)},
			},
		}, nil, nil
	}

	out, err := json.Marshal(user)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(out)},
		},
	}, nil, nil
}
