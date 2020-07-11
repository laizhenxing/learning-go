package main

func main() {
	i := GetValue()
	// 只有接口类型才能使用类型选择
	switch i.(type) {	// cannot type switch on non-interface value i (type int)
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}

// wrong : return int
func GetValue() interface{} {
	return 1
}
