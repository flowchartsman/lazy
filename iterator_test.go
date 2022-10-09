package lazy

import (
	"fmt"
)

func ExampleFuncIterator() {
	myIter := FuncIterator(0, func(i int) (int, error) {
		if i < 8 {
			return i + 2, nil
		}
		return i, Done
	})

	for myIter.Next() {
		fmt.Println(myIter.Val())
	}
	if myIter.Err() != nil {
		fmt.Println("err:" + myIter.Err().Error())
	}

	// Output:
	// 2
	// 4
	// 6
	// 8
}
