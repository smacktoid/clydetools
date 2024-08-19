package clydetools

import (
	"testing"
	"time"
)

func TestGetCurrentSeasonYearReturnsThisYearAfterAugust(t *testing.T) {
	expectedYear := 2024
	curentInstant := time.Date(expectedYear, 8, 01, 00, 00, 00, 0, time.UTC)

	CurrentTime = func() time.Time {
		return curentInstant
	}

	year := GetCurrentSeasonYear()
	if year != "2024" {
		t.Fatalf(`GetCurrentSeasonYear() should return the current year if it is invoked in August but got %s`, year)
	}
}

func TestGetCurrentSeasonYearReturnsLastYearBeforeAugust(t *testing.T) {
	expectedYear := 2024
	curentInstant := time.Date(expectedYear, 7, 01, 00, 00, 00, 0, time.UTC)

	CurrentTime = func() time.Time {
		return curentInstant
	}

	year := GetCurrentSeasonYear()
	if year != "2023" {
		t.Fatalf(`GetCurrentSeasonYear() should return the previous year if it is invoked before August but got %s`, year)
	}
}

func TestGetFixturesFailsIfNoAPIKeyIsSet(t *testing.T) {
	_, actual := GetFixtures()

	Expected := "CLYDETOOLS_API_KEY is not set"
	if actual.Error() != Expected {
		t.Errorf("Error actual = %v, and Expected = %v.", actual, Expected)
	}
}
