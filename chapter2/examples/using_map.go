package examples


import "runtime"


func UsingMapWithPointersDemo() {
	 N := 4000000
	myMap := make(map[int]*int)
	for i:=0; i<N; i++ {
		value := int(i)
		myMap[value] = &value
	}
	runtime.GC()
	_ = myMap[0]
}


func UsingMapWithoutPointersDemo() {
	N := 4000000
	myMap := make(map[int] int)

	for i:=0; i< N; i++ {
		value := int(i)
		myMap[value] = value

	}

	runtime.GC()
	_ = myMap[0]
}

func SplittingMapDemo() {
	var N = 40000000
	split := make([]map[int]int, 200)
	for i := range split {
		split[i] = make(map[int] int)
	}

	for i := 0; i < N; i++ {
		value := int(i)
		split[i%200][value] = value
	}

	runtime.GC()
	_ = split[0][0]
}
