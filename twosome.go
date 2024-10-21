package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ { 
            sum := nums[i] + nums[j] 
            if sum == target {
                return []int{i, j} 
            }
        }
    }
    return nil 
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    if result != nil {
        fmt.Printf("Indices: %v\n", result)
    } else {
        fmt.Println("No solution found")
    }
}
