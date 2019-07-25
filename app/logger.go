package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	
	"github.com/sirupsen/logrus"
)

const RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z07:00"

// initLog init log setting
func initLogger() {
	// newGenericLogger()
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: RFC3339NanoFixed,
		DisableColors:   false,
		FullTimestamp:   true,
	})

	logrus.Info("logger construction succeeded")
}