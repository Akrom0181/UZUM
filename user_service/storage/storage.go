package storage

import (
	"context"
	us "microservice/genproto/user_service"
)

type StorageI interface {
	CloseDB()
	Customer() CustomerRepoI
	User() UserRepoI
	Seller() SellerRepoI
	Branch() BranchRepoI
	Shop() ShopRepoI
}

type CustomerRepoI interface {
	Create(ctx context.Context, req *us.CreateCustomer) (*us.Customer, error)
	GetByID(ctx context.Context, req *us.CustomerPrimaryKey) (*us.Customer, error)
	GetList(ctx context.Context, req *us.GetListCustomerRequest) (*us.GetListCustomerResponse, error)
	Update(ctx context.Context, req *us.UpdateCustomer) (*us.Customer, error)
	Delete(ctx context.Context, req *us.CustomerPrimaryKey) error
}

type UserRepoI interface {
	Create(ctx context.Context, req *us.CreateUs) (*us.Us, error)
	GetByID(ctx context.Context, req *us.UsPrimaryKey) (*us.Us, error)
	GetList(ctx context.Context, req *us.GetListUsRequest) (*us.GetListUsResponse, error)
	Update(ctx context.Context, req *us.UpdateUs) (*us.Us, error)
	Delete(ctx context.Context, req *us.UsPrimaryKey) error
}

type SellerRepoI interface {
	Create(ctx context.Context, req *us.CreateSeller) (*us.Seller, error)
	GetByID(ctx context.Context, req *us.SellerPrimaryKey) (*us.Seller, error)
	GetList(ctx context.Context, req *us.GetListSellerRequest) (*us.GetListSellerResponse, error)
	Update(ctx context.Context, req *us.UpdateSeller) (*us.Seller, error)
	Delete(ctx context.Context, req *us.SellerPrimaryKey) error
}

type BranchRepoI interface {
	Create(ctx context.Context, req *us.CreateBranch) (*us.Branch, error)
	GetByID(ctx context.Context, req *us.BranchPrimaryKey) (*us.Branch, error)
	GetList(ctx context.Context, req *us.GetListBranchRequest) (*us.GetListBranchResponse, error)
	Update(ctx context.Context, req *us.UpdateBranch) (*us.Branch, error)
	Delete(ctx context.Context, req *us.BranchPrimaryKey) error
}

type ShopRepoI interface {
	Create(ctx context.Context, req *us.CreateShop) (*us.Shop, error)
	GetById(ctx context.Context, req *us.ShopPrimaryKey) (*us.Shop, error)
	GetList(ctx context.Context, req *us.GetListShopRequest) (*us.GetListShopResponse, error)
	Update(ctx context.Context, req *us.UpdateShop) (*us.Shop, error)
	Delete(ctx context.Context, req *us.ShopPrimaryKey) error
}
