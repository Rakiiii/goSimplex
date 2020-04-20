package gosimplex

import ( 
	"testing"
	"fmt"
	//"bufio"
	//"os"
	//"log"
	"errors"
)

type testfunc struct{
	c int
}
func (t *testfunc)Count(v *Vector)(float64,error){
	if len(*v) != 2{
		return 0.0,errors.New("Wrong amount of cofs")
	}else{
		t.c++
		return (*v)[0]*(*v)[0] + (*v)[0]*(*v)[1] + (*v)[1]*(*v)[1] - 6.0*(*v)[0] - 9.0*(*v)[1],nil
	}
}

type testsimplexgenerator struct{}

func (t testsimplexgenerator)GenerateSimplex(Systemer)[]FuncPair{
	return []FuncPair{FuncPair{Value:&Vector{0.0,0.0},Result:0},FuncPair{Value:&Vector{1.0,0.0},Result:-5},FuncPair{Value:&Vector{0.0,1.0},Result:-8}}
}

func (t *testfunc)CalculationAmount()int{
	return t.c
}

func TestRunNM(t *testing.T){
	fmt.Println("Start TestRunNM")

	checkFlag := true

	system := NewGenerializedSystem()
	system.Init(&testfunc{},VoidRestriction{},true,2)

	alg := NewNelderMidAlgorithm()
	alg.DefaultInit()

	simplexGenerator := testsimplexgenerator{}
	cond := NewItterationConditioner()
	cond.Init(10)

	res := alg.Run(system,simplexGenerator,cond)
	f := testfunc{}
	v,_ :=f.Count(res)
	
	resVec := Vector{0.9957275390625,3.9764404296875}
	resFunc := -20.999326035380363

	if !res.Cmp(&resVec){
		t.Error("Wrong vector:",res.String()," expected:",resVec.String())
		checkFlag = false
	}
	if v != resFunc{
		t.Error("Wrong func value:",v," expected:",resFunc)
		checkFlag = false
	}

	if checkFlag {
		fmt.Println("TestRunNM=[ok]")
	}
}