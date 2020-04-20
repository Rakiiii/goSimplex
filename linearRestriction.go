package gosimplex

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"errors"
	"io"
)

//Signum describe signum in restriction
type Signum int

//Type Signum
const (
	More Signum = 2
	MoreOrEquals Signum= 1
	Equals Signum = 0
	LessOrEquals Signum = -1
	Less Signum = -2
	WrongSignum Signum = -3
	)
var	(
	UnknownSinumError = errors.New("Unknown signum")
)

//InitSignum reads signum from string
func InitSignum(sign string)Signum{
	switch sign{
	case "=":
		return Equals
	case ">":
		return More
	case ">=":
		return MoreOrEquals
	case "<":
		return Less
	case "<=":
		return LessOrEquals
	default:
		return WrongSignum
	}
}

//LinearRestriction describe linear restriction implements GenerializedRestriction interface
type LinearRestriction struct{
	cofs []float64
	value float64
	sign Signum
}

type LinearRestrictions []LinearRestriction

func (r *LinearRestrictions)Check(v *Vector)(bool,error){
	for _,s := range (*r){
		flag,err := s.Check(v)
		if err != nil{
			return false,err
		}
		if !flag{
			return false,nil
		}
	}
	return true,nil
}

func (r *LinearRestrictions)Copy()*LinearRestrictions{
	new := make(LinearRestrictions,len(*r))
	for i,v := range (*r){
		new[i] = v.Copy()
	}
	return &new
}


//NewLinearRestriction return void LinearRestriction
func NewLinearRestriction()LinearRestriction{
	return LinearRestriction{cofs:nil,value:0.0,sign:WrongSignum}
}

//Init initialized restriction with cofs @c, free value @v and signum @s
func (l *LinearRestriction)Init(c []float64,v float64, s Signum){
	l.cofs = make([]float64,len(c))
	copy(l.cofs,c)
	l.value = v
	l.sign = s
}


//ReadRestrictionFromConsole reads system from stdin
func (l *LinearRestriction)ReadRestrictionFromConsole(am int){
	l.ReadRestrictionFromBuffer(am,bufio.NewReader(os.Stdin))
}

//ReadRestrictionFromBuffer reads restriction from buffered Reader @consoleReader for function with cofs amount @am
func (l *LinearRestriction)ReadRestrictionFromBuffer(am int,consoleReader *bufio.Reader){
	//consoleReader := bufio.NewReader(buf)

	var amountOfCofs int

	//read cofs amount
	for true{
		fmt.Println("Введите старший номер переменной в ограничении")

		readedString,err := consoleReader.ReadString('\n')
		if err != nil{
			fmt.Println("ошибка при чтении:",err.Error())
			os.Exit(1)
		}

		readedString = subStringBefore(readedString,NEXTLINE)
		
		amountOfCofs,err = strconv.Atoi(readedString)
		if err != nil{
			fmt.Println("Необходимо ввести целое число\nПопробуйте еще раз")
			continue
		}

		switch{
		case amountOfCofs <= 0:
			fmt.Println("Необходимо ввести строго положительное число\nПопробуйте еще раз")
			continue
		case amountOfCofs > am:
			fmt.Println("Число переменных в ограничении не может быть больше числа переменых в исходной функции\nПопробуйте еще раз")
			continue
		}

		break
	}

	//read cofs
	l.cofs = make([]float64,am)
	for cofNumber := 0; cofNumber < amountOfCofs;cofNumber++{
		for true{
			fmt.Println("Введите коэффицент при ",cofNumber+1,"-ой переменной")
			readedString,err := consoleReader.ReadString('\n')
			if err != nil{
				fmt.Println("ошибка при чтении:",err.Error())
				os.Exit(1)
			}

			readedString = subStringBefore(readedString,NEXTLINE)
	
			l.cofs[cofNumber],err = strconv.ParseFloat(readedString,64)
			if err != nil {
				fmt.Println("Вам необходимо ввести число\nПопробуйте еще раз")
				continue
			}
	
			break
		}
	}

	//filled non used cofs with zeros
	if amountOfCofs < am{
		for i := amountOfCofs;i < am;i++{
			l.cofs[i] = 0.0
		}
	}

	//read free value
	for true{
		fmt.Println("Введите свободный член")

		readedString,err := consoleReader.ReadString('\n')
		if err != nil{
			fmt.Println("ошибка при чтении:",err.Error())
			os.Exit(1)
		}

		readedString = subStringBefore(readedString,NEXTLINE)
	
		l.value,err = strconv.ParseFloat(readedString,64)
		if err != nil {
			fmt.Println("Вам необходимо ввести число\nПопробуйте еще раз")
			continue
		}
		break
	}

	//read signum
	for true{
		fmt.Println("Введите знак, допустимые значения для ввода:<,<=,=,>,>=")

		readedString,err := consoleReader.ReadString('\n')
		if err != nil{
			if err != io.EOF{
				fmt.Println("ошибка при чтении:",err.Error())
				os.Exit(1)
			}
		}

		readedString = subStringBefore(readedString,NEXTLINE)

		s := InitSignum(readedString)
		if s == WrongSignum{
			fmt.Println("Введен не корректный знак\nПопробуйте еще раз")
			continue
		}else{
			l.sign = s
			break
		}
	}
}

//Sign returns signum
func (l *LinearRestriction)Sign()Signum{
	return l.sign
}

//Cofs return (copy) array of float64 described cofs
func (l *LinearRestriction)Cofs()[]float64{
	ret := make([]float64,len(l.cofs))
	copy(ret,l.cofs)
	return ret
}

//Value returns free value
func (l *LinearRestriction)Value()float64{
	return l.value
}

//AmountOfCofs returns amount of cofs
func( l *LinearRestriction)AmountOfCofs()int{
	return len(l.cofs)
}

//Check implements GenerializedRestriction interface
func (l *LinearRestriction)Check(variabales *Vector)(bool,error){
	switch{
	case len(*variabales) > (len(l.cofs)):
		return false,errors.New("Too many variables for counting function")
	case len(*variabales) < (len(l.cofs)):
		return false,errors.New("Not enough variables for counting function")
	}

	result := 0.0
	for num,cof := range (*variabales){
		result += cof*l.cofs[num]
	}

	var ret bool
	switch l.sign{
	case Equals:
		if result == l.value{
			ret = true
		}else{
			ret = false
		}
	case More:
		if result > l.value{
			ret = true
		}else{
			ret = false
		}
	case MoreOrEquals:
		if result >= l.value{
			ret = true
		}else{
			ret = false
		}
	case Less:
		if result < l.value{
			ret = true
		}else{
			ret = false
		}
	case LessOrEquals:
		if result <= l.value{
			ret = true
		}else{
			ret = false
		}
	default:
		return false,errors.New("Unknown signum")
	}

	return ret,nil
}

//Copy returns copy of restriction
func (l *LinearRestriction)Copy()LinearRestriction{
	ret := NewLinearRestriction()
	ret.Init(l.cofs,l.value,l.sign)
	return ret
}