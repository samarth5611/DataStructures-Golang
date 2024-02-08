func containsDuplicate(nums []int) bool {
    mp := map[int]bool{}

    for _, num := range(nums){
        if mp[num] == true{
            return true
        }
        mp[num] = true
    }
    return false
}