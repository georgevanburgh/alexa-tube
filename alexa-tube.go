package main

import (
	"os"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

const (
	HELP_RESPONSE  = "Try asking me about the tube status for a particular line"
	ERROR_RESPONSE = "I'm sorry - I can't help you with that"
)

var Applications = map[string]interface{}{
	"/echo/tfl": alexa.EchoApplication{ // Route
		AppID:    os.Getenv("ALEXA_APP_ID"),
		OnIntent: EchoIntentHandler,
		OnLaunch: EchoIntentHandler,
	},
}

func main() {
	port, portEnvar := os.LookupEnv("HTTP_PLATFORM_PORT")
	if !portEnvar {
		port = "3000" // For testing
	}

	alexa.Run(Applications, port)
}

func EchoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	switch echoReq.GetIntentName() {
	case "AMAZON.HelpIntent":
		echoResp.OutputSpeech(HELP_RESPONSE).EndSession(false)
		break
	case "AMAZON.StopIntent":
		echoResp.EndSession(true)
		break
	case "GetLineStatus":
		station, err := echoReq.GetSlotValue("Line")
		if err != nil {
			echoResp.OutputSpeech(ERROR_RESPONSE)
		} else {
			echoResp.OutputSpeech("You asked me about the " + station + " line").EndSession(true)
		}
		break
	default:
		echoResp.OutputSpeech(ERROR_RESPONSE)
		break
	}
}
