package main

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

func main() {
	var smtp, email, pass string
	var port int

	fmt.Println("Welcome Gopher ʕ◔ϖ◔ʔ")

	fmt.Println("Enter your settings:")

	fmt.Print("SMTP: ")
	fmt.Scan(&smtp)

	fmt.Printf("port: ")
	if _, err := fmt.Scanf("%v", &port); err != nil {
		panic(err)
	}

	fmt.Print("email: ")
	fmt.Scan(&email)

	fmt.Print("password: ")
	fmt.Scan(&pass)

	var remember bool
	fmt.Printf("\n%s\n", "Want to remember? (t/f)")
	if _, err := fmt.Scan(&remember); err != nil {
		panic(err)
	}

	if remember {
		fmt.Println("TRUE")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", "zeucxb@gmail.com", "zeu-x@hotmail.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/Users/zeucxb/Downloads/11702749.png")
	m.Attach("/Users/zeucxb/Downloads/photo.png")

	// d := gomail.NewDialer("smtp.gmail.com", 587, "zeucxb", "")
	d := gomail.NewDialer(smtp, port, email, pass)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
