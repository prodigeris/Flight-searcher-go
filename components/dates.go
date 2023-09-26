package components

import "time"

func nextFridayAndSunday(today time.Time) (friday time.Time, sunday time.Time) {
	daysUntilFriday := 5 - int(today.Weekday())
	if daysUntilFriday <= 0 {
		daysUntilFriday += 7
	}
	friday = today.Add(time.Hour * 24 * time.Duration(daysUntilFriday))

	sunday = friday.Add(time.Hour * 24 * 2)

	return friday, sunday
}
