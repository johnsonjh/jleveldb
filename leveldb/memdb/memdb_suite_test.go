package memdb

import (
	"testing"

	"github.com/johnsonjh/jleveldb/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
