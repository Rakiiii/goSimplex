package gosimplex

import (
	"testing"
	"runtime"
	"fmt"
)

func TestInitNEXTLINE(t *testing.T){
	fmt.Println("Start TestInitNEXTLINE")

	initNEXTLINE()

	if runtime.GOOS == "windows" && NEXTLINE != "\r\n"{
		t.Error("Wrong windows init")
	} 
	if runtime.GOOS != "windows" && NEXTLINE != "\n"{
		t.Error("Wrong normal os init")
	}
	fmt.Println("TestInitNEXTLINE=[ok]")
}

func TestSubStringBefore(t *testing.T){
	fmt.Println("Start TestSubstringBefore")

	str := "kjhln\njonj\n"
	right := "kjhln"

	if right != subStringBefore(str,"\n"){
		t.Error("Wrong substring:",subStringBefore(str,"\n")," expected:",right)
	}else{
		fmt.Println("TestSubStringBefore=[ok]")
	}
}