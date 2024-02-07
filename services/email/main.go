package email

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/spf13/viper"
)

type MG struct {
	mailgun *mailgun.MailgunImpl
	sender  string
}

func NewMailgun() *MG {
	mailgunApiKey := viper.GetString("mailgun_api_key")
	mailgunDomain := viper.GetString("mailgun_domain")
	mailgunSendingAddress := viper.GetString("mailgun_sending_address")

	return &MG{
		mailgun: mailgun.NewMailgun(mailgunDomain, mailgunApiKey),
		sender:  mailgunSendingAddress,
	}
}

func (m *MG) SendMagicLink(recipient, magicKey string) (string, string) {
	subject := "Login to Drink Around the World"

	message := m.mailgun.NewMessage(m.sender, subject, "", recipient)
	body := fmt.Sprintf(`
<html>
<body>
	<h1>Login to Drink Around the World</h1>
  <p>%s</p>
</body>
</html>
`, magicKey)

	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := m.mailgun.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	return resp, id
}
