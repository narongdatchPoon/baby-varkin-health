package utils

import (
	"baby-varkin-health/initializers"
	"baby-varkin-health/models"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func Record(activity models.Activities, userMsg string, replyToken string, bot *messaging_api.MessagingApiAPI) (*messaging_api.ReplyMessageResponse, error) {

	// Attempt to create the activity
	err := initializers.DB.Create(&activity).Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			stringReturn := "error: duplicate reply token"
			return Reply(bot, replyToken, stringReturn)
		}
		stringReturn := "error: unable to record activity"
		return Reply(bot, replyToken, stringReturn)
	} else {

		// If the message contains "stockmilk"
		if strings.Contains(userMsg, "stockmilk") {
			var stockmilkCount resultCount
			initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").Where("actity_type = ?", "stockmilk").Find(&stockmilkCount)
			lot := math.Ceil(float64(stockmilkCount.Count) / 10)
			no := stockmilkCount.Count % 10
			stringReturn := fmt.Sprintf("น้องวารับทราบ\n %.0f/%d LOT:%.0f No. %d", lot, no, lot, no)
			return Reply(bot, replyToken, stringReturn)

		} else {
			parts := strings.Split(userMsg, " ")
			stringTitle := "น้องวารับทราบ " + GetIcon(parts[0], parts[1])
			data := "data|" + activity.ReplyToken
			deleteID := "del " + strconv.FormatUint(uint64(activity.ID), 10)
			json := DateTimePicker(stringTitle, data, deleteID)

			return ReplyFlexMessage(json, bot, replyToken)
		}
	}

}

func HandlePostbackEvent(event webhook.PostbackEvent, bot *messaging_api.MessagingApiAPI) {
	postbackData := event.Postback.Data
	replyToken := event.ReplyToken

	// Check if the postback data is from the DateTime Picker

	if strings.Contains(postbackData, "data|") {
		parts := strings.Split(postbackData, "|")
		getReplyToken := parts[1]
		var getActivity models.Activities
		initializers.DB.Where("reply_token = ?", getReplyToken).First(&getActivity)
		dateFormat := "2006-01-02T15:04:05"
		selectedDatetimeStr := event.Postback.Params["datetime"]
		selectedDatetime, err := time.Parse(dateFormat, selectedDatetimeStr)
		if err != nil {
			log.Printf("Error parsing datetime: %v", err)
			selectedDatetime = time.Time{} // or use a default value
		}
		getActivity.CreatedAt = selectedDatetime
		initializers.DB.Updates(&getActivity)

		replyMessage := fmt.Sprintf("แก้ไขวันเวลาเสร็จแล้ว คร๊าบบบ: \n%s\n%s", GetIcon(getActivity.ActityType, getActivity.ActityValue), selectedDatetimeStr)

		_, checkError := Reply(bot, replyToken, replyMessage)
		if checkError != nil {
			log.Println("Error replying to postback:", err)
		}
	} else {
		replyMessage := "UPDATE ERROR"
		_, err := Reply(bot, replyToken, replyMessage)
		if err != nil {
			log.Println("Error replying to postback:", err)
		}
	}
}

func Reply(bot *messaging_api.MessagingApiAPI, replyToken string, message string) (*messaging_api.ReplyMessageResponse, error) {
	replyMessageResponse, err := bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: message,
				},
			},
		})
	return replyMessageResponse, err
}
