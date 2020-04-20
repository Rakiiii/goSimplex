package gosimplex
//import "fmt"

//SimplexGenerator interface for generating simplex
type SimplexGenerator interface{
	//GenerateSimplex must return start simplex for @GenerializedSystem
	GenerateSimplex(Systemer)[]FuncPair
}

//RandomGenerator wrap for generationg random simplex
type RandomGenerator struct{}

//GenerateSimplex implemants SimplexGenerator interface with random simplex set
func (r RandomGenerator)GenerateSimplex(system Systemer)[]FuncPair{
	n := system.Size()+1
	simplex := make([]FuncPair,n)

	for i,_ := range simplex{
		vec := NewRandomVector(n-1)
		isOk,err := system.Restrict().Check(vec)
		handleFatalError(err)
		for !isOk || CheckVectorContainment(vec,simplex){
			vec = NewRandomVector(n-1)
			isOk,err = system.Restrict().Check(vec)
			handleFatalError(err)
		} 
		simplex[i].Value = vec
		simplex[i].Result,err = system.Func().Count(vec)
		handleFatalError(err)
	}
	return simplex
}