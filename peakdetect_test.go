package peakdetect

import "testing"

func TestPeakDetect(t *testing.T) {
	v := [...]float64{0.0, 1.0, 2.0, 1.0, 0.0, -1.0, 0.0, 3.0, 0.0}
	mini, minv, maxi, maxv := PeakDetect(v[:], 1.0)

	if len(mini) != 1 {
		t.Fail()
	}

	if len(maxi) != 2 {
		t.Fail()
	}

	if minv[0] != -1 {
		t.Fail()
	}

	if maxv[0] != 2 {
		t.Fail()
	}
}
