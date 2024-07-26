package twiliotext

import (
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Config struct {
	TwilioAccountSid string
	TwilioAuthToken  string
	TwilioNumber     string
}

type TwilioText struct {
	config *Config
}

type TwilioMethod interface {
	SendMessage(body, to string) (*twilioApi.ApiV2010Message, error)
	CheckMessage(sid string) (*twilioApi.ApiV2010Message, error)
}

func NewTwilioText(config *Config) TwilioMethod {
	return &TwilioText{config: config}
}

// Username: os.Getenv("TWILIO_ACC_SID"),
// Password: os.Getenv("TWILIO_AUTH_TOKEN"),
// params.SetFrom(os.Getenv("TWILIO_NUMBER"))
func (t *TwilioText) SendMessage(body, to string) (*twilioApi.ApiV2010Message, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: t.config.TwilioAccountSid,
		Password: t.config.TwilioAuthToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetBody(body)
	params.SetFrom(t.config.TwilioNumber)
	params.SetTo(to)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TwilioText) CheckMessage(sid string) (*twilioApi.ApiV2010Message, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: t.config.TwilioAccountSid,
		Password: t.config.TwilioAuthToken,
	})

	params := &twilioApi.FetchMessageParams{}
	params.SetPathAccountSid(t.config.TwilioNumber)

	resp, err := client.Api.FetchMessage(sid, params)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
