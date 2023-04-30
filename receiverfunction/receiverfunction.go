package receiverfunction


type Person struct {
	Name string
	Age  int
}

//SetName value-based receiver function
func (p Person) SetName(name string) {
	println("address of person object in RF", &p)
	p.Name = name
}

//SetAge pointer-based receiver function
func (p *Person) SetAge(age int) {
	println("address of person object in RF", p)
	p.Age = age
}
