// Command client is a small, self-contained example of an MCP client that talks
// to this service's MCP server. It connects over the Streamable HTTP transport,
// lists the available tools, and calls the sample "get_user" tool.
//
// Run the service first (e.g. `make run`), then in another terminal:
//
//	go run ./internal/mcp/client            # defaults to http://localhost:8081, user 1
//	go run ./internal/mcp/client -user 42
//	go run ./internal/mcp/client -endpoint http://localhost:8081 -user 1
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	endpoint := flag.String("endpoint", "http://localhost:8081", "MCP server endpoint")
	userID := flag.Uint("user", 1, "user_id to pass to the get_user tool")
	flag.Parse()

	ctx := context.Background()

	// Create a client and connect to the server over Streamable HTTP.
	client := mcp.NewClient(&mcp.Implementation{
		Name:    "{{ cookiecutter.root_directory_name }}-mcp-client",
		Version: "1.0.0",
	}, nil)

	session, err := client.Connect(ctx, mcp.NewStreamableClientTransport(*endpoint, nil))
	if err != nil {
		log.Fatalf("connect to %s: %v", *endpoint, err)
	}
	defer session.Close()

	// 1. Discover the tools the server exposes.
	tools, err := session.ListTools(ctx, &mcp.ListToolsParams{})
	if err != nil {
		log.Fatalf("list tools: %v", err)
	}

	fmt.Println("Available tools:")
	for _, t := range tools.Tools {
		fmt.Printf("  - %s: %s\n", t.Name, t.Description)
	}

	// 2. Call the sample get_user tool.
	res, err := session.CallTool(ctx, &mcp.CallToolParams{
		Name:      "get_user",
		Arguments: map[string]any{"user_id": *userID},
	})
	if err != nil {
		log.Fatalf("call get_user: %v", err)
	}

	for _, c := range res.Content {
		if tc, ok := c.(*mcp.TextContent); ok {
			fmt.Printf("get_user result: %s\n", tc.Text)
		}
	}

	if res.IsError {
		log.Fatal("get_user returned an error result")
	}
}
