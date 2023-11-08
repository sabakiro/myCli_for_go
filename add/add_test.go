package add_test

import (
	"slimple_cli/add"
	"testing"
)

func TestCreateFile(t *testing.T) {
	add.Instance.Run("ciallo")
}
