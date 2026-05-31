package processor

import "testing"

func BenchmarkProcessImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessImage(1)
	}
}
