package utils

import (
	"baby-varkin-health/initializers"
	"baby-varkin-health/models"
	"fmt"
	"math"
	"strings"
)

func Recorod(activity models.Activities, userMsg string, replyToken string) string {

	var checkToken resultCount
	initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").
		Where("reply_token = ? ", replyToken).
		Find(&checkToken)
	if checkToken.Count > 1 {
		return "error"
	} else {
		initializers.DB.Create(&activity)
	}

	// activity.Save()
	if strings.Contains(userMsg, "stockmilk") {
		var stockmilkCount resultCount
		initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").Where("actity_type = ?", "stockmilk").Find(&stockmilkCount)
		lot := math.Ceil(float64(stockmilkCount.Count) / 10)
		no := stockmilkCount.Count % 10
		stringReturn := fmt.Sprintf("น้องวารับทราบ\n %.0f/%d LOT:%.0f No. %d", lot, no, lot, no)
		// 1 No. 5

		return stringReturn
	} else {
		return "น้องวารับทราบ " + GetIcon(userMsg, "")
	}

}
