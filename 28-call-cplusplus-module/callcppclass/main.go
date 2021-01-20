// main.go

package main

//#include "person_cgo.h"
import "C"

type CGO_Person_T = C.Person_T

func CGO_NewPerson(name string, age int) *CGO_Person_T {
	return C.NewPerson(C.CString(name), C.int(age))
}

func CGO_PersonSay(person *CGO_Person_T) {
	C.PersonSay(person)
}

func CGO_DeletePerson(person *CGO_Person_T) {
	C.DeletePerson(person)
}

func main() {
	person := CGO_NewPerson("zhangsan", 20)
	CGO_PersonSay(person)
	CGO_DeletePerson(person)
}
