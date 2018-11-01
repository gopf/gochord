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
    expected := chordInfo{ "C", []stringInfo{{0, -1, nil}, {0, -1, nil}, {0, -1, nil}, {3, -1, nil}}}
    validateChord(t, parseChord("{C frets 0 0 0 3}"), expected)
}

func Test_simple_ukulele_chord_2(t *testing.T) {
    expected := chordInfo{ "Em", []stringInfo{{0, -1, nil}, {4, -1, nil}, {3, -1, nil}, {2, -1, nil}}}
    validateChord(t, parseChord("{Em frets 0 4 3 2}"), expected)
}

func Test_simple_guitar_chord(t *testing.T) {
    expected := chordInfo{ "Em", []stringInfo{{0, -1, nil}, {2, -1, nil}, {2, -1, nil}, {0, -1, nil}, {0, -1, nil}, {0, -1, nil}}}
    validateChord(t, parseChord("{Em frets 0 2 2 0 0 0}"), expected)
}

func Test_root_note(t *testing.T) {
    expected := chordInfo{ "Em", []stringInfo{{0, -1, nil}, {4, 4, nil}, {3, -1, nil}, {2, -1, nil}}}
    validateChord(t, parseChord("{Em frets 0 +4 3 2}"), expected)
}

func Test_muted_string(t *testing.T) {
    expected := chordInfo{ "Ab", []stringInfo{{-1, -1, nil}, {3, -1, nil}, {4, -1, nil}, {3, -1, nil}}}
    validateChord(t, parseChord("{Ab frets x 3 4 3}"), expected)
}
