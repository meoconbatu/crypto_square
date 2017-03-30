package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

const testVersion = 2

func Encode(pt string) string {
	pt = Normalize(pt)
	lenPt := len(pt)
	if lenPt == 0 {
		return ""
	}
	_, c := GetSizeOfRectangle(lenPt)
	lenCt := lenPt + c - 1
	ct := make([]byte, lenCt)
	iCt := 0
	for i := 0; i < c; i++ {
		for j := i; j < lenPt; j = j + c {
			ct[iCt] = pt[j]
			iCt = iCt + 1
		}
		if i+1 < c {
			ct[iCt] = byte(' ')
			iCt = iCt + 1
		}
	}
	return string(ct)
}
func GetSizeOfRectangle(lenPt int) (int, int) {
	var c float64
	r := math.Sqrt(float64(lenPt))
	floorR := math.Floor(math.Sqrt(float64(lenPt)))
	if r >= floorR+0.5 {
		r = floorR + 1
		c = r
	} else if r == floorR {
		r = floorR
		c = r
	} else {
		r = floorR
		c = r + 1
	}
	return int(r), int(c)
}
func Normalize(pt string) string {
	pt = strings.ToLower(pt)
	search := regexp.MustCompile("[^a-z0-9]")
	pt = search.ReplaceAllStringFunc(pt, func(s string) string {
		return string("")
	})
	return pt
}
