package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4452810766903-4453773811655-1RgTCGwnPUlZldoHHh97Z6PC")
	os.Setenv("CHANNEL_ID", "C04DJN77P46")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"Dailyblogram.png"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			File:     fileArr[i],
			Channels: channelArr,
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", &err)
			return
		}

		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URLPrivate)
	}
}
