/*
Source:
https://www.hackerrank.com/challenges/beautiful-triplets/problem

Given a sequence of integers a, a triplet (a[i],a[j],a[k]) is beautiful if:
* i < j < k
* a[k] - a[j] = a[j] - a[i] = d

Given an increasing sequenc of integers and the value of d, count the number of beautiful triplets in the sequence.

For example, the sequence arr=[2,2,3,4,5] and d=1. There are three beautiful triplets, by index: 
[i,j,k]=[0,2,3],[1,2,3],[2,3,4]. To test the first triplet, arr[j]-arr[i]=3-2=1 and arr[k]-arr[j]=4-3=1.

Function Description:
Complete the beautifulTriplets function in the editor below. It must return an integer that represents the number of beautiful triplets in the sequence.
beautifulTriplets has the following parameters:
d: an integer
arr: an array of integers, sorted ascending

Input Format:
The first line contains 2 space-separated integers n and d, the length of the sequence and the beautiful difference.
The second line contains n space-separated integers arr[i].

Output Format:
Print a single line denoting the number of beautiful triples in the sequence.
*/

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the beautifulTriplets function below.
func beautifulTriplets(d int32, arr []int32) int32 {
    var count int32

    for i := 0; i < len(arr) - 2; i++ {
        for j := i+1; j < len(arr) - 1; j++ {
            if arr[j] - arr[i] > d {
                j = len(arr)
            } else {
                for k := j+1; k < len(arr); k++ {
                    if arr[k] - arr[j] > d {
                        k = len(arr)
                    } else {
                        if arr[j] - arr[i] == d {
                            if arr[k] - arr[j] == d {
                            count++
                            }
                        }
                    }
                }
            }
        }
    }

    return count
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nd := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nd[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    dTemp, err := strconv.ParseInt(nd[1], 10, 64)
    checkError(err)
    d := int32(dTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := beautifulTriplets(d, arr)

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
