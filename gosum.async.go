/* This file is part of project-euler

Copyright (C) 2019, James Lee <jamesl33info@gmail.com>.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>. */

package main

import (
    "bufio"
    "fmt"
    "math/big"
    "os"
    "runtime"
    "strconv"
    "time"
)

func NumLines(filename string) int {
    numLines := 0

    if file, err := os.Open(filename); err == nil {
        scanner := bufio.NewScanner(file)

        for scanner.Scan() {
            numLines++
        }
    }

    return numLines
}

func SumInPortionOfFile(filename string, start, end int, result chan <- *big.Int) {
    sum := big.NewInt(0)

    if file, err := os.Open(filename); err == nil {
        lineNum := 0
        scanner := bufio.NewScanner(file)

        for scanner.Scan() {
            if lineNum >= start + end {
                break
            } else if lineNum >= start {
                number, _ := strconv.Atoi(scanner.Text())
                sum.Add(sum, big.NewInt(int64(number)))
            }

            lineNum++
        }
    }

    result <- sum
}

func main() {
    start := time.Now()

    numRoutines := runtime.GOMAXPROCS(0)
    resultChan := make(chan *big.Int, numRoutines)

    numLines := NumLines("/tmp/numbers.txt")
    linesPerRoutine := numLines / numRoutines

    for i := 0; i < numRoutines - 1; i++ {
        go SumInPortionOfFile("/tmp/numbers.txt", i * linesPerRoutine, linesPerRoutine, resultChan)
    }

    linesRemaining := numLines / numRoutines

    go SumInPortionOfFile("/tmp/numbers.txt", (numRoutines - 1) * linesPerRoutine, linesPerRoutine + linesRemaining, resultChan)

    sum := big.NewInt(0)

    for i := 0; i < numRoutines; i++ {
        sum.Add(sum, <- resultChan)
    }

    fmt.Printf("Total: %s, solution took %s", sum.String(), time.Since(start))
}
