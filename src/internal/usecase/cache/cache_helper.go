package cacheservice

import (
	"event-booking/src/internal/core/model"
	"fmt"
)

func getEventTicketsKey(event *model.Event) string {
	if event == nil {
		return ""
	}

	return fmt.Sprintf("event:%d:tickets", event.EventId)
}
