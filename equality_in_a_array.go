/*
Source:
https://www.hackerrank.com/challenges/equality-in-a-array/problem
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

type Array struct {
    a []int32
}

func (a *Array) Init() {
    a.a = make([]int32, 101)
}

func (a *Array) Append(n int32) {
    a.a[n] = a.a[n] + 1
}

func (a *Array) Max() int32 {
    var max int32
    for _, v := range a.a {
        if v > max {
            max = v
        }
    }
    return max
}

// Complete the equalizeArray function below.
func equalizeArray(arr []int32) int32 {
    counters := Array{}
    counters.Init()
    for _, v := range arr {
    counters.Append(v)
    }

    m := counters.Max()
    return int32(len(arr)) - m
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := equalizeArray(arr)

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
