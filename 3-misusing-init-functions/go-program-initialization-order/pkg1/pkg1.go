package pkg1

import (
	"fmt"

	_ "demo/pkg2"
)

const (
	c1  = "c1"
	c2  = "c2"
	pkg = "pkg1"
)

var (
	_  = constInitCheck(c1)
	_  = constInitCheck(c2)
	v1 = variableInit("v1")
	v2 = variableInit("v2")
)

func init() {
	fmt.Printf("%s: first-init  func invoked\n", pkg)
}

func init() {
	fmt.Printf("%s: second-init func invoked\n", pkg)
}

func main() {
	fmt.Printf("%s: main        func invoked\n", pkg)
}

func variableInit(name string) string {
	fmt.Printf("%s: variable %s has been initialized\n", pkg, name)
	return name
}

func constInitCheck(name string) string {
	if name != "" {
		fmt.Printf("%s: const %s has been initialized\n", pkg, name)
		return name
	}

	fmt.Printf("%s: const %s has not been initialized\n", pkg, name)
	return name
}
