package gosimplex

//Systemer interface representing optimisation system
type Systemer interface{
	Func()GenerializedFunctioner
	Restrict() GenerializedRestrictioner
	IsMin() bool
	Size() int
}


//GenerializedSystem describe generialized optimisation system with any type of restrictions and functions
type GenerializedSystem struct{
	fun GenerializedFunctioner
	restrict GenerializedRestrictioner
	isMin bool
	size int
}


//NewGenerializedSystem returns void GenerializedSystem
func NewGenerializedSystem()*GenerializedSystem{
	return &GenerializedSystem{fun:nil,restrict:nil,isMin:false,size:0}
}

//Init initialize generialized optimization system
func (s *GenerializedSystem)Init(f GenerializedFunctioner,r GenerializedRestrictioner,b bool,size int){
	s.fun = f
	s.restrict = r 
	s.isMin = b
	s.size = size
}

func (s *GenerializedSystem)Func()GenerializedFunctioner{
	return s.fun
}

func (s *GenerializedSystem)Restrict()GenerializedRestrictioner{
	return s.restrict
}

func (s *GenerializedSystem)IsMin()bool{
	return s.isMin
}

func (s *GenerializedSystem)Size()int{
	return s.size
}


