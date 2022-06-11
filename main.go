package main

import (
	"log"

	"github.com/sirijagadeesh/sendMailUsingGmailAPI/config"
	"github.com/sirijagadeesh/sendMailUsingGmailAPI/email"
)

func main() {
	log.Println("Hello World!!")

	if err := config.Load(); err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", config.GmailAPI())
	sender, err := email.GmailSender()
	if err != nil {
		log.Fatalln()
	}
	if err := sender.Send(
		"Test Gmail Local <xxxxxxxxxx@gmail.com>",
		[]string{"xxxxxxxxx@gmail.com", "xxxxxxxx@gmail.com"},
		"Test mail from Gmail API from Local PC",
		`
		Hello xxxxxx,
		Good evening .. we are sending mail using Gmail API with golang

		Regards,
		xxxxxx xxxxx
		`); err != nil {
		log.Fatalln(err)
	}

}
