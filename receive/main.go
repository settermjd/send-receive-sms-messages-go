package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/twilio/twilio-go/twiml"
)

// Handle requests where no response, effectively, is sent back to Twilio. More
// specifically, the TwiML contained in the body of the response does not provide
// further instructions to Twilio to take any action. Given that, no response will
// be sent to the sender of the incoming SMS.
func handleNoResponse(w http.ResponseWriter, r *http.Request) {
	twimlResult, err := twiml.Messages([]twiml.Element{
		&twiml.MessagingMessage{},
	})
	fmt.Println("Sending response without a body.")
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(err.Error()))
	} else {
		// w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(twimlResult))
	}
}

// Handle requests where a response is sent back to Twilio, instructing it to
// send a reply SMS to the sender of the original SMS.
//
// If the response body is "never gonna", regarless of case, then a response,
// based on "Never Gonna Give You Up" by Rick Astley, will be sent as the body
// of the SMS.
func handleSendResponse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(err.Error()))
	}

	body := r.PostForm.Get("Body")

	fmt.Printf("Request received with body: \"%s\"\n", body)
	fmt.Println("Sending response with a body.")

	const defaultResponse = "I just wanna tell you how I'm feeling - Gotta make you understand"
	options := [6]string{
		"give you up",
		"let you down",
		"run around and desert you",
		"make you cry",
		"say goodbye",
		"tell a lie, and hurt you",
	}

	message := &twiml.MessagingMessage{}
	if strings.ToLower(body) == "never gonna" {
		message.Body = options[rand.Intn(len(options)-1)]
	} else {
		message.Body = defaultResponse
	}

	twimlResult, err := twiml.Messages([]twiml.Element{
		message,
	})
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(err.Error()))
	} else {
		// w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(twimlResult))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/receive/no-response", handleNoResponse)
	mux.HandleFunc("/receive/with-response", handleSendResponse)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
