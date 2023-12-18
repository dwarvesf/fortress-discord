package cryptoutils

import "testing"

func TestStringBigIntToStringCurrency(t *testing.T) {
	type args struct {
		givenValue    string
		givenDecimals int
		expResult     string
	}

	tcs := map[string]args{
		"Given empty string": {
			givenValue:    "",
			givenDecimals: 18,
			expResult:     "0",
		},
		"Given 0": {
			givenValue:    "0",
			givenDecimals: 18,
			expResult:     "0",
		},
		"Given 1": {
			givenValue:    "1",
			givenDecimals: 18,
			expResult:     "0",
		},
		"Given 100000000000000000000000000000000": {
			givenValue:    "100000000000000000000000000000000",
			givenDecimals: 18,
			expResult:     "100,000,000,000,000",
		},
		"Given decimals 0": {
			givenValue:    "100000000000000000000000000000000",
			givenDecimals: 0,
			expResult:     "100,000,000,000,000,000,000,000,000,000,000",
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			result := StringBigIntToStringCurrency(tc.givenValue, tc.givenDecimals)
			if result != tc.expResult {
				t.Errorf("Expected %s, got %s", tc.expResult, result)
			}
		})
	}
}
