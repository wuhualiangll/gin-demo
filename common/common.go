package common

import (
	"fmt"
	"time"
)

func UniToTime(timesTamp int) string {
	fmt.Println(timesTamp)
	t := time.Unix(int64(timesTamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Println(str1, str2 string) string {
	return str1 + str2
}
