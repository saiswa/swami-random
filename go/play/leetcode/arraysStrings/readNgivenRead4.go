/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var solution = func(read4 func([]byte) int) func([]byte, int) int {
    // implement read below.
    return func(buf []byte, n int) int {
        // Solution without buffer
        copiedChars :=0
        readChars :=4

        for (copiedChars < n && readChars == 4) {
            readChars = read4(buf[copiedChars:])
            copiedChars += readChars
        }
        return int(math.Round(math.Min(float64(n), float64(copiedChars))))

//         // Solution with buffer
//         copiedChars :=0
//         readChars :=4
//         buf4 := make([]byte, 4)

//         for (copiedChars < n && readChars == 4) {
//             readChars = read4(buf4)
//             for i:=0; i<readChars; i++ {
//                 if copiedChars == n {
//                     return copiedChars
//                 }
//                 buf[copiedChars] = buf4[i]
//                 copiedChars++
//             }
//         }
        // return copiedChars

        // Original solution
        // // Read 4 chars at a time. Terminate if end of file or n reached.
        // fmt.Printf("Starting %s %d\n", buf, n)
        // charsRead := 0
        // isFinished := false
        // for !isFinished {
        //     if charsRead >= n {
        //         isFinished = true
        //         break
        //     }
        //     fmt.Println("charsRead", charsRead, "n", n, "buf", buf)
        //     bufCharsRead := 0
        //     charsToRead := n - charsRead
        //     if n - charsRead > 4 {
        //         bufCharsRead = read4(buf[charsRead:])
        //     } else {
        //         buf4 := make([]byte, 4)
        //         buf4CharsRead := read4(buf4)
        //         fmt.Println("buf4CharsRead", buf4CharsRead, "charsToRead", charsToRead)
        //         actualCharsToRead := min(charsToRead, buf4CharsRead)
        //         for i := 0; i<actualCharsToRead; i++ {
        //             buf[charsRead] = buf4[i]
        //             charsRead++
        //         }
        //     }
        //     fmt.Println("bufCharsRead", bufCharsRead, "charsRead", charsRead, "n", n, "buf", buf)
        //     charsRead += bufCharsRead
        //     if bufCharsRead < 4 {
        //         isFinished = true
        //         break
        //     }
        // }
        // fmt.Printf("Final: %s %d\n", buf, n)
        // return charsRead
    }
}

// func min(a int, b int) int {
//     if a < b {
//         return a
//     } else {
//         return b
//     }
// }
