package outputport

import (
	"context"
	"time"
)

type Cache interface {
	SetInt(ctx context.Context, key string, value int, expiration time.Duration) error
}
