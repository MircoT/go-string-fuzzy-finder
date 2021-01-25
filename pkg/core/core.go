package core

import (
	"fmt"
	"math"
	"sort"

	"github.com/MircoT/go-string-fuzzy-finder/pkg/alg"
)

const (
	defaultThreshold  = 0.8
	defaultSimilarRes = 6
)

var (
	errNoValidThreshold          = fmt.Errorf("not a valid threshold, it must be 0.0 < x < 1.0")
	errNoValidSimilarResultValue = fmt.Errorf("not a valid value, it must be > 0")
)

type SimpleFinder struct {
	threshold float64
	alg       func(target string, other string) int
	n         int
}

type bufElm struct {
	str  string
	diff int
}

type byAlg struct {
	target string
	buffer []bufElm
	finder SimpleFinder
}

func (s byAlg) Len() int           { return len(s.buffer) }
func (s byAlg) Swap(i, j int)      { s.buffer[i], s.buffer[j] = s.buffer[j], s.buffer[i] }
func (s byAlg) Less(i, j int) bool { return s.buffer[i].diff < s.buffer[j].diff }

func (f *SimpleFinder) Init(args ...interface{}) {
	threshold := defaultThreshold
	alg := alg.Levenshtein

	switch len(args) {
	case 2:
		alg = args[1].(func(target string, other string) int)

		fallthrough
	case 1:
		threshold = args[0].(float64)
	}

	f.threshold = threshold
	f.alg = alg
	f.n = defaultSimilarRes
}

func (f SimpleFinder) ratio(target string, value string, distance int) float64 {
	sumLens := float64(len(target) + len(value))

	return (sumLens - float64(distance)) / sumLens
}

// BestMatch returns the string most similar to target.
func (f SimpleFinder) BestMatch(target string, strings []string) (result string, err error) {
	bestMatch := math.MaxInt64

	for _, value := range strings {
		curString := value

		curDiff := f.alg(target, curString)

		if curDiff < bestMatch {
			bestMatch = curDiff
			result = curString
		}

		// log.Println(bestMatch, curString, curDiff, result, f.ratio(target, curString, curDiff))
	}

	return result, err
}

// Similars order the strings by similarity with target and returns the top n.
func (f SimpleFinder) Similars(target string, strings []string) (results []string, err error) {
	results = make([]string, 0)

	buffer := make([]bufElm, len(strings))

	for idx, value := range strings {
		curDiff := f.alg(target, value)
		buffer[idx].diff = curDiff
		buffer[idx].str = value
	}

	sort.Sort(byAlg{
		target: target,
		buffer: buffer,
		finder: f,
	})

	for _, elm := range buffer {
		curRatio := f.ratio(target, elm.str, elm.diff)

		if curRatio >= f.threshold {
			results = append(results, elm.str)
		}

		if len(results) == f.n {
			break
		}
	}

	return results, err
}

func (f *SimpleFinder) SetSimilarResultNum(n int) error {
	if n <= 0 {
		return errNoValidSimilarResultValue
	}

	f.n = n

	return nil
}

func (f *SimpleFinder) SetAlg(alg func(target string, other string) int) error {
	f.alg = alg

	return nil
}

func (f *SimpleFinder) SetMinThreshold(val float64) error {
	if val >= 1.0 || val <= 0.0 {
		return errNoValidThreshold
	}

	f.threshold = val

	return nil
}
