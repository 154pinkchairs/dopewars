package helpers

// This file contains utilities and data structures for interacting with the environmental variables and system metrics

import (
	"encoding/binary"
	"encoding/json"
	"os"
	"os/exec"
	"runtime"

	"github.com/lxn/win"
)

func getSettings() (int, error) {
	var settings map[string]interface{}
	file, err := os.Open("../settings.json")
	if err != nil {
		return 0, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&settings)
	if err != nil {
		return 0, err
	}
	return int(settings["Monitor Number"].(float64)), nil
}

func getMaxX(x int) int {
	if runtime.GOOS == "windows" {
		x = int(win.GetSystemMetrics(win.SM_CXSCREEN))
		return x
	} else if runtime.GOOS == "linux" {
		mointorno, err := getSettings()
		width, err := exec.Command("xdpyinfo", "-display", ":"+string(mointorno), "|", "grep", "dimensions:", "|", "awk", "'{print $2}'").Output()
		if err != nil {
			panic(err)
		}
		x = int(binary.BigEndian.Uint32(width))
		return x
	} else {
		width, err := exec.Command("system_profiler", "SPDisplaysDataType", "|", "grep", "Resolution:", "|", "awk", "'{print $2}'").Output()
		if err != nil {
			panic(err)
		}
		x = int(binary.BigEndian.Uint32(width))
		return x
	}
	return x
}

func getMaxY(y int) int {
	if runtime.GOOS == "windows" {
		y = int(win.GetSystemMetrics(win.SM_CYSCREEN))
		return y
	} else if runtime.GOOS == "linux" {
		mointorno, err := getSettings()
		// we must also separate by the "x" in the string, using awk -F "x""
		height, err := exec.Command("xdpyinfo", "-display", ":"+string(mointorno), "|", "grep", "dimensions:", "|", "awk", "-F", "x", "'{print $2}'").Output()
		if err != nil {
			panic(err)
		}
		y = int(binary.BigEndian.Uint32(height))
		return y
	} else {
		height, err := exec.Command("system_profiler", "SPDisplaysDataType", "|", "grep", "Resolution:", "|", "awk", "'{print $2}'").Output()
		if err != nil {
			panic(err)
		}
		y = int(binary.BigEndian.Uint32(height))
		return y
	}
	return y
}
