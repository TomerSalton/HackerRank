/*
Source:
https://www.hackerrank.com/challenges/encryption/problem

First, the spaces are removed from the text. Let L be the length of this text.
Then, characters are written into a grid, whose rows and columns have the following constraints:

floor(sqrt(L)) <= row <= column <= ceil(sqrt(L))

For example, the sentence s = 'if man was meant to stay on the ground god would have given us roots', after removing spaces is 54 characters long.
sqrt(54) is between 7 and 8, so it is written in the form of a grid with 7 rows and 8 columns.

ifmanwas
meanttos
tayonthe
groundgo
dwouldha
vegivenu
sroots

* Ensure that L <= rows x columns
* If multiple grids satisfy the above conditions, choose the one with the minimum area, i.e. rows X columns.

The encoded message is obtained by displaying the characters in a column,
inserting a space, and then displaying the next column and inserting a space, and so on.
For example, the encoded message for the above rectangle is:
imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn sseoau

You will be given a message to encode and print.

Function Description
Complete the encryption function in the editor below. It should return a single string composed as described.
encryption has the following parameter(s):
* s: a string to encrypt

Input Format
One line of text, the string s

Output Format
Print the encoded message on one line as described.
*/

package main

import (
	"math"
	"strings"
)

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func trim(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	return s
}

func calc(s string) (int, int) {
	L := len(s)
	sqr := math.Sqrt(float64(L))

	floor := int(math.Floor(sqr))
	ceil := int(math.Ceil(sqr))

	if L <= floor*ceil {
		return floor, ceil
	}
	return ceil, ceil
}

func createTable(rows, cols int, s string) []string {
	t := make([]string, rows)
	for i := 0; i < rows; i++ {
		start := cols * i
		end := min(cols*(i+1), len(s))
		t[i] = s[start:end]
	}
	return t
}

func readTable(t []string, cols int, rows int) string {
	var encrypted string
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if j < rows-1 {
				encrypted += t[j][i : i+1]
			} else if i < len(t[j]) {
				encrypted += t[j][i : i+1]
			}
		}
		encrypted += " "
	}
	return encrypted
}

// Complete the encryption function below.
func encryption(s string) string {
	// trim spaces from string
	s = trim(s)

	// claculte rows and oclumns
	rows, cols := calc(s)

	// fill string in table
	t := createTable(rows, cols, s)

	// read from table
	res := readTable(t, cols, rows)

	// return encrypted string
	return res
}
