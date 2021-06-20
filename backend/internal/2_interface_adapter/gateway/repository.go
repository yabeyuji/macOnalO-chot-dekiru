package gateway

import (
	"context"

	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("interface_adapter", "gateway")
}

type (
	Gateway struct {
		ToRefrigerator ToRefrigerator
		ToFreezer      ToFreezer
		ToShelf        ToShelf
	}

	// ToRefrigerator ...
	ToRefrigerator interface {
		UpdateVegetables(ctx context.Context, items map[string]int) error
		UpdateIngredients(ctx context.Context, items map[string]int) error
	}

	// ToFreezer ...
	ToFreezer interface {
		UpdatePatties(ctx context.Context, items map[string]int) error
	}

	// ToShelf ...
	ToShelf interface {
		UpdateBans(ctx context.Context, items map[string]int) error
	}
)

// NewGateway ...
func NewGateway(toRefrigerator ToRefrigerator, toFreezer ToFreezer, toShelf ToShelf) *Gateway {
	return &Gateway{
		ToRefrigerator: toRefrigerator,
		ToFreezer:      toFreezer,
		ToShelf:        toShelf,
	}
}
