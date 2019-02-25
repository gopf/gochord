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
    expected := chordInfo{ "C", []stringInfo{{999, []int{0}}, {999, []int{0}}, {999, []int{0}}, {999, []int{3}}}}
    validateChord(t, parseChord("{C frets 0 0 0 3}"), expected)
}

func Test_simple_ukulele_chord_2(t *testing.T) {
    expected := chordInfo{ "Em", []stringInfo{{999, []int{0}}, {999, []int{4}}, {999, []int{3}}, {999, []int{2}}}}
    validateChord(t, parseChord("{Em frets 0 4 3 2}"), expected)
}

func Test_simple_guitar_chord(t *testing.T) {
    expected := chordInfo{ "Em", []stringInfo{{999, []int{0}}, {999, []int{2}}, {999, []int{2}}, {999, []int{0}}, {999, []int{0}}, {999, []int{0}}}}
    validateChord(t, parseChord("{Em frets 0 2 2 0 0 0}"), expected)
}

func Test_root_note(t *testing.T) {
    expected := chordInfo{ "Em", []stringInfo{{999, []int{0}}, {4, []int{4}}, {999, []int{3}}, {999, []int{2}}}}
    validateChord(t, parseChord("{Em frets 0 +4 3 2}"), expected)
}

func Test_muted_string(t *testing.T) {
    expected := chordInfo{ "Ab", []stringInfo{{999, []int{-1}}, {999, []int{3}}, {999, []int{4}}, {999, []int{3}}}}
    validateChord(t, parseChord("{Ab frets x 3 4 3}"), expected)
}

func Test_multiple_frets(t *testing.T) {
    expected := chordInfo{ "major-1", []stringInfo{{999, []int{-1}}, {1, []int{1,3}}, {999, []int{1,2,4}}, {4, []int{1,3,4}}}}
    validateChord(t, parseChord("{major-1 frets x +1/3 1/2/4 1/3/+4}"), expected)
}
