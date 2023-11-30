package main

import (
	"log"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	
	"net/http"

  "github.com/gin-gonic/gin"
)

func main() {
	bot, err := linebot.New("8d105cbe79e183a3a3918534cdd5ff05", "RyQHRJDG8gIFup1G+79Wxxir3wUNSl8OPJ1heITYMrRymsfpbmfW114aLoe4961KiuQoCoQMViJOMLGUOOeCg9jozHVboRQPp3blGo1NN4oNypHC8bqLnsxOWgf9j66+cZLPtabMGPcFNYxNcJJ8kAdB04t89/1O/w1cDnyilFU=")
	log.Println(err)
	r := gin.Default()

	var messages = linebot.NewTextMessage("I AM BOT!")
	_, err = bot.PushMessage("U167fd8501f75f99cb52c097ac1fe28b2", messages).Do()
	if err != nil {
		// Do something when some bad happened
		log.Println(err)
	}
	
	r.POST("/webhook", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			// Do something when something bad happened.
			
		}
		
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				// Do Something...
				log.Println(event.Type)
				switch message := event.Message.(type) {
					case *linebot.TextMessage: 
					{
						log.Println(message.Text)
						log.Println(event.Source.UserID)
						if message.Text == "hello" {
							var messages = linebot.NewTextMessage("Hello, world!")
							replyToken := event.ReplyToken

							flexContainer, err := linebot.UnmarshalFlexMessageJSON(makeFlexMessage())
							var flexMessage = linebot.NewFlexMessage("เมนูอาหารครับ", flexContainer)
							if err != nil {
								log.Println(err)
							}

							_, err = bot.ReplyMessage(replyToken, flexMessage, messages).Do()
							if err != nil {
								// Do something when some bad happened
								log.Println(err)
							}
							// _, err := bot.PushMessage(event.Source.UserID, messages).Do()
							// if err != nil {
							// 	// Do something when some bad happened
							// 	log.Println(err)
							// }
						}
					}
					case *linebot.StickerMessage: 
					{
						log.Println(message.Keywords)
					}
				}
			} else if event.Type == linebot.EventTypePostback {
				log.Println(event.Postback.Data)
				if event.Postback.Data == "Burger" {
					replyToken := event.ReplyToken
					var messages = linebot.NewTextMessage("รับออเดอร์ " + event.Postback.Data + " เรียบร้อยครับ" )
					_, err = bot.ReplyMessage(replyToken, messages).Do()
							if err != nil {
								// Do something when some bad happened
								log.Println(err)
							}
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})

  })

  r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func makeFlexMessage() []byte {
	return []byte(`
	{
		"type": "bubble",
		"hero": {
			"type": "image",
			"url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_2_restaurant.png",
			"size": "full",
			"aspectRatio": "20:13",
			"aspectMode": "cover",
			"action": {
				"type": "uri",
				"uri": "https://linecorp.com"
			}
		},
		"body": {
			"type": "box",
			"layout": "vertical",
			"spacing": "md",
			"action": {
				"type": "uri",
				"uri": "https://linecorp.com"
			},
			"contents": [
				{
					"type": "text",
					"text": "Brown's Burger",
					"size": "xl",
					"weight": "bold"
				},
				{
					"type": "box",
					"layout": "vertical",
					"spacing": "sm",
					"contents": [
						{
							"type": "box",
							"layout": "baseline",
							"contents": [
								{
									"type": "icon",
									"url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/restaurant_regular_32.png"
								},
								{
									"type": "text",
									"text": "$100.5",
									"weight": "bold",
									"margin": "sm",
									"flex": 0
								},
								{
									"type": "text",
									"text": "400kcl",
									"size": "md",
									"align": "end",
									"color": "#aaaaaa"
								}
							]
						},
						{
							"type": "box",
							"layout": "baseline",
							"contents": [
								{
									"type": "icon",
									"url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/restaurant_large_32.png"
								},
								{
									"type": "text",
									"text": "$15.5",
									"weight": "bold",
									"margin": "sm",
									"flex": 0
								},
								{
									"type": "text",
									"text": "550kcl",
									"size": "md",
									"align": "end",
									"color": "#aaaaaa"
								}
							]
						}
					]
				},
				{
					"type": "text",
					"text": "Sauce, Onions, Pickles, Lettuce & Cheese",
					"wrap": true,
					"color": "#aaaaaa",
					"size": "xxs"
				}
			]
		},
		"footer": {
			"type": "box",
			"layout": "vertical",
			"contents": [
				{
					"type": "button",
					"style": "primary",
					"color": "#905c44",
					"margin": "xxl",
					"action": {
						"type": "postback",
						"label": "Buy Now!",
						"data": "Burger"
					}
				}
			]
		}
	}`)
}