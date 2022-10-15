package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	ransomArr, magazineArr := [26]int{}, [26]int{}
	for i := 0; i < len(ransomNote); i++ {
		ransomArr[ransomNote[i]-'a']++
	}
	for i := 0; i < len(magazine); i++ {
		magazineArr[magazine[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if ransomArr[i] > magazineArr[i] {
			return false
		}
	}
	return true
}
