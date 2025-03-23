// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateOrder(ctx context.Context, arg CreateOrderParams) (int32, error)
	CreateOrderProduct(ctx context.Context, arg CreateOrderProductParams) error
	CreateProduct(ctx context.Context, arg CreateProductParams) (int32, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (int32, error)
	DeleteCheapProducts(ctx context.Context, price pgtype.Numeric) (int64, error)
	DeleteOrder(ctx context.Context, idOrderMain int32) error
	DeleteUser(ctx context.Context, email string) error
	GetOrderByUserID(ctx context.Context, idUserF int32) ([]GetOrderByUserIDRow, error)
	GetProductsByPriceRange(ctx context.Context, arg GetProductsByPriceRangeParams) ([]SchemaProduct, error)
	GetUserSpendingStats(ctx context.Context) ([]GetUserSpendingStatsRow, error)
	GetUsersByPassword(ctx context.Context, password string) ([]SchemaUser, error)
	UpdateOrderTotal(ctx context.Context) error
	UpdateProductPrice(ctx context.Context, arg UpdateProductPriceParams) error
	UpdateUserName(ctx context.Context, arg UpdateUserNameParams) error
}

var _ Querier = (*Queries)(nil)
