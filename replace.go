package main

import (
		"regexp"
		"fmt"
		"os"
		"strings"
		"strconv"
		"io/ioutil"
		)

func main() {
	expr := regexp.MustCompile("\\d{1,2}:\\d{2}(:\\d{2}|)")
	seconds := 0
	allInput,_ := ioutil.ReadAll(os.Stdin)
	input := string(allInput)
	fmt.Println(expr.ReplaceAllStringFunc(input, func (in string) string { 
			times := strings.Split(in, ":")
			toAdd := 0
			for _, timePart := range times {
				timePartInt,_ := strconv.Atoi(timePart)
				toAdd = toAdd * 60 + timePartInt
			}
			defer func() { seconds = seconds + toAdd }()
			secs := seconds % 60
			mins := (seconds / 60) % 60
			hours := (seconds / 60 / 60)
			if hours > 0 {
				return fmt.Sprintf("%s %d:%02d:%02d", in, hours, mins, secs)
			} else {
				return fmt.Sprintf("%s %d:%02d", in, mins, secs)
			}
		}))
	fmt.Println("Thanks to: https://github.com/arranubels/albumCounter")
}
