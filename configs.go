package main

import (
	"fmt"
	"strings"

	"github.com/parakeet-nest/parakeet/enums/option"
)

type ConfigKey string

const (
	// Configs
	ConfigUltraPrecise     ConfigKey = "Ultra-Precise"
	ConfigMirostatBalanced ConfigKey = "Mirostat-Balanced"
	ConfigMirostat2Dynamic ConfigKey = "Mirostat2-Dynamic"
	ConfigCreativeHigh     ConfigKey = "Creative-High"
	ConfigAnalytical       ConfigKey = "Analytical"
)

// Configs map declaration
var Configs = map[ConfigKey]map[string]interface{}{
	ConfigCreativeHigh: {
		option.Temperature:     0.9,
		option.TopP:            0.9,
		option.TopK:            60,
		option.RepeatPenalty:   1.05,
		option.PresencePenalty: 0.4,
	},
	ConfigUltraPrecise: {
		option.Temperature:     0.2,
		option.TopP:            0.4,
		option.TopK:            30,
		option.RepeatPenalty:   1.2,
		option.PresencePenalty: 0.3,
		option.NumPredict:      1024,
	},
	ConfigMirostatBalanced: {
		option.Mirostat:        1,
		option.MirostatTau:     5.0,
		option.MirostatEta:     0.3,
		option.Temperature:     0.4,
		option.RepeatPenalty:   1.1,
		option.PresencePenalty: 0.3,
	},
	ConfigMirostat2Dynamic: {
		option.Mirostat:        2,
		option.MirostatTau:     5.0,
		option.MirostatEta:     0.2,
		option.Temperature:     0.6,
		option.RepeatPenalty:   1.15,
		option.PresencePenalty: 0.25,
	},
	ConfigAnalytical: {
		option.Mirostat:        1,
		option.MirostatTau:     3.5,
		option.MirostatEta:     0.12,
		option.Temperature:     0.4,
		option.RepeatPenalty:   1.15,
		option.PresencePenalty: 0.2,
	},
}

// Get all configs as a slice
func AllConfigs() []ConfigKey {
	configs := make([]ConfigKey, 0, len(Configs))
	for c := range Configs {
		configs = append(configs, c)
	}
	return configs
}

func ParseConfigs(input string) ([]ConfigKey, error) {
	if input == "" {
		return nil, nil
	}

	items := strings.Split(input, ",")
	result := make([]ConfigKey, 0, len(items))

	for _, item := range items {
		key := ConfigKey(strings.TrimSpace(item))
		if _, exists := Configs[key]; !exists {
			return nil, fmt.Errorf("invalid config key: %s", item)
		}
		result = append(result, key)
	}

	return result, nil
}
