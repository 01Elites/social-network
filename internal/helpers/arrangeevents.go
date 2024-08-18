package helpers

import "social-network/internal/models"

func ArrangeEvents(events []models.Event) []models.Event {
	for i := 0; i < len(events)-1; i++ {
		for j := 0; j < len(events); j++ {
			if events[j].EventTime.Before(events[i].EventTime) {
				events[i], events[j] = events[j], events[i]
			}
		}
	}
	return events
}
