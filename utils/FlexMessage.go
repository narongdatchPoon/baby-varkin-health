package utils

import (
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func ReturnFlexMessage() string {
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
              "label": "0.5 ออนซ์",
              "text": "drinkmilk 0.5"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
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
              "label": "อุนจิ",
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
              "label": "0.5 ออนซ์",
              "text": "pumpmilk 0.5"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
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
              "label": "ซ้าย",
              "text": "milkmilk left"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": " ขวา",
              "text": "milkmilk right"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "หยุด",
              "text": "milkmilk stop"
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
            "text": "เก็บนม"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "0.5 ออนซ์",
              "text": "stockmilk 0.5"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "1 ออนซ์",
              "text": "stockmilk 1"
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
              "text": "stockmilk 2"
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
              "text": "stockmilk 3"
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
              "text": "stockmilk 4"
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
              "text": "stockmilk 5"
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
              "text": "stockmilk 6"
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



func FlexMessageSummary() string {
	return `{
  "type": "carousel",
  "contents": [
  {
      "type": "bubble",
      "size": "micro",
      "hero": {
        "type": "image",
        "url": "https://vos.line-scdn.net/card-type-message-image-2024/307axhqi/1724129109711-f5b1FCb3sCoC66tepiitWuejyYfybKRYFCTmrFsB7C9MwiH4vA",
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
            "text": "ดูประวัติทั้งหมด"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "History",
              "text": "summarysum history-all"
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
            "text": "กินนม"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "สรุป",
              "text": "summarysum drinkmilk"
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
              "label": "สรุป",
              "text": "summarysum diaper"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "ประวัติ",
              "text": "summarysum history-diaper"
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
              "label": "สรุป",
              "text": "summarysum sleep-daily"
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
              "label": "สรุป",
              "text": "summarysum takeabath"
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
            "text": "ปั้มนม"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "สรุป",
              "text": "summarysum pumpmilk-daily"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "ประวัติ",
              "text": "summarysum history-pumpmilk"
            },
            "margin": "xs",
            "height": "sm",
            "style": "primary"
          },
          {
            "type": "button",
            "action": {
              "type": "message",
              "label": "สรุปเก็บนม",
              "text": "summarysum stock"
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
              "label": "สรุป",
              "text": "summarysum milkmilk"
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

func DateTimePicker(title string, data string, deleteID string) string {
	return fmt.Sprintf(`{
    "type": "bubble",
    "body": {
      "type": "box",
      "layout": "vertical",
      "contents": [
        
        {
          "type": "text",
          "weight": "bold",
          "size": "xl",
          "text": "%s"
        },
        {
          "type": "text",
          "weight": "regular",
          "size": "xs",
          "text": "***Remark หากเป็นเวลาปัจจุบันไม่ต้องเลือก"
        }
      ]
    },
    "footer": {
      "type": "box",
      "layout": "vertical",
      "spacing": "sm",
      "contents": [
        {
          "type": "button",
          "style": "link",
          "height": "sm",
          "action": {
            "type": "datetimepicker",
            "label": "แก้ไข วัน เวลา",
            "mode": "datetime",
            "data": "%s",
            "color": "#ffdd00"
          }
        },
        {
          "type": "button",
          "action": {
            "type": "message",
            "label": "ลบ กดผิด",
            "text": "%s"
          },
          "height": "sm",
          "color": "#EF4455"
        }
      ],
      "flex": 0
    }
  }`, title, data, deleteID)
}

type FlexMessageWrapper struct {
	*linebot.FlexMessage
}

func (f *FlexMessageWrapper) GetType() string {
	return "flex"
}

func ReplyFlexMessage(json string, bot *messaging_api.MessagingApiAPI, replyToken string) (*messaging_api.ReplyMessageResponse, error) {
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
}
