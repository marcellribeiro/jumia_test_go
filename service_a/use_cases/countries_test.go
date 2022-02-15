package use_cases

import (
	"log"
	"testing"
)

func gotWantCountry(text string, want string, t *testing.T) {
	got := GetCountryByPhone(text)
	log.Printf("Text: %v / Want: %v / Got: %v\n", text, want, got)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestCountriesWithSuccess(t *testing.T) {
	gotWantCountry("237 123 123 4", "Cameroon", t)
	gotWantCountry("55 98765 5566", "", t)
	gotWantCountry("+25167899883", "Ethiopia", t)
	gotWantCountry("(212)666777895", "Morocco", t)
	gotWantCountry("(258)66677 7895", "Mozambique", t)
	gotWantCountry("256-234556677", "Uganda", t)
}

func TestCountriesWithError(t *testing.T) {
	gotWantCountry("5587609808", "Brasil", t)
}
