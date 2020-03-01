package main

import (
	"testing"
)

func Benchmark_sequencedArchivator(b *testing.B) {
	for i:=0; i< b.N; i++{
		SeqArchivator([]string{
			"1.txt",
			"2.txt",
			"3.txt",
			"4.txt",
			"5.txt",
			"6.txt",
			"7.txt",
			"8.txt",
			"9.txt",
			"10.txt",
		})
	}
}

func Benchmark_concurrentArchivator(b *testing.B) {
	for i:=0; i< b.N; i++{
		ConArchivator([]string{
			"1.txt",
			"2.txt",
			"3.txt",
			"4.txt",
			"5.txt",
			"6.txt",
			"7.txt",
			"8.txt",
			"9.txt",
			"10.txt",
		})
	}
}