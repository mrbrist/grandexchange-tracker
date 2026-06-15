package main

import (
	"context"
	"cron/internal/database"
	"fmt"
	"log"
	"time"
)

func GenerateWeekTimestamps() []int64 {
	now := time.Now().Truncate(time.Hour)

	timestamps := make([]int64, 24*7)
	for i := 0; i < len(timestamps); i++ {
		timestamps[i] = now.Add(-time.Duration(len(timestamps)-1-i) * time.Hour).Unix()
	}

	return timestamps
}

func GetMissingTimestamps(DB *database.Queries) {
	timestamps, err := DB.GetUniqueTimestamps(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	allTimestamps := GenerateWeekTimestamps()

	var result []int64

	const tolerance = 3600

	for _, target := range allTimestamps {
		found := false

		for _, existing := range timestamps {
			diff := target - existing
			if diff < 0 {
				diff = -diff
			}

			if diff <= tolerance {
				found = true
				break
			}
		}

		if !found {
			result = append(result, target)
		}
	}

	fmt.Println(len(allTimestamps))
	fmt.Println(len(timestamps))
	fmt.Println(len(result))
}
