package main

import (
	"github.com/rajcspsg/mastering_go/chapter2/examples"
	"runtime"
)

func main() {
	var mem runtime.MemStats
	examples.PrintStats(mem)
	//examples.UsingMapWithoutPointersDemo()
	examples.SplittingMapDemo()
	examples.PrintStats(mem)
}
