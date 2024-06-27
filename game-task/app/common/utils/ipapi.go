package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IPInfo struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Country string `json:"country"`
	Region  string `json:"regionName"`
	City    string `json:"city"`
}

func GetIPInfo(ip string) (*IPInfo, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get IP info: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ipInfo IPInfo
	if err := json.Unmarshal(body, &ipInfo); err != nil {
		return nil, err
	}

	if ipInfo.Status != "success" {
		return nil, fmt.Errorf("failed to get IP info: %v", ipInfo.Message)
	}

	return &ipInfo, nil
}
