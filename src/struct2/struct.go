package struct2

import "fmt"

type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}

type Cat struct {
	Name    string
	Age     uint8
	Address string
}

func (cat *Cat) Grow() {
	cat.Age++
}

func (cat *Cat) Move(newAddress string) (address string) {
	address, cat.Address = cat.Address, newAddress
	return
}

type Animal interface {
	Grow()
	Move(string) string
}

// func main() {
// 	cat := Cat{"Little C", 2, "In the house"}
// 	animal, ok := interface{}(&cat).(Animal)
// 	fmt.Printf("%v, %v\n", ok, animal)

// }

func (person *Person) Move(address string) (oldAddress string) {
	oldAddress = person.Address
	person.Address = address
	return
}

func StructHandle() {
	p := Person{Name: "Robert", Gender: "Male", Age: 33}
	oldAddress := p.Move("San Francisco")
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Address)
}
