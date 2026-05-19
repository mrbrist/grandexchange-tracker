package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Items []struct {
	Examine  string `json:"examine"`
	ID       int    `json:"id"`
	Members  bool   `json:"members"`
	Lowalch  int    `json:"lowalch,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Value    int    `json:"value,omitempty"`
	Highalch int    `json:"highalch,omitempty"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
}

func GetItemMappingData(appName string) (Items, error) {
	url := "https://prices.runescape.wiki/api/v1/osrs/mapping"

	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	fmt.Printf("Fetching Item Mapping Data from %s...\n", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Recommended by the RuneScape Wiki API
	req.Header.Set("User-Agent", appName)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mapping data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var items Items

	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return items, nil
}
