package notification

import (
	. "github.com/gyf841010/pz-infra/logging"

	"github.com/aws/aws-sdk-go/service/ses"
)

type SESClient interface {
	SendEmail(input *ses.SendEmailInput) (*ses.SendEmailOutput, error)
	//	SendRawEmail(input *ses.SendRawEmailInput) (*ses.SendRawEmailOutput, error)
}

type MockSESClient struct {
}

func (*MockSESClient) SendEmail(input *ses.SendEmailInput) (*ses.SendEmailOutput, error) {
	Log.Debug("mock send email")
	messageId := "mock_message_id"
	return &ses.SendEmailOutput{MessageId: &messageId}, nil
}

//func (*MockSESClient)SendRawEmail(input *ses.SendRawEmailInput) (*ses.SendRawEmailOutput, error)  {
//	log.Debugf("mock send row email")
//	messageId := "mock_raw_message_id"
//	return &ses.SendRawEmailOutput{MessageId: &messageId}, nil
//}
