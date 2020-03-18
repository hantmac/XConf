// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: config/config.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Config service

type ConfigService interface {
	CreateApp(ctx context.Context, in *AppRequest, opts ...client.CallOption) (*AppResponse, error)
	QueryApp(ctx context.Context, in *AppRequest, opts ...client.CallOption) (*AppResponse, error)
	DeleteApp(ctx context.Context, in *AppRequest, opts ...client.CallOption) (*Response, error)
	ListApps(ctx context.Context, in *Request, opts ...client.CallOption) (*AppsResponse, error)
	CreateCluster(ctx context.Context, in *ClusterRequest, opts ...client.CallOption) (*ClusterResponse, error)
	QueryCluster(ctx context.Context, in *ClusterRequest, opts ...client.CallOption) (*ClusterResponse, error)
	DeleteCluster(ctx context.Context, in *ClusterRequest, opts ...client.CallOption) (*Response, error)
	ListClusters(ctx context.Context, in *AppRequest, opts ...client.CallOption) (*ClustersResponse, error)
	CreateNamespace(ctx context.Context, in *NamespaceRequest, opts ...client.CallOption) (*NamespaceResponse, error)
	DeleteNamespace(ctx context.Context, in *NamespaceRequest, opts ...client.CallOption) (*Response, error)
	QueryNamespace(ctx context.Context, in *NamespaceRequest, opts ...client.CallOption) (*NamespaceResponse, error)
	ListNamespaces(ctx context.Context, in *ClusterRequest, opts ...client.CallOption) (*NamespacesResponse, error)
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...client.CallOption) (*Response, error)
	ReleaseConfig(ctx context.Context, in *ReleaseRequest, opts ...client.CallOption) (*Response, error)
	ListReleaseHistory(ctx context.Context, in *NamespaceRequest, opts ...client.CallOption) (*ReleaseHistoryResponse, error)
	Rollback(ctx context.Context, in *ReleaseRequest, opts ...client.CallOption) (*Response, error)
	Read(ctx context.Context, in *QueryConfigRequest, opts ...client.CallOption) (*ConfigResponse, error)
	Watch(ctx context.Context, in *Request, opts ...client.CallOption) (Config_WatchService, error)
}

type configService struct {
	c    client.Client
	name string
}

func NewConfigService(name string, c client.Client) ConfigService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "config"
	}
	return &configService{
		c:    c,
		name: name,
	}
}

