package service

import (
	"context"
	"github.com/fastmall/cart/api"
)

type CartService struct {
	api.UnimplementedCartServiceServer
}

func (c CartService) GetCustomerCart(ctx context.Context, request *api.GetCustomerCartRequest) (*api.GetCustomerCartResponse, error) {
	res := api.GetCustomerCartResponse{
		CustomerId: request.CustomerId,
		List:       nil,
	}
	return &res, nil
}

func (c CartService) mustEmbedUnimplementedCartServiceServer() {
	panic("implement me")
}
