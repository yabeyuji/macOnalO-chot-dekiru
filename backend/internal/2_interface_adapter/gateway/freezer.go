package gateway

import (
	"context"
)

// GetPatties ...
func (gw *Gateway) GetPatties(ctx context.Context, requestPatties map[string]int) error {
	err := gw.ToFreezer.UpdatePatties(ctx, requestPatties)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}
