/*
Source:
https://www.hackerrank.com/challenges/electronics-shop/problem

Problem:
Monica wants to buy a keyboard and a USB drive from her favorite electronics store. The store has several models of each. Monica wants to spend as much as possible for the 2 items, given her budget.

Given the price lists for the store's keyboards and USB drives, and Monica's budget, find and print the amount of money Monica will spend. If she doesn't have enough money to both a keyboard and a USB drive, print -1 instead. She will buy only the two required items.

For example, suppose she has b = 60 to spend. Three types of keyboards cost keyboards = [40,50,60]. Two USB drives cost drives = [5,8,12]. She could purchase a 40 + 12 = 52, or a 50 + 8 = 59. She chooses the latter. She can't buy more than 2 items so she can't spend exactly 60.

Function Description:
Complete the getMoneySpent function in the editor below. It should return the maximum total price for the two items within Monica's budget, or -1 if she cannot afford both items.

getMoneySpent has the following parameter(s):

keyboards: an array of integers representing keyboard prices
drives: an array of integers representing drive prices
b: the units of currency in Monica's budget

Input Format:
The first line contains three space-separated integers b, n, and m, her budget, the number of keyboard models and the number of USB drive models.
The second line contains n space-separated integers keyboards, the prices of each keyboard model.
The third line contains m space-separated integers drives, the prices of the USB drives.

Constraints:
* 1 <= n, m <= 1000
* 1 <= b <= 10^6
* The price of each item is in the inclusive rane 1-10^6.
It is guaranteed that each type is 1, 2, 3, 4, or 5.

Output Format:
Print a single integer denoting the amount of money Monica will spend. If she doesn't have enough money to buy one keyboard and one USB drive, print -1 instead.
*/

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "sort"
    "strconv"
    "strings"
)

type Sortable struct {
    arr []int32
}

func (s Sortable) Len() int {
    return len(s.arr)
}

func (s Sortable) Less(i, j int) bool {
    return s.arr[i] <= s.arr[j]
}

func (s Sortable) Swap(i, j int) {
    s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}

func (s Sortable) Val(i int) int32 {
    return s.arr[i]
}

func (s Sortable) Last() int32 {
    return s.arr[len(s.arr) - 1]
}

func (s Sortable) Array() []int32 {
    return s.arr
}

/*
 * Complete the getMoneySpent function below.
 */
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
    var max int32 = -1

    // Prepare structs for sorting
    k := Sortable{arr: keyboards}
    d := Sortable{arr: drives}

    // Sort the arrays    
    sort.Sort(k)
    sort.Sort(d)

    // Naive case -> cannot afford
    if k.Val(0) + d.Val(0) > b {
        return max
    }

    // Naive case -> can afford most expensive
    if k.Last() + d.Last() < b {
        return k.Last() + d.Last()
    }

    for _, keyboard := range k.Array() {
        
        // Complete loop when keyboards are too expensive
        if b <= keyboard {
            return max
        }

        for _, drive := range d.Array() {

            // Optimal return value
            if drive + keyboard == b {
                return b
            }
            
            // Complete loop when drives are too expensive
            if drive + keyboard > b {
                break
            }

            // If affordable and more expensive than current max, update max    
            if drive + keyboard > max {
                    max = drive + keyboard
            }
        }
    }
    return max
}
