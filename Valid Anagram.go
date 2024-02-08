func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    } 

    maps := map[string]int{}

    for i:=0; i<len(s); i++{
        maps[string(s[i])]++
        maps[string(t[i])]--
    }

    for _, value := range(maps){
        if value != 0{
            return false
        }
    }
    return true

}