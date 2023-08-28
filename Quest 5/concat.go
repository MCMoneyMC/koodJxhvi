package piscine

func Concat(str1 string, str2 string) string {
	return string(append([]rune(str1), []rune(str2)...))
}
