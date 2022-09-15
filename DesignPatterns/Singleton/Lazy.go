package Singleton

type singleton struct {
}

var ins *singleton

func GetInsOr() *singleton {
	if ins == nil {
		ins = &singleton{}
	}

	return ins
}
