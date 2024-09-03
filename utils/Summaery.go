package utils

import (
	"baby-varkin-health/initializers"
	"baby-varkin-health/models"
	"fmt"
	"math"
	"sort"
	"time"
)

type resultSumFloat struct {
	Sum float32
}
type resultSum struct {
	Sum int
}
type resultCount struct {
	Count int
}
type SleepSummary struct {
	SleepDay     time.Time
	TotalMinutes float32
}

func HistoryAll(today string) string {
	var historyAll []models.Activities
	initializers.DB.Model(&models.Activities{}).
		Order("created_at desc").
		Limit(100).
		Find(&historyAll)
	stringSummary := fmt.Sprintf("----------------------\nวันที่ %s\n-------HISTORY-------\n", today)

	// Sort by ID if needed
	sort.Slice(historyAll, func(i, j int) bool {
		return historyAll[i].ID < historyAll[j].ID
	})
	for _, activity := range historyAll {
		icon := GetIcon(activity.ActityType, activity.ActityValue)
		stringSummary += fmt.Sprintf("%d: %s  %s\n", activity.ID, ConvertDateTime(activity.CreatedAt, "2006-01-02 15:04"), icon)
	}

	stringSummary += fmt.Sprintf("-------HISTORY-------\n-----Delete Recorod------\n-------del {ID}-------\n")
	return stringSummary
}

func DrinkMilk(startOfDay time.Time, endOfDay time.Time, today string) string {
	var sumDrinkMilk resultSumFloat
	initializers.DB.Model(&models.Activities{}).Select("SUM(CAST(actity_value AS Float)) as Sum").Where("created_at >= ? AND created_at < ? and actity_type = ?", startOfDay, endOfDay, "drinkmilk").Find(&sumDrinkMilk)
	var lastThreeActivities []models.Activities
	initializers.DB.Model(&models.Activities{}).
		Where("created_at >= ? AND created_at < ? AND actity_type = ?", startOfDay, endOfDay, "drinkmilk").
		Order("created_at desc").
		Limit(10).
		Find(&lastThreeActivities)

	var drinkMilkSummary string

	// Loop through the activities to build the summary
	for _, activity := range lastThreeActivities {
		drinkMilkSummary += fmt.Sprintf("%s 🍼 %s ออนซ์\n", ConvertDateTime(activity.CreatedAt, "2006-01-02 15:04"), activity.ActityValue)
	}

	// Add the total sum of drink milk
	drinkMilkSummary += fmt.Sprintf("----------------------\n")
	drinkMilkSummary += fmt.Sprintf("🍼 รวม %.2f ออนซ์", sumDrinkMilk.Sum)

	Textsummarysum := fmt.Sprintf(`----------------------
วันที่ %s
----------🧑‍🍼----------
ดื่มนม 10 ครั้งล่าสุด 🧑‍🍼
%s
----------------------`, today, drinkMilkSummary)
	return Textsummarysum
}

func Diaper(startOfMonth time.Time, endOfMonth time.Time, startOfDay time.Time, endOfDay time.Time, today string) string {
	var countMonth resultCount
	initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").Where("created_at >= ? AND created_at < ? and actity_type = ? ", startOfMonth, endOfMonth, "pampers").Find(&countMonth)

	var lastPee models.Activities
	initializers.DB.Where("actity_type = ? and actity_value = ?", "pampers", "pee").Order("created_at desc").First(&lastPee)

	var lastPoo models.Activities
	initializers.DB.Where("actity_type = ? and actity_value = ?", "pampers", "poopoo").Order("created_at desc").First(&lastPoo)
	var countPee resultCount
	initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").Where("created_at >= ? AND created_at < ? and actity_type = ? and actity_value = ?", startOfDay, endOfDay, "pampers", "pee").Find(&countPee)

	var countPooPoo resultCount
	initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").Where("created_at >= ? AND created_at < ? and actity_type = ? and actity_value = ?", startOfDay, endOfDay, "pampers", "poopoo").Find(&countPooPoo)
	peeTime := ConvertDateTime(lastPee.CreatedAt, "2006-01-02 15:04")
	pooTime := ConvertDateTime(lastPoo.CreatedAt, "2006-01-02 15:04")
	var Textsummarysum = fmt.Sprintf(`------------------------
วันที่ %s
------💩----💦--------
1 เดือนใช้ไป %d ชิ้น
----------💦----------
ซิ่งฉ่อง ครั้งล่าสุด 💦
%s
ซิ่งฉ่องรวม %d ครั้ง
----------💩----------
อุนจิ ครั้งล่าสุด 💩
%s
อุนจิรวม %d ครั้ง
------------------------`, today, countMonth.Count, peeTime, countPee.Count, pooTime, countPooPoo.Count)

	return Textsummarysum
}

