// Copyright 2018 Nikola Trubitsyn. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package zw

import "testing"

func TestHasZeroWidthCharacters(t *testing.T) {
	expected := true
	actual := HasZeroWidthCharacters("HERE>\u200C<HERE")
	if actual != expected {
		t.Errorf("String was incorrect, got: %t, want: %t", actual, expected)
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
	actual := RemoveZeroWidthCharacters("HERE><HERE")
	if actual != expected {
		t.Errorf("String was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestRemoveZeroWidthSpace(t *testing.T) {
	expected := "HERE><HERE"
	actual := RemoveZeroWidthCharacters("HERE>\u200b<HERE")
	if actual != expected {
		t.Errorf("String was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestRemoveZeroWidthNoBreakSpace(t *testing.T) {
	expected := "HERE><HERE"
	actual := RemoveZeroWidthCharacters("HERE>\uFEFF<HERE")
	if actual != expected {
		t.Errorf("String was incorrect, got: %s, want: %s.", actual, expected)
	}
}

func TestRemoveZeroWidthJoiner(t *testing.T) {
	expected := "HERE><HERE"
	actual := RemoveZeroWidthCharacters("HERE>\u200D<HERE")
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
