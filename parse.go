package main

import (
    "regexp"
    "strconv"
)

var chordPattern = regexp.MustCompile("\\{(?P<name>[\\w]+)[\\s]+frets[\\s]+(?P<frets>(?:[x\\d\\s]+))\\}")
var whitespacePattern = regexp.MustCompile("[\\s]+")

type stringInfo struct {
    mainFret int
    extraFrets []int
}

type chordInfo struct {
    name string
    strings []stringInfo
}

func matchRegexp(re *regexp.Regexp, s string) map[string]string {
    match := re.FindStringSubmatch(s)
    names := re.SubexpNames()
    if match == nil {
        return nil
    } else {
	    match, names = match[1:], names[1:]
	    result := make(map[string]string, len(match))
	    for i, _ := range names {
		    result[names[i]] = match[i]
	    }
	    return result
    }
}


func parseChord(s string) chordInfo {
    fields := matchRegexp(chordPattern, s)
    var strings []stringInfo
    for _, fret := range whitespacePattern.Split(fields["frets"], -1) {
        if fret == "x" {
            strings = append(strings, stringInfo{-1, nil})
        } else if fretnum, err := strconv.Atoi(fret); err == nil {
            strings = append(strings, stringInfo{fretnum, nil})
        }
    }
    return chordInfo{ fields["name"], strings }
}
