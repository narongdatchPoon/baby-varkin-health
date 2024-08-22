package handler

import (
	"baby-varkin-health/initializers"
	"baby-varkin-health/models"
	"baby-varkin-health/utils"
	"errors"
	"fmt"
	"log"
	"math"
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

	} else if strings.Contains(userMsg, "drinkmilk") || strings.Contains(userMsg, "pampers") || strings.Contains(userMsg, "sleep") || strings.Contains(userMsg, "wakeup") || strings.Contains(userMsg, "takeabath") || strings.Contains(userMsg, "pumpmilk") || strings.Contains(userMsg, "milkmilk") {
		parts := strings.Split(userMsg, " ")
		activity := models.Activities{
			ActityType:  parts[0],
			ActityValue: parts[1],
		}
		activity.Save()
		return Reply(bot, replyToken, "น้องวารับทราบ")
	} else if userMsg == "Summary" {
		startOfDay := time.Now().UTC().Truncate(24 * time.Hour)
		endOfDay := startOfDay.Add(24 * time.Hour)

		var sumDrinkMilk resultSum
		initializers.DB.Model(&models.Activities{}).Select("SUM(CAST(actity_value AS INTEGER)) as Sum").Where("created_at >= ? AND created_at < ? and actity_type = ?", startOfDay, endOfDay, "drinkmilk").Find(&sumDrinkMilk)

		var countPee resultCount
		initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").Where("created_at >= ? AND created_at < ? and actity_type = ? and actity_value = ?", startOfDay, endOfDay, "pampers", "pee").Find(&countPee)

		var countPooPoo resultCount
		initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").Where("created_at >= ? AND created_at < ? and actity_type = ? and actity_value = ?", startOfDay, endOfDay, "pampers", "poopoo").Find(&countPooPoo)

		var lastDrinkMilk models.Activities
		initializers.DB.Where("actity_type = ?", "drinkmilk").Order("created_at desc").First(&lastDrinkMilk)

		var lastPee models.Activities
		initializers.DB.Where("actity_type = ? and actity_value = ?", "pampers", "pee").Order("created_at desc").First(&lastPee)

		var lastPoo models.Activities
		initializers.DB.Where("actity_type = ? and actity_value = ?", "pampers", "poopoo").Order("created_at desc").First(&lastPoo)

		var sleepActivity, wakeUpActivity models.Activities
		strSleep := ""

		err := initializers.DB.Where("actity_type = ?", "sleep").Order("created_at desc").First(&sleepActivity).Error
		err1 := initializers.DB.Where("actity_type = ?", "wakeup").Order("created_at desc").First(&wakeUpActivity).Error

		if err != nil || err1 != nil {
			strSleep = fmt.Sprintf(`ไม่มีข้อมูลการนอน`)
		} else {
			duration := wakeUpActivity.CreatedAt.Sub(sleepActivity.CreatedAt)
			hours := int(duration.Hours())
			min := int(duration.Minutes()) % 60
			startTime := convertDateTime(sleepActivity.CreatedAt, "15:04")
			endTime := convertDateTime(wakeUpActivity.CreatedAt, "15:04")
			strSleep = fmt.Sprintf(`นอนตั้งแต่ %s - %s 
💤 %d ชม. %d นาที`, startTime, endTime, hours, min)
		}

		// err2 := initializers.DB.Where("actity_type = ?", "takeabath").Order("created_at desc").First(&wakeUpActivity).Error
		var takeabathCount resultCount
		initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").Where("created_at >= ? AND created_at < ? and actity_type = ? and actity_value = ?", startOfDay, endOfDay, "takeabath", "1").Find(&takeabathCount)

		var sumPumpMilk resultSum
		initializers.DB.Model(&models.Activities{}).Select("SUM(CAST(actity_value AS INTEGER)) as Sum").Where("created_at >= ? AND created_at < ? and actity_type = ?", startOfDay, endOfDay, "pumpmilk").Find(&sumPumpMilk)

		var countPumpMilk resultCount
		initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").Where("created_at >= ? AND created_at < ? and actity_type = ?", startOfDay, endOfDay, "pumpmilk").Find(&countPumpMilk)

		var countMilkMilk resultCount
		initializers.DB.Model(&models.Activities{}).Select("Count(actity_value) as Count").Where("created_at >= ? AND created_at < ? and actity_type = ?", startOfDay, endOfDay, "milkmilk").Find(&countMilkMilk)

		today := convertDateTime(startOfDay, "2006-01-02")
		milkTime := convertDateTime(lastDrinkMilk.CreatedAt, "2006-01-02 15:04")
		peeTime := convertDateTime(lastPee.CreatedAt, "2006-01-02 15:04")
		pooTime := convertDateTime(lastPoo.CreatedAt, "2006-01-02 15:04")
		var avgPumpMilk string
		_avgPumpMilk := float64(sumPumpMilk.Sum) / float64(countPumpMilk.Count)
		if math.IsNaN(_avgPumpMilk) {
			avgPumpMilk = "0"
		} else {
			fmt.Sprintf("%.2f", _avgPumpMilk)
		}

		var TextSummary = fmt.Sprintf(`----------------------
สรุป ของ น้องวา !!! 
วันที่ %s
----------🧑‍🍼----------
ดื่มนม ครั้งล่าสุด 🧑‍🍼
เวลา %s ไปแล้ว %s ออนซ์
🍼รวม %d ออนซ์
----------🚽----------
ซิ่งฉ่อง ครั้งล่าสุด 🚽
%s 
ซิ่งฉ่องรวม %d ครั้ง
----------💩----------
อุนจิ ครั้งล่าสุด 💩
%s
อุนจิรวม %d ครั้ง
----------💤----------
%s
----------🛁----------
วันนี้อาบน้ำ🛁 %d ครั้ง
----------🍼----------
รวม น้ำนมที่ปั้ม วันนี้
🍼 %d ออนซ์
🍼 %d ครั้ง
เฉลี่ยน %s ออนซ์
----------🤰----------
เข้าเต้า %d ครั้ง
----------------------
----------------------`, today, milkTime, lastDrinkMilk.ActityValue, sumDrinkMilk.Sum, peeTime, countPee.Count, pooTime, countPooPoo.Count, strSleep, takeabathCount.Count, sumPumpMilk.Sum, countPumpMilk.Count, avgPumpMilk, countMilkMilk.Count)
		return Reply(bot, replyToken, TextSummary)
	} else {
		return Reply(bot, replyToken, "'"+userMsg+"'  น้องวา- ไม่เข้าใจ งับบบบ")
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
func convertDateTime(valTime time.Time, format string) string {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return ""
	}
	return valTime.In(loc).Format(format)
}

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler).Methods("POST")

}
