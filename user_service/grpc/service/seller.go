package service

import (
	"context"
	"microservice/config"
	"microservice/genproto/user_service"
	"microservice/grpc/client"
	"microservice/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type SellerService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewSellerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *SellerService {
	return &SellerService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *SellerService) Create(ctx context.Context, req *user_service.CreateSeller) (*user_service.Seller, error) {

	f.log.Info("---CreateSeller--->>>", logger.Any("req", req))

	resp, err := f.strg.Seller().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateSeller--->>>", logger.Error(err))
		return &user_service.Seller{}, err
	}

	return resp, nil
}

func (f *SellerService) GetByID(ctx context.Context, req *user_service.SellerPrimaryKey) (*user_service.Seller, error) {
	f.log.Info("---GetSingleSeller--->>>", logger.Any("req", req))

	resp, err := f.strg.Seller().GetByID(ctx, req)
	if err != nil {
		f.log.Error("---GetSingleSeller--->>>", logger.Error(err))
		return &user_service.Seller{}, err
	}

	return resp, nil
}

func (f *SellerService) GetList(ctx context.Context, req *user_service.GetListSellerRequest) (*user_service.GetListSellerResponse, error) {

	f.log.Info("---GetAllSellers--->>>", logger.Any("req", req))

	resp, err := f.strg.Seller().GetList(ctx, req)
	if err != nil {
		f.log.Error("---GetAllSellers--->>>", logger.Error(err))
		return &user_service.GetListSellerResponse{}, err
	}

	return resp, nil
}

func (f *SellerService) Update(ctx context.Context, req *user_service.UpdateSeller) (*user_service.Seller, error) {
	f.log.Info("---UpdateSeller--->>>", logger.Any("req", req))

	resp, err := f.strg.Seller().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateSeller--->>>", logger.Error(err))
		return &user_service.Seller{}, err
	}

	return resp, nil
}

func (f *SellerService) Delete(ctx context.Context, req *user_service.SellerPrimaryKey) (*user_service.Empty2, error) {
	f.log.Info("---DeleteSeller--->>>", logger.Any("req", req))

	err := f.strg.Seller().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteSeller--->>>", logger.Error(err))
		return &user_service.Empty2{}, err
	}

	return &user_service.Empty2{}, nil
}
