package gosimplex

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"errors"
	"io"
)

type Signum int

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

type LinearRestriction struct{
	cofs []float64
	value float64
	sign Signum
}

func NewLinearRestriction()LinearRestriction{
	return LinearRestriction{cofs:nil,value:0.0,sign:WrongSignum}
}

func (l *LinearRestriction)Init(c []float64,v float64, s Signum){
	l.cofs = make([]float64,len(c))
	copy(l.cofs,c)
	l.value = v
	l.sign = s
}

func (l *LinearRestriction)ReadRestrictionFromConsole(am int){
	l.ReadRestrictionFromBuffer(am,os.Stdin)
}

func (l *LinearRestriction)ReadRestrictionFromBuffer(am int,buf io.Reader){
	consoleReader := bufio.NewReader(buf)

	var amountOfCofs int

	for true{
		fmt.Println("Введите количество переменных в ограничении")

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

		if amountOfCofs < am{
			for i := amountOfCofs;i < am;i++{
				l.cofs[i] = 0.0
			}
		}

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

func (l *LinearRestriction)Sign()Signum{
	return l.sign
}

func (l *LinearRestriction)Cofs()[]float64{
	ret := make([]float64,len(l.cofs))
	copy(ret,l.cofs)
	return ret
}

func (l *LinearRestriction)Value()float64{
	return l.value
}

func( l *LinearRestriction)AmountOfCofs()int{
	return len(l.cofs)
}

func (l *LinearRestriction)CheckRestriction(variabales []float64)(bool,error){
	switch{
	case len(variabales) > (len(l.cofs)):
		return false,errors.New("Too many variables for counting function")
	case len(variabales) < (len(l.cofs)):
		return false,errors.New("Not enough variables for counting function")
	}

	result := 0.0
	for num,cof := range variabales{
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

func (l *LinearRestriction)Copy()LinearRestriction{
	ret := NewLinearRestriction()
	ret.Init(l.cofs,l.value,l.sign)
	return ret
}