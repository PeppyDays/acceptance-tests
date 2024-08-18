package interactions

var GreetPrefix = "Hello, "
var GreetDefaultName = "world"
var CursePrefix = "Go to hell, "
var CurseDefaultName = "you"
var CursePostfix = "!"

func Greet(name string) string {
	if name == "" {
		name = GreetDefaultName
	}
	return GreetPrefix + name
}

func Curse(name string) string {
	if name == "" {
		name = CurseDefaultName
	}
	return CursePrefix + name + CursePostfix
}
