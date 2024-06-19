package service

import (
	"context"
	"microservice/config"
	"microservice/genproto/user_service"
	"microservice/grpc/client"
	"microservice/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type UsService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewUsService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *UsService {
	return &UsService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (f *UsService) Create(ctx context.Context, req *user_service.CreateUs) (*user_service.Us, error) {

	f.log.Info("---CreateSystem_User--->>>", logger.Any("req", req))

	resp, err := f.strg.User().Create(ctx, req)
	if err != nil {
		f.log.Error("---CreateSystem_User--->>>", logger.Error(err))
		return &user_service.Us{}, err
	}

	return resp, nil
}

func (f *UsService) GetByID(ctx context.Context, req *user_service.UsPrimaryKey) (*user_service.Us, error) {
	f.log.Info("---GetSingleSystem_User--->>>", logger.Any("req", req))

	resp, err := f.strg.User().GetByID(ctx, req)
	if err != nil {
		f.log.Error("---GetSingleSystem_User--->>>", logger.Error(err))
		return &user_service.Us{}, err
	}

	return resp, nil
}

func (f *UsService) GetList(ctx context.Context, req *user_service.GetListUsRequest) (*user_service.GetListUsResponse, error) {

	f.log.Info("---GetAllSystem_User--->>>", logger.Any("req", req))

	resp, err := f.strg.User().GetList(ctx, req)
	if err != nil {
		f.log.Error("---GetAllSystem_User--->>>", logger.Error(err))
		return &user_service.GetListUsResponse{}, err
	}

	return resp, nil
}

func (f *UsService) Update(ctx context.Context, req *user_service.UpdateUs) (*user_service.Us, error) {
	f.log.Info("---UpdateSystem_User--->>>", logger.Any("req", req))

	resp, err := f.strg.User().Update(ctx, req)
	if err != nil {
		f.log.Error("---UpdateSystem_User--->>>", logger.Error(err))
		return &user_service.Us{}, err
	}

	return resp, nil
}

func (f *UsService) Delete(ctx context.Context, req *user_service.UsPrimaryKey) (*user_service.Empty1, error) {
	f.log.Info("---DeleteSystem_User--->>>", logger.Any("req", req))

	err := f.strg.User().Delete(ctx, req)
	if err != nil {
		f.log.Error("---DeleteSystem_User--->>>", logger.Error(err))
		return &user_service.Empty1{}, err
	}

	return &user_service.Empty1{}, nil
}
