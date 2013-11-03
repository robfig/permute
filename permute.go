package main

import "fmt"

func permute(todo string) []string {
	return helper(nil, []rune(todo))
}

func helper(done, todo []rune) []string {
	if len(todo) == 0 {
		return []string{string(done)}
	}

	var out []string
	for i, ch := range todo {
		out = append(out,
			helper(append(done, ch), remove(todo, i))...)
	}
	return out
}

func main() {
	for _, str := range permute("") {
		fmt.Println(str)
	}
}

func remove(in []rune, i int) []rune {
	var out []rune
	out = append(out, in[:i]...)
	out = append(out, in[i+1:]...)
	return out
}
