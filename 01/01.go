package main

import (
	"fmt"
	"container/list"
)


func main(){
	iLowerLimit := 1
	iUpperLimit := 1000

	var aList list.List

	for i:=iLowerLimit; i < iUpperLimit; i++{
		if i % 3 == 0{
			aList.PushBack(i)
		}else if i % 5 == 0 {
			aList.PushBack(i)
		}

	}

	PrintSumOfList(&aList)
}

func PrintSumOfList(aList *list.List){
	var iSum int
	iSum = 0

	for iVal:=aList.Front(); iVal != nil ; iVal=iVal.Next(){
		iValue := iVal.Value.(int)
		iSum += iValue
	}

	fmt.Println("The sum of the values is ", iSum)
}
