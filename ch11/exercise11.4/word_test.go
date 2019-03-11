// Answer for Exercicse 11.1: Modify randomPalindrome to exercise
// IsPalindrome's handling of punctuation and spaces.
// WARNING: the conversion r -> upper -> lower doesn't preserve
// the value of r in some cases, e.g., µ Μ, ſ S, ı I

package word_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
	"unicode"
)

// IsPalindrome reports whether s reads the same forward and backward.
// Letter case is ignored, as are non-letters.
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomNoisyPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x200)) // random rune up to \u99
		runes[i] = r
		r1 := r
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			r = unicode.ToUpper(r)
			if unicode.ToLower(r) != r1 {
				fmt.Printf("cap? %c %c\n", r1, r)
			}
		}
		runes[n-1-i] = r
	}
	return "?" + string(runes) + "!"
}

func TestRandomNoisyPalindrome(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	n := 0
	for i := 0; i < 1000; i++ {
		p := randomNoisyPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsNoisyPalindrome(%q) = false", p)
			n++
		}
	}
	fmt.Fprintf(os.Stderr, "fail = %d\n", n)
}
