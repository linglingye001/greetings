package test

type Test struct {
	Name string
}

func (t *Test) Hello() string {
	return "Hello, " + t.Name
}


