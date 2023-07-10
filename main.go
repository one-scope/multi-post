package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/one-scope/multi-post/discord"
	"github.com/one-scope/multi-post/slack"
)

const (
	version = "Version 1.0.0"
)

func main() {
	tShowVersion := flag.Bool("version", false, "Show version")
	tGroupID := flag.String("channel", "default", "Set channel")
	tConfigFile := flag.String("config", "config.yaml", "Set config file")
	flag.Parse()
	if *tShowVersion {
		fmt.Fprintln(os.Stdout, version)
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
	var botByServiceName = make(map[string]Bot)
	for tKey, tService := range tConfig.ServiceByID {
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
		botByServiceName[tKey] = tBot
	}

	//バリデーション
	tGroups, tOK := tConfig.GroupByID[*tGroupID]
	if !tOK {
		fmt.Fprintln(os.Stderr, "not found \""+*tGroupID+"\" channels. please specify channels exactly")
		os.Exit(1)
	}
	//投稿処理
	for _, tGroup := range tGroups {
		if tError := botByServiceName[tGroup.Service].SendMessage(tGroup.Channel, tContent); tError != nil {
			fmt.Fprintln(os.Stderr, tError)
			os.Exit(1)
		}
	}
}
