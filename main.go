// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"math/rand"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
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
				if message.Text =="還錢啦"{
					rand.Seed(time.Now().Unix())
					Audio := []string{
      						"https://s19.aconvert.com/convert/p3r68-cdx67/j3ook-eu935.mp3",
     						"https://s19.aconvert.com/convert/p3r68-cdx67/3gir2-me8bj.mp3",
      						"https://s19.aconvert.com/convert/p3r68-cdx67/hm037-7jp07.mp3",
     				 		"https://s19.aconvert.com/convert/p3r68-cdx67/ymjd8-per4w.mp3",
      						"https://s19.aconvert.com/convert/p3r68-cdx67/4dojm-7k8y6.mp3",
						"https://s19.aconvert.com/convert/p3r68-cdx67/mf34t-w9oi3.mp3",
     					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewAudioMessage(Audio[rand.Intn(len(Audio))],1)).Do(); err != nil {
					log.Print(err)
					}
				}
			}
		}
	}
}
