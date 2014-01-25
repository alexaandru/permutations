package permutation

import (
    "fmt"
    "testing"
)

func TestEqual(t *testing.T) {
    a, b := &Permutation{1, 2, 3}, &Permutation{1, 2, 3}

    if !a.Equal(b) {
        t.Error(fmt.Sprintf("Expected %+v to Equal() %+v, it did not", a, b))
    }
}

func TestReverse(t *testing.T) {
    a, b := &Permutation{1, 2, 3, 4, 5, 6}, &Permutation{1, -5, -4, -3, -2, 6}
    b.Reverse(1, 4)

    if !a.Equal(b) {
        t.Error(fmt.Sprintf("Expected %+v to Equal() %+v, it did not", a, b))
    }
}

func TestIndex(t *testing.T) {
    a := &Permutation{1, 2, 3}

    if a.Index(3) != 2 {
        t.Error(fmt.Sprintf("Expected %+v Index() to return 2", a))
    }
}

func TestSortByReversal(t *testing.T) {
    t.SkipNow()
}

func TestCountBreakPoints(t *testing.T) {
    t.SkipNow()
}

func TestString(t *testing.T) {
    t.SkipNow()
}

func TestIdentityPermutation(t *testing.T) {
    t.SkipNow()
}

func TestLoadPermutationFromFile(t *testing.T) {
    t.SkipNow()
}
