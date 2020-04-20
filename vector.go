package gosimplex

import (
	"fmt"
	"strconv"
	"math/rand"
)

type Vector []float64


func NewVector(size int)*Vector{
	n := make(Vector,size)
	return &n	
}

func NewUnitVector(size int)*Vector{
	vec := NewVector(size)
	for i,_ := range (*vec){
		(*vec)[i] = 1.0
	}
	return vec
}

//NewZeroVector returns pointer to new Vector of size @size filled with zeros
func NewZeroVector(size int)*Vector{
	vec := NewVector(size)
	for i,_ := range (*vec){
		(*vec)[i] = 0.0
	}
	return vec
}

func NewRandomVector(size int)*Vector{
	vec := NewVector(size)
	for i,_ := range (*vec){
		(*vec)[i] = rand.NormFloat64()
	}
	return vec
}

//Add returns result of @v +@x
func (v *Vector)Add(x *Vector)*Vector{
	if len(*v) != len(*x){
		return nil
	}
	result := NewVector(len(*v))
	for i,_ := range (*v){
		(*result)[i] = (*v)[i] + (*x)[i]
	}
	return result
}

//Sub returns sub of v-x 
func (v *Vector)Sub(x *Vector)*Vector{
	if len(*v) != len(*x){
		return nil
	}
	result := NewVector(len(*v))
	for i,_ := range (*v){
		(*result)[i] = (*v)[i] - (*x)[i]
	}
	return result
}

func (v *Vector)TimesFloat(t float64)*Vector{
	result := NewVector(len(*v))

	for i,f := range *v{
		(*result)[i] = f*t
	}
	return result
}

func (v *Vector)DivFloat(t float64)*Vector{
	result := NewVector(len(*v))

	for i,f := range *v{
		(*result)[i] = f/t
	}
	return result
}

func (v *Vector)String()string{
	str := "("
	for _,v := range (*v){
		str += strconv.FormatFloat(v,'f',-1,64) +","
	}
	str = str[:len(str)-1]
	str += ")"
	return str
}

func (v *Vector)Cmp(v2 *Vector)bool{
	if len(*v) != len(*v2){
		return false
	}

	for i,j := range (*v){
		if j != (*v2)[i]{
			return false
		}
	}
	return true
}

func (v *Vector)Println(){
	fmt.Println(v.String())
}



