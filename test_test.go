package test_test

import (
	"testing"

	test "github.com/advanderveer/go-test"
)

func TestEquals(t *testing.T) {
	test.Equals(t, 1, 1)
	test.OkEquals(t, 1)(1, nil)

}
