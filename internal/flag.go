package internal

func ReadFromArgs(args []string, name string) ([]string, string) {
	left := []string{}
	foundKey := false
	val := ""
	for _, v := range args {
		if foundKey {
			val = v
			foundKey = false
			continue
		}
		if v == "-"+name || v == "--"+name {
			foundKey = true
			continue
		}

		left = append(left, v)
	}

	if foundKey {
		val = "true"
	}

	return left, val
}
