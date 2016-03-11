package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

// SSL/TLS Email Example

func main() {
	app := cli.NewApp()

	app.Version = "0.0.2"
	app.Name = "Quck SMTP TLS Checker"
	app.Usage = "./checktls --n smtp.gmail.com"
	app.Flags = []cli.Flag{
		// Core Client Configuration
		cli.StringFlag{
			Name:  "port, p",
			Value: "587",
			Usage: "Optional server port. Default 587",
		},
		cli.StringFlag{
			Name:  "hostnames, n",
			Value: "",
			Usage: "One or more SMTP hostnames comma delimited.",
		},
	}
	app.Action = func(c *cli.Context) {

		if c.String("hostnames") == "" {
			log.Fatalf("Error: --hostnames is a required field\n")
			return
		}

		if c.String("port") == "" {
			log.Fatalf("Error: --port cannot be an empty string\n")
			return
		}

		hosts := strings.Split(c.String("hostnames"), ",")
		for i := range hosts {
			log.Printf("checking: %s:%s", hosts[i], c.String("port"))
			//host = fmt.Sprintf("%s%s, ", row, fields[i])
			tlsTest(hosts[i], c.String("port"))
		}

	}
	app.Run(os.Args)
}

func tlsTest(host string, port string) {

	smtpserver := host + ":" + port

	config := &tls.Config{ServerName: host}

	c, err := smtp.Dial(smtpserver)
	if err != nil {
		log.Printf("Could not connect to %s:%s\n", host, port)
		log.Printf("\x1b[31;1mError\x1b[0m  \"%v\"\n", err)
		return
	}

	err = c.StartTLS(config)
	if err != nil {
		errorMsg := fmt.Sprintf("\x1b[31;1mError:\x1b[0m [%s:%s] failed with error message\n\t\x1b[31;1m%s %s\x1b[0m", host, port, host, err)

		log.Println(errorMsg)
	} else {
		log.Println("✔ ", host, " certificate is good")
	}

}
