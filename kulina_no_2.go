func findTheDissimilarity(s string, t string) byte {
	j, perbedaan := len(t), t[len(t)-1]
	for i := 0; i < j-1; i++ {
		perbedaan ^= s[i]
		perbedaan ^= t[i]
	}
	return perbedaan
}
