package main

import (
	"log"
	"strings"
	"time"
)

type LoRaPreset struct {
	Name            string
	Description     string
	SpreadingFactor int
	Bandwidth       int
	CodingRate      int
	TxPower         int
	ATCommands      []string
}

var presets = map[string]LoRaPreset{
	"short_fast": {
		Name:        "Short Range & Fast",
		Description: "SF 7, 125kHz, CR 5, 14dBm",
		ATCommands: []string{
			"AT+SF=7",
			"AT+BW=125000",
			"AT+CR=5",
			"AT+TP=14",
		},
	},
	"long_slow": {
		Name:        "Long Range & Slow",
		Description: "SF 12, 125kHz, CR 8, 20dBm",
		ATCommands: []string{
			"AT+SF=12",
			"AT+BW=125000",
			"AT+CR=8",
			"AT+TP=20",
		},
	},
}

func GetPreset(name string) (LoRaPreset, bool) {
	preset, ok := presets[name]
	return preset, ok
}

func GetPresetNames() []string {
	names := make([]string, 0, len(presets))
	for name := range presets {
		names = append(names, name)
	}
	return names
}

func (sc *SerialConnection) ApplyPreset(preset LoRaPreset) {
	if sc.port == nil || !sc.connected {
		return
	}

	if !sc.enterATModeWithResponse() {
		return // Failed to enter AT mode
	}

	for _, cmd := range preset.ATCommands {
		if !sc.sendATCommandWithResponse(cmd) {
			break // Stop if a command fails
		}
	}

	sc.exitATModeWithResponse()
}

func (sc *SerialConnection) enterATModeWithResponse() bool {
	if sc.port == nil || !sc.connected {
		return false
	}

	sc.SendMessage("+++") // No line ending for entry usually, but let's stick to what's common
	time.Sleep(100 * time.Millisecond)
	return true
}

func (sc *SerialConnection) sendATCommandWithResponse(cmd string) bool {
	if sc.port == nil || !sc.connected {
		return false
	}

	sc.SendMessage(cmd + "\r\n")
	time.Sleep(100 * time.Millisecond)
	return true
}

func (sc *SerialConnection) exitATModeWithResponse() {
	if sc.port == nil || !sc.connected {
		return
	}

	sc.SendMessage("AT+EXIT\r\n")
	time.Sleep(100 * time.Millisecond)
}
