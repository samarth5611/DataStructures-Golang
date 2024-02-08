func groupAnagrams(strs []string) [][]string {
    res := [][]string{}
    hashMap := map[string][]string{}

    for _, curString := range(strs){
        hashString := sortString(curString)
        hashMap[hashString] = append(hashMap[hashString], curString)
    }

    for _, value := range(hashMap){
        res = append(res, value)
    }
    return res
}

func sortString(s string) string{
    newString := strings.Split(s , "")
    sort.Strings(newString)
    return strings.Join(newString , "")
}
