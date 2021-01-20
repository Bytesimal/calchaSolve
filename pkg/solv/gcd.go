/*
 * Copyright © 2021 NeuroByte Tech. All rights reserved.
 *
 * NeuroByte Tech is the Developer Company of Rohan Mathew.
 *
 * Project: calchaSolve
 * File Name: gcd.go
 * Last Modified: 20/01/2021, 15:39
 */

package solv

import "math"

// Uses Euclid's GCD algorithm to calculate HCF between two numbers faster than brute-forcing every number in
// the range of either n1 or n2.
func GCD(n1, n2 int64) int64 {
	// handle negative inputs
	n1 = int64(math.Abs(float64(n1)))
	n2 = int64(math.Abs(float64(n2)))

	for n1 != n2 {
		if n1 > n2 {
			n1 -= n2
		} else {
			n2 -= n1
		}
	}

	return n1
}
