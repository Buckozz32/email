package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/emersion/go-smtp"
)

var address = "host"

func init() {
	flag.StringVar(&address, "host", address, "Listen address")
}

type backend struct {}

func (backend *backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &session{}, nil
}


func (s *session) AutoPlain(username, password string) error {
	return nil
}

func (s *session) Mail(from string, opts *smtp.MailOptions) error {
	return nil
}

func (s *session) Rcpt(to string, opts *smtp.RcptOptions) error {
	return nil
}

func (s *session) Data(r io.Reader) error {
	return nil
}

func (s *session) Reset() {}

func (s *session) Logout() error {
	return nil
}

func main() {
	flag.Parse()

	s := smtp.NewServer(&backend{})

	s.Addr = address
	s.Domain = "localhost"
	s.AllowInsecureAuth = true
	s.Debug = os.Stdout

	log.Println("Starting SMTP server at", address)
	log.Fatal(s.ListenAndServe())
}
