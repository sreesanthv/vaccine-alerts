# vaccine-alerts
Send notification for available vaccine slots through following medium,
- Slack
- Telegram

## Slack Integration
- Create slack channel for alerts.
- Create a Slack App.
- Create a webhook URL which is connected to slack channel.

## Setup
- Export the followin env variables
    - COWIN_URL (https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByDistrict)
    - ALERT_DAYS (7): How many days from today should be checked.
    - SLACK_WEBHOOK_URL (): Slack Webhook URL integrated with slack channel.
    - COWIN_DISTRICT_ID: 302,305
        - Malappuram - 302
        - Kozhikode - 305
    - COWIN_FIRST_DOSE_ONLY (false): Only get alert for first dose.
    - COWIN_SECOND_DOSE_ONLY (false): Only get alert for second dose.
    - COWIN_FREE_VACCINE_ONLY (false): Only get alerts for free (cost) vaccines only.
    - MIN_AGE (): Only get alerts satisfying minumum age.
    - TELEGRAM_BOT_TOKEN - Token of bot with write message permission to chat.
    - TELEGRAM_CHAT_ID - Chat ID of channel
- go build
- ./vaccine-alerts telegram