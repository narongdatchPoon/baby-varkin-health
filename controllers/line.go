package controllers

import (
	"baby-varkin-health/initializers"
	"baby-varkin-health/models"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "test"})
}

func WebHook(c *gin.Context) {

	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	bot, err := messaging_api.NewMessagingApiAPI(
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("/callback called...")

	cb, err := webhook.ParseRequest(channelSecret, c.Request)
	if err != nil {
		log.Printf("Cannot parse request: %+v\n", err)
		if errors.Is(err, webhook.ErrInvalidSignature) {
			log.Println("/WriteHeader 400")
		} else {
			log.Println("/WriteHeader 500")
		}
		return
	}

	log.Println("Handling events...")
	for _, event := range cb.Events {
		log.Printf("/callback called%+v...\n", event)

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

	c.JSON(http.StatusOK, gin.H{"data": "test"})
}

type result struct {
	Sum int
}

func returnFlexMessage() string {
	return `{
  "type": "carousel",
  "contents": [
    {
      "type": "bubble",
      "size": "micro",
      "hero": {
        "type": "image",
        "url": "https://vos.line-scdn.net/card-type-message-image-2024/307axhqi/1724144356175-kAHKfFqGbwPMISYjeSTEPYTaX7fRtgYUpMoM7QbaUCFLIQ3pLK",
        "size": "full",
        "aspectMode": "cover",
        "aspectRatio": "320:213"
      },
      "body": {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "weight": "bold",
            "size": "sm",
            "wrap": true,
            "text": "กินนม กี่ ออนซ์"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "1 ออนซ์",
              "text": "drinkmilk 1"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "2 ออนซ์",
              "text": "drinkmilk 2"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "3 ออนซ์",
              "text": "drinkmilk 3"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "4 ออนซ์",
              "text": "drinkmilk 4"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "5 ออนซ์",
              "text": "drinkmilk 5"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "6 ออนซ์",
              "text": "drinkmilk 6"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          }
        ],
        "spacing": "sm",
        "paddingAll": "13px"
      }
    },
    {
      "type": "bubble",
      "size": "micro",
      "hero": {
        "type": "image",
        "url": "https://vos.line-scdn.net/card-type-message-image-2024/307axhqi/1724128350647-B2bJrPglMPiS2oeNiYfRGcoDAosKQ6djJ3rXBddP2lbg1wZCJl",
        "size": "full",
        "aspectMode": "cover",
        "aspectRatio": "320:213",
        "margin": "none",
        "position": "relative"
      },
      "body": {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "weight": "bold",
            "size": "sm",
            "wrap": true,
            "text": "แพมเพิส"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "ชิ้งฉ่อง",
              "text": "pampers pee"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "อึคึ",
              "text": "pampers poopoo"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          }
        ],
        "spacing": "sm",
        "paddingAll": "13px"
      }
    },
    {
      "type": "bubble",
      "size": "micro",
      "hero": {
        "type": "image",
        "url": "https://vos.line-scdn.net/card-type-message-image-2024/307axhqi/1724128979240-QK4QnCIToKmxBiQFFEVFWNYZai21uKmLBsLlNvaOnHWhcojGPY",
        "size": "full",
        "aspectMode": "cover",
        "aspectRatio": "320:213",
        "margin": "none",
        "position": "relative"
      },
      "body": {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "weight": "bold",
            "size": "sm",
            "wrap": true,
            "text": "นอนเถอะ น้าาาา"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "นอนแย้ว ~ เย้",
              "text": "sleep 1"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          }
        ],
        "spacing": "sm",
        "paddingAll": "13px"
      }
    },
    {
      "type": "bubble",
      "size": "micro",
      "hero": {
        "type": "image",
        "url": "https://vos.line-scdn.net/card-type-message-image-2024/307axhqi/1724129109711-f5b1FCb3sCoC66tepiitWuejyYfybKRYFCTmrFsB7C9MwiH4vA",
        "size": "full",
        "aspectMode": "cover",
        "aspectRatio": "320:213",
        "margin": "none",
        "position": "relative"
      },
      "body": {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "weight": "bold",
            "size": "sm",
            "wrap": true,
            "text": "ตื่นๆ มีเรื่องแล้ว"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "ตื่นแย้ว ~ เย้",
              "text": "wakeup 1"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          }
        ],
        "spacing": "sm",
        "paddingAll": "13px"
      }
    },
    {
      "type": "bubble",
      "size": "micro",
      "hero": {
        "type": "image",
        "url": "https://vos.line-scdn.net/card-type-message-image-2024/307axhqi/1724128836193-NSpyCH7AwNYYqKHXbsP9CYP5zNUi5OGEqIXm7fIAH1PFqLuDiQ",
        "size": "full",
        "aspectMode": "cover",
        "aspectRatio": "320:213",
        "margin": "none",
        "position": "relative"
      },
      "body": {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "weight": "bold",
            "size": "sm",
            "wrap": true,
            "text": "อาบน้ำกันเถอะ !!!"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "อาบน้ำ ~ เย้",
              "text": "takeabath 1"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          }
        ],
        "spacing": "sm",
        "paddingAll": "13px"
      }
    },
    {
      "type": "bubble",
      "size": "micro",
      "hero": {
        "type": "image",
        "url": "https://vos.line-scdn.net/card-type-message-image-2024/307axhqi/1724128499848-OHcCpDE9mOGXc76hJsjoeiwQzQLAFwGoYWjFa3OAPrlky2BlxC",
        "size": "full",
        "aspectMode": "cover",
        "aspectRatio": "320:213"
      },
      "body": {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "weight": "bold",
            "size": "sm",
            "wrap": true,
            "text": "ปั้มนมได้ กี่ ออนซ์"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "1 ออนซ์",
              "text": "pumpmilk 1"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "2 ออนซ์",
              "text": "pumpmilk 2"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "3 ออนซ์",
              "text": "pumpmilk 3"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "4 ออนซ์",
              "text": "pumpmilk 4"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "5 ออนซ์",
              "text": "pumpmilk 5"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "6 ออนซ์",
              "text": "pumpmilk 6"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          }
        ],
        "spacing": "sm",
        "paddingAll": "13px"
      }
    },
    {
      "type": "bubble",
      "size": "micro",
      "hero": {
        "type": "image",
        "url": "https://vos.line-scdn.net/card-type-message-image-2024/307axhqi/1724128718742-GxPw9QTJkORIQMxILTbRhMMwNU9KzueqZEJEVu7qqDRkd6Zr6h",
        "size": "full",
        "aspectMode": "cover",
        "aspectRatio": "320:213",
        "margin": "none",
        "position": "relative"
      },
      "body": {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "weight": "bold",
            "size": "sm",
            "wrap": true,
            "text": "จะเอาเต้า - เต้า"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "เต้า - เต้า - เต้า",
              "text": "milkmilk 1"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          }
        ],
        "spacing": "sm",
        "paddingAll": "13px"
      }
    }
  ]
}`
}

func ReplyMessage(bot *messaging_api.MessagingApiAPI, replyToken string, userMsg string) (*messaging_api.ReplyMessageResponse, error) {
	if userMsg == "Menu" {
		json := returnFlexMessage()
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
		// initializers.DB.Exec("DEALLOCATE ALL")
		// initializers.DB.Create(&activity)
		activity.Save()

		// log.Println(userMsg)

		// if err != nil {
		// 	return ReplyErrorMessage(bot, replyToken, "'"+userMsg+"'  น้องวา- ไม่เข้าใจ งับบบบ")
		// }
		return ReplyErrorMessage(bot, replyToken, "น้องวารับทราบ")
	} else if userMsg == "Summary" {
		var sumDrinkMilk result
		// initializers.DB.Exec("DEALLOCATE ALL")
		initializers.DB.Model(&models.Activities{}).Select("SUM(CAST(actity_value AS INTEGER)) as Sum").Where("actity_type = ?", "drinkmilk").Find(&sumDrinkMilk)
		// SELECT name, sum(age) as total FROM `users` WHERE name LIKE "group%" GROUP BY `name` LIMIT 1

		log.Println("drinkmilk===>", sumDrinkMilk)
		var TestSend = fmt.Sprintf(`น้องวารับทราบ 
ดื่มนมไปแล้ว %d ออนซ์`, sumDrinkMilk)
		return ReplyErrorMessage(bot, replyToken, TestSend)
	} else {
		return ReplyErrorMessage(bot, replyToken, "'"+userMsg+"'  น้องวา- ไม่เข้าใจ งับบบบ")
	}

}

func ReplyErrorMessage(bot *messaging_api.MessagingApiAPI, replyToken string, message string) (*messaging_api.ReplyMessageResponse, error) {
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

type FlexMessageWrapper struct {
	*linebot.FlexMessage
}

func (f *FlexMessageWrapper) GetType() string {
	return "flex"
}
