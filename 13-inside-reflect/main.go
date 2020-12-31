package main

func main() {
	safeAssert()

	caseType(nil)    // you are nothing
	caseType(10)     // number 11
	caseType("test") // string test suffix

	//detectType(nil) // panic: reflect: call of reflect.Value.Type on zero Value
	detectType(10)
	detectType("test")
}
