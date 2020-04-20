package gosimplex

//GenerializedFunctioner describes generialized function
type GenerializedFunctioner interface{
	//Count must return value of function in point @Vector
	Count(*Vector)(float64,error)
	//CalculationAmount must return amount of function calculation
	CalculationAmount()int
}