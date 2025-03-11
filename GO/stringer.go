package main

import "fmt"

//go:generate stringer -type=Age

type Age int

const (
	CHILDERN Age = iota
	ADOLESCENTS
	ADULTS
)

//func (a Age) String() string {
//	switch a {
//	case CHILDERN:
//		return "childern"
//	case ADOLESCENTS:
//		return "adolescents"
//	case ADULTS:
//		return "adults"
//	default:
//		return ""
//	}
//}

func main() {
	fmt.Println(CHILDERN.String())
}
