package service

import (
	"context"
	"microservice/config"
	"microservice/genproto/user_service"
	"microservice/grpc/client"
	"microservice/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type BranchService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
}

func NewBranchService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *BranchService {
	return &BranchService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (b *BranchService) Create(ctx context.Context, req *user_service.CreateBranch) (*user_service.Branch, error) {
	b.log.Info("---CreateBranch--->>>", logger.Any("req", req))

	resp, err := b.strg.Branch().Create(ctx, req)
	if err != nil {
		b.log.Error("---CreateBranch--->>>", logger.Error(err))
		return &user_service.Branch{}, err
	}

	return resp, nil
}

func (b *BranchService) GetByID(ctx context.Context, req *user_service.BranchPrimaryKey) (*user_service.Branch, error) {
	b.log.Info("---GetSingleBranch--->>>", logger.Any("req", req))

	resp, err := b.strg.Branch().GetByID(ctx, req)
	if err != nil {
		b.log.Error("---GetSingleBranch--->>>", logger.Error(err))
		return &user_service.Branch{}, err
	}

	return resp, nil
}

func (b *BranchService) GetList(ctx context.Context, req *user_service.GetListBranchRequest) (resp *user_service.GetListBranchResponse, err error) {
	b.log.Info("---GetAllBranch--->>>", logger.Any("req", req))

	resp, err = b.strg.Branch().GetList(ctx, req)
	if err != nil {
		b.log.Error("---GetAllBranch--->>>", logger.Error(err))
		return &user_service.GetListBranchResponse{}, err
	}

	return resp, nil
}

func (b *BranchService) Update(ctx context.Context, req *user_service.UpdateBranch) (resp *user_service.Branch, err error) {
	b.log.Info("---UpdateBranch--->>>", logger.Any("req", req))

	resp, err = b.strg.Branch().Update(ctx, req)
	if err != nil {
		b.log.Error("---UpdateBranch--->>>", logger.Error(err))
		return &user_service.Branch{}, err
	}

	return resp, nil
}

func (b *BranchService) Delete(ctx context.Context, req *user_service.BranchPrimaryKey) (resp *user_service.Empty3, err error) {
	b.log.Info("---DeleteBranch--->>>", logger.Any("req", req))

	err = b.strg.Branch().Delete(ctx, req)
	if err != nil {
		b.log.Error("---DeleteBranch--->>>", logger.Error(err))
		return &user_service.Empty3{}, err
	}

	return &user_service.Empty3{}, nil
}
