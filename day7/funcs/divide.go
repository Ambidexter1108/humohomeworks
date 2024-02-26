package funcs



func Divide(masive []int) (chet, nechet []int) {
	for _, num := range masive {
		if num%2 == 0 {
			chet = append(chet, num)
		} else {
			nechet = append(nechet, num)
		}

	}
	return chet, nechet
}


