package main

import (
	"runtime"

	"github.com/rajcspsg/mastering_go/chapter2/examples"
)

func main() {
	var mem runtime.MemStats
	examples.PrintStats(mem)
	//examples.UsingMapWithoutPointersDemo()
	examples.SplittingMapDemo()
	examples.PrintStats(mem)

}
