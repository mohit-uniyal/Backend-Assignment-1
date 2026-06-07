package inputport

import (
	"context"
	"event-booking/src/internal/core/dto"
)

type EventsUsecase interface {
	CreateEvent(ctx context.Context, event *dto.Event) (*dto.Event, error)
}
