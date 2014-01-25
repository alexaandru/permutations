package utils

import (
    u "alexaandru/utils"
    "fmt"
    "strconv"
    "strings"
)

type Permutation []int

const Verbose, Quiet = true, false

func (p *Permutation) Equal(other *Permutation) bool {
    if len(*p) != len(*other) {
        return false
    }

    for k, v := range *p {
        if v != (*other)[k] {
            return false
        }
    }

    return true
}

func (p *Permutation) Reverse(n, m int) {
    for i := 0; i <= (m-n)/2; i++ {
        (*p)[n+i], (*p)[m-i] = -(*p)[m-i], -(*p)[n+i]
    }
}

func (p Permutation) Index(val int) int {
    for k, v := range p {
        if v == val || v == -val {
            return k
        }
    }

    return -1
}

func (p Permutation) SortByReversal(verbose bool) (dist int) {
    iter := func(a, b int) {
        p.Reverse(a, b)
        dist++
        if verbose {
            fmt.Println(p)
        }
    }

    for i := 0; i < len(p); i++ {
        if p[i] != i+1 {
            iter(i, p.Index(i+1))
            if p[i] == -(i + 1) {
                iter(i, i)
            }
        }
    }

    return
}

func (p Permutation) CountBreakPoints() (count int) {
    for i := 1; i < len(p); i++ {
        a, b := u.AbsInt(p[i-1]), u.AbsInt(p[i])
        if a+1 != b && a-1 != b {
            count++
        }
    }

    count++ // no idea why

    return
}

func (p Permutation) String() string {
    out := []string{}
    for _, v := range p {
        curr := ""
        if v > 0 {
            curr += "+"
        }

        curr += strconv.Itoa(v)
        out = append(out, curr)
    }

    return "(" + strings.Join(out, " ") + ")"
}

func IdentityPermutation(n int) (per Permutation) {
    per = make(Permutation, n)

    for i := 0; i < n; i++ {
        per[i] = i + 1
    }

    return
}

func LoadPermutationFromFile(fname string) (per Permutation) {
    data := u.LoadFile(fname)
    tokens := strings.Split(data[1:len(data)-1], " ")
    xs := make([]int, len(tokens))
    for k, v := range tokens {
        xs[k] = u.ParseInt(v)
    }

    return Permutation(xs)
}
