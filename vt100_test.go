package term

import (
	"reflect"
	"testing"
)

func TestEscapeSequenceAdd(t *testing.T) {
	var attrs escapeSequence
	attrs.Add(boldOn)
	attrs.Add(redFg)
	if expected := (escapeSequence{boldOn, redFg}); !reflect.DeepEqual(attrs, expected) {
		t.Fatalf("attrs = %v, want %v", attrs, expected)
	}
}

func TestEscapeSequenceAdjust(t *testing.T) {
	for _, tc := range []struct {
		was, is  bool
		expected escapeSequence
	}{
		{false, false, nil},
		{false, true, escapeSequence{boldOn}},
		{true, false, escapeSequence{boldOff}},
		{true, true, nil},
	} {
		var attrs escapeSequence
		attrs.Adjust(tc.was, tc.is, boldOn, boldOff)
		if !reflect.DeepEqual(attrs, tc.expected) {
			t.Fatalf("attrs.Adjust(%t, %t, ...) -> %v, want %v", tc.was, tc.is, attrs, tc.expected)
		}
	}
}

func TestEscapeSequenceBytes(t *testing.T) {
	for _, tc := range []struct {
		attrs    escapeSequence
		expected []byte
	}{
		{nil, nil},
		{escapeSequence{boldOn}, []byte("\033[1m")},
		{escapeSequence{boldOn, redFg}, []byte("\033[1;31m")},
	} {
		if got := tc.attrs.Bytes(); !reflect.DeepEqual(got, tc.expected) {
			t.Fatalf("%v.Bytes() = %v, want %v", tc.attrs, got, tc.expected)
		}
	}
}
