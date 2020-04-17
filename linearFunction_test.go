package gosimplex

import (
	"testing"
	"fmt"
	"os"
	"log"
)

func TestNewLinearFunction(t *testing.T){
	fmt.Println("Start TestNewLinearFunction")

	test := NewLinearFunction()

	if test.cofs != nil{
		t.Error("wrong cofs init")
		log.Panic()
	}
	if test.isMin{
		t.Error("wrong cofs init")
		log.Panic()
	}

	fmt.Println("TestNewLinearFunction=[ok]")
}

func TestCofsLF(t *testing.T){
	fmt.Println("Start TestCofsLF")

	cofsR := []float64{2.0,3.5,4.5,0.5}
	test := LinearFunction{cofs:cofsR,}

	checkFlag := true
	for i,v := range test.Cofs(){
		if cofsR[i] != v{
			t.Error("Wrong value:",v," at position:",i," expected:",cofsR[i] )
			checkFlag = false
		}
	}

	cpy := test.Cofs()
	cpy[3] = 0.0

	if test.cofs[3] == 0.0{
		t.Error("Wrong copy return from Cofs()")
		checkFlag = false
	}

	if checkFlag{
		fmt.Println("TestCofsLF=[ok]")
	}
}

func TestIsMin(t *testing.T){
	fmt.Println("Start TestIsMin")

	test := LinearFunction{isMin:true,}

	if !test.IsMin(){
		t.Error("Wrong return from IsMin")
	}else{
		fmt.Println("TestIsMin=[ok]")
	}

}

func TestInitLF(t *testing.T){
	fmt.Println("Start TestInitLF")

	test := NewLinearFunction()

	cofs := []float64{0.0,1.1,2.2,3.3}
	isMin := true

	test.Init(cofs,isMin)

	checkFlag := true

	for i,v := range test.Cofs(){
		if cofs[i] != v{
			t.Error("Wrong value:",v," at position:",i," expected:",cofs[i] )
			checkFlag = false
		}
	}

	if !test.IsMin(){
		t.Error("Wrong return from IsMin")
		checkFlag = false
	}

	cofs[3] = 0.0

	if test.cofs[3] == 0.0{
		t.Error("Wrong copy work with cofs")
		checkFlag = false
	}

	if checkFlag{
		fmt.Println("TestInitLF=[ok]")
	}
}

func TestValueLF(t *testing.T){
	fmt.Println("Start TestValueLF")

	cofsR := []float64{2.0,3.5,4.5,0.5}
	test := LinearFunction{cofs:cofsR,}

	checkFlag := true

	if test.Value() != 0.5{
		t.Error("Wrong value")
		checkFlag = false
	}

	if checkFlag{
		fmt.Println("TestValueLF=[ok]")
	}
}

func TestAmountOfCofsLF(t *testing.T){
	fmt.Println("Start TestAmountOfCofsLF")

	test := LinearFunction{cofs:[]float64{1.0,0.0,0.0},isMin:true}

	if test.AmountOfCofs() != 2{
		t.Error("Wrong amount of cofs:",test.AmountOfCofs()," expected:2")
	}else{
		fmt.Println("TestAmountOfCofsLF=[ok]")
	}
}

func TestCopyLF(t *testing.T){
	fmt.Println("Start TestCopyLF")

	test := NewLinearFunction()

	cofs := []float64{0.0,1.1,2.2,3.3}
	isMin := true

	test.Init(cofs,isMin)

	cpy := test.Copy()

	checkFlag := true

	if cpy.isMin != test.isMin{
		t.Error("Wrong copy of isMin")
		checkFlag = false
	}

	for i,v := range test.cofs{
		if cpy.cofs[i] != v{
			t.Error("Wrong cof:",cpy.cofs[i]," at position:",i," expected:",v)
			checkFlag = false
		}
	}

	cpy.cofs[2] = 0.0

	if test.cofs[2] == 0.0{
		t.Error("Original was changed by copy")
		checkFlag = false
	}

	if checkFlag{
		fmt.Println("TestCopyLF=[ok]")
	}
}

func TestCountFunction(t *testing.T){
	fmt.Println("Start TestCountFunction")

	test := LinearFunction{cofs:[]float64{1.0,2.0,0.0},isMin:true}

	res,err := test.CountFunction([]float64{-2.0,3.0})
	if err != nil{
		log.Panic(err)
	}
	if res != 4.0{
		t.Error("Wrong function result:",res," expected:4.0")
	}else{
		fmt.Println("TestCountFunction=[ok]")
	}
}

func TestReadFunctionFromBuffer(t *testing.T){
	
	fmt.Println("Start TestReadFunctionFromBuffer")

	sub := os.Stdout

	var err error
	os.Stdout,err = os.Create("testing/TestReadFunctionFromBuffer.out")
	if err != nil{
		log.Panic(err)
	}

	file,err := os.Open("testing/ReadFunctionFromBuffer")
	if err != nil{
		log.Panic(err)
	}

	test := NewLinearFunction()
	test.ReadFunctionFromBuffer(file)

	os.Stdout = sub

	cofs := []float64{0.3,-0.6,3.0,6.0}

	checkFlag := true
	for i,v := range test.cofs{
		if v != cofs[i]{
			t.Error("Wrong value:",v," at position:",i," expected:",cofs[i])
			checkFlag = false
		}
	}

	if !test.isMin{
		t.Error("Wrong isMin value")
		checkFlag = false
	}

	if checkFlag {
		fmt.Println("TestReadFunctionFromBuffer=[ok]")
	}
}