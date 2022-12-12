package http

import (
	"fmt"
	netHeader "net/http"
	"regexp"
	"strconv"
	"strings"
)

type AcceptSpec struct {
	Value string
	Q     float64
}

func GetAccepts(header netHeader.Header) ([]AcceptSpec, error) {
	var acceptSpecs []AcceptSpec
	headerAcceptItems := header["Accept"]
	if len(headerAcceptItems) == 0 {
		return acceptSpecs, nil
	}
	commaRegex := regexp.MustCompile(",")
	semicolonRegex := regexp.MustCompile(";")
	assignRegex := regexp.MustCompile("=")
	spaceBetweenCharsRegex := regexp.MustCompile(`.\s.`)
	qRegex := regexp.MustCompile(`q=[\d+]\.[\d]*`)

	for _, accept := range headerAcceptItems {
		acceptItems := commaRegex.Split(accept, -1)
		for _, item := range acceptItems {
			itemElements := semicolonRegex.Split(item, -1)
			if len(itemElements) == 0 {
				continue
			}
			spec := AcceptSpec{Q: 1.0}
			for _, element := range itemElements {
				element := strings.TrimSpace(element)
				if qRegex.MatchString(element) {
					values := assignRegex.Split(element, -1)
					if len(values) != 2 {
						continue
					}
					q, err := strconv.ParseFloat(values[1], 64)
					if err == nil {
						spec.Q = q
					} else {
						return acceptSpecs, err
					}
				} else if spaceBetweenCharsRegex.MatchString(element) {
					return acceptSpecs, fmt.Errorf("invalid element in header Accept: '%s'", element)
				} else {
					spec.Value = strings.TrimSpace(element)
				}
			}
			if len(spec.Value) > 0 {
				acceptSpecs = append(acceptSpecs, spec)
			}
		}
	}
	return acceptSpecs, nil
}

// sortAccept Holds accepted response types
type sortAccept struct {
	specs []AcceptSpec
	prefs []string
}

func (s sortAccept) Len() int {
	return len(s.specs)
}

// We want to sort by descending order of suitability: higher quality
// to lower quality, and preferred to less preferred.
func (s sortAccept) Less(i, j int) bool {
	switch {
	case s.specs[i].Q == s.specs[j].Q:
		return indexOf(s.prefs, s.specs[i].Value) < indexOf(s.prefs, s.specs[j].Value)
	default:
		return s.specs[i].Q > s.specs[j].Q
	}
}

func (s sortAccept) Swap(i, j int) {
	s.specs[i], s.specs[j] = s.specs[j], s.specs[i]
}

// This exists so we can search short slices of strings without
// requiring them to be sorted. Returning the len value if not found
// is so that it can be used directly in a comparison when sorting (a
// `-1` would mean "not found" was sorted before found entries).
func indexOf(ss []string, search string) int {
	for i, s := range ss {
		if s == search {
			return i
		}
	}
	return len(ss)
}
