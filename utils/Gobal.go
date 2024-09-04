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

func GetIcon(activity_type string, activity_value string) string {
	var icon string
	if activity_type == "pampers" {
		if activity_value == "pee" {
			icon = "💦"
		} else {
			icon = "💩"
		}
	} else if activity_type == "drinkmilk" {
		icon = "🍼 " + activity_value
	} else if activity_type == "sleep" {
		icon = "💤 sleep"
	} else if activity_type == "wakeup" {
		icon = "💤 wake"
	} else if activity_type == "takeabath" {
		icon = "🛁"
	} else if activity_type == "pumpmilk" {
		icon = "⛽ " + activity_value
	} else if activity_type == "milkmilk" {
		icon = "🤱🏻 " + activity_value
	} else if activity_type == "stockmilk" {
		icon = "📦 " + activity_value
	} else {
		icon = activity_type
	}

	return icon
}
