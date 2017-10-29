package main

import (
	"fmt"
	"strings"
	"log"
	"net/http"
	"os"
	"math/rand"
	
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	if _, err := bot.PushMessage("U2c68fd429a99dceccc8956571baa7d00", linebot.NewTextMessage("hello")).Do(); err != nil {
		log.Print(err)
	}
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				//var txt = Send(message.Text);
				//rand.Seed(99)
				answers := []string{"洞么參在此誰敢放肆","就跟你說要有外野了",
						    "洞么參剛到步校","打手槍次數是零",
						    "欸幹嘛不要這樣啊","都欺負我","欸我不是二分隊的啊",
						    "我沒有要簽啊","小心我把你們埋進土裡",
						    "013 洞妖三 1800我告訴我家狗我想尿尿它要帶我去，預計1830回家"}
				//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!"+txt+"  "+event.Source.UserID+"   "+event.ReplyToken)).Do(); err != nil {
				//	log.Print(err)
				//}
				if strings.Contains(message.Text, "哈哈哈"){
					var txt =answers[rand.Intn(len(answers))] 
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(txt)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
		
		if event.Type == linebot.EventTypeFollow {
			var text = "Hi!請輸入一個人名，就會出現意想不到的結果喔!"
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text)).Do(); err != nil {
				log.Print(err)
			}
		}
	}
}
