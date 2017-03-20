package glog

import (
	"os"
	"testing"
)

func TestPlacebo(t *testing.T) {
	t.Parallel()
	if !true {
		t.FailNow()
	}
}

func TestLevel(t *testing.T) {
	os.Unsetenv("LOG_LEVEL")
	if level() != Error {
		t.FailNow()
	}
	os.Setenv("LOG_LEVEL", "DeBuG")
	if level() != Debug {
		t.FailNow()
	}
	os.Setenv("LOG_LEVEL", "info")
	if level() != Info {
		t.FailNow()
	}
	os.Setenv("LOG_LEVEL", "notice")
	if level() != Notice {
		t.FailNow()
	}
	os.Setenv("LOG_LEVEL", "WarNing")
	if level() != Warning {
		t.FailNow()
	}
	os.Setenv("LOG_LEVEL", "ErrOR")
	if level() != Error {
		t.FailNow()
	}
	os.Setenv("LOG_LEVEL", "critical")
	if level() != Critical {
		t.FailNow()
	}
	os.Setenv("LOG_LEVEL", "Alert")
	if level() != Alert {
		t.FailNow()
	}
	os.Setenv("LOG_LEVEL", "Emergency")
	if level() != Emergency {
		t.FailNow()
	}
	os.Setenv("LOG_LEVEL", "BadInput")
	if level() != Error {
		t.FailNow()
	}
}
