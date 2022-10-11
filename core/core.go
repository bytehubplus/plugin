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

package core

import (
	"context"
	"fmt"

	"github.com/bytehubplus/plugin/protos/core"
	plugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type Processer interface {
	Process(req []byte) ([]byte, error)
}

type Process struct {
}

type ProcessPlugin struct {
	plugin.Plugin
	Impl Processer
}

// GRPCServer should register this plugin for serving with the
// given GRPCServer. Unlike Plugin.Server, this is only called once
// since gRPC plugins serve singletons.
func (p *ProcessPlugin) GRPCServer(b *plugin.GRPCBroker, s *grpc.Server) error {
	core.RegisterBaseServer(s, &ProcessServer{Impl: p.Impl})
	return nil
}

// GRPCClient should return the interface implementation for the plugin
// you're serving via gRPC. The provided context will be canceled by
// go-plugin in the event of the plugin process exiting.
func (p *ProcessPlugin) GRPCClient(ctx context.Context, b *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &ProcessClient{client: core.NewBaseClient(c)}, nil
}

type ProcessClient struct {
	client core.BaseClient
}

func (p *ProcessClient) Process(req []byte) ([]byte, error) {
	resp, err := p.client.Process(context.Background(), &core.Request{Request: req})
	if err != nil {
		return nil, err
	}

	return resp.Response, nil
}

type ProcessServer struct {
	Impl Processer
}

func (p *ProcessServer) Process(ctx context.Context, req *core.Request) (*core.Response, error) {
	response, err := p.Impl.Process(req.Request)
	return &core.Response{Response: response}, err
}

func (p *Process) Process(req []byte) ([]byte, error) {
	return []byte(fmt.Sprintf("%s%s", req, "richzhao")), nil
}
