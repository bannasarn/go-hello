package couple

func Slice(s string) []string {

	var result []string
	patchStr := s
	lenStr := 2

	if len(s)%lenStr != 0 {
		patchStr += "*"
	}

	for i := 0; i < len(patchStr); i += lenStr {
		result = append(result, patchStr[i:i+lenStr])
	}

	return result
}
