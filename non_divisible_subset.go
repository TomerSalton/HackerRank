/*
Source:
https://www.hackerrank.com/challenges/non-divisible-subset/problem

Problem:
Given a set of distinct integers, print the size of a maximal subset of S where the sum of any 2 numbers in S' is not evenly divisible by k.
For example, the array S = [19,10,12,24,25,22] and k=4.
One of the arrays that can be created is S'[0] = [10, 12, 25]. 
Another is S'[1] = [19,22,24].
After testing all permutations, the maximum length solution array has 3 elements.

Function Description
Complete the nonDivisibleSubset function in the editor below.
It should return an integer representing the length of the longest subset of S meeting the criteria.

nonDivisibleSubset has the following parameter(s):
* S: an array of integers
* k: an integer

Input Format:
The first line contains 2 space-seperated integers, n and k, the number of values in S and the non factor.
The second line contains n space-seperated integers describing S[i], the unique values of the set.

Output Format:
The size of the largest possible susbset (S')
*/

package main

import (
    "bufio"
    "fmt"
    "io"
    "math"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */

func nonDivisibleSubset(k int32, s []int32) int32 {
    // init
    var res int32
    var counters []int32
    for i := 0; i < int(k); i++ {
        counters = append(counters, 0)
    }

    // count instances of each modulo
    for _, v := range s {
        mod := int(math.Mod(float64(v), float64(k)))
        counters[mod] = counters[mod] + 1
    }

    // find the max of each pair
    var modIndex int32
    for modIndex = 0; modIndex < int32(len(counters))/2 + 1; modIndex++ {

        // only one modulo 0 can attend the array
        // e.g.: k = 10, then 10 + 20 % 10 = 0.
        if modIndex == 0 {
            if counters[0] > 0 {
                res++
            }
            continue
        }

        // if k is even, only one modulo of half the k can attend the array
        // e.g: k = 10, you cant have 2 5's since 5 + 5 % 10 = 0.
        pair := int32(len(counters)) - modIndex
        if modIndex == pair {
            if counters[pair] > 0 {
                res++
            }
            continue
        }

        res = res + int32(math.Max(float64(counters[modIndex]), float64(counters[pair])))
    }

    return res           
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var s []int32

    for i := 0; i < int(n); i++ {
        sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
        checkError(err)
        sItem := int32(sItemTemp)
        s = append(s, sItem)
    }

    result := nonDivisibleSubset(k, s)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
