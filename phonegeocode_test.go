package phonegeocode

import (
	"testing"
)

func TestThingsThatAreFound(t *testing.T) {
	cases := []struct {
		areaCode, number, country string
	}{
		{"44", "+4479991113332", "GB"},
		{"1807", "+1807142342342", "CA"},
		{"1", "+1781234555552", "US"},
		{"34","+3462233455552", "ES"},
		{"353", "+3538523523455", "IE"},
		{"1", "+19425242343333", "US"},
		{"1", "+1786822211132", "US"},
		{"34", "+34634343434443", "ES"},
		{"1905", "+190572354235235", "CA"},
	}

	p := New()

	for _, tc := range cases {
		areaCode, cc, err := p.Country(tc.number)
		if err != nil {
			t.Errorf("Not expecting number '%s' to yield an error; got %v", tc.number, err)
		}
		if cc != tc.country {
			t.Errorf("Number '%s' did not match expected CC '%s'; got '%s'", tc.number, tc.country, cc)
		}
		if areaCode != tc.areaCode {
			t.Errorf("Number '%s' did not match expected areaCode '%s'; got '%s'", tc.areaCode, tc.areaCode, cc)
		}
	}
}

func TestThingsThatAreNotFound(t *testing.T) {
	cases := []struct {
		number string
	}{
		{"+9915155235325"},
		{"+898235265"},
	}

	p := New()

	for _, tc := range cases {
		_, _, err := p.Country(tc.number)
		if err == nil {
			t.Errorf("Expecting number '%s' to yield an error", tc.number)
		}
	}
}
