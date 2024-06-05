package main

import (
	"testing"
)

func TestSanitize(t *testing.T) {
	unsanitizedComment := "Uncrewed, yes. <a href=\"https:&#x2F;&#x2F;en.wikipedia.org&#x2F;wiki&#x2F;Boeing_Orbital_Flight_Test_2\" rel=\"nofollow\">https:&#x2F;&#x2F;en.wikipedia.org&#x2F;wiki&#x2F;Boeing_Orbital_Flight_Test_2</a>"
	expectedComment := "Uncrewed, yes. https://en.wikipedia.org/wiki/Boeing_Orbital_Flight_Test_2"
	sanitizedComment := sanitize(unsanitizedComment)

	// check if sanitized comment is equal to expected comment
	if sanitizedComment != expectedComment {
		t.Errorf("Expected %q, got %q", expectedComment, sanitizedComment)
	}
}
