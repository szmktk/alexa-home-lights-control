package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

// TODO remove secret
const voiceMonkeyToken = "hard_coded_token"

var monkeys = map[string]string{
	"brightness_10":  "kl110study10",
	"brightness_25":  "kl110study25",
	"brightness_50":  "kl110study50",
	"brightness_75":  "kl110study75",
	"brightness_100": "kl110study100",
}

func trigger(monkeyKey string) {
	monkeyName, exists := monkeys[monkeyKey]
	if !exists {
		fmt.Fprintf(os.Stderr, "❌ Unknown monkey key: %s\n", monkeyKey)
		os.Exit(1)
	}

	url := fmt.Sprintf("https://api-v2.voicemonkey.io/trigger?token=%s&device=%s", voiceMonkeyToken, monkeyName)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Request failed: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("✅ Triggered routine: %s\n", monkeyKey)
	} else {
		fmt.Fprintf(os.Stderr, "❌ Failed: HTTP %d\n", resp.StatusCode)
		os.Exit(1)
	}
}

func main() {
	brightnessFlag := flag.Int("brightness", 0, "Set brightness (10, 25, 50, 75, 100)")

	flag.Parse()

	if *brightnessFlag != 0 {
		key := fmt.Sprintf("brightness_%d", *brightnessFlag)
		trigger(key)
	} else {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
