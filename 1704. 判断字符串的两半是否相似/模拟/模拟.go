package main

func main() {

}

func halvesAreAlike(s string) bool {
	former, latter := 0, 0
	set := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	strBytes := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		if _, ok := set[strBytes[i]]; ok { // 是元音
			former++
		}
	}
	for i := len(s) / 2; i < len(s); i++ {
		if _, ok := set[strBytes[i]]; ok { // 是元音
			latter++
		}
	}
	return former == latter
}
