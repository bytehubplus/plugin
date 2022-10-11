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
