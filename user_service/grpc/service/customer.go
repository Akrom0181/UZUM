package service

import (
	"context"
	"microservice/config"
	"microservice/genproto/user_service"
	"microservice/grpc/client"
	"microservice/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type UserService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *UserService {
	return &UserService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *UserService) Create(ctx context.Context, req *user_service.CreateCustomer) (resp *user_service.Customer, err error) {

	f.log.Info("---CreateCustomer--->>>", logger.Any("req", req))

	resp, err = f.strg.Customer().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateCustomer--->>>", logger.Error(err))
		return &user_service.Customer{}, err
	}

	return resp, nil
}

func (f *UserService) GetByID(ctx context.Context, req *user_service.CustomerPrimaryKey) (resp *user_service.Customer, err error) {
	f.log.Info("---GetSingleCustomer--->>>", logger.Any("req", req))

	resp, err = f.strg.Customer().GetByID(ctx, req)
	if err != nil {
		f.log.Error("---GetSingleCustomer--->>>", logger.Error(err))
		return &user_service.Customer{}, err
	}

	return resp, nil
}

func (f *UserService) GetList(ctx context.Context, req *user_service.GetListCustomerRequest) (resp *user_service.GetListCustomerResponse, err error) {

	f.log.Info("---GetAllCustomer--->>>", logger.Any("req", req))

	resp, err = f.strg.Customer().GetList(ctx, req)
	if err != nil {
		f.log.Error("---GetAllCustomer--->>>", logger.Error(err))
		return &user_service.GetListCustomerResponse{}, err
	}

	return resp, nil
}

func (f *UserService) Update(ctx context.Context, req *user_service.UpdateCustomer) (resp *user_service.Customer, err error) {
	f.log.Info("---UpdateCustomer--->>>", logger.Any("req", req))

	resp, err = f.strg.Customer().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateCustomer--->>>", logger.Error(err))
		return &user_service.Customer{}, err
	}

	return resp, nil
}

func (f *UserService) Delete(ctx context.Context, req *user_service.CustomerPrimaryKey) (resp *user_service.Empty, err error) {
	f.log.Info("---DeleteCustomer--->>>", logger.Any("req", req))

	err = f.strg.Customer().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteCustomer--->>>", logger.Error(err))
		return &user_service.Empty{}, err
	}

	return &user_service.Empty{}, nil
}
