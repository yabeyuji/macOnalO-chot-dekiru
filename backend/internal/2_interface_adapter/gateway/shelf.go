package gateway

import (
	"context"
)

// GetBans ...
func (gw *Gateway) GetBans(ctx context.Context, requestBans map[string]int) error {
	err := gw.ToShelf.UpdateBans(ctx, requestBans)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
