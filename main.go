package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/one-scope/multi-post/discord"
	"github.com/one-scope/multi-post/slack"
)

var (
	CONFIG_FILE = "config.yaml"
	VERSION     = "Version 1.0.0"
)

func main() {
	tVersion := flag.Bool("version", false, "Show version")
	tChannel := flag.String("channel", "default", "Set channel")
	flag.Parse()
	if *tVersion {
		fmt.Fprintln(os.Stdout, VERSION)
	}
	tContent := flag.Arg(0)

	tConfig, tError := ReadConfig(CONFIG_FILE)
	if tError != nil {
		fmt.Fprintln(os.Stderr, tError)
		os.Exit(1)
	}

	//Botのセットアップ
	//このswitchのタイプ分けいる？もうちょっとスマートに書けへんの
	for key, tService := range tConfig.Services {
		switch tService.Type {
		case "discord":
			tDiscord := discord.Bot{}
			if tError := tDiscord.SetCredentials(tService.Credentials); tError != nil {
				fmt.Fprintln(os.Stderr, tError)
				os.Exit(1)
			}
			ServiceMap[key] = &tDiscord
		case "slack":
			tSlack := slack.Bot{}
			if tError := tSlack.SetCredentials(tService.Credentials); tError != nil {
				fmt.Fprintln(os.Stderr, tError)
				os.Exit(1)
			}
			ServiceMap[key] = &tSlack
		}
	}

	//投稿処理
	//変数名がまどろっこしいこれあかんと思う
	for _, channel := range tConfig.Channels[*tChannel] {
		if tError := ServiceMap[channel.Service].SendMessage(channel.Channel, tContent); tError != nil {
			fmt.Fprintln(os.Stderr, tError)
			os.Exit(1)
		}
	}
}
