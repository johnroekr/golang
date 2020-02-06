package myPackage

import "fmt"

var pop map[string]string
var PopList map[string]string

func init() {
	pop = make(map[string]string)
	pop["Adele"] = "Hello"
	pop["Alicia Keys"] = "Fallin'"
	pop["John Legend"] = "All of Me"

	PopList = pop
}

func GetMusic(singer string) string {
	return pop[singer]
}

func getKeys() {
	for _, kv := range pop {
		fmt.Println(kv)
	}
}
