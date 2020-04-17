package gosimplex

import ( 
	"strings"
	"runtime"
)

var NEXTLINE string

func initNEXTLINE(){
	if runtime.GOOS == "windows"{
		NEXTLINE = "\r\n"
	}else{
		NEXTLINE = "\n"
	}
}

func subStringBefore(str string, character string)string{
	pos := strings.Index(str,character)
	if pos == -1{
		return str
	}
	return str[0:pos]
}