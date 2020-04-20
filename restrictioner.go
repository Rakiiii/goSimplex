package gosimplex

//GenerializedRestrictioner interface for genearilization any restriction
type GenerializedRestrictioner interface{
	//Check must return true if restrion is ok with @Vector
	Check(*Vector)(bool,error)
}

type VoidRestriction struct{}

func (s VoidRestriction)Check(v *Vector)(bool,error){
	return true,nil
}