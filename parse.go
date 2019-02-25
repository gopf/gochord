package main

import (
    "regexp"
    "strconv"
)

var chordPattern = regexp.MustCompile("\\{(?P<name>[^\\s]+)[\\s]+frets[\\s]+(?P<frets>(?:[\\d\\s\\+x/]+))\\}")
var rootPattern = regexp.MustCompile("\\+(?P<fretnum>[\\d]+)")
var fretPattern = regexp.MustCompile("(?P<fretnum>[\\d]+)")
var whitespacePattern = regexp.MustCompile("[\\s]+")
var fretSeparatorPattern = regexp.MustCompile("/")

type stringInfo struct {
    rootFret int
    frets []int
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
    for _, string := range whitespacePattern.Split(fields["frets"], -1) {
        var rootFret int = 999
        var frets []int
        for _, fret := range fretSeparatorPattern.Split(string, -1) {
            if fret == "x" {
                frets = append(frets, -1)
                continue
            }

            rootMatch := matchRegexp(rootPattern, fret)
            if rootValue, ok := rootMatch["fretnum"]; ok {
                if fretnum, err := strconv.Atoi(rootValue); err == nil {
                    frets = append(frets, fretnum)
                	rootFret = fretnum
                    continue
                }
            }

            fretMatch := matchRegexp(fretPattern, fret)
            if fretValue, ok := fretMatch["fretnum"]; ok {
                if fretnum, err := strconv.Atoi(fretValue); err == nil {
	                frets = append(frets, fretnum)
                    continue
                }
            }

            // no match
        }
	    strings = append(strings, stringInfo{rootFret, frets})
    }
    return chordInfo{ fields["name"], strings }
}
