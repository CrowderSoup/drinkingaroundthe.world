package email

import (
	"context"
	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type MG struct {
	mailgun *mailgun.MailgunImpl
	sender  string
}

func NewMailgun(domain, apiKey, senderAddress string) *MG {
	return &MG{
		mailgun: mailgun.NewMailgun(domain, apiKey),
		sender:  senderAddress,
	}
}

func (m *MG) SendMagicLink(recipient, magicKey string) error {
	subject := "Login to Drink Around the World"

	message := m.mailgun.NewMessage(m.sender, subject, "", recipient)
	// TODO: Need to read in the template from a file. Also need to make it use the token in a link
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
	_, _, err := m.mailgun.Send(ctx, message)

	if err != nil {
		return err
	}

	return nil
}
