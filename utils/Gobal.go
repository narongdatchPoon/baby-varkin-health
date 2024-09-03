package utils

import (
	"fmt"
	"time"
)

func ConvertDateTime(valTime time.Time, format string) string {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return ""
	}
	return valTime.In(loc).Format(format)
}
