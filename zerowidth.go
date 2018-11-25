// Copyright 2018 Nikola Trubitsyn. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package zerowidth implements functions to remove zero-width characters from strings.
package zerowidth

import "strings"

const (
	// ZWSP represents zero-width space.
	ZWSP = '\u200B'

	// ZWNBSP represents zero-width no-break space.
	ZWNBSP = '\uFEFF'

	// ZWJ represents zero-width joiner.
	ZWJ = '\u200D'

	// ZWNJ represents zero-width non-joiner.
	ZWNJ = '\u200C'

	empty = ""
)

var replacer = strings.NewReplacer(string(ZWSP), empty,
	string(ZWNBSP), empty,
	string(ZWJ), empty,
	string(ZWNJ), empty)

// HasZeroWidthCharacters reports whether string s contains zero-width characters.
func HasZeroWidthCharacters(s string) bool {
	return strings.ContainsRune(s, ZWSP) ||
		strings.ContainsRune(s, ZWNBSP) ||
		strings.ContainsRune(s, ZWJ) ||
		strings.ContainsRune(s, ZWNJ)
}

// RemoveZeroWidthCharacters removes all zero-width characters from string s.
func RemoveZeroWidthCharacters(s string) string {
	return replacer.Replace(s)
}

// RemoveZeroWidthSpace removes zero-width space characters from string s.
func RemoveZeroWidthSpace(s string) string {
	return strings.Replace(s, string(ZWSP), empty, -1)
}

// RemoveZeroWidthNoBreakSpace removes zero-width no-break space characters from string s.
func RemoveZeroWidthNoBreakSpace(s string) string {
	return strings.Replace(s, string(ZWNBSP), empty, -1)
}

// RemoveZeroWidthJoiner removes zero-width joiner characters from string s.
func RemoveZeroWidthJoiner(s string) string {
	return strings.Replace(s, string(ZWJ), empty, -1)
}

// RemoveZeroWidthNonJoiner removes zero-width non-joiner characters from string s.
func RemoveZeroWidthNonJoiner(s string) string {
	return strings.Replace(s, string(ZWNJ), empty, -1)
}
