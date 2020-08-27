/*
You are given a number of sticks of varying lengths. You will iteratively cut the sticks into smaller sticks, discarding the shortest pieces until there are none left. At each iteration you will determine the length of the shortest stick remaining, cut that length from each of the longer sticks and then discard all the pieces of that shortest length. When all the remaining sticks are the same length, they cannot be shortened so discard them.

Given the lengths of  sticks, print the number of sticks that are left before each iteration until there are none left.

For example, there are n=3 sticks of lengths arr = [1,2,3]. The shortest stick length is 1, so we cut that length from the longer two and discard the pieces of length 1. Now our lengths are arr = [1,2]. Again, the shortest stick is of length 1, so we cut that amount from the longer stick and discard those pieces. There is only one stick left, arr = [1], so we discard that stick. Our lengths are answer = [3,2,1].

Function Description

Complete the cutTheSticks function in the editor below. It should return an array of integers representing the number of sticks before each cut operation is performed.

cutTheSticks has the following parameter(s):
* arr: an array of integers representing the length of each stick

Input Format
The first line contains a single integer n, the size of arr.
The next line contains n space-separated integers, each an arr[i] where each value represents the length of the i'th stick.

Output Format
For each operation, print the number of sticks that are present before the operation on separate lines.
package main
*/

func minimum(arr []int32) int32 {
    m := arr[0]
    for _, v := range arr {
        if v < m {
            m = v
        }
    }
    return m
}

func cut(m int32, arr []int32) []int32 {
    var new []int32
    for _, v := range arr {
        v = v - m
        if v > 0 {
            new = append(new, v)
        }
    }
    return new
}

// Complete the cutTheSticks function below.
func cutTheSticks(arr []int32) []int32 {
    var res []int32
    for len(arr) > 0 {
        res = append(res, int32(len(arr)))
        m := minimum(arr)
        arr = cut(m, arr)
    }
        return res
}
