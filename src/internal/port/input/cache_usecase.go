package inputport

import "context"

type CacheUsecase interface {
	PopulateEvents(ctx context.Context) error
}
