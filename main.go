package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	gomail "gopkg.in/gomail.v2"
)

func main() {
	var (
		config config
		mail   mail
	)

	fmt.Println("Welcome Gopher ʕ◔ϖ◔ʔ")

	if _, err := os.Stat(".config"); os.IsNotExist(err) {
		fmt.Println("Enter your settings:")

		fmt.Print("SMTP: ")
		fmt.Scan(&config.SMTP)

		fmt.Printf("port: ")
		if _, err := fmt.Scanf("%v", &config.Port); err != nil {
			panic(err)
		}

		fmt.Print("email: ")
		fmt.Scan(&config.Email)

		fmt.Print("password: ")
		fmt.Scan(&config.Pass)

		var remember bool
		fmt.Printf("\n%s\n", "Want to remember? (t/f)")
		if _, err := fmt.Scan(&remember); err != nil {
			panic(err)
		}

		if remember {
			if c, err := json.Marshal(config); err == nil {
				if err := ioutil.WriteFile(".config", c, 0644); err != nil {
					panic(err)
				}
			} else {
				panic(err)
			}
		}
	}

	if file, err := ioutil.ReadFile(".config"); err == nil {
		if err = json.Unmarshal(file, &config); err != nil {
			panic(err)
		}
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.Email)
	m.SetHeader("To", mail.Headers.To...)
	m.SetAddressHeader("Cc", mail.Headers.Cc.Adress, mail.Headers.Cc.Name)
	m.SetHeader("Subject", mail.Headers.Subject)
	m.SetBody("text/html", mail.Body.Text)

	for _, file := range mail.Attach {
		m.Attach(file)
	}

	// m := gomail.NewMessage()
	// m.SetHeader("From", config.Email)
	// m.SetHeader("To", "zeucxb@gmail.com", "zeu-x@hotmail.com")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	// m.SetHeader("Subject", "Hello!")
	// m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/Users/zeucxb/Downloads/11702749.png")
	// m.Attach("/Users/zeucxb/Downloads/photo.png")

	d := gomail.NewDialer(config.SMTP, config.Port, config.Email, config.Pass)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

type config struct {
	SMTP  string `json:"smtp"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
	Port  int    `json:"port"`
}

type mail struct {
	Headers struct {
		To []string `json:"to"`
		Cc struct {
			Adress string `json:"adress"`
			Name   string `json:"name"`
		} `json:"cc"`
		Subject string `json:"subject"`
	} `json:"headers"`
	Body struct {
		Text string `json:"text"`
		File string `json:"file"`
	} `json:"body"`
	Attach []string `json:"attach"`
	Mails  []string `json:"mails"`
}
