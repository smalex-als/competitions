package main

import "fmt"

func main() {
	res := fileNaming([]string{"doc", "doc", "image", "doc(1)", "doc"})
	fmt.Println(res)
}

func fileNaming(names []string) []string {
	res := make([]string, len(names))
	files := make(map[string]bool)
	for i := 0; i < len(names); i++ {
		if _, ok := files[names[i]]; ok {
			for j := 1; j < 1000; j++ {
				newname := fmt.Sprintf("%s(%d)", names[i], j)
				if _, ok2 := files[newname]; !ok2 {
					files[newname] = true
					res[i] = newname
					break
				}
			}
		} else {
			files[names[i]] = true
			res[i] = names[i]
		}
	}
	return res
}
