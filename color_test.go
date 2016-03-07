// +build !windows,!plan9,!netbsd,!openbsd

package log

import (
	"testing"
)

func TestUseColor(t *testing.T) {
	t.Parallel()

	var b bool
	b = useColor()
	if b != true && b != false {
		t.FailNow()
	}
}
