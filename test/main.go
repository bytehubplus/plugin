package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bytehubplus/plugin/core"
	plugin "github.com/hashicorp/go-plugin"
)

func main() {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: core.HandshakeConfig,
		Plugins:         core.PluginMap,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC,
		},
		Cmd: exec.Command("./coreplugin"),
	})

	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		fmt.Printf("Loading plugin failed : %s\n", err.Error())
		os.Exit(1)
	}

	raw, err := rpcClient.Dispense("core")
	if err != nil {
		fmt.Printf("Loading plugin failed : %s\n", err.Error())
		os.Exit(1)
	}

	processer := raw.(core.Processer)

	response, err := processer.Process([]byte("good evening"))
	fmt.Printf("Get response from plugin:%s\n", response)
}
