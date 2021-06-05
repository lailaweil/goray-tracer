package utils

import "testing"

func TestFloatEquals(t *testing.T) {

	cases := []struct{
		name string
		numberA float64
		numberB float64
		areEqual bool
	}{
		{
			name:     "Equal/1.8-1.8",
			numberA:  1.8,
			numberB:  1.8,
			areEqual: true,
		},
		{
			name:     "Equal/0.000000002-0",
			numberA:  0.000000002,
			numberB:  0,
			areEqual: true,
		},
		{
			name:     "Equal/0.99998-0.99999",
			numberA:  0.99998,
			numberB:  0.99999,
			areEqual: true,
		},
		{
			name:     "Equal/0.99999-0.99998",
			numberA:  0.99999,
			numberB:  0.99998,
			areEqual: true,
		},
		{
			name:     "Not_Equal/0.99990-0.99970",
			numberA:  0.99990,
			numberB:  0.99970,
			areEqual: false,
		},
		{
			name:     "Not_Equal/0.99970-0.99990",
			numberB:  0.99970,
			numberA:  0.99990,
			areEqual: false,
		},
		{
			name:     "Not_Equal/1.7-1.8",
			numberB:  1.7,
			numberA:  1.8,
			areEqual: false,
		},
	}

	for _, c := range cases {
		result := FloatEquals(c.numberA, c.numberB)

		if result != c.areEqual {
			t.Errorf("%s - failed comparing float numbers: expected (%v) but got (%v)", c.name, c.areEqual, result)
		}
	}
}
