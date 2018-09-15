package main

import (
	. "fmt"
	. "./keys"
	"time"
)

func main(){
	api := GetTwitterApi()
	text := ChooseTweet()
	tweethour := time.Now().Hour()
	if tweethour == 10 {
		tweet, err := api.PostTweet(text, nil)
		if err != nil {
			panic(err)
		}
		Print(tweet.Text)
	}
}

