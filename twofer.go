package main

import "fmt"

func ShareWith(name string) string {
	if name != "" {
		x := (fmt.Sprintf("One for %v, one for me", name))
	} else {
		x := string(fmt.Sprintf("One for you, one for me"))
	}
//	return fmt.Println(x)
	return x.String()
}
func main() {
	ShareWith("jenny")
}
