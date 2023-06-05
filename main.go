package main

import (
	"github.com/NubeIO/flow-framework/module/shared"
	"github.com/NubeIO/system-module/logger"
	"github.com/NubeIO/system-module/pkg"
	"github.com/hashicorp/go-plugin"
)

func ServePlugin() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins:         plugin.PluginSet{"system-module": &shared.NubeModule{Impl: &pkg.Module{}}},
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}

func main() {
	logger.SetLogger("INFO")
	ServePlugin()
}
