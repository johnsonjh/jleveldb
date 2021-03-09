package iterator_test

import (
	"testing"

	"github.com/johnsonjh/jleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
