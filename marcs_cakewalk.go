/*
Source:
https://www.hackerrank.com/challenges/marcs-cakewalk/problem
*/

package main

import (
    "bufio"
    "fmt"
    "io"
    "math"
    "os"
    "sort"
    "strconv"
    "strings"
)

type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}

type Array struct {
    arr []int32
}

func (a Array) Len() int {
    return len(a.arr)
}

func (a Array) Less(i, j int) bool {
    return a.arr[i] <= a.arr[j]
}

func (a Array) Swap(i, j int) {
    tmp := a.arr[i]
    a.arr[i] = a.arr[j]
    a.arr[j] = tmp
}

func (a Array) Val(i int) int32 {
    return a.arr[i]
}

// Complete the marcsCakewalk function below.
func marcsCakewalk(calorie []int32) int64 {
    var d int64
    var mul float64

    a := Array{
        arr: calorie,
    }
    sort.Sort(a)

    for i := a.Len() - 1; 0 <= i; i-- {
        d = d + int64(float64(a.Val(i)) * math.Pow(2, mul))
        mul++
    }

    return d
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

    calorieTemp := strings.Split(readLine(reader), " ")

    var calorie []int32

    for i := 0; i < int(n); i++ {
        calorieItemTemp, err := strconv.ParseInt(calorieTemp[i], 10, 64)
        checkError(err)
        calorieItem := int32(calorieItemTemp)
        calorie = append(calorie, calorieItem)
    }

    result := marcsCakewalk(calorie)

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
