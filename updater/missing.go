package main

import (
	"context"
	"cron/internal/database"
	"fmt"
	"log"
)

func GetMissingTimestamps(DB *database.Queries) {
	timestamps, err := DB.GetUniqueTimestamps(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(timestamps)
}
