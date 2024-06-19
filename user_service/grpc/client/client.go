package client

import (
	"fmt"
	"microservice/config"
	"microservice/genproto/catalog_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	CategoryService() catalog_service.CategoryServiceClient
	// ProductCategoryService() catalog_service.CategoryServiceClient

}

type grpcClients struct {
	categoryService catalog_service.CategoryServiceClient
	// productCategoryService catalog_service.CategoryServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	connCategoryService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.CatalogServicePort, cfg.CatalogServiceHost),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	)
	if err != nil {
		return nil, err
	}


	// connProductCategoryService, err := grpc.Dial(
	// 	fmt.Sprintf("%s%s", cfg.CatalogServicePort, cfg.CatalogServiceHost),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// 	grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(52428800), grpc.MaxCallSendMsgSize(52428800)),
	// )
	// if err != nil {
	// 	return nil, err
	// }

	return &grpcClients{
		categoryService: catalog_service.NewCategoryServiceClient(connCategoryService),
		// productCategoryService: catalog_service.CategoryServiceServer(connProductCategoryService),
	}, nil
}

func (g *grpcClients) CategoryService() catalog_service.CategoryServiceClient {
	return g.categoryService
}




































// 	connUser, err := grpc.Dial(
// 		fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort),
// 		grpc.WithTransportCredentials(insecure.NewCredentials()))

// 	if err != nil {
// 		return nil, fmt.Errorf("user service dial host: %s port:%s err: %s",
// 			cfg.UserServiceHost, cfg.UserServicePort, err)
// 	}

// 	return &GrpcClient{
// 		cfg: cfg,
// 		connections: map[string]interface{}{
// 			"user_service": user_service.NewUserServiceClient(connUser),
// 		},
// 	}, nil
// }

// func (g *GrpcClient) UserService() user_service.UserServiceClient {
// 	return g.connections["user_service"].(user_service.UserServiceClient)
// }

// // Implement the Customer() method
// func (g *GrpcClient) User() user_service.UserServiceClient {
// 	return g.UserService() // Return the customer client
// }

// // type ServiceManagerI interface{}

// // type grpcClients struct{}

// // func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
