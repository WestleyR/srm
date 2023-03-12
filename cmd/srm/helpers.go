package cmd

// Helpers

func XOR(in ...bool) bool {
	trues := 0

	for _, x := range in {
		if x {
			trues++
		}
	}

	return trues == 1
}

func ZeroOrOne(in ...bool) bool {
	trues := 0

	for _, x := range in {
		if x {
			trues++
		}
	}

	return trues > 1
}
