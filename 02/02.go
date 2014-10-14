package main

import (
	"container/list"
	"fmt"
)

var aFibonnacciList list.List

func main(){
	iUpperLimit := 4000000
	fibonacciNumbers(iUpperLimit, &aFibonnacciList)
	printEvenSequence(&aFibonnacciList)
}


func fibonacciNumbers(iUpperLimit int, aList *list.List)  {
	iX := 0
	iY := 1
	iZ := iX + iY

	for iZ < iUpperLimit {
		iZ = iX + iY
		// add the generated value to the list
		if iZ < iUpperLimit{
			aList.PushBack(iZ)
		}

		iX = iY
		iY = iZ
	}
}


func printEvenSequence(aList *list.List){
	var iSum int
	iSum = 0
	for iVal:=aList.Front(); iVal!=nil; iVal = iVal.Next(){
//		fmt.Println(iVal.Value)
		iParsedVal := iVal.Value.(int)
		if iParsedVal % 2 == 0 {
			iSum += iParsedVal
		}

	}

	fmt.Println("sum of val is ", iSum)
}
