package main

import "testing"

func TestMain(t *testing.T) {
	t.Run("test for test", func(t *testing.T) {
		if 1 != 1 {
			t.Fail()
		}
	})
}
