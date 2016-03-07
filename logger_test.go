package log

import (
	"io/ioutil"
	"testing"
)

func init() {
	stderr = ioutil.Discard
}

func TestLogger(t *testing.T) {
	l := Logger{}

	args := []interface{}{"one", struct{ Two int }{Two: 2}, 3}
	for s := range severities {
		for _, c := range []bool{false, true} {
			for a := 0; a < len(args); a++ {
				color = c
				severity = Severity(s)
				l.Log("message", Severity(s), args[a:]...)
				l.Emergency("message", args[a:]...)
				l.Alert("message", args[a:]...)
				l.Critical("message", args[a:]...)
				l.Error("message", args[a:]...)
				l.Warning("message", args[a:]...)
				l.Notice("message", args[a:]...)
				l.Info("message", args[a:]...)
				l.Debug("message", args[a:]...)
			}
		}
	}
}

func BenchmarkLogger(b *testing.B) {
	l := Logger{}
	args := []interface{}{"one", struct{ Two int }{Two: 2}, 3}

	for n := 0; n < b.N; n++ {
		l.Log("message", Error, args)
	}
}
