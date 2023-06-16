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
    credentials: ./another-slack.env
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
### slack
```
{
    "Token": "xxxxxx"
}
```

