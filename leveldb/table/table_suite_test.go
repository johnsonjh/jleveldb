package table

import (
	"testing"

	"github.com/johnsonjh/jleveldb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
