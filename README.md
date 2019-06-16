# gcal-converter

convet some calendar to google calendar

## how to use

### development

local

```
$ go run main.go
```

deploy

```
$ gcloud functions deploy <APP_NAME> --entry-point HelloPubSub --runtime go111 --trigger-topic <TOPIC_NAME> --source=./app
```

