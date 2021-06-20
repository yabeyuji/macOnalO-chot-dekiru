package presenter

import (
	"context"

	"backend/internal/4_enterprise_business_rule/entity"
)

// Shipment ...
func (prsntr *Presenter) Shipment(ctx context.Context, order *entity.Order) error {
	// 商品の出荷
	err := prsntr.ToShipment.PutProducts(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	// 商品の出荷記録
	err = prsntr.ToShipment.WriteLog(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
