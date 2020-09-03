/*
Source:
https://www.hackerrank.com/challenges/library-fine/problem

Problem:
Your local library needs your help! Given the expected and actual return dates for a library book, create a program that calculates the fine (if any).
The fee structure is as follows:
- If the book is returned on or before the expected return date, no fine will be charged (i.e.: fine = 0).
- If the book is returned after the expected return day but still within the same calendar month and year as the expected return date, fine = 15 * (number of days late).
- If the book is returned after the expected return month but still within the same calendar year as the expected return date, fine = 500 * (number of months late).
- If the book is returned after the calendar year in which it was expected, there is a fixed fine of 10,000.
Charges are based only on the least precise measure of lateness.
For example, whether a book is due January 1, 2017 or December 31, 2017, if it is returned January 1, 2018, that is a year late and the fine would be 10,000.

Function Description:
Complete the libraryFine function in the editor below. It must return an integer representing the fine due.
libraryFine has the following parameter(s):
* d1, m1, y1: returned date day, month and year
* d2, m2, y2: due date day, month and year

Input Format:
The 1st line contains 3 space-separated integers, d1,m1,y1, denoting the respective day, month, and year on which the book was returned.
The 2nd line contains 3 space-separated integers, d2,m2,y2 denoting the respective day, month, and year on which the book was due to be returned.

Constraints
1 <= d1, d2 <= 31
1 <= m1, m2 <= 12
1 <= y1, y2 <= 3000

Output Format:
Print a single integer denoting the library fine for the book received as input.
*/

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "time"
)
func date(y int32, m int32, d int32) time.Time {
    return time.Date(int(y), time.Month(m), int(d), 0, 0, 0, 0, time.UTC)
}

// Complete the libraryFine function below.
func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
    act := date(y1, m1, d1)
    due := date(y2, m2, d2)

    days := act.Sub(due).Hours() / 24

    if y1 > y2 {
        return 10000
    }

    if y2 == y1 && m2 == m1 && d2 < d1 {
            return int32(15 * days)
    } 


    if y2 == y1 && m2 < m1 {
        return int32(500 * (m1 - m2))
    }
    
    return 0
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    d1M1Y1 := strings.Split(readLine(reader), " ")

    d1Temp, err := strconv.ParseInt(d1M1Y1[0], 10, 64)
    checkError(err)
    d1 := int32(d1Temp)

    m1Temp, err := strconv.ParseInt(d1M1Y1[1], 10, 64)
    checkError(err)
    m1 := int32(m1Temp)

    y1Temp, err := strconv.ParseInt(d1M1Y1[2], 10, 64)
    checkError(err)
    y1 := int32(y1Temp)

    d2M2Y2 := strings.Split(readLine(reader), " ")

    d2Temp, err := strconv.ParseInt(d2M2Y2[0], 10, 64)
    checkError(err)
    d2 := int32(d2Temp)

    m2Temp, err := strconv.ParseInt(d2M2Y2[1], 10, 64)
    checkError(err)
    m2 := int32(m2Temp)

    y2Temp, err := strconv.ParseInt(d2M2Y2[2], 10, 64)
    checkError(err)
    y2 := int32(y2Temp)

    result := libraryFine(d1, m1, y1, d2, m2, y2)

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
