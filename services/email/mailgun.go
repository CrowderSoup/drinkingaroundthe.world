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
	urlString := fmt.Sprintf(`https://drinkaroundthe.world/auth/verify?m=%s`, magicKey)
	anchorTag := fmt.Sprintf(`<a href="%s" target="_blank">Click here</a>`, urlString)
	body := fmt.Sprintf(`
<html>
<body>
	<h1>Login to Drink Around the World</h1>
  <p>
    %s to log in to Drink Around the World. If you're unable to click the link, copy the URL below:
  </p>
  <p>
    %s
  </p>
</body>
</html>
`, anchorTag, urlString)

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
