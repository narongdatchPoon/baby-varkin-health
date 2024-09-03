package handler

import (
	"baby-varkin-health/initializers"
	"baby-varkin-health/models"
	"baby-varkin-health/utils"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"

	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func init() {
	initializers.ConnectDB()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	tokenChannel := os.Getenv("LINE_CHANNEL_TOKEN")
	bot, err := messaging_api.NewMessagingApiAPI(
		tokenChannel,
	)
	if err != nil {
		log.Fatal(err)
	}

	cb, err := webhook.ParseRequest(channelSecret, r)
	if err != nil {
		log.Printf("Cannot parse request: %+v\n", err)
		if errors.Is(err, webhook.ErrInvalidSignature) {
			log.Println("/WriteHeader 400")
			w.WriteHeader(400)
		} else {
			log.Println("/WriteHeader 500")
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range cb.Events {
		switch e := event.(type) {
		case webhook.MessageEvent:
			switch message := e.Message.(type) {
			case webhook.TextMessageContent:
				_, err := ReplyMessage(bot, e.ReplyToken, message.Text)
				if err != nil {
					log.Print("err")
				} else {
					log.Println("Sent text reply.")
				}

			default:
				log.Printf("Unsupported message content: %T\n", e.Message)
			}
		default:
			log.Printf("Unsupported message: %T\n", event)
		}
	}
	w.WriteHeader(http.StatusOK)
}

func ReplyMessage(bot *messaging_api.MessagingApiAPI, replyToken string, userMsg string) (*messaging_api.ReplyMessageResponse, error) {
	if userMsg == "Menu" {
		json := utils.ReturnFlexMessage()
		//Unmarshal JSON
		flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(json))
		if err != nil {
			log.Println(err)
		}
		// New Flex Message
		flexMessage := linebot.NewFlexMessage("FlexWithJSON", flexContainer)

		// Wrap the Flex Message
		flexMessageWrapper := &FlexMessageWrapper{FlexMessage: flexMessage}

		// Reply Message
		replyMessageResponse, err := bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					flexMessageWrapper,
				},
			})

		//  flexMessage).Do()
		if err != nil {
			log.Println(err)
		}
		return replyMessageResponse, err

	} else if strings.Contains(userMsg, "summarysum") {
		parts := strings.Split(userMsg, " ")
		loc, _ := time.LoadLocation("Asia/Bangkok")
		now := time.Now().In(loc)
		startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
		// startOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		// startOfDay := time.Now().UTC().Truncate(24 * time.Hour)
		endOfDay := startOfDay.Add(24 * time.Hour)

		startOfMonth := time.Now().UTC().Truncate(24 * 30 * time.Hour)
		endOfMonth := startOfDay.Add(24 * 30 * time.Hour)
		today := utils.ConvertDateTime(startOfDay, "2006-01-02")
		if parts[1] == "history-all" {
			stringSummary := utils.HistoryAll(today)
			return Reply(bot, replyToken, stringSummary)
		} else if parts[1] == "drinkmilk" {
			Textsummarysum := utils.DrinkMilk(startOfDay, endOfDay, today)
			return Reply(bot, replyToken, Textsummarysum)
		} else if parts[1] == "diaper" {
			Textsummarysum := utils.Diaper(startOfMonth, endOfMonth, startOfDay, endOfDay, today)
			return Reply(bot, replyToken, Textsummarysum)
		} else if parts[1] == "history-diaper" {
			Textsummarysum := utils.HistoryDiaper(startOfMonth, endOfMonth, startOfDay, endOfDay, today)
			return Reply(bot, replyToken, Textsummarysum)
		} else if parts[1] == "sleep-daily" {
			Textsummarysum := utils.SleepDaily(startOfMonth, endOfMonth, startOfDay, endOfDay, today)
			return Reply(bot, replyToken, Textsummarysum)
		} else if strings.Contains(userMsg, "takeabath") {
			Textsummarysum := utils.Takeabath(startOfMonth, endOfMonth, startOfDay, endOfDay, today)
			return Reply(bot, replyToken, Textsummarysum)
		} else if strings.Contains(userMsg, "pumpmilk-daily") {
			Textsummarysum := utils.PumpmilkDaily(startOfMonth, endOfMonth, startOfDay, endOfDay, today)
			return Reply(bot, replyToken, Textsummarysum)
		} else if strings.Contains(userMsg, "history-pumpmilk") {
			Textsummarysum := utils.HistoryPumpmilk(startOfMonth, endOfMonth, startOfDay, endOfDay, today)
			return Reply(bot, replyToken, Textsummarysum)
		} else if strings.Contains(userMsg, "stock") {
			Textsummarysum := utils.Stockmilk(startOfMonth, endOfMonth, startOfDay, endOfDay, today)
			return Reply(bot, replyToken, Textsummarysum)
		} else {
			return Reply(bot, replyToken, "'"+userMsg+"'  น้องวา- ไม่เข้าใจ งับบบบ")
		}

	} else if strings.Contains(userMsg, "drinkmilk") ||
		strings.Contains(userMsg, "pampers") ||
		strings.Contains(userMsg, "sleep") ||
		strings.Contains(userMsg, "wakeup") ||
		strings.Contains(userMsg, "takeabath") ||
		strings.Contains(userMsg, "pumpmilk") ||
		strings.Contains(userMsg, "milkmilk") ||
		strings.Contains(userMsg, "stockmilk") {
		parts := strings.Split(userMsg, " ")
		activity := models.Activities{
			ActityType:  parts[0],
			ActityValue: parts[1],
			ReplyToken:  replyToken,
		}
		initializers.DB.Create(&activity)
		// activity.Save()
		if strings.Contains(userMsg, "stockmilk") {
			var stockmilkCount resultCount
			initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").Where("actity_type = ?", "stockmilk").Find(&stockmilkCount)
			lot := math.Ceil(float64(stockmilkCount.Count) / 10)
			no := stockmilkCount.Count % 10
			stringReturn := fmt.Sprintf("น้องวารับทราบ\n %.0f/%d LOT:%.0f No. %d", lot, no, lot, no)
			// 1 No. 5

			return Reply(bot, replyToken, stringReturn)
		} else {
			return Reply(bot, replyToken, "น้องวารับทราบ")
		}

	} else if userMsg == "Summary" {
		json := utils.FlexMessageSummary()
		//Unmarshal JSON
		flexContainer, err := linebot.UnmarshalFlexMessageJSON([]byte(json))
		if err != nil {
			log.Println(err)
		}
		// New Flex Message
		flexMessage := linebot.NewFlexMessage("FlexWithJSON", flexContainer)

		// Wrap the Flex Message
		flexMessageWrapper := &FlexMessageWrapper{FlexMessage: flexMessage}

		// Reply Message
		replyMessageResponse, err := bot.ReplyMessage(
			&messaging_api.ReplyMessageRequest{
				ReplyToken: replyToken,
				Messages: []messaging_api.MessageInterface{
					flexMessageWrapper,
				},
			})

		//  flexMessage).Do()
		if err != nil {
			log.Println(err)
		}
		return replyMessageResponse, err
	} else {
		parts := strings.Split(userMsg, " ")
		act_id, _ := strconv.Atoi(parts[1])
		if parts[0] == "del" {

			initializers.DB.Delete(&models.Activities{}, act_id)

			return Reply(bot, replyToken, "'"+userMsg+"'  น้องวา- ลบข้อมูลให้แล้วงับบบ")
		} else {

			return Reply(bot, replyToken, "'"+userMsg+"'  น้องวา- ไม่เข้าใจ งับบบบ---")
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

type resultSumFloat struct {
	Sum float32
}
type resultSum struct {
	Sum int
}
type resultCount struct {
	Count int
}

type FlexMessageWrapper struct {
	*linebot.FlexMessage
}

func (f *FlexMessageWrapper) GetType() string {
	return "flex"
}

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler).Methods("POST")

}
