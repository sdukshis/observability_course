#!/bin/bash

curl -X POST \
     -H 'Content-Type: application/json' \
     -d '{"chat_id": "-1001908039645", "text": "This is a test from curl [link](https://google.com)", "disable_notification": true, "parse_mode": "Markdown"}' \
     https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/sendMessage
