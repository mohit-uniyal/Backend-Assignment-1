package outputport

import (
	"context"
	"event-booking/src/internal/core/model"
)

type EventsRepo interface {
	CreateEvent(ctx context.Context, event *model.Event) (int, error)
}
