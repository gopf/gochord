package main

import (
    "regexp"
    "strconv"
)

var chordPattern = regexp.MustCompile("\\{(?P<name>[\\w]+)[\\s]+frets[\\s]+(?P<frets>(?:[\\d\\s\\+x]+))\\}")
var rootPattern = regexp.MustCompile("\\+(?P<fretnum>[\\d]+)")
var fretPattern = regexp.MustCompile("(?P<fretnum>[\\d]+)")
var whitespacePattern = regexp.MustCompile("[\\s]+")

type stringInfo struct {
    mainFret int
    rootFret int
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
            strings = append(strings, stringInfo{-1, -1, nil})
            continue
        }

        rootMatch := matchRegexp(rootPattern, fret)
        if rootValue, ok := rootMatch["fretnum"]; ok {
            if fretnum, err := strconv.Atoi(rootValue); err == nil {
                strings = append(strings, stringInfo{fretnum, fretnum, nil})
                continue
            }
        }

        fretMatch := matchRegexp(fretPattern, fret)
        if fretValue, ok := fretMatch["fretnum"]; ok {
            if fretnum, err := strconv.Atoi(fretValue); err == nil {
                strings = append(strings, stringInfo{fretnum, -1, nil})
                continue
            }
        }

        // no match
    }
    return chordInfo{ fields["name"], strings }
}
