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

/*
Test cases
"abc"
[1,2,1]
"abcde"
[1,2,1]
""
[1,2]
"abc"
[1,5]
*/

var solution = func(read4 func([]byte) int) func([]byte, int) int {
    totalCharsCopied := 0
    buf4 := make([]byte, 4)
    totalBufCharsRemaining := 0
    bufStartIndex := 0
    return func(buf []byte, n int) int {
        fmt.Printf("Starting. Reading %d chars. Buffer: %s. bufStartIndex: %d.\n",
                   n, buf4, bufStartIndex)
        charsRead := 4
        // chars copied in this session
        charsCopied := 0
        fmt.Println("totalBufCharsRemaining", totalBufCharsRemaining)
        // First, read from the buffer
        for totalBufCharsRemaining>0 && charsCopied < n {
            buf[charsCopied] = buf4[bufStartIndex]
            charsCopied++
            bufStartIndex++
            totalBufCharsRemaining--
            fmt.Printf("Read from buffer %s %s bufStartIndex:%d totalBufCharsRemaining:%d\n",
                       buf, buf4, bufStartIndex, totalBufCharsRemaining)
        }
        for (charsCopied < n && charsRead == 4) {
            charsRead = read4(buf4)
            fmt.Println(charsRead)
            totalBufCharsRemaining = charsRead
            bufStartIndex = 0
            for i:=0; i<charsRead; i++ {
                if charsCopied == n {
                    break;
                }
                fmt.Printf("buffer char %s\n", string(buf4[i]))
                buf[charsCopied] = buf4[i]
                charsCopied++
                totalBufCharsRemaining--
                bufStartIndex++
                fmt.Printf("Read from main %s %s bufStartIndex:%d totalBufCharsRemaining:%d\n",
                       buf, buf4, bufStartIndex, totalBufCharsRemaining)
            }
        }
        fmt.Println("===========Done===========\n")
        if n < charsCopied {
            totalCharsCopied += n
            return n
        } else {
            totalCharsCopied += charsCopied
            return charsCopied
        }
    }
}
