package examples

import (
	"fmt"
)

const SIZE = 15

type Node struct {
	Value int
	Next *Node
}

type HashTable struct {
	Table map[int] *Node
	Size int
}

func HashFunction(i, size int) int {
	return (i % size)
}

func insertIntoHashTable(hash *HashTable, value int) int  {
	index := HashFunction(value, hash.Size)
	element := Node { Value: value, Next: hash.Table[index] }
	hash.Table[index] = &element
	return index
}

func traverseHashTable(hash *HashTable) {
	for k := range hash.Table {
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Printf("%s \n", "nil")
		}
	}
}

func LookupHashTable(hash *HashTable, value int) bool {
	index := HashFunction(value, hash.Size)

	if hash.Table[index] != nil {
		t := hash.Table[index]

	    for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func HashTableDemo() {
	table := make(map[int] *Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Number of Spaces: ", hash.Size)

	for i:=0; i< 120; i++ {
		insertIntoHashTable(hash, i)
	}

	traverseHashTable(hash)

	fmt.Println("Looking up started")

	for i:= 100; i < 130; i++ {
		fmt.Println(i, " result is " , LookupHashTable(hash, i))
	}
}
