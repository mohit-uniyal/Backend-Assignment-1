package inputport

import (
	"context"
	"event-booking/src/internal/core/model"
)

type CacheUsecase interface {
	PopulateEvents(ctx context.Context) error
	SetEvent(ctx context.Context, event *model.Event) error
}
