package solv

// Uses Euclid's GCD algorithm to calculate HCF between two numbers faster than brute-forcing every number in
// the range of either n1 or n2.
func GCD(n1, n2 int64) int64 {
	for n1 != n2 {
		if n1 > n2 {
			n1 -= n2
		} else {
			n2 -= n1
		}
	}

	return n1
}
