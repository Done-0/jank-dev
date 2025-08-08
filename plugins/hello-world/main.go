package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-plugin"

	"github.com/Done-0/jank/pkg/plugin/consts"

	jank "github.com/Done-0/jank/pkg/plugin"
)

type HelloPlugin struct{}

func (p *HelloPlugin) Execute(ctx context.Context, method string, args map[string]string) (map[string]string, error) {
	switch method {
	case "greet":
		name := args["name"]
		if name == "" {
			name = "World"
		}
		return map[string]string{"message": fmt.Sprintf("Hello, %s!", name)}, nil
	case "info":
		return map[string]string{
			"plugin":  "hello-world",
			"version": "1.0.0",
			"status":  "ready",
		}, nil
	case "echo":
		return args, nil
	default:
		return map[string]string{"error": "unknown method: " + method}, fmt.Errorf("unknown method: %s", method)
	}
}

func (p *HelloPlugin) HealthCheck(ctx context.Context) error {
	return nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: jank.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			consts.PluginTypeHandler: jank.NewGRPCPlugin(&HelloPlugin{}),
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
