package main

import(
	"fmt"
	"reflect"
)


func main(){
	// Arrays

	var languages = [5]string{"Java", "Javascript", "Typescript" , "Go is awesome" , "Rust"}

/*	for i := 0 ; i < len(languages) ; i++ {
		option := languages[i]
		fmt.Println(i,option)
	}

	for i , option := range languages{
		fmt.Println(i,option)
	}
*/
	// Slices

	allLangs := languages[:]
	fmt.Println(reflect.TypeOf(allLangs).Kind())

	frameworks := []string{"React", "Vue", "Angular", "Svelte",
	"Laravel", "Django", "Flask", "Fiber",}

	var jsframeworks = []string{}
	jsframeworks = frameworks[0:4:4]
	frameworks = append(frameworks,"Meteor")

	fmt.Printf("All Frameworks : %v \n" , frameworks)
	fmt.Printf("Js Frameworks : %v \n" , jsframeworks)

	firstReleases := map[string]int {
		"C": 1972, "C++": 1985, "Java": 1996,
		"Python": 1991, "JavaScript": 1996, "Go": 2012,
	}

	for k,v :=  range firstReleases {
		fmt.Printf("%s was first released in the year %d\n", k,v)

	}

	//Control Flow
	x := []float64{2.15, 3.14, 42.0, 29.5}
	fmt.Println(average(x))   // 19.197499999999998
	fmt.Println(average2(x))

	count := 1
	for count < 5{
		count++
	}
	fmt.Println(count)
}


func average(x []float64) (avg float64){
	total := 0.0
	if len(x) == 0 {
		avg = 0
	} else{
		for _,v := range x{
			total +=v
		}
	avg = total / float64(len(x))
	}
	return 
}

func average2(x []float64) (avg float64){
	total := 0.0

	switch len(x) {
	case 0:
		avg = 0
	
	default:
		for _,v := range x{
			total += v
		}
		avg = total / float64(len(x))
	}
// 	return
// }

