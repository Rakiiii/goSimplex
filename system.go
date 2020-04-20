package gosimplex

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"errors"
)

var (
	CofsError = errors.New("Wrong amount of cofs in restriction")
)

type System struct{
	optimisationFunction LinearFunction
	//restrictions []LinearRestriction
	restrictions LinearRestrictions
}

func NewSystem()System{
	return System{optimisationFunction:NewLinearFunction(),restrictions:nil}
}

func (s *System)Init(f LinearFunction, r []LinearRestriction)error{
	for _,lr := range r{
		if lr.AmountOfCofs() != f.AmountOfCofs(){
			return errors.New("Wrong amount of cofs in restriction")
		}
	}
	s.optimisationFunction = f.Copy()
	s.restrictions = make([]LinearRestriction,len(r))
	for num,res := range r{
		s.restrictions[num] = res.Copy()
	}
	return nil
}

func (s *System)ReadSystemFromConsole(){
	s.ReadSystemFromBuffer(bufio.NewReader(os.Stdin))
}

func (s *System)ReadSystemFromBuffer(consoleReader *bufio.Reader){

	s.optimisationFunction.ReadFunctionFromBuffer(consoleReader)

	var amountOfRestrictions int
	
	for true{
		fmt.Println("Введите количество ограничений в системе")

		readedString,err := consoleReader.ReadString('\n')
		if err != nil{
			fmt.Println("ошибка при чтении:",err.Error())
			os.Exit(1)
		}

		readedString = subStringBefore(readedString,NEXTLINE)

		amountOfRestrictions,err = strconv.Atoi(readedString)
		if err != nil {
			fmt.Println("Вам необходимо ввести целое число\nПопробуйте еще раз")
			continue
		}

		if amountOfRestrictions <= 0{
			fmt.Println("Количество ограничений может быть только строго положительным\nПопробуйте еще раз")
			continue
		}
		break
	}

	s.restrictions = make([]LinearRestriction,amountOfRestrictions)

	for num := 0 ; num < amountOfRestrictions; num ++{
		fmt.Println("Инициализация ",num+1,"-ого ограничения")
		s.restrictions[num].ReadRestrictionFromBuffer(s.optimisationFunction.AmountOfCofs(),consoleReader)
	}
}

func(s *System)Copy()System{
	ret := NewSystem()
	ret.Init(s.optimisationFunction,s.restrictions)
	return ret
}

func (s *System)AmountOfRestrictions()int{
	return len(s.restrictions)
}

func (s *System)Function()*LinearFunction{
	return &s.optimisationFunction
}

func (s *System)Restriction(i int)*LinearRestriction{
	if i > len(s.restrictions){
		return nil
	}else{
		return &s.restrictions[i]
	}
}

func (s *System)Restrict()GenerializedRestrictioner{
	return s.restrictions.Copy()
}

func (s *System)Func()GenerializedFunctioner{
	return &(s.optimisationFunction)
}

func (s *System)Size() int{
	return s.optimisationFunction.AmountOfCofs()
}

func (s *System)IsMin()bool{
	return s.optimisationFunction.IsMin()
}