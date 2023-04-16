package util

func StrTrimCommonSuffix(a, b string) (string, string, string) {
	aLen := len(a)
	bLen := len(b)
	var i int
	for i = 0; i < aLen && i < bLen && a[aLen-i-1] == b[bLen-i-1]; i++ {
	}
	return a[:aLen-i], b[:bLen-i], a[aLen-i:]
}
