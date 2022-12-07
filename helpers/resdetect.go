//go:build darwin || linux

package helpers

// This file contains utilities and data structures for interacting with the environmental variables and system metrics
import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

/* this feature (settings menu/display selection) is not yet implemented. Import encoding/json to use this function
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
	dispno := int(settings["Monitor Number"].(float64))
	return dispno, nil
}
*/

func GetMaxX() int {
	var x int
	var width []byte
	var err error
	if runtime.GOOS == "linux" {
		dispno, _ := strconv.Atoi(os.Getenv("DISPLAY"))
		if err != nil {
			dispno = 0
		}
		//wrapper around xdpyinfo -display :0 | grep dimensions: | awk '{print $2}' | awk -F x '{print $1}'
		width, _ = exec.Command("xdpyinfo", "-display", ":"+fmt.Sprintf("%d", dispno), "|", "grep", "dimensions:", "|", "awk", "'{print $2}'", "|", "awk", "-F", "x", "'{print $1}'").Output()
		if err != nil {
			//use the helper Shell script ./width.sh
			width, err = exec.Command("./helpers/width.sh").Output()
			if err != nil {
				log.Printf("Error getting width: %v", err)
				log.Printf("Using default width of 1152")
				width = []byte{0x04, 0x60} // 1152
			}
		}
	} else {
		width, err = exec.Command("system_profiler", "SPDisplaysDataType", "|", "grep", "Resolution:", "|", "awk", "'{print $2}'").Output()
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
	}
	//width is a byte slice that contains the ascii table values (in dec notation) for each digit of the width
	//we must convert it to an int
	xstr := bytes.NewBuffer(width).String()
	x, _ = strconv.Atoi(xstr)
	if x == 0 {
		x = 1152
		log.Println("Failed to get screen width, using default value of 1152")
	}
	return x
}

func GetMaxY() int {
	var y int
	var height []byte
	var err error
	if runtime.GOOS == "linux" {
		dispno, _ := strconv.Atoi(os.Getenv("DISPLAY"))
		height, _ = exec.Command("xdpyinfo", "-display", ":"+fmt.Sprintf("%d", dispno), "|", "grep", "dimensions:", "|", "awk", "-F", "x", "'{print $2}'").Output()
		if err != nil {
			//use the helper Shell script ./height.sh
			height, err = exec.Command("./helpers/height.sh").Output()
			if err != nil {
				log.Printf("Error getting height: %v", err)
				log.Printf("Using default height of 864")
				height = []byte{0x03, 0x60} // 864
			}
		}
	} else {
		// macOS
		height, err = exec.Command("system_profiler", "SPDisplaysDataType", "|", "grep", "Resolution:", "|", "awk", "'{print $2}'").Output()
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
	}
	ystr := bytes.NewBuffer(height).String()
	y, _ = strconv.Atoi(ystr)
	if y == 0 {
		y = 864
		log.Println("Failed to get screen height, using default value of 864")
	}
	return y
}
