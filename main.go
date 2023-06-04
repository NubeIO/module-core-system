package main

import (
	"github.com/NubeIO/flow-framework/module/shared"
	"github.com/NubeIO/system-module/logger"
	"github.com/hashicorp/go-plugin"
)

func ServePlugin() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins:         plugin.PluginSet{},
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}

func main() {
	logger.SetLogger("INFO")
	ServePlugin()
}
