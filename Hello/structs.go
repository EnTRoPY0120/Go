package main

// import "fmt"

// type stack struct {
// 	index int
// 	data [5]int
// }

// func (s *stack) push(k int) {
// 	s.data[s.index] = k
// 	s.index++
// }

// func (s *stack) pop() int{
// 	s.index--
// 	return s.data[s.index]
// }

// func main(){
// 	s := new(stack)
// 	s.push(23)
// 	s.push(14)
// 	// An error in the pop method
// 	s.pop()
// 	s.pop()
// 	fmt.Printf("The stack is %v\n", *s)
// }

//type Person struct {
//	Name string
//}
//
//func setNameByReference(person *Person, name string) {
//	person.Name = name
//}
//
//func setNameByValue(person Person, name string) Person {
//	person.Name = name
//	return person
//}
//
//func main() {
//	//	v := Person{Name: "Alice"}
//	//	p := &v
//	//	p.Name = "John"
//	//	fmt.Println(v)
//	var person Person = Person{Name: "John"}
//	setNameByReference(&person, "Doe")
//
//	fmt.Println(person)
//
//	person = setNameByValue(person, "Alice")
//
//	fmt.Println(person)
//}
