package grpc_client

import (
	"fmt"
	"log"
	pc "microservice/genproto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"microservice/config"
)

// GrpcClientI ..90.
type GrpcClientI interface {
	UserService() pc.CustomerServiceClient
	SystemUserService() pc.UsServiceClient
	SellerServive() pc.SellerServiceClient
	BranchService() pc.BranchServiceClient
	ShopService() pc.ShopServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port:%s err: %s",
			cfg.UserServiceHost, cfg.UserServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"user_service": pc.NewCustomerServiceClient(connUser),
			"system_user":  pc.NewUsServiceClient(connUser),
			"seller":       pc.NewSellerServiceClient(connUser),
			"branch":       pc.NewBranchServiceClient(connUser),
			"shop":			pc.NewShopServiceClient(connUser),
		},
	}, nil
}

func (g *GrpcClient) UserService() pc.CustomerServiceClient {
	client, ok := g.connections["user_service"].(pc.CustomerServiceClient)
	if !ok {
		log.Println("failed to assert type for user_service")
		return nil
	}
	return client
}

func (g *GrpcClient) SystemUserService() pc.UsServiceClient {
	client, ok := g.connections["system_user"].(pc.UsServiceClient)
	if !ok {
		log.Println("failed to assert type for system_user")
		return nil
	}
	return client
}

func (g *GrpcClient) SellerService() pc.SellerServiceClient {
	client, ok := g.connections["seller"].(pc.SellerServiceClient)
	if !ok {
		log.Println("failed to assert type for seller")
		return nil
	}
	return client
}

func (g *GrpcClient) BranchService() pc.BranchServiceClient {
	client, ok := g.connections["branch"].(pc.BranchServiceClient)
	if !ok {
		log.Println("failed to assert type for branch")
		return nil
	}
	return client
}

func (g *GrpcClient) ShopService() pc.ShopServiceClient {
	client, ok := g.connections["shop"].(pc.ShopServiceClient)
	if !ok {
		log.Println("failed to assert type for shop")
		return nil
	}
	return client
}


