package main

import (
	"github.com/BradErz/monorepo/services/tasks/web"
	"github.com/sirupsen/logrus"
)

func main() {
	lgr := logrus.New()
	srv, err := web.New(logrus.NewEntry(lgr), ":50051")
	if err != nil {
		logrus.WithError(err).Fatal("failed to listen on port")
	}
	if err := srv.Run(); err != nil {
		logrus.WithError(err).Fatal("failed to run server")
	}
}
