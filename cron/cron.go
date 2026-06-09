package main

import (
	"context"
	"cron/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type ItemPriceHistoryDataPoint struct {
	AvgHighPrice    int `json:"avgHighPrice"`
	HighPriceVolume int `json:"highPriceVolume"`
	AvgLowPrice     int `json:"avgLowPrice"`
	LowPriceVolume  int `json:"lowPriceVolume"`
}

type ItemPriceHistory1h struct {
	Data      map[string]ItemPriceHistoryDataPoint `json:"data"`
	Timestamp int64                                `json:"timestamp"`
}

func StartScheduling(appName string, DB *database.Queries) error {
	s, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	fmt.Println("Creating scheduler job...")

	ts, err := DB.GetLatestTimestamp(context.Background())
	if err != nil {
		return err
	}

	// now := time.Now().Unix()

	lastUpdate := time.Unix(ts, 0)
	fmt.Printf("%slast update:%s %d\n", ColorRed, ColorNone, lastUpdate.Unix())
	fmt.Printf("%snow:%s %d\n", ColorRed, ColorNone, time.Now().Unix())

	shouldRunNow := time.Now().Unix()-lastUpdate.Unix() >= int64(time.Hour.Seconds())*2

	fmt.Printf("%stime since update:%s %d\n", ColorRed, ColorNone, time.Now().Unix()-lastUpdate.Unix())

	_, err = s.NewJob(
		gocron.DurationJob(time.Hour),
		gocron.NewTask(func() error {
			fmt.Printf("%sRUNNING PRICE HISTORY JOB%s\n", ColorPurple, ColorNone)
			return UpdatePriceHistory1h(appName, DB)
		}),
	)
	if err != nil {
		return err
	}

	if shouldRunNow {
		fmt.Printf("%sData is stale → running initial update immediately%s\n", ColorYellow, ColorNone)
		go func() {
			_ = UpdatePriceHistory1h(appName, DB)
		}()
	} else {
		fmt.Printf("%sData is fresh → scheduler will handle next run%s\n", ColorYellow, ColorNone)
	}

	fmt.Println("Starting scheduler...")
	s.Start()

	fmt.Println("Scheduler started successfully")

	return nil
}

func UpdatePriceHistory1h(appName string, DB *database.Queries) error {
	url := "https://prices.runescape.wiki/api/v1/osrs/1h"

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	fmt.Printf("Fetching Item Price History 1h Data from %s...\n", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Recommended by the RuneScape Wiki API
	req.Header.Set("User-Agent", appName)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch history data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var pricehistory1h ItemPriceHistory1h

	if err := json.NewDecoder(resp.Body).Decode(&pricehistory1h); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	for id, item := range pricehistory1h.Data {
		id, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
		}
		err = DB.AddItemHistory(context.Background(), database.AddItemHistoryParams{
			ItemID:         int32(id),
			PriceTimestamp: pricehistory1h.Timestamp,
			AvgHighPrice:   int32(item.AvgHighPrice),
			AvgLowPrice:    int32(item.AvgLowPrice),
			HighVolume:     int64(item.HighPriceVolume),
			LowVolume:      int64(item.LowPriceVolume),
		})
		if err != nil {
			// Disable error from flooding terminal since its always duplicate error
			// fmt.Printf("failed to add history to database: %s", err)
			continue
		}
	}
	fmt.Printf("Added new data to database with timestamp: %d\n", pricehistory1h.Timestamp)
	fmt.Printf("%swaiting for next update in 1h...%s\n", ColorGreen, ColorNone)
	return nil
}
