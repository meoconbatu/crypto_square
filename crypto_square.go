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
	r, c := GetSizeOfRectangle(lenPt)
	lenCt := lenPt + int(c) - 1
	ct := make([]rune, lenCt)

	aa := int(r)*int(c) - lenPt
	aaa := int(c) - aa
	var xx, xy, j int
	for i, a := range pt {
		xx = int(math.Mod(float64(i), c))
		xy = int(math.Floor(float64(i) / c))
		j = xx*(int(r)+1) + xy
		if xx >= aaa && aa > 1 {
			j = j - (xx - aaa)
		}
		ct[j] = a
		if xy == 0 && j > 0 {
			ct[j-1] = rune(' ')
		}
	}
	return string(ct)
}
func GetSizeOfRectangle(lenPt int) (float64, float64) {
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
	return r, c
}
func Normalize(pt string) string {
	pt = strings.ToLower(pt)
	search := regexp.MustCompile("[^a-z0-9]")
	pt = search.ReplaceAllStringFunc(pt, func(s string) string {
		return string("")
	})
	return pt
}
