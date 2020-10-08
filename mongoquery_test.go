package mongoquery

import (
	"testing"
)

func TestQuery(t *testing.T) {
	a := Query("coll2", 60)
	aExpected := 3
	if a != aExpected {
		t.Errorf("expected %d, got %d", aExpected, a)
	}
}
