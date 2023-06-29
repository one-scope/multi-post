package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/one-scope/multi-post/discord"
	"github.com/one-scope/multi-post/slack"
)

var (
	VERSION = "Version 1.0.0"
)

func main() {
	tShowVersion := flag.Bool("version", false, "Show version")
	tGroup := flag.String("channel", "default", "Set channel")
	tConfigFile := flag.String("config", "config.yaml", "Set config file")
	flag.Parse()
	if *tShowVersion {
		fmt.Fprintln(os.Stdout, VERSION)
		return
	}
	tContent := flag.Arg(0)
	if tContent == "" {
		fmt.Fprintln(os.Stderr, "please specify message")
		os.Exit(1)
	}

	tConfig, tError := readConfig(*tConfigFile)
	if tError != nil {
		fmt.Fprintln(os.Stderr, tError)
		os.Exit(1)
	}

	//Botのセットアップ
	for tKey, tService := range tConfig.Services {
		var tBot Bot
		switch tService.Type {
		case "discord":
			tBot = &discord.Bot{}
		case "slack":
			tBot = &slack.Bot{}
		}
		if tError := tBot.SetCredentials(tService.Credentials); tError != nil {
			fmt.Fprintln(os.Stderr, tError)
			os.Exit(1)
		}
		//discordのみ終了処理
		if tDiscordBot, tOK := tBot.(*discord.Bot); tOK {
			tDiscordBot.Close()
		}
		ServiceMap[tKey] = tBot
	}

	//バリデーション
	tChannels, tOK := tConfig.Groups[*tGroup]
	if !tOK {
		fmt.Fprintln(os.Stderr, "not found \""+*tGroup+"\" channels. please specify channels exactly")
		os.Exit(1)
	}
	//投稿処理
	for _, tGroup := range tChannels {
		if tError := ServiceMap[tGroup.Service].SendMessage(tGroup.Channel, tContent); tError != nil {
			fmt.Fprintln(os.Stderr, tError)
			os.Exit(1)
		}
	}
}
