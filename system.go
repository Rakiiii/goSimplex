package gosimplex

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"io"
)

type System struct{
	optimisationFunction LinearFunction
	restrictions []LinearRestriction
}

func NewSystem()System{
	return System{optimisationFunction:NewLinearFunction(),restrictions:nil}
}

func (s *System)Init(f LinearFunction, r []LinearRestriction){
	s.optimisationFunction = f.Copy()
	s.restrictions = make([]LinearRestriction,len(r))
	for num,res := range r{
		s.restrictions[num] = res.Copy()
	}
}

func (s *System)ReadSystemFromConsole(){
	s.ReadSystemFromBuffer(os.Stdin)
}

func (s *System)ReadSystemFromBuffer(buf io.Reader){

	s.optimisationFunction.ReadFunctionFromBuffer(buf)

	consoleReader := bufio.NewReader(buf)

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
		s.restrictions[num].ReadRestrictionFromBuffer(s.optimisationFunction.AmountOfCofs(),buf)
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

func (s *System)Function()LinearFunction{
	return s.optimisationFunction
}

func (s *System)Restriction(i int)LinearRestriction{
	if i > len(s.restrictions){
		return NewLinearRestriction()
	}else{
		return s.restrictions[i]
	}
}