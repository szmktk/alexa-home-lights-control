package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

const voiceMonkeyToken = "YOUR_VOICEMONKEY_API_TOKEN"

var monkeys = map[string]string{
	"on":             "your_monkey_name_for_on",
	"off":            "your_monkey_name_for_off",
	"brightness_25":  "your_monkey_name_for_25_brightness",
	"brightness_50":  "your_monkey_name_for_50_brightness",
	"brightness_75":  "your_monkey_name_for_75_brightness",
	"brightness_100": "your_monkey_name_for_100_brightness",
}

func trigger(monkeyKey string) {
	monkeyName, exists := monkeys[monkeyKey]
	if !exists {
		fmt.Fprintf(os.Stderr, "❌ Unknown monkey key: %s\n", monkeyKey)
		os.Exit(1)
	}

	url := fmt.Sprintf("https://api.voicemonkey.io/trigger?token=%s&monkey=%s", voiceMonkeyToken, monkeyName)

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
	onFlag := flag.Bool("on", false, "Turn the light on")
	offFlag := flag.Bool("off", false, "Turn the light off")
	brightnessFlag := flag.Int("brightness", 0, "Set brightness (25, 50, 75, 100)")

	flag.Parse()

	if *onFlag {
		trigger("on")
	} else if *offFlag {
		trigger("off")
	} else if *brightnessFlag != 0 {
		key := fmt.Sprintf("brightness_%d", *brightnessFlag)
		trigger(key)
	} else {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
