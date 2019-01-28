package main

var (
	tsks = make(map[string]interface{})
)

func register() {
	tsks["add"] = add
	tsks["div"] = div
	tsks["mul"] = mul
	tsks["sub"] = sub
}

func add(arg1, arg2 float64) (float64, error) {
	return arg1 + arg2, nil
}
func div(arg1, arg2 float64) (float64, error) {
	return arg1 / arg2, nil
}
func mul(arg1, arg2 float64) (float64, error) {
	return arg1 * arg2, nil
}
func sub(arg1, arg2 float64) (float64, error) {
	return arg1 - arg2, nil
}