func HistoryDiaper(startOfMonth time.Time, endOfMonth time.Time, startOfDay time.Time, endOfDay time.Time, today string) string {
	var lastThreeActivities []models.Activities
	initializers.DB.Model(&models.Activities{}).
		Where("created_at >= ? AND created_at < ? AND actity_type = ?", startOfMonth, endOfMonth, "pampers").
		Order("created_at desc").
		Find(&lastThreeActivities)

	var stringSummary string

	// Loop through the activities to build the summary
	for _, activity := range lastThreeActivities {
		icon := ""
		if activity.ActityValue == "pee" {
			icon = "💦"
		} else {
			icon = "💩"
		}

		stringSummary += fmt.Sprintf("%s  %s\n", ConvertDateTime(activity.CreatedAt, "2006-01-02 15:04"), icon)
	}

	var Textsummarysum = fmt.Sprintf(`----------------------	
วันที่ %s
------💩----💦--------
%s----------------------`, today, stringSummary)

	return Textsummarysum
}

func SleepDaily(startOfMonth time.Time, endOfMonth time.Time, startOfDay time.Time, endOfDay time.Time, today string) string {
	var sleepActivity, wakeUpActivity models.Activities
	strSleep := ""
	err := initializers.DB.Where("actity_type = ?", "sleep").Order("created_at desc").First(&sleepActivity).Error
	err1 := initializers.DB.Where("actity_type = ?", "wakeup").Order("created_at desc").First(&wakeUpActivity).Error

	if err != nil || err1 != nil {
		strSleep = fmt.Sprintf(`ไม่มีข้อมูลการนอน~`)
	} else if wakeUpActivity.CreatedAt.Sub(sleepActivity.CreatedAt).Minutes() < 0 {
		strSleep = fmt.Sprintf(`ไม่มีข้อมูลการนอน~`)
	} else {
		duration := wakeUpActivity.CreatedAt.Sub(sleepActivity.CreatedAt)
		hours := int(duration.Hours())
		min := int(duration.Minutes()) % 60
		startTime := ConvertDateTime(sleepActivity.CreatedAt, "15:04")
		endTime := ConvertDateTime(wakeUpActivity.CreatedAt, "15:04")
		strSleep = fmt.Sprintf(`ล่าสุด นอนตั้งแต่ %s - %s
💤 %d ชม. %d นาที`, startTime, endTime, hours, min)

	}

	var sleepSummaries []SleepSummary
	initializers.DB.Raw(`
				SELECT 
    DATE(sleep.created_at) AS sleep_day,
    SUM(EXTRACT(EPOCH FROM (wakeup.created_at - sleep.created_at)) / 60) AS total_minutes
FROM 
    activities sleep
JOIN LATERAL (
    SELECT *
    FROM activities wakeup
    WHERE wakeup.actity_type = 'wakeup'
      AND wakeup.created_at > sleep.created_at
    ORDER BY wakeup.created_at ASC
    LIMIT 1
) wakeup ON true
WHERE 
    sleep.actity_type = 'sleep'
GROUP BY 
    sleep_day
ORDER BY 
    sleep_day DESC
LIMIT 7;
			`).Scan(&sleepSummaries)

	// Build summary string
	var stringSummary string
	for _, summary := range sleepSummaries {
		hours := int(summary.TotalMinutes / 60)
		minutes := int(summary.TotalMinutes) % 60
		stringSummary += fmt.Sprintf("%s :=> %d ชม. %d นาที\n", ConvertDateTime(summary.SleepDay, "2006-01-02"), hours, minutes)
	}

	var Textsummarysum = fmt.Sprintf(`-----------------------	
วันที่ %s
----------💤----------
%s
----------💤----------
7 วันย้อนหลัง นอนไปกี่ชม
----------------------
%s----------------------`, today, strSleep, stringSummary)

	return Textsummarysum
}

func Takeabath(startOfMonth time.Time, endOfMonth time.Time, startOfDay time.Time, endOfDay time.Time, today string) string {
	var takeabathCount resultCount
	initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").Where("created_at >= ? AND created_at < ? and actity_type = ? and actity_value = ?", startOfDay, endOfDay, "takeabath", "1").Find(&takeabathCount)
	var Textsummarysum = fmt.Sprintf(`-----------------------	
วันที่ %s
----------🛁----------
วันนี้อาบน้ำ🛁 %d ครั้ง
----------🛁----------`, today, takeabathCount.Count)

	return Textsummarysum
}

