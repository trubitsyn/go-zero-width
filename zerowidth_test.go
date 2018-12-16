// Copyright 2018 Nikola Trubitsyn. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package zerowidth

import "testing"

func TestHasZeroWidthCharacters(t *testing.T) {
	samples := []string{"\u200B", "\u200C", "\u200D", "\uFEFF"}
	for _, s := range samples {
		expected := true
		actual := HasZeroWidthCharacters(s)
		if actual != expected {
			t.Errorf("String was incorrect, got: %t, want: %t", actual, expected)
		}
	}
}

func TestHasNoZeroWidthCharacters(t *testing.T) {
	expected := false
	actual := HasZeroWidthCharacters("HERE>string<HERE")
	if actual != expected {
		t.Errorf("String was incorrect, got: %t, want: %t", actual, expected)
	}
}

func TestRemoveZeroWidthCharacters(t *testing.T) {
	expected := "HERE><HERE"
	actual := RemoveZeroWidthCharacters("HERE>\u200B\u200C\u200D\uFEFF<HERE")
	if actual != expected {
		t.Errorf("String was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestRemoveZeroWidthSpace(t *testing.T) {
	expected := "HERE><HERE"
	actual := RemoveZeroWidthSpace("HERE>\u200B<HERE")
	if actual != expected {
		t.Errorf("String was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestRemoveZeroWidthNoBreakSpace(t *testing.T) {
	expected := "HERE><HERE"
	actual := RemoveZeroWidthNoBreakSpace("HERE>\uFEFF<HERE")
	if actual != expected {
		t.Errorf("String was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestRemoveZeroWidthJoiner(t *testing.T) {
	expected := "HERE><HERE"
	actual := RemoveZeroWidthJoiner("HERE>\u200D<HERE")
	if actual != expected {
		t.Errorf("String was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestRemoveZeroWidthNonJoiner(t *testing.T) {
	expected := "HERE><HERE"
	actual := RemoveZeroWidthNonJoiner("HERE>\u200C<HERE")
	if actual != expected {
		t.Errorf("String was incorrect, got: %s, want: %s.", actual, expected)
	}
}
