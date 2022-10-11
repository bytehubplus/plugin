package core

import (
	plugin "github.com/hashicorp/go-plugin"
)

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "CORE_PLUGIN",
	MagicCookieValue: "RICH",
}

var PluginMap = map[string]plugin.Plugin{
	"core": &ProcessPlugin{},
}
