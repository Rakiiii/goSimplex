package gosimplex

import (
	"testing"
	"fmt"
	"log"
)

func TestNewZeroVector(t *testing.T){
	fmt.Println("Start TestNewZeroVector")

	v := NewZeroVector(3)
	if len(*v) != 3{
		t.Error("Wrong vector size")
		log.Panic()
	}

	for _,j := range *v{
		if j != 0.0{
			t.Error("Wrong value:",j)
		}
	}

	fmt.Println("TestNewZeroVector=[ok]")
}

func TestNewUnitVector(t *testing.T){
	fmt.Println("Start TestNewUnitVector")


	v := NewUnitVector(3)
	if len(*v) != 3{
		t.Error("Wrong vector size")
		log.Panic()
	}

	for _,j := range *v{
		if j != 1.0{
			t.Error("Wrong value:",v)
		}
	}

	fmt.Println("TestNewUnitVector=[ok]")
}

func TestVectorOpperations(t *testing.T){
	fmt.Println("Start TestVectorOpperations")

	checkFlag := true

	v1 := NewUnitVector(3)
	v2 := NewUnitVector(3)

	v3 := v1.Add(v2)

	if len(*v3) != 3{
		t.Error("Wrong length after add")
		checkFlag = false
	}
	for _,f := range *v3{
		if f != 2.0{
			t.Error("Wrong value after plus:",f)
			checkFlag = false
		}
	}

	v4 := v1.Sub(v2)
	if len(*v4) != 3{
		t.Error("Wrong length after sub")
		checkFlag = false
	}
	for _,f := range *v4{
		if f != 0.0{
			t.Error("Wrong value after sub:",f)
			checkFlag = false
		}
	}

	v5 := v3.DivFloat(2.0)
	if len(*v5) != 3{
		t.Error("Wrong length after DivFloat")
		checkFlag = false
	}
	for _,f := range *v5{
		if f != 1.0{
			t.Error("Wrong value after DivFloat:",f)
			checkFlag = false
		}
	}

	v6 := v3.TimesFloat(2.0)
	if len(*v6) != 3{
		t.Error("Wrong length after TimesFloat")
		checkFlag = false
	}
	for _,f := range *v6{
		if f != 4.0{
			t.Error("Wrong value after TimesFloat:",f)
			checkFlag = false
		}
	}

	if checkFlag{
		fmt.Println("TestVectorOpperations=[ok]")
	}
}

func TestString(t *testing.T){
	fmt.Println("Start TestString")

	v := NewUnitVector(4)
	str := "(1,1,1,1)"

	if v.String() != str{
		t.Error("Wrong string return:",v.String()," expected:",str)
	}else{
		fmt.Println("TestString=[ok]")
	}
}
