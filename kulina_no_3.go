func romanToInt(s string) int {
	romawi := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	hasil := 0
	if len(s) == 1 {
		return romawi[s[0]]
	}
	for i := 1; i < len(s); i++ {
		if romawi[s[i-1]] >= romawi[s[i]] {
			hasil += romawi[s[i-1]]
		} else {
			hasil += romawi[s[i]] - romawi[s[i-1]]
			i++
		}
		if i == len(s)-1 {
			hasil += romawi[s[i]]
		}
	}
	return hasil
}
