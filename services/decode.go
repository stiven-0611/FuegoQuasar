package services

import (
	"fmt"
	"strings"
)

func GetMessage(messages ...[]string) (msg string) {

	if messages == nil {
		fmt.Println("err")
		return ""
	}

	_, min := getMaxMin(messages)

	var decodeMsg []string
	var msm int = 0

	for i := 0; i < min; i++ {
		for _, msg := range messages {
			msm = len(msg) - min
			if len(msg[i+msm]) > 0 {
				decodeMsg = append(decodeMsg, msg[i+msm])
				break
			}
		}
	}

	return strings.Join(decodeMsg, " ")
}

func getMaxMin(messages [][]string) (int, int) {
	var min int = len(messages[0])
	var max int = 0
	for _, msg := range messages {
		if len(msg) < min {
			min = len(msg)
		}

		if len(msg) > max {
			max = len(msg)
		}
	}
	return max, min
}
