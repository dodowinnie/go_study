package cal
import (
	// "fmt"
	"testing"
)

func TestGetSub (t *testing.T){
	res := getSub(10,3)
	
	t.Logf("test result=%v", res)
}