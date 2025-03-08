package _map

func GetKeys(mapContent map[string]string) []string {
	var keys []string
	for key := range mapContent {
		keys = append(keys, key)
	}
	return keys
}
