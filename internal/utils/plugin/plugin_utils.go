// Package plugin 插件工具函数
// 创建者：Done-0
// 创建时间：2025-08-08
package plugin

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Done-0/jank/configs"
)

// GenerateBinaryPath 生成插件二进制文件路径
func GenerateBinaryPath(pluginPath, pluginID, configBinary string) string {
	if configBinary == "" {
		configs, err := configs.GetConfig()
		if err != nil {
			log.Fatalf("failed to get config: %v", err)
		}

		return filepath.Join(pluginPath, configs.PluginConfig.PluginBinDir, pluginID)
	}
	return filepath.Join(pluginPath, configBinary)
}

// CheckBinaryExists 检查二进制文件是否存在
func CheckBinaryExists(binaryPath string) bool {
	_, err := os.Stat(binaryPath)
	return err == nil
}

// CheckMainFileExists 检查 main.go 文件是否存在
func CheckMainFileExists(pluginPath string) bool {
	configs, err := configs.GetConfig()
	if err != nil {
		log.Printf("failed to get config: %v", err)
		return false
	}

	mainFile := filepath.Join(pluginPath, configs.PluginConfig.PluginMainFile)
	_, err = os.Stat(mainFile)
	return err == nil
}

// EnsureBinDirectory 确保 bin 目录存在
func EnsureBinDirectory(pluginPath string) error {
	configs, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	binDir := filepath.Join(pluginPath, configs.PluginConfig.PluginBinDir)
	return os.MkdirAll(binDir, 0755)
}

// RunGoModTidy 在指定目录执行 go mod tidy
func RunGoModTidy(pluginPath string) error {
	configs, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	cmd := exec.Command(configs.PluginConfig.GoCommand, "mod", "tidy")
	cmd.Dir = pluginPath
	_, err = cmd.CombinedOutput()
	return err
}

// CompileGoPlugin 编译 Go 插件（跨平台兼容）
func CompileGoPlugin(pluginPath, outputPath string) error {
	configs, err := configs.GetConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	buildArgs := append([]string{configs.PluginConfig.GoBuildCommand, "-o", outputPath}, configs.PluginConfig.PluginMainFile)
	cmd := exec.Command(configs.PluginConfig.GoCommand, buildArgs...)
	cmd.Dir = pluginPath

	// 设置跨平台编译环境变量
	env := os.Environ()
	if !configs.PluginConfig.CGOEnabled {
		env = append(env, configs.PluginConfig.CGOEnvVar)
	}
	cmd.Env = env

	_, err = cmd.CombinedOutput()
	return err
}

// GenerateOutputPath 生成编译输出路径
func GenerateOutputPath(configBinary, pluginID string) string {
	if configBinary == "" {
		configs, err := configs.GetConfig()
		if err != nil {
			log.Fatalf("failed to get config: %v", err)
		}

		return filepath.Join(configs.PluginConfig.PluginBinDir, pluginID)
	}
	return configBinary
}
