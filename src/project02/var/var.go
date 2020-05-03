package main
import "fmt"
import "project02/test"

// 全局变量
var n1 = 1
var (
	n2 = 1
	n3 = 2
)

func main() {
	// 变量定义
	var i, i1, i2 int
	var j, j1, j2 = 10, 11, 12
	k, k1, k2 := 10, 11, 12

	fmt.Println("i=", i, "i1=", i1, "i2=", i2)
	fmt.Println("j=", j, "j1=", j1, "j2=", j2)
	fmt.Println("k=", k, "k1=", k1, "k2=", k2)
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	{
		var i = 10
		fmt.Println("i=", i)
	}
	fmt.Println("test=", test.Test)
}
