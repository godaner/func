package shell
import "fmt"

func main() {
	s := []int{88, 1, 55, 45, 4546, 564, 68, 13, 1, 3654, 896}
	s = Shell(s,len(s)/3)
	fmt.Println(s)
}

// Shell
func Shell(nums []int, dk int) (res []int) {
	for { // 计算dk
		dk = dk / 2
		for i := 0; i < dk; i++ { // 循环各个分组
			for j := i; j < len(nums); j = j + dk {//循环每个分组

			}
		}
		if dk == 1 {
			return nums
		}
	}

}
