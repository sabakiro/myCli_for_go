package watch_test

import (
	"slimple_cli/watch"
	"testing"
)

func TestWatch(t *testing.T) {
	watch.Instance.Run()
}
// func TestAddFolder(t *testing.T) {
// 	watch.Instance.AddFolder(".")
// }
