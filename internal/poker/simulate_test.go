package poker

import "testing"

func BenchmarkSimulate(b *testing.B) {
	for b.Loop() {
		Simulate(100_000)
	}
}

func BenchmarkSimulateConcurrent(b *testing.B) {
	for b.Loop() {
		SimulateConcurrent(100_000)
	}
}
