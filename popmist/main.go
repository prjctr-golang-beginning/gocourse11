package main

// 1. Entry in nil Map
// 2. Nil Pointer Dereference
// 3. Multiple-value in single-value context
// 4. Unchangeable Array values
// 5. Shadow variable
// 6. Unexpected new-line
// 7. Unaltered strings
// 8. Print favorite band name ABBA
// 9. Missing Copy
// 10. Append issue
// 11. Unexpected ++

func main() {
	//a1 := [2]int{1, 2}
	//a2 := []int{1, 2}

	//Change1(a1)
	//Change2(a2)

	//fmt.Println(a1)
	//fmt.Println(a2)

	//i := 0
	//if i == 0 {
	//	i := 1
	//	i++
	//}
	//fmt.Println(i)

	//s2 := []string{
	//	`one`,
	//	`two`,
	//	`three`,
	//}
	//fmt.Println(s2)

	//str100 := []rune{'H', 'e', 'l', 'l', 'o'}
	//str100[0] = 'h'
	//fmt.Println(string(str100))

	//res := strings.TrimRight(`ABBA`, `BA`)
	//res := strings.Trim(`  ABBA `, ` `)
	//fmt.Println(res)

	//var src, dst []string
	//src := []string{`A`, `B`, `C`}
	//dst := make([]string, len(src))
	//copy(dst, src)
	//fmt.Println(dst)

	//i := 0
	//fmt.Println(i++)
	//fmt.Println(++i)

	//b := []byte{'a', 'b'}
	//b1 := append(b, 'c')
	//b2 := append(b, 'd')
	//fmt.Println(string(b1)) // abc
	//fmt.Println(string(b2)) // abd
}

func Change1(a [2]int) {
	a[0] = 10
}

func Change2(a []int) {
	a[0] = 10
}
