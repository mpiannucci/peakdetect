package peakdetect

import "math"

// Finds the local maxima and minima ("peaks") in the vector
// Returns four arrays -> min indices, min values, max indices, max values
//
// A point is considered a maximum peak if it has the maximal
// value, and was preceded (to the left) by a value lower by
// DELTA.
func PeakDetect(vals []float64, delta float64) ([]int, []float64, []int, []float64) {
	minx := make([]int, 0, 0)
	maxx := make([]int, 0, 0)
	minvals := make([]float64, 0, 0)
	maxvals := make([]float64, 0, 0)

	xvals := make([]int, len(vals), len(vals))
	for i := 0; i < len(vals); i++ {
		xvals[i] = i
	}

	if delta < 0 {
		return minx, minvals, maxx, maxvals
	}

	mn := math.Inf(1)
	mx := math.Inf(-1)
	mnpos := int(math.NaN())
	mxpos := int(math.NaN())

	lookForMax := true

	for i, val := range vals {
		if val > mx {
			mx = val
			mxpos = i
		}
		if val < mn {
			mn = val
			mnpos = i
		}

		if lookForMax {
			if val < mx-delta {
				maxx = append(maxx, mxpos)
				maxvals = append(maxvals, mx)

				mn = val
				mnpos = i
				lookForMax = false
			}
		} else {
			if val > mn+delta {
				minx = append(minx, mnpos)
				minvals = append(minvals, mn)

				mx = val
				mxpos = i
				lookForMax = true
			}
		}
	}

	return minx, minvals, maxx, maxvals
}
