package main

import "fmt"

func twoSum() []int {
	nums := []int{2, 7, 11, 15}
	target := 9

	m := make(map[int]int)

	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}
	}
	return nil
}

func main() {
	result := twoSum()
	fmt.Println(result)
}