func (c *configService) CreateApp(ctx context.Context, in *AppRequest, opts ...client.CallOption) (*AppResponse, error) {
	req := c.c.NewRequest(c.name, "Config.CreateApp", in)
	out := new(AppResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) QueryApp(ctx context.Context, in *AppRequest, opts ...client.CallOption) (*AppResponse, error) {
	req := c.c.NewRequest(c.name, "Config.QueryApp", in)
	out := new(AppResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) DeleteApp(ctx context.Context, in *AppRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Config.DeleteApp", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) ListApps(ctx context.Context, in *Request, opts ...client.CallOption) (*AppsResponse, error) {
	req := c.c.NewRequest(c.name, "Config.ListApps", in)
	out := new(AppsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) CreateCluster(ctx context.Context, in *ClusterRequest, opts ...client.CallOption) (*ClusterResponse, error) {
	req := c.c.NewRequest(c.name, "Config.CreateCluster", in)
	out := new(ClusterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) QueryCluster(ctx context.Context, in *ClusterRequest, opts ...client.CallOption) (*ClusterResponse, error) {
	req := c.c.NewRequest(c.name, "Config.QueryCluster", in)
	out := new(ClusterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) DeleteCluster(ctx context.Context, in *ClusterRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Config.DeleteCluster", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) ListClusters(ctx context.Context, in *AppRequest, opts ...client.CallOption) (*ClustersResponse, error) {
	req := c.c.NewRequest(c.name, "Config.ListClusters", in)
	out := new(ClustersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) CreateNamespace(ctx context.Context, in *NamespaceRequest, opts ...client.CallOption) (*NamespaceResponse, error) {
	req := c.c.NewRequest(c.name, "Config.CreateNamespace", in)
	out := new(NamespaceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) DeleteNamespace(ctx context.Context, in *NamespaceRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Config.DeleteNamespace", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) QueryNamespace(ctx context.Context, in *NamespaceRequest, opts ...client.CallOption) (*NamespaceResponse, error) {
	req := c.c.NewRequest(c.name, "Config.QueryNamespace", in)
	out := new(NamespaceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) ListNamespaces(ctx context.Context, in *ClusterRequest, opts ...client.CallOption) (*NamespacesResponse, error) {
	req := c.c.NewRequest(c.name, "Config.ListNamespaces", in)
	out := new(NamespacesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Config.UpdateConfig", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) ReleaseConfig(ctx context.Context, in *ReleaseRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Config.ReleaseConfig", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) ListReleaseHistory(ctx context.Context, in *NamespaceRequest, opts ...client.CallOption) (*ReleaseHistoryResponse, error) {
	req := c.c.NewRequest(c.name, "Config.ListReleaseHistory", in)
	out := new(ReleaseHistoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) Rollback(ctx context.Context, in *ReleaseRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Config.Rollback", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) Read(ctx context.Context, in *QueryConfigRequest, opts ...client.CallOption) (*ConfigResponse, error) {
	req := c.c.NewRequest(c.name, "Config.Read", in)
	out := new(ConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configService) Watch(ctx context.Context, in *Request, opts ...client.CallOption) (Config_WatchService, error) {
	req := c.c.NewRequest(c.name, "Config.Watch", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &configServiceWatch{stream}, nil
}

type Config_WatchService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ConfigResponse, error)
}

type configServiceWatch struct {
	stream client.Stream
}

func (x *configServiceWatch) Close() error {
	return x.stream.Close()
}

func (x *configServiceWatch) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *configServiceWatch) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *configServiceWatch) Recv() (*ConfigResponse, error) {
	m := new(ConfigResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Config service

type ConfigHandler interface {
	CreateApp(context.Context, *AppRequest, *AppResponse) error
	QueryApp(context.Context, *AppRequest, *AppResponse) error
	DeleteApp(context.Context, *AppRequest, *Response) error
	ListApps(context.Context, *Request, *AppsResponse) error
	CreateCluster(context.Context, *ClusterRequest, *ClusterResponse) error
	QueryCluster(context.Context, *ClusterRequest, *ClusterResponse) error
	DeleteCluster(context.Context, *ClusterRequest, *Response) error
	ListClusters(context.Context, *AppRequest, *ClustersResponse) error
	CreateNamespace(context.Context, *NamespaceRequest, *NamespaceResponse) error
	DeleteNamespace(context.Context, *NamespaceRequest, *Response) error
	QueryNamespace(context.Context, *NamespaceRequest, *NamespaceResponse) error
	ListNamespaces(context.Context, *ClusterRequest, *NamespacesResponse) error
	UpdateConfig(context.Context, *UpdateConfigRequest, *Response) error
	ReleaseConfig(context.Context, *ReleaseRequest, *Response) error
	ListReleaseHistory(context.Context, *NamespaceRequest, *ReleaseHistoryResponse) error
	Rollback(context.Context, *ReleaseRequest, *Response) error
	Read(context.Context, *QueryConfigRequest, *ConfigResponse) error
	Watch(context.Context, *Request, Config_WatchStream) error
}

func RegisterConfigHandler(s server.Server, hdlr ConfigHandler, opts ...server.HandlerOption) error {
	type config interface {
		CreateApp(ctx context.Context, in *AppRequest, out *AppResponse) error
		QueryApp(ctx context.Context, in *AppRequest, out *AppResponse) error
		DeleteApp(ctx context.Context, in *AppRequest, out *Response) error
		ListApps(ctx context.Context, in *Request, out *AppsResponse) error
		CreateCluster(ctx context.Context, in *ClusterRequest, out *ClusterResponse) error
		QueryCluster(ctx context.Context, in *ClusterRequest, out *ClusterResponse) error
		DeleteCluster(ctx context.Context, in *ClusterRequest, out *Response) error
		ListClusters(ctx context.Context, in *AppRequest, out *ClustersResponse) error
		CreateNamespace(ctx context.Context, in *NamespaceRequest, out *NamespaceResponse) error
		DeleteNamespace(ctx context.Context, in *NamespaceRequest, out *Response) error
		QueryNamespace(ctx context.Context, in *NamespaceRequest, out *NamespaceResponse) error
		ListNamespaces(ctx context.Context, in *ClusterRequest, out *NamespacesResponse) error
		UpdateConfig(ctx context.Context, in *UpdateConfigRequest, out *Response) error
		ReleaseConfig(ctx context.Context, in *ReleaseRequest, out *Response) error
		ListReleaseHistory(ctx context.Context, in *NamespaceRequest, out *ReleaseHistoryResponse) error
		Rollback(ctx context.Context, in *ReleaseRequest, out *Response) error
		Read(ctx context.Context, in *QueryConfigRequest, out *ConfigResponse) error
		Watch(ctx context.Context, stream server.Stream) error
	}
	type Config struct {
		config
	}
	h := &configHandler{hdlr}
	return s.Handle(s.NewHandler(&Config{h}, opts...))
}

type configHandler struct {
	ConfigHandler
}

func (h *configHandler) CreateApp(ctx context.Context, in *AppRequest, out *AppResponse) error {
	return h.ConfigHandler.CreateApp(ctx, in, out)
}

func (h *configHandler) QueryApp(ctx context.Context, in *AppRequest, out *AppResponse) error {
	return h.ConfigHandler.QueryApp(ctx, in, out)
}

func (h *configHandler) DeleteApp(ctx context.Context, in *AppRequest, out *Response) error {
	return h.ConfigHandler.DeleteApp(ctx, in, out)
}

func (h *configHandler) ListApps(ctx context.Context, in *Request, out *AppsResponse) error {
	return h.ConfigHandler.ListApps(ctx, in, out)
}

func (h *configHandler) CreateCluster(ctx context.Context, in *ClusterRequest, out *ClusterResponse) error {
	return h.ConfigHandler.CreateCluster(ctx, in, out)
}

func (h *configHandler) QueryCluster(ctx context.Context, in *ClusterRequest, out *ClusterResponse) error {
	return h.ConfigHandler.QueryCluster(ctx, in, out)
}

func (h *configHandler) DeleteCluster(ctx context.Context, in *ClusterRequest, out *Response) error {
	return h.ConfigHandler.DeleteCluster(ctx, in, out)
}

func (h *configHandler) ListClusters(ctx context.Context, in *AppRequest, out *ClustersResponse) error {
	return h.ConfigHandler.ListClusters(ctx, in, out)
}

func (h *configHandler) CreateNamespace(ctx context.Context, in *NamespaceRequest, out *NamespaceResponse) error {
	return h.ConfigHandler.CreateNamespace(ctx, in, out)
}

func (h *configHandler) DeleteNamespace(ctx context.Context, in *NamespaceRequest, out *Response) error {
	return h.ConfigHandler.DeleteNamespace(ctx, in, out)
}

func (h *configHandler) QueryNamespace(ctx context.Context, in *NamespaceRequest, out *NamespaceResponse) error {
	return h.ConfigHandler.QueryNamespace(ctx, in, out)
}

func (h *configHandler) ListNamespaces(ctx context.Context, in *ClusterRequest, out *NamespacesResponse) error {
	return h.ConfigHandler.ListNamespaces(ctx, in, out)
}

func (h *configHandler) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, out *Response) error {
	return h.ConfigHandler.UpdateConfig(ctx, in, out)
}

func (h *configHandler) ReleaseConfig(ctx context.Context, in *ReleaseRequest, out *Response) error {
	return h.ConfigHandler.ReleaseConfig(ctx, in, out)
}

func (h *configHandler) ListReleaseHistory(ctx context.Context, in *NamespaceRequest, out *ReleaseHistoryResponse) error {
	return h.ConfigHandler.ListReleaseHistory(ctx, in, out)
}

func (h *configHandler) Rollback(ctx context.Context, in *ReleaseRequest, out *Response) error {
	return h.ConfigHandler.Rollback(ctx, in, out)
}

func (h *configHandler) Read(ctx context.Context, in *QueryConfigRequest, out *ConfigResponse) error {
	return h.ConfigHandler.Read(ctx, in, out)
}

func (h *configHandler) Watch(ctx context.Context, stream server.Stream) error {
	m := new(Request)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ConfigHandler.Watch(ctx, m, &configWatchStream{stream})
}

type Config_WatchStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ConfigResponse) error
}

type configWatchStream struct {
	stream server.Stream
}

func (x *configWatchStream) Close() error {
	return x.stream.Close()
}

func (x *configWatchStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *configWatchStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *configWatchStream) Send(m *ConfigResponse) error {
	return x.stream.Send(m)
}
