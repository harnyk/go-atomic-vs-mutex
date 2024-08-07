package avm_test

import (
	"avm"
	"testing"
)

func TestSynchronizeWithMutexes(t *testing.T) {
	if avm.SynchronizeWithMutexes(1e6) != 0 {
		t.Error("error")
	}
}

func TestSynchronizeWithAtomic(t *testing.T) {
	if avm.SynchronizeWithAtomic(1e6) != 0 {
		t.Error("error")
	}
}

func BenchmarkSynchronizeWithMutexes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		avm.SynchronizeWithMutexes(1e8)
	}
}

func BenchmarkSynchronizeWithAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		avm.SynchronizeWithAtomic(1e8)
	}
}
