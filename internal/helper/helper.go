package helper

import (
	"log"
	"os"
	"slices"
	"time"
)

// envLookup returns the value of the given environment variable
func EnvLookup(key string) string {
	val, ok := os.LookupEnv(key)

	if !ok && key != "PROVIDER" {
		log.Fatalf("failed to retrieve %s\n", key)
	}

	if val == "" && key != "PROVIDER" {
		log.Fatalf("failed to find value for %s\n", key)
	}

	return val
}

/*
SortByDate sorts the given dates in ascending order
TODO: Refactor to use generics to support additional types
*/

func SortByDate(dates []string) ([]time.Time, error) {
	var timeDates []time.Time

	for _, date := range dates {
		timeDate, err := time.Parse(time.DateOnly, date)
		if err != nil {
			return nil, err
		}
		timeDates = append(timeDates, timeDate)
	}

	slices.SortFunc(timeDates, func(a, b time.Time) int {
		return a.Compare(b)
	})

	return timeDates, nil
}
