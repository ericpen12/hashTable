package HashTable

func HashFunction(str string)  int {
	var sum int = 0
	for i, s := range str {
		ii := int(i) + 1
		ss := int(s)
		sum += ss*ii
	}
	sum = sum % 2069
	return sum
}