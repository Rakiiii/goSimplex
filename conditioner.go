package gosimplex

//Conditioner interface for stoping algorithms
type Conditioner interface{
	//CheckCondition is returns false ? algorithm must be stoped
	CheckCondition()bool
}

//ItterationConditioner stop algorithm after @max amount of its itterarions
type ItterationConditioner struct{
	itterator int
	max int
}

//NewItterationConditioner returns void ItterationConditioner
func NewItterationConditioner()*ItterationConditioner{
	return &ItterationConditioner{max:0,itterator:0}
}

//Init initialize ItterationConditioner with max amount of itterations of @m
func (c *ItterationConditioner)Init(m int){
	c.itterator = 0
	c.max = m
}

//CheckCondition implements Conditioner interface
func (c *ItterationConditioner)CheckCondition()bool{
	c.itterator++
	return c.itterator <= c.max
}

//FunctionConditioner stop algorithm after @max amount of function @f calculation
type FunctionConditioner struct{
	max int
	f GenerializedFunctioner
}

//NewFunctionConditioner returns void FunctionConditioner
func NewFunctionConditioner()*FunctionConditioner{
	return &FunctionConditioner{max:0,f:nil}
}

//Init initialized FunctionConditioner with funtion @fun and max amount of calculation @m
func (c *FunctionConditioner)Init(m int,fun GenerializedFunctioner){
	c.f = fun
	c.max = m
}

//CheckCondition implements Conditioner interface
func (c *FunctionConditioner)CheckCondition()bool{
	return c.f.CalculationAmount() < c.max
}



