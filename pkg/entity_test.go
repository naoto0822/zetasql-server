package pkg

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	t.Run("IsValid is true", func(t *testing.T) {
		r := Request{
			SQL: "SELECT 1",
		}

		expect := true
		actual := r.IsValid()
		if expect != actual {
			t.Errorf("failed to TestIsValid, expect: %+v ,actual: %+v, input: %+v", expect, actual, r)
		}
	})

	t.Run("IsValid is false", func(t *testing.T) {
		r := Request{
			SQL: "",
		}

		expect := false
		actual := r.IsValid()
		if expect != actual {
			t.Errorf("failed to TestIsValid, expect: %+v ,actual: %+v, input: %+v", expect, actual, r)
		}
	})
}
