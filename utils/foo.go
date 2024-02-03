package utils

import (
	"fmt"
	"time"
)

// show time in string format
// and also println a tag version
func ShowTime() string {
	t := time.Now().Format("15:04:05")
	fmt.Println("Time --> v4", t)
	return t
}
