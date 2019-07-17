# Google Colud FunctionでGolangのojichatを動かす

## command

とりあえずアジアリージョンで動かす。PubSubのほうがキレイそうだけどHTTP Triggerで

```
gcloud functions deploy ojichat --runtime go111 --trigger-http --region asia-northeast1 --entry-point Ojichat --set-env-vars VERIFICATION_TOKEN=your_token
```

