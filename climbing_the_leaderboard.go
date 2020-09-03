/*
Source:
https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem

Alice is playing an arcade game and wants to climb to the top of the leaderboard and wants to track her ranking. The game uses Dense Ranking, so its leaderboard works like this:

* The player with the highest score is ranked number 1 on the leaderboard.
* Players who have equal scores receive the same ranking number, and the next player(s) receive the immediately following ranking number.

For example, the four players on the leaderboard have high scores of 100, 90, 90, and 80.
Those players will have ranks 1, 2, 2, and 3, respectively.
If Alice's scores are 70, 80 and 105, her rankings after each game are 4th, 3rd and 1st.

Function Description:
Complete the climbingLeaderboard function in the editor below.
It should return an integer array where each element res[j] represents Alice's rank after the jth game.

climbingLeaderboard has the following parameter(s):
* scores: an array of integers that represent leaderboard scores
* alice: an array of integers that represent Alice's scores

Input Format:
The first line contains an integer n, the number of players on the leaderboard.
The next line contains n space-separated integers scores[i], the leaderboard scores in decreasing order.
The next line contains an integer, m, denoting the number games Alice plays.
The last line contains m space-separated integers alice[j], the game scores.

Output Format:
Print m integers. The jth integer should indicate Alice's rank after playing the jth game.
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

func normalize(scores []int32) []int32 {
    var a []int32
    for i := len(scores) - 1; 0 <= i; i-- {
        if len(a) == 0 || a[len(a) - 1] != scores[i] {
            a = append(a, scores[i])
        }
    }
    return a
}

func findPosition(leaderboard []int32, score int32, first int) int32 {
    for i := first; i < len(leaderboard); i++ {
        if score < leaderboard[i] {
            return int32(len(leaderboard) - (i-1))
        }
        if score == leaderboard[i] {
            return int32(len(leaderboard) - i)
        }
    }
    return 1
}

// Complete the climbingLeaderboard function below.
func climbingLeaderboard(scores []int32, alice []int32) []int32 {
    var first int
    var res []int32
    arr := normalize(scores)

    for _, v := range alice {
        pos := findPosition(arr, v, first)
        res = append(res, pos)
        first = len(arr) - int(pos) + 1
    }
    
    return res

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024*1024*3)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    scoresTemp := strings.Split(readLine(reader), " ")

    var scores []int32

    for i := 0; i < int(scoresCount); i++ {
        scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
        checkError(err)
        scoresItem := int32(scoresItemTemp)
        scores = append(scores, scoresItem)
    }

    aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    aliceTemp := strings.Split(readLine(reader), " ")

    var alice []int32

    for i := 0; i < int(aliceCount); i++ {
        aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
        checkError(err)
        aliceItem := int32(aliceItemTemp)
        alice = append(alice, aliceItem)
    }

    result := climbingLeaderboard(scores, alice)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
