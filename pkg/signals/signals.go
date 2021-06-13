package signals

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func Interrupt(cancel context.CancelFunc, logger *logrus.Entry) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case s := <-c:
		logger.Infof("Caught %s signal. Exiting.", s)
		cancel()
		return
	}
}
