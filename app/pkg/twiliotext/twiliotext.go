package twiliotext

import (
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"
)

func SendMessage(body, to string) (*twilioApi.ApiV2010Message, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACC_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetBody(body)
	params.SetFrom(os.Getenv("TWILIO_NUMBER"))
	params.SetTo(to)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CheckMessage(sid string) (*twilioApi.ApiV2010Message, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACC_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	params := &twilioApi.FetchMessageParams{}
	params.SetPathAccountSid(os.Getenv("TWILIO_ACC_SID"))

	resp, err := client.Api.FetchMessage(sid, params)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
