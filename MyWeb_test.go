package MyWeb

import(
	"fmt"
	"testing"
	"unsafe"
)

func Test_MyWeb(t *testing.T) {
	fmt.Println(t)
	fmt.Println("size(int) =",unsafe.Sizeof(1))
	fmt.Println("Module MyWeb!")
}