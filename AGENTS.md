# Send & Receive SMS Messages with Go

This sample demonstrates how to send and receive SMS messages using Twilio's Programmable Messaging API and Go.

## Environment Variables

Copy `.env.example` to `.env`. Never commit `.env`.

```bash
cp send/.env.example send/.env
```

| Variable | Where to find | Format |
| -------- | ------------- | ------ |
| `TWILIO_ACCOUNT_SID` | Console homepage or Admin dropdown (top right) → Account Management → Keys & Credentials → API Keys & Tokens | Starts with `AC` |
| `TWILIO_AUTH_TOKEN` | Console homepage or Admin dropdown (top right) → Account Management → Keys & Credentials → API Keys & Tokens → click to reveal | 32-char string. Treat as a password. |
| `TWILIO_PHONE_NUMBER` | Console → Phone Numbers → Manage → Active Numbers | E.164 format: `+15551234567` |
| `RECIPIENT` | The phone number you want to send SMS to | E.164 format: `+15551234567` |

## Commands

```bash
# Install dependencies (send module)
cd send && go mod tidy

# Install dependencies (receive module)
cd receive && go mod tidy

# Run the send module (sends one SMS)
cd send && go run main.go

# Run the receive module (HTTP server on port 4000)
cd receive && go run main.go

# Expose webhooks locally
# Requires ngrok — install and authenticate at https://ngrok.com before running
ngrok http 4000
# Set the resulting URL + /receive/with-response as the webhook in Twilio Console
```

## Project Structure

- `send/main.go` — CLI app that sends a single SMS via Twilio REST API
- `send/.env.example` — environment variable template for the send module
- `receive/main.go` — HTTP server handling incoming SMS webhooks, returns TwiML responses

## Agent Boundaries

**Always:**
- Confirm `send/.env` is configured before running any command
- Use the Environment Variables section to guide the user to each credential — don't ask them to find values without direction
- Confirm the app is running before asking the user to test it

**Never:**
- Run the app with missing or placeholder credentials
- Hardcode credentials or phone numbers in source files
- Skip the `cp send/.env.example send/.env` step

## Verify It's Working

**Send:** After configuring `send/.env` and running `cd send && go run main.go`, the number in `RECIPIENT` should receive an SMS from `TWILIO_PHONE_NUMBER`.

**Receive:** Start the server (`cd receive && go run main.go`), expose it with ngrok, and configure the ngrok URL + `/receive/with-response` as the incoming message webhook for your Twilio number in the Console. Text that number and expect an automated reply.

## Twilio Resources

- [Twilio Console](https://console.twilio.com) — credentials, phone numbers, webhook configuration
- [Programmable Messaging docs](https://www.twilio.com/docs/sms) — SMS API reference
- [twilio-go SDK](https://www.twilio.com/docs/libraries/go) — Go SDK documentation
- [TwiML for Messaging](https://www.twilio.com/docs/messaging/twiml) — TwiML reference for SMS responses
