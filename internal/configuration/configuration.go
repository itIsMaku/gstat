package configuration

import (
	"encoding/json"
	"os"
)

type Config struct {
	HistoryDir string         `json:"historyDirectory"`
	Interval   IntervalConfig `json:"interval"`
}

type IntervalConfig struct {
	Seconds int      `json:"seconds"`
	Targets []Target `json:"targets"`
}

type Target struct {
	Protocol  string `json:"protocol"`
	Target    string `json:"target"`
	Reachable bool   `json:"reachable"`
}

func CreateConfig(fileName string) error {
	defaultConfig := Config{
		HistoryDir: "history",
		Interval: IntervalConfig{
			Seconds: 60,
			Targets: []Target{
				{
					Protocol:  "tcp",
					Target:    "example.com:80",
					Reachable: true,
				},
			},
		},
	}

	if _, err := os.Stat(fileName); err == nil {
		return nil
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(defaultConfig); err != nil {
		return err
	}

	return nil
}

func LoadConfig(fileName string) (*Config, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
