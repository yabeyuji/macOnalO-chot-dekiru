package presenter

import (
	"context"
)

// Monitor ...
func (prsntr *Presenter) UpdateOrders(ctx context.Context, orderNumber string, phase string) {
	// オーダー情報更新
	prsntr.ToMonitor.UpdateOrders(ctx, orderNumber, phase)
	return
}
