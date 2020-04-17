package gosimplex

import (
	"fmt"
	"log"
	"testing"
	"os"
)

func TestInitSignum(t *testing.T){
	fmt.Println("Start TestInitSignum")

	str := []string{">","<","=",">=","<=","=="}
	checkFlag := true

	if InitSignum(str[0]) != More{
		t.Error("Wrong init for more:",InitSignum(str[0]))
		checkFlag = false
	}

	if InitSignum(str[1]) != Less{
		t.Error("Wrong init for less:",InitSignum(str[1]))
		checkFlag = false
	}

	if InitSignum(str[2]) != Equals{
		t.Error("Wrong init for equals:",InitSignum(str[2]))
		checkFlag = false
	}

	if InitSignum(str[3]) != MoreOrEquals{
		t.Error("Wrong init for MoreOrEquals:",InitSignum(str[3]))
		checkFlag = false
	}

	if InitSignum(str[4]) != LessOrEquals{
		t.Error("Wrong init for LessOrEquals:",InitSignum(str[4]))
		checkFlag = false
	}

	if InitSignum(str[5]) != WrongSignum{
		t.Error("Wrong init for WrongSignum:",InitSignum(str[5]))
		checkFlag = false
	}


	if checkFlag{
		fmt.Println("TestInitSignum=[ok]")
	}
}

func TestNewLinearRestriction(t *testing.T){
	fmt.Println("Start TestNewLinearRestriction")

	test := NewLinearRestriction()

	if test.sign != WrongSignum || test.cofs != nil || test.value != 0.0{
		t.Error("Wrong Init")
	}else{
		fmt.Println("NewLinearRestriction=[ok]")
	}
}

func TestInitLR(t *testing.T){
	fmt.Println("Start TestInitLR")

	test := NewLinearRestriction()

	cofs := []float64{0.0,1.1,2.2}
	test.Init(cofs,5.5,InitSignum("<="))

	checkFlag := true

	for i,v := range test.cofs{
		if v != cofs[i]{
			t.Error("Wrong value:",v," at position:",i," expected:",cofs[i])
			checkFlag = false
		}
	}

	cofs[0] = 3.0
	if test.cofs[0] == 3.0{
		t.Error("Copy changed with original")
		checkFlag = false
	}

	if test.value != 5.5{
		t.Error("Wrong value:",test.value," expected:5.5")
		checkFlag = false
	}

	if test.sign != LessOrEquals{
		t.Error("Wrong sign:",test.sign," expected:",LessOrEquals)
	}

	if checkFlag{
		fmt.Println("TestInitLR=[ok]")
	}
}

func TestCopyLR(t *testing.T){
	fmt.Println("Start TestCopyLR")

	test := NewLinearRestriction()

	cofs := []float64{0.0,1.1,2.2}
	test.Init(cofs,5.5,InitSignum("<="))

	cpy := test.Copy()

	checkFlag := true

	for i,v := range cpy.cofs{
		if v != test.Cofs()[i]{
			t.Error("Wrong value:",v," at position:",i," expected:",test.Cofs()[i])
			checkFlag = false
		}
	}

	cpy.cofs[0] = 3.0
	if test.cofs[0] == cpy.cofs[0]{
		t.Error("Copy changed with original")
		checkFlag = false
	}

	if test.value != cpy.value{
		t.Error("Wrong value:",cpy.value," expected:",test.value)
		checkFlag = false
	}

	if test.sign != cpy.sign{
		t.Error("Wrong sign:",cpy.sign," expected:",test.sign)
	}

	if checkFlag{
		fmt.Println("TestCopyLR=[ok]")
	}
}

func TestReadRestrictionFromBuffer(t *testing.T){
	
	fmt.Println("Start TestReadRestrictionFromBuffer")

	sub := os.Stdout

	var err error
	os.Stdout,err = os.Create("testing/TestRestrictionFromBuffer.out")
	if err != nil{
		log.Panic(err)
	}

	file,err := os.Open("testing/ReadRestrictionFromBuffer")
	if err != nil{
		log.Panic(err)
	}

	test := NewLinearRestriction()
	test.ReadRestrictionFromBuffer(4,file)

	os.Stdout = sub

	checkFlag := true

	cofs := []float64{3.5,-4.0,2.0,0.0}
	value := 1.0
	
	if test.Sign() != LessOrEquals{
		t.Error("Wrong signum:",test.Sign()," expected:",LessOrEquals)
		checkFlag = false
	}

	if test.Value() != value{
		t.Error("Wrong value:",test.Value()," expected:",value)
		checkFlag = false
	}

	if test.AmountOfCofs() != 4{
		t.Error("Wrong amount of cofs:",test.AmountOfCofs()," expected:",4)
	}

	for i,v := range test.Cofs(){
		if v != cofs[i]{
			t.Error("Wrong cof:",v," at position:",i," expected:",cofs[i])
			checkFlag = false
		}
	}

	if checkFlag{
		fmt.Println("TestReadRestrictionFromBuffer=[ok]")
	}
}