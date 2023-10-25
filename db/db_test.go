package db

import (
	"testing"

	"github.com/warpgr/tdba_common/utils"
)

func TestXxx(t *testing.T) {
	t.Cleanup(func() {
		utils.CleanupLogFiles()
	})
}
