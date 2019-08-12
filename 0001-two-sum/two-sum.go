package leetcode

func twoSum(nums []int, target int) []int {
    result := make([]int, 2, 2)
    for index1, num1 := range nums {
        for index2, num2 := range nums {
            if num1 + num2 == target && index1 != index2 {
                result[0] = index1
                result[1] = index2
                return result
            }
        }
    }
    return result
}
