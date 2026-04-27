package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	// Load the environment variables from the .env file in the current directory.
	// See .env.example for the required environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialise a new Twilio Rest Client to send an SMS to the recipient
	client := twilio.NewRestClient()
	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("RECIPIENT"))
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody("This is the ship that made the Kessel Run in fourteen parsecs?")

	// Attempt to send an SMS to the recipient, logging both success and failure
	_, err = client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}
