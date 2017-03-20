// +build !solaris,!linux,!darwin,!dragonfly,!freebsd

package glog

import (
	"testing"
)

func TestUseColor(t *testing.T) {
	t.Parallel()

	var b bool
	b = useColor()
	if b != false {
		t.FailNow()
	}
}
