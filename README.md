# Google Colud FunctionでGolangのojichatを動かす

## やりたいこと

Google Cloud Function はnode.jsやpythonだけでなく、Goの関数を実行できる。
[ojichat](https://github.com/greymd/ojichat)のオリジナルはGoで開発されており、これをそのままCloud Functionで動作させたい。

slackからslashコマンドで呼び出すこととした

## GCP側の準備

アカウントを作って `gcloud` コマンドを使えるようにしておく

## Slack slash command

<https://api.slack.com/slash-commands>

スラッシュコマンドでojichatを呼び出す。 APIの作成はapi.slackドメインであってワークスペースのドメインではないので注意。

スラッシュコマンドを作成すると、Verification Tokenが払い出される。これを認証で使う。

## Golangでの実装

コード参照。go.modの作成が必須なので、golang 1.11以降で開発すること。

Verification Tokenは環境変数に格納する。

ojichat自体は `-e` , `-p` オプションを持っているが、とりあえず引数は固定にした。Nameオプションのみ有効。

slackの認証を細かに行ったり受信結果をパースしたり、という部分はライブラリの恩恵を受ける。
<https://github.com/nlopes/slack>

## デプロイ

とりあえずアジアリージョンで動かす。PubSubのほうがキレイそうだけどHTTP Triggerで。

```
gcloud functions deploy ojichat --runtime go111 --trigger-http --region asia-northeast1 --entry-point Ojichat --set-env-vars VERIFICATION_TOKEN=your_token
```

