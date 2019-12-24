package main

func interfaceMap2stringMap(input map[string]interface{})  map[string]string {
	output := make(map[string]string)
	for key, value := range input {
		output[key] = value.(string)
	}
	return output
}

func interfaceList2stringList(input []interface{})  []string {
	output := make([]string, 0)
	for _, value := range input {
		output = append(output,value.(string))
	}
	return output
}
