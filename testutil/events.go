package testutil

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func FindEventsByType(events []sdk.Event, eventType string) (foundEvents []sdk.Event, found bool) {
	for _, event := range events {
		if event.Type == eventType {
			foundEvents = append(foundEvents, event)
		}
	}
	found = len(foundEvents) > 0
	return foundEvents, found
}

func EventHasAttributes(event sdk.Event, attributes map[string]string) bool {
	for key, value := range attributes {
		found := false
		for _, attr := range event.Attributes {
			if attr.Key == key && attr.Value == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
