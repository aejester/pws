package api

func AlertArrayDifference(oldAlerts *[]Alert, newAlerts *[]Alert) *[]Alert {
	var diff []Alert

	if oldAlerts == nil {
		diff = *newAlerts
	} else {
		for _, newAlert := range *newAlerts {
			canAppend := true
			for _, oldAlert := range *oldAlerts {
				if newAlert.Description == oldAlert.Description {
					canAppend = false
				}
			}
			if canAppend {
				diff = append(diff, newAlert)
			}
		}
	}

	return &diff
}
