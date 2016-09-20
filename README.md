# peakdetect

Converted from MATLAB script at http://billauer.co.il/peakdet.html

### Usage

Download and install the repo
```bash
go get github.com/mpiannucci/peakdetect
```

Call it on a slice
```go
import "github.com/mpiannucci/peakdetect"
...
v := [...]float64{0.0,1.0,2.0,1.0,0.0,-1.0,0.0,3.0,0.0}
mini, minv, maxi, maxv := peakdetect.PeakDetect(v[:], 1.0)
```

### Documentation

https://godoc.org/github.com/mpiannucci/peakdetect

### License

Converted by Matthew Iannucci

Original by Eli Billauer, 3.4.05 (Explicitly not copyrighted).
This function is released to the public domain; Any use is allowed.
