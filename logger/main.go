package logger

import (
	"os"

	"github.com/op/go-logging"
)

var Log = logging.MustGetLogger("example")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func InitLogger() {
	// You could set this to any `io.Writer` such as a file
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	//backend1Leveled := logging.AddModuleLevel(backend1, format)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend1Leveled := logging.AddModuleLevel(backend1Formatter)
	//backend1Leveled.SetLevel(logging.ERROR, "")
	//backend1Formatter.SetLevel(logging.ERROR, "")
	logging.SetBackend(backend1Leveled)

	file, err := os.OpenFile("logging.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		backend2 := logging.NewLogBackend(file, "", 0)
		backend2Formatter := logging.NewBackendFormatter(backend2, format)
		logging.SetBackend(backend1Formatter, backend2Formatter)
	} else {
		Log.Info("Failed to log to file, using default stderr")
		logging.SetBackend(backend1Formatter)
	}
}

func main() {
	//
	// // Set the backends to be used.
	// fmt.Printf("t1 =  %s", time.Since(start))
	// start3 := time.Now()
	// log.Info("info")
	// fmt.Printf("t3 =  %s", time.Since(start3))
	//
	// start4 := time.Now()
	// log.Error("err")
	// fmt.Printf("t4 =  %s", time.Since(start4))
	//
	// log.Warning("warning")
	//
	// log.Critical("crit")
	//
	// fmt.Printf("%s", time.Since(start))

}
