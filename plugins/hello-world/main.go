package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/Done-0/jank/pkg/plugin/consts"

	jank "github.com/Done-0/jank/pkg/plugin"
)

type PluginConfig struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Repository  string `json:"repository"`
	Type        string `json:"type"`
}

type HelloPlugin struct {
	config *PluginConfig
}

func loadConfig() (*PluginConfig, error) {
	data, err := os.ReadFile(consts.PluginConfigFile)
	if err != nil {
		return nil, err
	}

	var config PluginConfig
	return &config, json.Unmarshal(data, &config)
}

func (p *HelloPlugin) Execute(ctx context.Context, method string, args map[string]string) (map[string]string, error) {
	switch method {
	case "greet":
		payload := args[consts.PluginPayloadKey]

		// 解析 JSON payload
		var params map[string]interface{}
		if err := json.Unmarshal([]byte(payload), &params); err != nil {
			return map[string]string{"error": "invalid JSON payload"}, fmt.Errorf("failed to parse payload: %w", err)
		}

		// 获取 name 参数
		name, exists := params["name"]
		if !exists {
			return map[string]string{"error": "missing 'name' parameter"}, fmt.Errorf("name parameter is required")
		}

		return map[string]string{"message": fmt.Sprintf("Hello, %s!", name)}, nil
	case "info":
		if p.config == nil {
			return map[string]string{"error": "config not loaded"}, fmt.Errorf("config not loaded")
		}
		return map[string]string{
			"id":          p.config.ID,
			"name":        p.config.Name,
			"version":     p.config.Version,
			"author":      p.config.Author,
			"description": p.config.Description,
			"repository":  p.config.Repository,
			"type":        p.config.Type,
			"status":      "ready",
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
	config, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %v\n", err)
		os.Exit(1)
	}

	logger := hclog.New(&hclog.LoggerOptions{
		Name:   config.Name,
		Output: os.Stderr,
		Level:  hclog.Info,
	})

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: jank.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			consts.PluginTypeHandler: jank.NewGRPCPlugin(&HelloPlugin{config: config}),
		},
		GRPCServer: plugin.DefaultGRPCServer,
		Logger:     logger,
	})
}
