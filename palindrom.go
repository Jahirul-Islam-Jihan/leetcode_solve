package main

import "fmt"

func isPalindrome(x int) bool {
	var temp = x
	 reverse:= 0
	 var digit int
	
	for temp != 0{
		digit = temp % 10
		reverse = (reverse * 10) + digit
		temp=temp/10
	}
	if(reverse==x && x>=0){
		return true
	}else{
		return false
	}
		
	}
	func main(){
		fmt.Println(isPalindrome(121))
	}