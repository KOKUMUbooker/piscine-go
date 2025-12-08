package main

import (
	"fmt"

	"piscine"
)

func main() {
	// a := []int{1, 2, 3, 4, 5, 6}
	// piscine.ForEach(piscine.PrintNbr, a)

	// a := []int{1, 2, 3, 4, 5, 6}
	// result := piscine.Map(piscine.IsPrime, a)
	// fmt.Println(result)

	// a1 := []string{"Hello", "how", "are", "you"}
	// a2 := []string{"This", "is", "4", "you"}

	// result1 := piscine.Any(piscine.IsNumeric, a1)
	// result2 := piscine.Any(piscine.IsNumeric, a2)

	// fmt.Println(result1)
	// fmt.Println(result2)

	// tab1 := []string{"Hello", "how", "are", "you"}
	// tab2 := []string{"This", "1", "is", "4", "you"}
	// answer1 := piscine.CountIf(piscine.IsNumeric, tab1)
	// answer2 := piscine.CountIf(piscine.IsNumeric, tab2)
	// fmt.Println(answer1)
	// fmt.Println(answer2)

	// a1 := []int{0, 1, 2, 3, 4, 5}
	// a2 := []int{0, 2, 1, 3}
	// a3 := []int{-862003, -857689, -749961, -501097, -478671, -406595, 207173, 535179}
	// a4 := []int{-156032, 98861, 177420, 483189, -614781, -129629, -395885, -128337}
	// a5 := []int{-663257, -377752, 189991, 363200, 417738, 483350, 694143, 833173}

	// result1 := piscine.IsSorted(piscine.F, a1)
	// result2 := piscine.IsSorted(piscine.F, a2)
	// result3 := piscine.IsSorted(piscine.F, a3)
	// result4 := piscine.IsSorted(piscine.F, a4)
	// result5 := piscine.IsSorted(piscine.F, a5)

	// fmt.Println(result1)
	// fmt.Println(result2)
	// fmt.Println(result3)
	// fmt.Println(result4)
	// fmt.Println(result5)

	// piscine.DescendComb()

	// a := []int{1, 2, 3, 1, 2, 3, 4}
	// unmatch := piscine.Unmatch(a)
	// fmt.Println(unmatch)

	// fmt.Println(piscine.FoodDeliveryTime("burger"))
	// fmt.Println(piscine.FoodDeliveryTime("chips"))
	// fmt.Println(piscine.FoodDeliveryTime("nuggets"))
	// fmt.Println(piscine.FoodDeliveryTime("burger") + piscine.FoodDeliveryTime("chips") + piscine.FoodDeliveryTime("nuggets"))

	// steps := piscine.CollatzCountdown(12)
	// fmt.Println(steps)

	// middle := piscine.Abort(2, 3, 8, 5, 7)
	// fmt.Println(middle)

	// summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	// summary := "Burger Burger Water Coffe    e Water Chips Carrot"
	// for index, element := range piscine.ShoppingSummaryCounter(summary) {
	// 	fmt.Println(index, "=>", element)
	// }

	// const N = 6
	// a := make([]string, N)
	// a[0] = "a"
	// a[2] = "b"
	// a[4] = "c"

	// for _, v := range a {
	// 	fmt.Println(v)
	// }

	// fmt.Println("Size after compacting:", piscine.Compact(&a))

	// for _, v := range a {
	// 	fmt.Println(v)
	// }

	// toConcat := []string{"Hello!", " How", " are", " you?"}
	// fmt.Println(piscine.Join(toConcat, ":"))

	// x := 5
	// y := &x
	// z := &y
	// a := &z

	// w := 2
	// b := &w

	// u := 7
	// e := &u
	// f := &e
	// g := &f
	// h := &g
	// i := &h
	// j := &i
	// c := &j

	// k := 6
	// l := &k
	// m := &l
	// n := &m
	// d := &n

	// fmt.Println(***a)
	// fmt.Println(*b)
	// fmt.Println(*******c)
	// fmt.Println(****d)

	// piscine.Enigma(a, b, c, d)

	// fmt.Println("After using Enigma")
	// fmt.Println(***a)
	// fmt.Println(*b)
	// fmt.Println(*******c)
	// fmt.Println(****d)

	// fmt.Println(piscine.DescendAppendRange(10, 5))
	// fmt.Println(piscine.DescendAppendRange(5, 10))

	// slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
	// fmt.Println(piscine.ShoppingListSort(slice))

	// fmt.Println(piscine.ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))

	// p4 := []string{"4th Place"}
	// p3 := []string{"3rd Place"}
	// p2 := []string{"2nd Place"}
	// p1 := []string{"1st Place"}

	// position := [][]string{p4, p3, p2, p1}
	// fmt.Println(piscine.PodiumPosition(position))

	// fmt.Println(piscine.ActiveBits(7))

	// a := []int{23, 123, 1, 11, 55, 93}
	// max := piscine.Max(a)
	// fmt.Println(max)

	// fmt.Print(piscine.LoafOfBread("deliciousbread"))
	// fmt.Print(piscine.LoafOfBread("This is a loaf of bread"))
	// fmt.Print(piscine.LoafOfBread(""))

	link := &piscine.List1{}

	piscine.ListPushBack(link, "Hello")
	piscine.ListPushBack(link, "man")
	piscine.ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}
