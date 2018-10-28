package main

import (
    "reflect"
    "testing"
)

func validateChord(t *testing.T, chord chordInfo, expectedChord chordInfo) {
    if !reflect.DeepEqual(chord, expectedChord) {
       t.Errorf("Chord was incorrect, got: %v, want: %v.", chord, expectedChord)
    }
}

func Test_simple_ukulele_chord_1(t *testing.T) {
    expected := chordInfo{ "C", []stringInfo{{0, nil}, {0, nil}, {0, nil}, {3, nil}}}
    validateChord(t, parseChord("{C frets 0 0 0 3}"), expected)
}

func Test_simple_ukulele_chord_2(t *testing.T) {
    expected := chordInfo{ "Em", []stringInfo{{0, nil}, {4, nil}, {3, nil}, {2, nil}}}
    validateChord(t, parseChord("{Em frets 0 4 3 2}"), expected)
}
