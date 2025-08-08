// Package plugin 插件工具函数
// 创建者：Done-0
// 创建时间：2025-08-08
package plugin

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Done-0/jank/pkg/plugin/consts"
)

// GenerateBinaryPath 生成插件二进制文件路径
func GenerateBinaryPath(pluginPath, pluginID, configBinary string) string {
	if configBinary == "" {
		return filepath.Join(pluginPath, consts.PluginBinDir, pluginID)
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
	mainFile := filepath.Join(pluginPath, consts.PluginMainFile)
	_, err := os.Stat(mainFile)
	return err == nil
}

// EnsureBinDirectory 确保 bin 目录存在
func EnsureBinDirectory(pluginPath string) error {
	binDir := filepath.Join(pluginPath, consts.PluginBinDir)
	return os.MkdirAll(binDir, 0755)
}

// RunGoModTidy 在指定目录执行 go mod tidy
func RunGoModTidy(pluginPath string) error {
	cmd := exec.Command(consts.GoCommand, consts.GoModTidyArgs...)
	cmd.Dir = pluginPath
	_, err := cmd.CombinedOutput()
	return err
}

// CompileGoPlugin 编译 Go 插件（跨平台兼容）
func CompileGoPlugin(pluginPath, outputPath string) error {
	buildArgs := append([]string{consts.GoBuildCommand, "-o", outputPath}, consts.PluginMainFile)
	cmd := exec.Command(consts.GoCommand, buildArgs...)
	cmd.Dir = pluginPath
	
	// 设置跨平台编译环境变量
	env := append(os.Environ(), 
		consts.CGODisabledEnv,
	)
	cmd.Env = env
	
	_, err := cmd.CombinedOutput()
	return err
}

// GenerateOutputPath 生成编译输出路径
func GenerateOutputPath(configBinary, pluginID string) string {
	if configBinary == "" {
		return filepath.Join(consts.PluginBinDir, pluginID)
	}
	return configBinary
}
