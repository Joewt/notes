package main

import "fmt"

type Sorter interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
func Sort(data Sorter) {
    for pass := 1; pass < data.Len(); pass++ {
	    for i := 0; i < data.Len()-pass; i++ {
	        if data.Less(i+1, i) {
	            data.Swap(i, i+1)
	        }
	    }
    }
}

func IsSorted(data Sorter) bool {
    n := data.Len()
    for i := n - 1; i > 0; i-- {
	if data.Less(i, i-1){
	    return false
	}
    }
    return true
}

type IntArray []int

func (p IntArray) Len() int             { return len(p) }
func (p IntArray) Less(i, j int) bool   { return p[i] < p[j]}
func (p IntArray) Swap(i, j int)        { p[i], p[j] = p[j],p[i]}


type StringArray []string

func (p StringArray) Len() int           { return len(p) }
func (p StringArray) Less(i, j int) bool { return p[i] < p[j] }
func (p StringArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Convenience wrappers for common cases
func SortInts(a []int)       { Sort(IntArray(a)) }
func SortStrings(a []string) { Sort(StringArray(a)) }

func IntsAreSorted(a []int) bool       { return IsSorted(IntArray(a)) }
func StringsAreSorted(a []string) bool { return IsSorted(StringArray(a)) }


func ints() {
    data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
    a := IntArray(data) //conversion to type IntArray
    Sort(a)
    if !IsSorted(a) {
        panic("fails")
    }
    fmt.Printf("The sorted array is: %v\n", a)
}

func strings() {
    data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
    a := StringArray(data)
    Sort(a)
    if !IsSorted(a) {
        panic("fail")
    }
    fmt.Printf("The sorted array is: %v\n", a)
}

type day struct {
    num       int
    shortName string
    longName  string
}

type dayArray struct {
    data []*day
}

func (p *dayArray) Len() int           { return len(p.data) }
func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

func days() {
    Sunday    := day{0, "SUN", "Sunday"}
    Monday    := day{1, "MON", "Monday"}
    Tuesday   := day{2, "TUE", "Tuesday"}
    Wednesday := day{3, "WED", "Wednesday"}
    Thursday  := day{4, "THU", "Thursday"}
    Friday    := day{5, "FRI", "Friday"}
    Saturday  := day{6, "SAT", "Saturday"}
    data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
    a := dayArray{data}
    Sort(&a)
    if !IsSorted(&a) {
        panic("fail")
    }
    for _, d := range data {
        fmt.Printf("%s ", d.longName)
    }
    fmt.Printf("\n")
}

func main() {
    ints()
    strings()
    days()
}