package main

import "fmt"

func main(){

	/************************************************************************************************************************************************************/
	//	ARRAYS ---> Homogenous collection of data
	/************************************************************************************************************************************************************/

	var arr[4]int	// by default 0 values are added, if datatype is false then default val is false, empty string in case of strings

	fmt.Println(len(arr))

	nums:=[4]int{1,2,3,4}

	fmt.Println(nums)

	/************************************************************************************************************************************************************/
	//	SLICE ---> Like vectors in cpp
	/************************************************************************************************************************************************************/

	var slice[]int

	fmt.Println(slice)
	fmt.Println(len(slice))


	var slice2 = make([]int,2)	// 2 is the default size

		fmt.Println(slice2)
		fmt.Println(len(slice2))

		// cap function --> maximum number of elements that can fit

		slice2 = append(slice2,2)	// adding values after initial t20 zeroes, which we declared above
		slice2 = append(slice2,12)
		slice2 = append(slice2,22)
		slice2 = append(slice2,21)

		fmt.Println(slice2)
		fmt.Println(len(slice2))


		//   ':' is slice operator, takes start, stop,  return array's value from start to end, start is inclusive

		fmt.Println(slice2[0:4])

}