package utils

import (
	"fmt"
	"time"
)

func ShowTime() string {
	t := time.Now().Format("15:04:05")
	fmt.Println("time: ", t)
	return t
}
