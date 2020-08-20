package examples

import (
	"fmt"
	"runtime"
	"time"
)

func MemStatsDemo() {
	var mem runtime.MemStats

	for i:=0; i<10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed")
		}
	}
	PrintStats(mem)

	for i:=0 ; i< 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed at ", i)
		}
		time.Sleep(5 * time.Second)
	}
	PrintStats(mem)
}

func PrintStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc", mem.Alloc)
	fmt.Println("mem.TotalAlloc", mem.TotalAlloc)
	fmt.Println("mem.heapAlloc", mem.HeapAlloc)
	fmt.Println("mem.NumGC", mem.NumGC)
	fmt.Println("-----------------------------")
}
