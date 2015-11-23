package testing


import (
	"testing"
)

var addTable = []struct {
	in []int
	out int
}{
	{[]int{1,2},3},
	{[]int{1,2,3,4},10},
}


func TestCadnAddNumbers(t *testing.T){
	t.Parallel()
	for _, entry := range addTable {
		result := Add(entry.in...)
		if result != entry.out {
			t.Error("Failed to add nubers as expected")
		}
	}





	result := Add(1,2)
//    time.Sleep(10 * time.Second)
	if result != 3 {
		t.Log("Failed to add")
		t.Fail()
	}

	result = Add(1,2,3,4)
	if result != 10 {
		t.Error("failed to run")
	}
}

func TestSubtractNumbers(t *testing.T){
	result := Subtract(10,5)

	if result != 5 {
		t.Log("Failed to subtract")
		t.Fail()
	}
}

func TestMultiplication(t *testing.T){
	t.Skip("not implamented")
}