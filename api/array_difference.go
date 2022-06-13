package api

func AlertArrayDifference(oldAlerts *[]Alert, newAlerts *[]Alert) *[]Alert {
	var diff []Alert

	for _, newAlert := range *newAlerts {
		for _, oldAlert := range *oldAlerts {
			if newAlert.Description != oldAlert.Description {
				diff = append(diff, newAlert)
			}
		}

	}

	return &diff
}
