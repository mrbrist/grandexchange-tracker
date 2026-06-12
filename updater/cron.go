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

var globalJob gocron.Job

func PrintNextRun() {
	nextRun, err := globalJob.NextRun()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(
		"Next run: %s\n",
		nextRun.Format(time.RFC822),
	)
}

func StartScheduling(appName string, DB *database.Queries) (gocron.Scheduler, error) {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	ts, err := DB.GetLatestTimestamp(context.Background())
	if err != nil {
		return nil, err
	}

	lastUpdate := time.Unix(ts, 0)

	job, err := s.NewJob(
		gocron.CronJob(
			//"*/10 * * * *", // 10 mins
			"0 * * * *", // 1 hr
			false,
		),
		gocron.NewTask(func() error {
			return UpdatePriceHistory1h(appName, DB)
		}),
	)
	if err != nil {
		return nil, err
	}

	globalJob = job

	shouldRunNow := time.Since(lastUpdate) >= 2*time.Hour

	if shouldRunNow {
		fmt.Printf("%sData is stale → running update immediately%s\n",
			ColorYellow, ColorNone)

		go func() {
			_ = UpdatePriceHistory1h(appName, DB)
		}()
	} else {
		fmt.Printf("%sData is recent → waiting...%s\n",
			ColorYellow, ColorNone)
	}

	s.Start()

	PrintNextRun()

	return s, nil
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
	PrintNextRun()
	return nil
}
