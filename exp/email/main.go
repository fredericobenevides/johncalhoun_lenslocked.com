package main

import (
	"fmt"

	"gopkg.in/mailgun/mailgun-go.v1"
)

func SendSimpleMessage() (string, error) {
	mg := mailgun.NewMailgun("sandboxffa740cdcf1f484db89220e5fb9eeeb6.mailgun.org",
		"bc1c88453b59152e79b818613c1a1343-87cdd773-8fae6019",
		"pubkey-fa480b2acf9d97213a48152b1c793282")
	m := mg.NewMessage(
		"Excited User <support@lenslocked.com>",
		"Hello",
		"Testing some Mailgun awesomeness!",
		"fredbene@gmail.com",
	)
	_, id, err := mg.Send(m)
	if err != nil {
		fmt.Println("Erro", err)
	}
	return id, err
}

func main() {
	SendSimpleMessage()

}
