package log

import "testing"

func TestPrintf(t *testing.T) {
	With("a", "b").Printf("%s=%s", "hello", "world")
}
