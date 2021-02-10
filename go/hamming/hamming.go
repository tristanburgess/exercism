/*
Package hamming exports a single function Distance which
calculates the Hamming distance between two DNA sequences.
*/
package hamming

import "fmt"

// Distance takes in two equal-lengthed input sequences of DNA letters as strings
// and outputs the Hamming distance between the two sequences. The Hamming distance is
// the number of positions at which the two sequences have different DNA letters.
func Distance(a, b string) (int, error) {
	aLen, bLen := len(a), len(b)
	if aLen != bLen {
		return 0, fmt.Errorf("Expected string arguments a and b to have the same lengths."+
			"Actual [a - val: %q - len: %d], [b - val: %q - len: %d", a, aLen, b, bLen)
	}

	hammingDist := 0
	for idx, letter := range a {
		if letter != rune(b[idx]) {
			hammingDist++
		}
	}

	return hammingDist, nil
}
