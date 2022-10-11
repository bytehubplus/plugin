package main

import (
	"github.com/bytehubplus/plugin/core"
	"github.com/hashicorp/go-plugin"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: core.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"core": &core.ProcessPlugin{Impl: &core.Process{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
