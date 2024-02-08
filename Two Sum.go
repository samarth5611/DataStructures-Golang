func twoSum(nums []int, target int) []int {
    mp := map[int]int{}

    for i := 0; i < len(nums); i++{
        rem := target - nums[i]
        index, ok := mp[rem]
        if ok{
            return []int{index, i}
        }
        mp[nums[i]] = i
    }
    return []int{0, 0}
}