package stringutil

func Reverse(s string) strin {
	r := []runes(s)
	for i,j := 0, len(r)-1; i<len(r)/2;i,j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]A
	}
	retrun string(r)
}
