package gosimplex

import "log"

//handleFatalError wrap handling fatal erros
func handleFatalError(err error){
	if err != nil{
		log.Panic(err)
	}
}