func PumpmilkDaily(startOfMonth time.Time, endOfMonth time.Time, startOfDay time.Time, endOfDay time.Time, today string) string {
	var sumPumpMilk resultSumFloat
	initializers.DB.Model(&models.Activities{}).
		Select("SUM(CAST(actity_value AS FLOAT)) as Sum").
		Where("created_at >= ? AND created_at < ? and actity_type = ?", startOfDay, endOfDay, "pumpmilk").
		Find(&sumPumpMilk)
	var lastThreeActivities []models.Activities
	initializers.DB.Model(&models.Activities{}).
		Where("created_at >= ? AND created_at < ? AND actity_type = ?", startOfDay, endOfDay, "pumpmilk").
		Order("created_at desc").
		Limit(30).
		Find(&lastThreeActivities)
	// Sort by ID if needed
	sort.Slice(lastThreeActivities, func(i, j int) bool {
		return lastThreeActivities[i].ID < lastThreeActivities[j].ID
	})
	stringSummary := ""
	for _, activity := range lastThreeActivities {
		icon := GetIcon(activity.ActityType, activity.ActityValue)
		stringSummary += fmt.Sprintf("%s %s\n", ConvertDateTime(activity.CreatedAt, "2006-01-02 15:04"), icon)
	}

	var sumStockMilk resultSumFloat
	initializers.DB.Model(&models.Activities{}).Select("SUM(CAST(actity_value AS FLOAT)) as Sum").Where("actity_type = ?", "stockmilk").Find(&sumStockMilk)
	var countStockMilkUsed resultCount
	initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").Where("actity_type = ? and actity_value = ?", "stockmilk", "0").Find(&countStockMilkUsed)

	oldLot := math.Ceil(float64(countStockMilkUsed.Count) / 10)
	if oldLot < 1 {
		oldLot = 1
	}
	oldNo := countStockMilkUsed.Count % 10
	if oldNo < 1 {
		oldNo = 1
	}

	var limit1StockMilk models.Activities
	initializers.DB.Model(&models.Activities{}).
		Where("actity_type = ? and actity_value <> ?", "stockmilk", "0").
		Order("created_at asc").
		First(&limit1StockMilk)

	var Textsummarysum = fmt.Sprintf(`-----------------------	
วันที่ %s
-------⛽ + 🍼----------
%s
----------------------
รวม 🍼 %.0f ครั้ง
----------------------
----------------------
----------------------
Stock Milk 🥶  🍼  📦
----------------------
รวมทั้งหมด %.0f ออน
----------------------
เก่าสุด %.0f/%d (Lot %.0f No. %d)
วันที่เก็บ %s 
🍼 %s
----------------------`, today, stringSummary, sumPumpMilk.Sum, sumStockMilk.Sum, oldLot, oldNo, oldLot, oldNo,
		ConvertDateTime(limit1StockMilk.CreatedAt, "2006-01-02 15:04"),
		limit1StockMilk.ActityValue)

	return Textsummarysum
}

func HistoryPumpmilk(startOfMonth time.Time, endOfMonth time.Time, startOfDay time.Time, endOfDay time.Time, today string) string {
	var lastThreeActivities []models.Activities
	initializers.DB.Model(&models.Activities{}).
		Where("actity_type = ?", "pumpmilk").
		Order("created_at desc").
		Limit(30).
		Find(&lastThreeActivities)
	// Sort by ID if needed
	sort.Slice(lastThreeActivities, func(i, j int) bool {
		return lastThreeActivities[i].ID < lastThreeActivities[j].ID
	})
	stringSummary := ""
	for _, activity := range lastThreeActivities {
		icon := GetIcon(activity.ActityType, activity.ActityValue)
		stringSummary += fmt.Sprintf("%s %s\n", ConvertDateTime(activity.CreatedAt, "2006-01-02 15:04"), icon)
	}

	var Textsummarysum = fmt.Sprintf(`-----------------------	
HISTORY
-------⛽ + 🍼----------
%s
-------⛽ + 🍼----------
%s
-----------------------`, today, stringSummary)

	return Textsummarysum
}

func Stockmilk(startOfMonth time.Time, endOfMonth time.Time, startOfDay time.Time, endOfDay time.Time, today string) string {
	var lastThreeActivities []models.Activities
	initializers.DB.Model(&models.Activities{}).
		Where("actity_type = ?", "stockmilk").
		Order("created_at desc").
		Limit(30).
		Find(&lastThreeActivities)
	// Sort by ID if needed
	// sort.Slice(lastThreeActivities, func(i, j int) bool {
	// 	return lastThreeActivities[i].ID < lastThreeActivities[j].ID
	// })
	stringSummary := ""
	for _, activity := range lastThreeActivities {
		icon := GetIcon(activity.ActityType, activity.ActityValue)
		stringSummary += fmt.Sprintf("%s %s\n", ConvertDateTime(activity.CreatedAt, "2006-01-02 15:04"), icon)
	}

	var Textsummarysum = fmt.Sprintf(`-----------------------	
HISTORY
-------🥶  🍼  📦----------
%s
-------🥶  🍼  📦----------
%s
-----------------------`, today, stringSummary)

	return Textsummarysum
}
