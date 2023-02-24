package logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/joint-online-judge/go-horse/pkg/config"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:03.000",
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		CallerFirst:     true,
	})
	if config.Conf.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
