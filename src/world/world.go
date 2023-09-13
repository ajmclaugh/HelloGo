package HelloGo

func AddWorldPointer(hello *string) {
	*hello = *hello + " World!"
}

func AddWorld(hello string) string {
	hello = hello + " World!"
	return hello
}

//https://blog.logrocket.com/benchmarking-golang-improve-function-performance/