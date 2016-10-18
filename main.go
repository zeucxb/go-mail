package main

import gomail "gopkg.in/gomail.v2"

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "zeucxb@gmail.com", "zeu-x@hotmail.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/Users/zeucxb/Downloads/11702749.png")
	m.Attach("/Users/zeucxb/Downloads/photo.png")

	d := gomail.NewDialer("smtp.gmail.com", 587, "zeucxb", "ntkjyfbeemkuihss")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
