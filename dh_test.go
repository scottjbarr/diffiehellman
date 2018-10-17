package diffiehellman

import (
	"reflect"
	"testing"
)

func TestDH(t *testing.T) {
	alice, bob := verify()

	if !reflect.DeepEqual(*alice, *bob) {
		t.Errorf("%v and %v should be equal", *alice, *bob)
	}
}
