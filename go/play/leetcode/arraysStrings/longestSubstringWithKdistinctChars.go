/*
"ba"
2
"ece"
2
"eceba"
2
"aa"
1
*/

func lengthOfLongestSubstringKDistinct(s string, k int) int {
    n := len(s)
    if n==0 || k==0 {
        return 0
    }
    // left & right pointers
    l, r := 0, 0
    hashmap := make(map[byte]int)

    maxLen := 1

    for r<n {
        // key is char, value is the index in the string
        hashmap[s[r]] = r
        r++

        if len(hashmap)>k {
            // What is the left most index of the char stored?
            lowestIndex := minValue(hashmap)
            delete(hashmap, s[lowestIndex])
            l = lowestIndex + 1
        }
        maxLen = max(maxLen, r-l)
    }
    return maxLen

//     // Find longest substring of k starting with index 0
//     // When you find the next unique character, find the longest substring with k-1
//     // Terminating condition: k=0 return 0
//     // k=1: find the largest repeating single character
//     // unique set
//     set := make(map[byte]int)
//     fmt.Println(s, k, len(s))
//     if k<=0 || len(s)==0 {
//         return 0
//     }
//     // longest string of 1 char
//     if len(s)==1 {
//         return 1
//     }
//     // len(s)>1
//     // max length from beginning, 1st non-beginning character index
//     nbi := -1
//     for i,_ := range(s) {
//          if i>0 && s[i]!=s[0] && nbi < 0 {
//              nbi = i
//          }
//         occurences,isExists := set[s[i]]
//         if isExists {
//             set[s[i]] = occurences + 1
//         } else {
//             if len(set)>k {
//                 // max length found
//                 // current length
//                 cl := i
//                 // length from rest of the string
//                 lr := lengthOfLongestSubstringKDistinct(s[nbi:], k)
//                 if cl > lr {
//                     return cl
//                 } else {
//                     return lr
//                 }
//             }
//             set[s[i]] = 1
//         }

//         // if s[i]==s[0] {
//         //     fmt.Println("bl set", s, s[1:], s[0], s[i], k, bl, i)
//         //     bl++
//         // } else {
//         //     // max length of later part of string with same chars
//         //     ll := lengthOfLongestSubstringKDistinct(s[i:], k)
//         //     // max length of later part of string with 1 char less
//         //     ll1 := lengthOfLongestSubstringKDistinct(s[i:], k-1)
//         //     fmt.Println("Finding longest substrings", s, k, bl, ll, ll1)
//         //     if bl+ll1 > ll {
//         //         fmt.Println("Earlier part of string longest", s, bl+ll1, k, bl, ll)
//         //         return bl+ll1
//         //     } else {
//         //         fmt.Println("Latter part of string longest", s, ll, k, bl, ll1)
//         //         return ll
//         //     }
//         // }
//     }
//     return 0
}

func minValue(hashmap map[byte]int) int {
    min := math.MaxInt32
    for _, val := range hashmap {
        if min > val {
            min = val
        }
    }
    return min
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
