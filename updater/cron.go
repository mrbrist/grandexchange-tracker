package main

import (
	"context"
	"cron/internal/database"
	"cron/scheduler"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
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

func StartScheduling(appName string, DB *database.Queries) (*scheduler.Scheduler, error) {
	s, err := scheduler.New()
	if err != nil {
		return nil, err
	}

	err = s.AddAndRunNow(
		"price-history-1h",
		"0 * * * *",
		func() error {
			return UpdatePriceHistory1h(appName, DB)
		},
		true,
	)
	if err != nil {
		return nil, err
	}

	err = s.AddAndRunNow(
		"get-missing-timestamps",
		"0 * * * *",
		func() error {
			return GetMissingTimestamps(DB, &globalMissing)
		},
		true,
	)
	if err != nil {
		return nil, err
	}

	err = s.Add(
		"update-missing-timestamp",
		"*/1 * * * *",
		func() error {
			if len(globalMissing) == 0 {
				return nil
			}

			ts := globalMissing[0]
			globalMissing = globalMissing[1:]

			return UpdatePriceHistoryForTimestamp(appName, DB, ts)
		},
	)
	if err != nil {
		return nil, err
	}

	s.Start()
	return s, nil
}

func UpdatePriceHistoryForTimestamp(appName string, DB *database.Queries, timestamp int64) error {
	// fmt.Println(timestamp)
	// return nil

	url := fmt.Sprintf("https://prices.runescape.wiki/api/v1/osrs/1h?timestamp=%d", timestamp)

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	fmt.Printf("Fetching %d Item Price History Data from %s...\n", timestamp, url)

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
	fmt.Printf("%sAdded new data to database with timestamp: %d%s\n", ColorPurple, pricehistory1h.Timestamp, ColorNone)
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
	fmt.Printf("%sAdded new data to database with timestamp: %d%s\n", ColorPurple, pricehistory1h.Timestamp, ColorNone)
	fmt.Printf("%swaiting for next update in 1h...%s\n", ColorGreen, ColorNone)
	return nil
}
