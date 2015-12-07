package smtpd

import (
	"errors"
	"net/smtp"
	"strings"
)

func ExampleServer() {
	var server *Server

	// No-op server. Accepts and discards
	server = &Server{}
	server.ListenAndServe("127.0.0.1:10025")

	// Relay server. Accepts only from single IP address and forwards using the Gmail smtp
	server = &Server{
		HeloChecker: func(peer Peer, name string) error {
			if !strings.HasPrefix(peer.Addr.String(), "42.42.42.42:") {
				return errors.New("Denied")
			}
			return nil
		},

		Handler: func(peer Peer, env Envelope) error {
			return smtp.SendMail(
				"smtp.gmail.com:587",
				smtp.PlainAuth(
					"",
					"username@gmail.com",
					"password",
					"smtp.gmail.com",
				),
				env.Sender,
				env.Recipients,
				env.Data,
			)
		},
	}

	server.ListenAndServe("127.0.0.1:10025")
}
