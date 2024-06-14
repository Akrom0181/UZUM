package client

import (
	"microservice/config"
)

type ServiceManagerI interface{}

type grpcClients struct{}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	return &grpcClients{}, nil
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
