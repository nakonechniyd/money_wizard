package strings

func StrOrStr(value string, defVal string) string {
	if value != "" {
		return value
	}
	return defVal
}
