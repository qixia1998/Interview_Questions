package main

func main() {
	var doubleM map[string]map[string]string
	// panic: assignment to entry in nil map
	//doubleM = make(map[string]map[string]string)
	v1 := make(map[string]string)
	v1["k1"] = "v1"
	doubleM["m1"] = v1

}
