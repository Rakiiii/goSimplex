package gosimplex

import (
	//"bufio"
	"strconv"
	//"io"
	"fmt"
)

type Simplex struct{
	syst System
	table [][]float64
	basic []int
	width int
	heigh int
}

func NewSimplex()Simplex{
	return Simplex{syst:NewSystem(),
		table:nil,
		basic:nil,
		width:-1,
		heigh:-1 }
}

func (s *Simplex)Init(st System){
	//copy system
	s.syst = st.Copy()

	source := setSourseTable(st)

	//count size of simplex table
	s.heigh = s.syst.AmountOfRestrictions()+1
	//f := s.syst.Function()
	s.width = s.syst.Function().AmountOfCofs()+1

	//init simplex table
	s.table = make([][]float64,s.heigh)
	for i,_ := range s.table{
		s.table[i] = make([]float64,s.width+s.heigh-1)
	}

	//init basis
	s.basic = make([]int,0)

	//fill simplex table
	for i := 0; i < s.heigh;i++{
		for j := 0;j < len(s.table[i]);j++{
			if j < s.width{
				s.table[i][j] = source[i][j]
			}else{
				s.table[i][j] = 0
			}

		}

		//set 1 elements in table at basis
		if (s.width+i) < len(s.table[i]){
			s.table[i][s.width+i] = 1
			s.basic = append(s.basic,s.width+i)
		}
	}

	s.width = len(s.table[0])
}

func (s *Simplex)Run()([]float64,[][]float64){
	var mainCol int
	var mainRow int

	for !s.isOptimumFound(){

		mainCol = s.findMainCol()
		mainRow = s.findMainRow(mainCol)

		s.basic[mainRow] = mainCol
		
		newTable := make([][]float64,s.heigh)

		for i,_ := range newTable{
			newTable[i] = make([]float64,s.width)
		}

		for j := 0; j < s.width; j++{
			newTable[mainRow][j] = s.table[mainRow][j] / s.table[mainRow][mainCol]
		}

		for i := 0; i < s.heigh ; i++{
			if i == mainRow{
				continue
			}
			for j := 0; j < s.width; j++{
				newTable[i][j] = s.table[i][j] - s.table[i][mainCol]*newTable[mainRow][j]	
			}
		}


		s.table = newTable
	}

	

	result := make([]float64,s.syst.Function().AmountOfCofs())

	for i,_ := range result{
		k := indexOf( i + 1 , s.basic)
		if k != -1{
			result[i] = s.table[k][0]
		}else{
			result[i] = 0
		}
	}

	return result,s.table
}

func (s *Simplex) findMainCol()int{
	mainCol := 1
	for j := 2; j < s.width; j++{
		if s.syst.Function().IsMin() {
			if s.table[s.heigh-1][j] > s.table[s.heigh-1][mainCol]{
				mainCol = j
			}
		}else{
			if s.table[s.heigh-1][j] < s.table[s.heigh-1][mainCol]{
				mainCol = j
			}
		}
	}

	return mainCol
}

func (s *Simplex)findMainRow(mainCol int)int{
	mainRow := 0

	for i := 0; i < s.heigh-1; i++{
		if s.table[i][mainCol] > 0{
			mainRow = i
			break
		}
	}

	for i := mainRow + 1; i < s.heigh-1; i ++{
		if (s.table[i][mainCol] > 0) && ( ( s.table[i][0] / s.table[i][mainCol] ) < ( s.table[mainRow][0] / s.table[mainRow][mainCol] ) ){
			mainRow = i
		}
	}

	return mainRow
}

func (s *Simplex)isOptimumFound()bool{
	for i := 1; i < s.width ; i++{
		if s.syst.Function().IsMin(){
			if s.table[s.heigh-1][i] > 0{
				return false
			}
		}else{
			if s.table[s.heigh-1][i] < 0{
				return false
			}			
		}
	}
	return true
}

func (s *Simplex)PrintTableToFmt(){
	str := ""
	for i,row := range s.table{
		if i+1 < len(s.table){
			str += strconv.Itoa(s.basic[i])+" "
		} else{
			str += "  "
		}
		
		for _,el := range row{
			str += strconv.FormatFloat(el,'f',-1,64) + " "
		}
		str += NEXTLINE
	}
	fmt.Println(str)
}

func PrintlnTable(table [][]float64){
	str := ""
	for _,row := range table{
		for _,el := range row{
			str += strconv.FormatFloat(el,'f',-1,64) + " "
		}
		str += NEXTLINE
	}
	fmt.Println(str)
}

func indexOf(n int,b []int)int{
	for i,v := range b{
		if v == n{
			return i
		}
	}
	return -1
}

func setSourseTable(syst System)[][]float64{
	result := make([][]float64,syst.AmountOfRestrictions()+1)
	for i := 0; i < syst.AmountOfRestrictions();i++{
		result[i] = append([]float64{-1.0*syst.Restriction(i).Value()}, syst.Restriction(i).Cofs()...)
	}
	result[len(result)-1] = make([]float64,syst.Function().AmountOfCofs()+1)
	result[len(result)-1][0] = syst.Function().Value()

	for i := 1; i < len(result[len(result)-1]);i++{
		result[len(result)-1][i] = -1.0*syst.Function().Cofs()[i-1]
	}

	return result
}

