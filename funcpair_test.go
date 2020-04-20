package gosimplex

import (
	"testing"
	"fmt"
)

func TestReverse(t *testing.T){
	fmt.Println("Start TestReverse")

	ar := []FuncPair{FuncPair{Result:0.0,Value:NewUnitVector(2)},FuncPair{Result:1.0,Value:NewUnitVector(2)},FuncPair{Result:2.0,Value:NewUnitVector(2)},FuncPair{Result:3.0,Value:NewUnitVector(2)}}
	rev := []FuncPair{FuncPair{Result:3.0,Value:NewUnitVector(2)},FuncPair{Result:2.0,Value:NewUnitVector(2)},FuncPair{Result:1.0,Value:NewUnitVector(2)},FuncPair{Result:0.0,Value:NewUnitVector(2)}}

	test := Reverse(ar)

	checkFlag := true
	for i,v := range rev{
		if v.Result != test[i].Result || !v.Value.Cmp(test[i].Value){
			t.Error("Wrong value:["+test[i].Value.String(),",",test[i].Result,"] at position:",i," expected:["+test[i].Value.String(),",",test[i].Result,"]")
			checkFlag = false
		}
	}

	if checkFlag{
		fmt.Println("TestReverse=[ok]")
	}
}