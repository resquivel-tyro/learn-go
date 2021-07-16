package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// when benchmark code is executed, it runs b.N times and measures how long it takes
// the framework will determine what is a "good" value to get decent results
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
// Sample benchmark output in console:
//goos: darwin
//goarch: amd64
//pkg: hello/iteration
//cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
//BenchmarkRepeat
//BenchmarkRepeat-12    	 9531441	       118.0 ns/op
//PASS