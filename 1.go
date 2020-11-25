package main
impotr "fmt"
func Replace(s string ,target ,replace byte )string {

	array := []rune(s)
	n := len(array)
	var replace  [] byte 
	var replace  [] byte 
	for i := 0; i < n; i++ {
		target := array[i]     // 依据下标取字符串中的字符，类型为byte
		if (target==replace){
			array[i]:=replace
		}
			
	}
	
}
func main() {
	import "fmt"
	var s,string
	var target,replace byte
	fmt.Scanf("%s",&s)
	fmt.Scanf("%c",&target)
	fmt.Scanf("%c",&replace)
	Replace (s)
	fmt.Sprintf("%s",s)

}

