# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A Twilio SMS messaging sample application in Go.
The repo contains two independent Go modules — one for sending SMS, one for receiving — that can be developed and deployed separately.

## Module Structure

- `send/` — CLI app that sends an SMS via the Twilio REST API (`github.com/twilio/twilio-go`)
- `receive/` — HTTP server (port `:4000`) that handles incoming SMS webhooks and returns TwiML responses

Each module has its own `go.mod`/`go.sum`.
There is no workspace file or shared module.

## Commands

All commands must be run from inside the relevant module directory.

```bash
# Run the send module
cd send && go run main.go

# Run the receive module
cd receive && go run main.go

# Tidy dependencies (run per module)
cd send && go mod tidy
cd receive && go mod tidy

# Build (run per module)
cd send && go build ./...
cd receive && go build ./...
```

There are no tests in this codebase (it is a sample application).

## Environment Configuration

Copy `send/.env.example` to `send/.env` and populate:

| Variable              | Description                           |
| --------------------- | ------------------------------------- |
| `TWILIO_ACCOUNT_SID`  | Twilio account SID                    |
| `TWILIO_AUTH_TOKEN`   | Twilio auth token                     |
| `TWILIO_PHONE_NUMBER` | Sender phone number (E.164 format)    |
| `RECIPIENT`           | Recipient phone number (E.164 format) |

The `send` module loads `.env` via `github.com/joho/godotenv`.
The `receive` module reads env vars directly (no `.env` loading).

## Receive Module Endpoints

- `POST /receive/no-response` — Accepts an SMS without replying
- `POST /receive/with-response` — Accepts an SMS and replies with a TwiML `<Message>`.
  Send "never gonna" to trigger a Rick Astley Easter egg (random lyric from `rickAstleyOptions` slice).

To test locally, expose port 4000 with `ngrok http 4000` and configure the resulting URL as your Twilio webhook.

## TwiML Responses

Both endpoints return XML in TwiML format using `github.com/twilio/twilio-go/twiml`.
The `receive` module constructs a `MessagingResponse` with a `Message` verb — not the SMS REST API.
