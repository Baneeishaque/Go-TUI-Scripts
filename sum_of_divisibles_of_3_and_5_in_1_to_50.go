package main
import (
	"fmt"
)
func main(){
	sum :=0
	for i:=1;i<50;i++{
		//fmt.Printf("%v\n",i)
		if i%3==0 || i%5==0{
			//fmt.Printf("%v\n",i)
			sum=sum+i
		}
	}
	fmt.Printf("Sum is %v",sum)
}