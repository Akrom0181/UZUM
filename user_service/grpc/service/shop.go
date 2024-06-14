package service

import (
	"context"
	"microservice/config"
	"microservice/genproto/user_service"
	"microservice/grpc/client"
	"microservice/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ShopService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewShopService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *ShopService {
	return &ShopService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (s *ShopService) Create(ctx context.Context, req *user_service.CreateShop) (*user_service.Shop, error) {
	s.log.Info("---CreateShop--->>>", logger.Any("req", req))

	resp, err := s.strg.Shop().Create(ctx, req)
	if err != nil {
		s.log.Error("---CreateShop--->>>", logger.Error(err))
		return &user_service.Shop{}, err
	}

	return resp, nil
}

func (s *ShopService) GetByID(ctx context.Context, req *user_service.ShopPrimaryKey) (*user_service.Shop, error) {
	s.log.Info("---GetSingleShop--->>>", logger.Any("req", req))

	resp, err := s.strg.Shop().GetById(ctx, req)
	if err != nil {
		s.log.Error("---GetSingleShop--->>>", logger.Error(err))
		return &user_service.Shop{}, err
	}

	return resp, nil
}

func (s *ShopService) GetList(ctx context.Context, req *user_service.GetListShopRequest) (*user_service.GetListShopResponse, error) {
	s.log.Info("---GetAllShops--->>>", logger.Any("req", req))

	resp, err := s.strg.Shop().GetList(ctx, req)
	if err != nil {
		s.log.Error("---GetAllShops--->>>", logger.Error(err))
		return &user_service.GetListShopResponse{}, err
	}

	return resp, nil
}

func (s *ShopService) Update(ctx context.Context, req *user_service.UpdateShop) (*user_service.Shop, error) {
	s.log.Info("---UpdateShop--->>>", logger.Any("req", req))

	resp, err := s.strg.Shop().Update(ctx, req)
	if err != nil {
		s.log.Error("---UpdateShop--->>>", logger.Error(err))
		return &user_service.Shop{}, err
	}

	return resp, nil
}

func (s *ShopService) Delete(ctx context.Context, req *user_service.ShopPrimaryKey) (*user_service.Empty4, error) {
	s.log.Info("---DeleteShop--->>>", logger.Any("req", req))

	err := s.strg.Shop().Delete(ctx, req)
	if err != nil {
		s.log.Error("---DeleteShop--->>>", logger.Error(err))
		return &user_service.Empty4{}, err
	}

	return &user_service.Empty4{}, nil
}
