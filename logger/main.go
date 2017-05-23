package logger

import (
	"os"

	"github.com/op/go-logging"
)

var Log = logging.MustGetLogger("example")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var errFormat = logging.MustStringFormatter(
	`%{color}%{time:2006/01/02 15:04:05.000} %{callpath} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

var infoFormat = logging.MustStringFormatter(
	`%{color}%{time:2006/01/02 15:04:05.000} %{shortfile} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

var format = logging.MustStringFormatter(
	`%{color}%{time:2006/01/02 15:04:05.000} %{shortfile} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func InitLogger() {

	errBackEnd := logging.NewLogBackend(os.Stderr, "", 0)
	errBackEndFormatter := logging.NewBackendFormatter(errBackEnd, errFormat)
	errLeveled := logging.AddModuleLevel(errBackEndFormatter)
	errLeveled.SetLevel(logging.ERROR, "")

	infoBackEnd := logging.NewLogBackend(os.Stderr, "", 0)
	infoBackEndFormatter := logging.NewBackendFormatter(infoBackEnd, infoFormat)
	infoLeveled := logging.AddModuleLevel(infoBackEndFormatter)
	infoLeveled.SetLevel(logging.INFO, "")

	file, err := os.OpenFile("logging.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		backend2 := logging.NewLogBackend(file, "", 0)
		backend2Formatter := logging.NewBackendFormatter(backend2, format)
		logging.SetBackend(infoLeveled, errLeveled, backend2Formatter)
	} else {
		Log.Info("Failed to log to file, using default stderr")
		logging.SetBackend(infoLeveled, errLeveled)
	}
}
