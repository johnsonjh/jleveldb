//go:build goleveldbtests
// +build goleveldbtests

package leveldb

import (
	"testing"

	"github.com/johnsonjh/jleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
