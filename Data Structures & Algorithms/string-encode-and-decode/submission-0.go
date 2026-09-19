type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder

	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteByte('#')
		sb.WriteString(str)
	}

	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	i := 0
	result := []string{}

	for i < len(encoded) {
        j := strings.IndexByte(encoded[i:], '#') + i
        length, _ := strconv.Atoi(encoded[i:j])
        start := j + 1
        end := start + length
        result = append(result, encoded[start:end])
        i = end
	}

	return result
}
