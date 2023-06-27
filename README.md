# multi-post
discord, slackの指定したチャンネルに同じ内容のテキストを同時に投稿します．\
以下のファイルを実行ファイルと同じディレクトリにおいて実行すると動きます．
- `config.yaml`
- credentialsで指定するファイル
# 使い方
## config.yamlの設定
以下は`config.yaml`の一例です．
```
services:
  onescope-discord:
    type: discord
    credentials: ./onescope-credentials.json
  another-slack:
    type: slack
    credentials: ./another-slack.json
channels:
  webdev:
    - service: onescope-discord
      channel: tech_webdev
    - service: another-slack
      channel: tech_js  
```

| key | description |
| --- | --- |
| `services` | 投稿先のサービス名 |
| `type` | `slack` or `discord` ，サービスの種類 |
| `credentials` | token等の認証情報を含んだファイル |
| `channels` | 同時投稿するチャンネル群の任意の名前 |
| `service` | `services`にあるサービスの中から指定 |
| `channel` | 投稿したいチャンネル名 |

## credentialsで指定するファイルの形式
jsonファイルで以下の様に設定
### discord
```
{
    "Token": "xxxxxx",
    "GuildID": "xxxxxx"
}
```
`Token`は[Discord Developer Portal](https://discord.com/developers/applications)にて取得．接頭辞の`Bot`は付けなくてよい．
`GuildID`はdiscordのサーバーID．
### slack
```
{
    "Token": "xoxp-xxxxxx"
}
```
`Token`はユーザートークンを用いる．BotのScopeは`chat:Write`のみで良い．

## 詳しい始め方
### discord
1. [Discord Developer Portal](https://discord.com/developers/applications)にアクセス．
1. New Applicationをクリックし，新しくBotを作る．
1. サイドバーからBotのタブにアクセス．`Token`を取得．前述したdiscordのjsonファイルの`Token`欄にペースト．
1. サイドバーからOAuth2/Generalのタブにアクセス．Client IDをコピーする．
1. `https://discordapp.com/oauth2/authorize?client_id=YOUR_CLIENT_ID&scope=bot&permissions=0`にアクセス．YOUR_CLIENT_IDをBotのClient IDに置き換える．
1. Botを導入したいサーバーを選択して認証する．
1. ツールを実行するとBotとして指定チャンネルに投稿してくれる．

### slack
1. [Slack API: Applications](https://api.slack.com/apps)にアクセス．
1. Create New Appをクリックして新しくBotを作る．この際，`From scratch`を選択．Botを導入したいワークスペースを選択する．
1. サイドバーからOAuth & Permissionsタブにアクセス．
1. User Token Scopesにある，Add an OAuth Scopeをクリック．少しスクロールして`chat:write`を選択．
1. OAuth Tokens for Your Workspaceにある，Install to Workspaceをクリック．
1. User OAuth Tokenから`Token`を取得．前述したslackのjsonファイルの`Token`欄にペースト．
1. ツールを実行すると上記の作業をしていたアカウントとして指定チャンネルに投稿してくれる．
