package mail

import "github/elliot9/ginExample/pkg/mailer"

type Service interface {
	Welcome(to, userName, verificationLink string) error
}

type service struct {
	mailer mailer.Mailer
}

func New(mailer mailer.Mailer) Service {
	return &service{mailer: mailer}
}
