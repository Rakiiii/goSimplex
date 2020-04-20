package gosimplex

import (
	"testing"
	"os"
	"log"
	"fmt"
	"bufio"
	"io"
)

func TestRun(t *testing.T){
	fmt.Println("Start TestRun")
	file,err := os.Open("testing/TestRun")
	if err != nil{
		log.Panic(err)
	}

	checkFlag := true

	syst := NewSystem()

	s := os.Stdout

	os.Stdout,err = os.Create("testing/TestRestrictionFromBuffer.out")
	if err != nil{
		log.Panic(err)
	}
	

	syst.ReadSystemFromBuffer(bufio.NewReader(file))
	os.Stdout = s

	simplex := NewSimplex()

	simplex.Init(syst)

	result,_ := simplex.Run()

	sub := os.Stdout
	os.Stdout,err = os.Create("testing/TestRunTableRes")
	if err != nil{
		log.Panic(err)
	}

	simplex.PrintTableToFmt()

	os.Stdout = sub

	if !compareFiles("testing/TestRunTable","testing/TestRunTableRes"){
		checkFlag = false
		t.Error("Wrong result table")
	}
	

	for i,v := range result{
		if v != 10.0{
			t.Error("Wrong result:",v," at position:",i," expected:10.0",)
			checkFlag = false
		}
	}

	if checkFlag{
		fmt.Println("TestRun=[ok]")
	}
}

func compareFiles(file1 string,file2 string)bool{
	str := 0
	f1,err := os.Open(file1)
	if err != nil{
		log.Panic(err)
	}
	defer f1.Close()
	f2,err := os.Open(file2)
	if err != nil{
		log.Panic(err)
	}
	defer f2.Close()
	r1 := bufio.NewReader(f1)
	r2 := bufio.NewReader(f2)

	for err != io.EOF{
		s1,err1 := r1.ReadString('\n')
		s2,err2 := r2.ReadString('\n')
		if err1 != nil{
			err = err1
		}
		if err2 != nil{
			err = err2
		}
		if s1 != s2{
			fmt.Println("Wrong string:",s1,"| expected:",s2,"| at position:",str)
			return false
		}
		str++
	}
	return true
}