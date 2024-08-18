package dsl

type Greeter interface {
	Greet(name string) (string, error)
}

type Meany interface {
	Curse(name string) (string, error)
}
