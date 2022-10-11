// The AGPLv3 License (AGPLv3)

// Copyright (c) 2022 ZHAO Zhenhua <zhao.zhenhua@gmail.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/bytehubplus/plugin/core"
	plugin "github.com/hashicorp/go-plugin"
)

func main() {

	log.SetOutput(ioutil.Discard)

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
