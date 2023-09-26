package components

import (
	"testing"
	"time"
)

func TestNextFridayAndSunday(t *testing.T) {
	today := time.Date(2023, time.September, 19, 0, 0, 0, 0, time.UTC)

	expectedFriday := time.Date(2023, time.September, 22, 0, 0, 0, 0, time.UTC)
	expectedSunday := time.Date(2023, time.September, 24, 0, 0, 0, 0, time.UTC)

	friday, sunday := nextFridayAndSunday(today)

	if friday != expectedFriday {
		t.Errorf("Next Friday is incorrect. Got: %s, Expected: %s", friday, expectedFriday)
	}

	if sunday != expectedSunday {
		t.Errorf("Next Sunday is incorrect. Got: %s, Expected: %s", sunday, expectedSunday)
	}
}
