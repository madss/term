package term

import (
	"reflect"
	"testing"
)

func TestEscapeSequenceAdd(t *testing.T) {
	seq := escapeSequence{kind: 'm'}
	seq.Add(boldOn)
	seq.Add(redFg)
	if expected := (escapeSequence{kind: 'm', attrs: []attribute{boldOn, redFg}}); !reflect.DeepEqual(seq, expected) {
		t.Fatalf("seq = %v, want %v", seq.attrs, expected)
	}
}

func TestEscapeSequenceAdjust(t *testing.T) {
	for _, tc := range []struct {
		was, is  bool
		expected escapeSequence
	}{
		{false, false, escapeSequence{kind: 'm'}},
		{false, true, escapeSequence{kind: 'm', attrs: []attribute{boldOn}}},
		{true, false, escapeSequence{kind: 'm', attrs: []attribute{boldOff}}},
		{true, true, escapeSequence{kind: 'm'}},
	} {
		seq := escapeSequence{kind: 'm'}
		seq.Adjust(tc.was, tc.is, boldOn, boldOff)
		if !reflect.DeepEqual(seq, tc.expected) {
			t.Fatalf("seq.Adjust(%t, %t, ...) -> %v, want %v", tc.was, tc.is, seq, tc.expected)
		}
	}
}

func TestEscapeSequenceBytes(t *testing.T) {
	for _, tc := range []struct {
		seq      escapeSequence
		expected []byte
	}{
		{escapeSequence{kind: 'm'}, nil},
		{escapeSequence{kind: 'm', attrs: []attribute{boldOn}}, []byte("\033[1m")},
		{escapeSequence{kind: 'm', attrs: []attribute{boldOn, redFg}}, []byte("\033[1;31m")},
	} {
		if got := tc.seq.Bytes(); !reflect.DeepEqual(got, tc.expected) {
			t.Fatalf("%v.Bytes() = %v, want %v", tc.seq, got, tc.expected)
		}
	}
}
