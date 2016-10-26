package main

import (
	"runtime"
	"testing"
	"time"
)

var sinkhole = ""

func BenchmarkWriteMapSize(b *testing.B) {
	println("Map Size", len(in))
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Second)
	}
}

func BenchmarkStringMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkhole = getFromMap(in[i%len(in)])
	}
}

func BenchmarkStringSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkhole = getFromSwitch(in[i%len(in)])
	}
}

func BenchmarkGC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.GC()
	}
}
