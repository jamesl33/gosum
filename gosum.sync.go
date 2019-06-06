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
    "strconv"
    "time"
)

func SumInFile(filename string) big.Int {
    sum := big.NewInt(0)

    if file, err := os.Open(filename); err == nil {
        scanner := bufio.NewScanner(file)

        for scanner.Scan() {
            number, _ := strconv.Atoi(scanner.Text())
            sum.Add(sum, big.NewInt(int64(number)))
        }
    }

    return *sum
}

func main() {
    start := time.Now()
    sum := SumInFile("/tmp/numbers.txt")
    fmt.Printf("Total: %s, solution took %s", sum.String(), time.Since(start))
}
