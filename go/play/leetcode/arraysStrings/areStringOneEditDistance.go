/*
Test cases
"ab"
"acb"
""
""
""
"a"
"a"
""
"a"
"A"
"abcc"
"accc"
"c"
"c"
*/
func isOneEditDistance(s string, t string) bool {
    ns, nt := len(s), len(t)
    // Ensure that s is shorter than t.
    if ns>nt {
        return isOneEditDistance(t, s)
    }
    if nt-ns > 1{
        return false
    }
    for i, _ := range(s) {
        if s[i]!=t[i] {
            if ns==nt {
                return s[i+1:]==t[i+1:]
            } else {
                return s[i:]==t[i+1:]
            }
        }
    }
    //Strings are same. If lengths are different by 1, return true.
    return nt-ns == 1
//     if len(s) == 1 && len(t) == 1 {
//         return s!=t
//     }
//     if len(s) == 0 || len(t) == 0 {
//         return math.Round(math.Abs(float64(len(s)-len(t)))) == 1
//     }
//     //left & right
//     ls, rs := getSplitString(s)
//     lt, rt := getSplitString(t)

//     //distance
//     dl := isOneEditDistance(ls, lt)
//     dr := isOneEditDistance(rs, rt)

//     return dl != dr

    // sIndex := 0
    // tIndex := 0
    // distance := 0
    // for (sIndex < len(s) || tIndex < len(t)) && distance < 2 {
    //     if sIndex < len(s) && tIndex < len(t) && s[sIndex] == t[tIndex] {
    //         sIndex++
    //         tIndex++
    //     } else {
    //         distance++
    //         if sIndex+1 < len(s) && tIndex < len(t) && s[sIndex+1] == t[tIndex] {
    //             sIndex++
    //         } else if tIndex+1 < len(t) && sIndex < len(s) && t[tIndex+1] == s[sIndex] {
    //             tIndex++
    //         } else {
    //             sIndex++
    //             tIndex++
    //         }
    //         // } else if tIndex+1<len(t) && sIndex+1<len(s) && s[sIndex+1] == t[tIndex+1] {
    //         //     sIndex++
    //         //     tIndex++
    //         // } else if sIndex < len(s) {
    //         //     sIndex++
    //         // } else if tIndex < len(t) {
    //         //     tIndex++
    //         // }
    //     }
    // }
    // return distance==1

}

func getSplitString(s string) (string, string) {
    ms := len(s)/2
    var ls,rs string
    if ms>0 {
        ls = s[0:ms]
        if ms>1 {
            rs = s[ms+1:len(s)-1]
        }
    }
    return ls,rs
}
