package permutation

import (
    "fmt"
    "github.com/alexaandru/utils"
    "strconv"
    "strings"
)

// Permutation holds an int permutation
type Permutation []int

// Verbose is an option
const Verbose = false

// Quiet is an option
const Quiet = false

// Equal compares two permutations for equality
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

// Reverse reverses and negates a permutation range (between n and m)
func (p *Permutation) Reverse(n, m int) {
    for i := 0; i <= (m-n)/2; i++ {
        (*p)[n+i], (*p)[m-i] = -(*p)[m-i], -(*p)[n+i]
    }
}

// Index returns the index corresponding to val
func (p Permutation) Index(val int) int {
    for k, v := range p {
        if v == val || v == -val {
            return k
        }
    }

    return -1
}

// SortByReversal sorts a permutation by reversing segments
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

// CountBreakPoints counts the breaks between contiguous ranges in the permutation
func (p Permutation) CountBreakPoints() (count int) {
    for i := 1; i < len(p); i++ {
        a, b := utils.AbsInt(p[i-1]), utils.AbsInt(p[i])
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

// IdentityPermutation returns the identity permutation
func IdentityPermutation(n int) (per Permutation) {
    per = make(Permutation, n)

    for i := 0; i < n; i++ {
        per[i] = i + 1
    }

    return
}

// LoadPermutationFromFile loads a permutation from file
func LoadPermutationFromFile(fname string) (per Permutation) {
    data := utils.LoadFile(fname)
    tokens := strings.Split(data[1:len(data)-1], " ")
    xs := make([]int, len(tokens))
    for k, v := range tokens {
        xs[k] = utils.ParseInt(v)
    }

    return Permutation(xs)
}
