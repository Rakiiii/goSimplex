package gosimplex


//NelderMidAlgorithm describe constatnts and start method for Nelder-Mid simplex-method
type NelderMidAlgorithm struct{
	alpha,beta,gamma float64
}

//NewNelderMidAlgorithm return void NelderMidAlgorithm
func NewNelderMidAlgorithm()*NelderMidAlgorithm{
	return &NelderMidAlgorithm{ alpha:0.0,
								beta:0.0,
								gamma:0.0}
}


//Init initialized Nelder-Mid constatnts
func (n *NelderMidAlgorithm)Init(a,b,g float64){
	n.alpha = a
	n.beta = b
	n.gamma = g
}

//DefaultInit default initialization for neledr mid algortihm constants
func (n *NelderMidAlgorithm)DefaultInit(){
	n.alpha = 1.0
	n.beta = 0.5
	n.gamma = 2.0
}

//Run start Nelder-Mid optimisation method for system @system with start simplex initialized with @sg and amount of itteration setted with @cond
//returns Vector of found optimum
func (n *NelderMidAlgorithm)Run(system Systemer,sg SimplexGenerator,cond Conditioner)*Vector{
	//init simplex
	simplex := sg.GenerateSimplex(system)

	//if condition is ok
	for cond.CheckCondition(){
		//set best order
		//best value is last,worst value at position 0
		simplex = QuicksortFuncPair(simplex)
		if system.IsMin(){
			simplex = Reverse(simplex)
		}

		var newVector *Vector = nil
		var newResult float64 = 0.0

		//found gravity center
		mid := foundGravityCenter(simplex[1:])

		//do reflection operation
		reflectedValue := mid.TimesFloat( 1 + n.alpha ).Sub( simplex[0].Value.TimesFloat( n.alpha ) )
		reflectedResult,err := system.Func().Count(reflectedValue)
		handleFatalError(err)

		//if reflection if ok
		if (reflectedResult < simplex[0].Result && system.IsMin()) || (reflectedResult > simplex[0].Result && !system.IsMin()){
			//try to do expansion
			expensionValue := mid.Add( reflectedValue.Sub(mid).TimesFloat( n.gamma ) )
			expensionResult,err := system.Func().Count( expensionValue )
			handleFatalError(err)

			//if expension is ok then move expended vector to simplex
			if (expensionResult < simplex[0].Result && system.IsMin()) || (expensionResult > simplex[0].Result && !system.IsMin()){
				newVector = expensionValue
				newResult = expensionResult
			}else{
				//otherwise move reflected value to simplex
				newVector = reflectedValue
				newResult = reflectedResult
			}

		}else{
			//if reflection is not working start contart
			contratcValue := mid.Add( simplex[0].Value.Sub( mid ).TimesFloat( n.beta ) )
			contractResult,err := system.Func().Count( contratcValue)
			handleFatalError(err)

			//if contartc is ok move contracted vector to simplex
			if (contractResult < simplex[ len(simplex)-1 ].Result && system.IsMin()) || (contractResult > simplex[ len(simplex)-1 ].Result && !system.IsMin()){
				newVector = contratcValue
				newResult = contractResult
			}
		}

		//if condtracting or reflection isn't working then start shrinking operation
		if newVector == nil{
			simplex = n.shrinkSimplex(simplex,system)

		}else{
			//move vector to simplex
			simplex[0].Value = newVector
			simplex[0].Result = newResult
		}
	}

	//found optimum
	simplex = QuicksortFuncPair(simplex)
	if system.IsMin(){
		simplex = Reverse(simplex)
	}

	
	return simplex[len(simplex)-1].Value
}

//foundGravityCenter returns vector that describe gravity center of simplex
func foundGravityCenter(simplex []FuncPair)*Vector{
	result := simplex[0].Value 
	for i := 1; i < len(simplex); i++{
		result = result.Add(simplex[i].Value)
	}
	return result.DivFloat(float64(len(simplex)))
}
//shrinkSimplex returns shrinked simplex
func (n *NelderMidAlgorithm)shrinkSimplex(simplex []FuncPair,system Systemer)[]FuncPair{
	newSimplex := make([]FuncPair,len(simplex))
	var err error
	for i := 0; i < len(simplex)-1 ; i++{
		newSimplex[i].Value = simplex[ len(simplex)-1 ].Value.Add( simplex[i].Value.Sub(simplex[ len(simplex)-1 ].Value).TimesFloat(n.beta))
		newSimplex[i].Result,err = system.Func().Count(newSimplex[i].Value)
		handleFatalError(err)
	}
	newSimplex[ len(simplex)-1 ].Value = simplex[ len(simplex)-1 ].Value
	newSimplex[ len(simplex)-1 ].Result = simplex[ len(simplex)-1 ].Result
	
	return newSimplex
}

