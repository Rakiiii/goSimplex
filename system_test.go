package gosimplex

import (
	"testing"
	"bufio"
	"fmt"
	"os"
	"log"
	"errors"
)

func TestNewSystem(t *testing.T){
	fmt.Println("Start TestNewSystem")

	test := NewSystem()

	checkFlag := true

	if test.optimisationFunction.cofs != nil{
		t.Error("Wrong optfuntion init in system")
		checkFlag = false
	} 

	if test.restrictions != nil{
		t.Error("Wrong restrictions init in system")
		checkFlag = false
	}

	if checkFlag{
		fmt.Println("TestNewSystem=[ok]")
	}
}

func TestInitS(t *testing.T){
	fmt.Println("Start TestInitS")
	checkFlag := true

	ts := NewSystem()
	tf := LinearFunction{cofs:[]float64{1.0,1.0,1.0},isMin:true}
	tr := []LinearRestriction{LinearRestriction{cofs:[]float64{1.0,1.0,1.0},value:1.0,sign:More}}
	if err := ts.Init(tf,tr); errors.Is(err,CofsError){
		t.Error("Wrong error:",err," expected:",CofsError)
		checkFlag = false
	}

	fcofs := []float64{0.3,-0.6,3.0,6.0}
	fisMin := true
	r1cofs := []float64{0.0,1.1,2.2}
	r1value := 4.0
	r1sign := More
	r2cofs := []float64{1.1,3.3,0.0}
	r2value := 0.2
	r2sign := Less

	f := NewLinearFunction()
	f.Init(fcofs,fisMin)

	r1 := NewLinearRestriction()
	r1.Init(r1cofs,r1value,r1sign)
	r2 := NewLinearRestriction()
	r2.Init(r2cofs,r2value,r2sign)
	rs := []LinearRestriction{r1,r2}

	s := NewSystem()
	if err := s.Init(f,rs);err != nil{
		t.Error("Init error:",err)
		checkFlag = false
		log.Panic()
	}

	for i,v := range s.Function().cofs{
		if v != fcofs[i]{
			t.Error("Wrong cofs:",v," at position:",i," expected:",fcofs[i])
			checkFlag = false
		}
	}

	if s.optimisationFunction.AmountOfCofs() != 3{
		t.Error("Wrong amount of cofs in function:",s.optimisationFunction.AmountOfCofs()," expected:3")
		checkFlag = false
	}

	if !s.optimisationFunction.isMin{
		t.Error("Wrong isMin value in function:false expected:true")
		checkFlag = false
	}

	if s.AmountOfRestrictions() != 2{
		t.Error("Wrong amount of restrictions:",s.AmountOfRestrictions()," expected:2")
		checkFlag = false
	}

	for i := 0; i < s.AmountOfRestrictions();i ++{
		if s.restrictions[i].AmountOfCofs() != rs[i].AmountOfCofs(){
			t.Error("Wrong amount of cofs:",s.restrictions[i].AmountOfCofs()," in restriction:",i," expected:",rs[i].AmountOfCofs())
			checkFlag = false
		}
		for j,v := range s.restrictions[i].cofs{
			if v != rs[i].cofs[j]{
				t.Error("Wrong value:",v," in restriction:",i," at position:",j," expected:",rs[i].cofs[j])
				checkFlag = false
			}
		}
	}

	if s.restrictions[0].sign != r1sign{
		t.Error("Wrong singum at 1 restriction:",s.restrictions[0].sign," expected:",r1sign)
		checkFlag = false
	}
	if s.restrictions[1].sign != r2sign{
		t.Error("Wrong singum at 2 restriction:",s.restrictions[1].sign," expected:",r2sign)
		checkFlag = false
	}

	if s.Restriction(0).Value() != r1value{
		t.Error("Wrong value at 1 restriction:",s.Restriction(0).Value()," expected:",r1value)
		checkFlag = false
	}
	if s.Restriction(1).Value() != r2value{
		t.Error("Wrong value at 2 restriction:",s.Restriction(1).Value()," expected:",r2value)
		checkFlag = false
	}


	if checkFlag{
		fmt.Println("TestInitS=[ok]")
	}
}

func TestSystemFromBuffer(t *testing.T){
	fmt.Println("Start TestSystemFromBuffer")

	sub := os.Stdout

	var err error
	os.Stdout,err = os.Create("testing/TestRestrictionFromBuffer.out")
	if err != nil{
		log.Panic(err)
	}

	checkFlag := true

	fcofs := []float64{0.3,-0.6,3.0,6.0}
	fisMin := true
	r1cofs := []float64{1.0,-2.0,0.0}
	r1value := 0.0
	r1sign := Less
	r2cofs := []float64{0.0,1.0,0.0}
	r2value := 0.0
	r2sign := More

	f := NewLinearFunction()
	f.Init(fcofs,fisMin)

	r1 := NewLinearRestriction()
	r1.Init(r1cofs,r1value,r1sign)
	r2 := NewLinearRestriction()
	r2.Init(r2cofs,r2value,r2sign)
	rs := []LinearRestriction{r1,r2}

	s := NewSystem()
	file,err := os.Open("testing/ReadSystemFromBuffer")
	if err != nil{
		log.Panic(err)
	}
	reader := bufio.NewReader(file)
	s.ReadSystemFromBuffer(reader)

	for i,v := range s.Function().cofs{
		if v != fcofs[i]{
			t.Error("Wrong cofs:",v," at position:",i," expected:",fcofs[i])
			checkFlag = false
		}
	}

	if s.optimisationFunction.AmountOfCofs() != 3{
		t.Error("Wrong amount of cofs in function:",s.optimisationFunction.AmountOfCofs()," expected:3")
		checkFlag = false
	}

	if !s.optimisationFunction.isMin{
		t.Error("Wrong isMin value in function:false expected:true")
		checkFlag = false
	}

	if s.AmountOfRestrictions() != 2{
		t.Error("Wrong amount of restrictions:",s.AmountOfRestrictions()," expected:2")
		checkFlag = false
	}

	for i := 0; i < s.AmountOfRestrictions();i ++{
		if s.restrictions[i].AmountOfCofs() != rs[i].AmountOfCofs(){
			t.Error("Wrong amount of cofs:",s.restrictions[i].AmountOfCofs()," in restriction:",i," expected:",rs[i].AmountOfCofs())
			checkFlag = false
		}
		for j,v := range s.restrictions[i].cofs{
			if v != rs[i].cofs[j]{
				t.Error("Wrong value:",v," in restriction:",i," at position:",j," expected:",rs[i].cofs[j])
				checkFlag = false
			}
		}
	}

	if s.restrictions[0].sign != r1sign{
		t.Error("Wrong singum at 1 restriction:",s.restrictions[0].sign," expected:",r1sign)
		checkFlag = false
	}
	if s.restrictions[1].sign != r2sign{
		t.Error("Wrong singum at 2 restriction:",s.restrictions[1].sign," expected:",r2sign)
		checkFlag = false
	}

	if s.Restriction(0).Value() != r1value{
		t.Error("Wrong value at 1 restriction:",s.Restriction(0).Value()," expected:",r1value)
		checkFlag = false
	}
	if s.Restriction(1).Value() != r2value{
		t.Error("Wrong value at 2 restriction:",s.Restriction(1).Value()," expected:",r2value)
		checkFlag = false
	}

	os.Stdout = sub

	if checkFlag{
		fmt.Println("TestSystemFromBuffer=[ok]")
	}
}