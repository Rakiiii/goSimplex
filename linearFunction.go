package gosimplex
import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"errors"
	"io"
)

const (
	MAX = "max"
	MIN = "min"
)
var (
	NotEnoughVariablesError = errors.New("Not enough variables for counting function")
	TooManyVariablesError = errors.New("Too many variables for counting function")
)

type LinearFunction struct{
	cofs []float64
	isMin bool
	counted int
}

func NewLinearFunction()LinearFunction{
	return LinearFunction{cofs:nil,isMin:false,counted:0,}
}

func (l *LinearFunction)Init(c []float64, s bool){
	l.cofs = make([]float64,len(c))
	copy(l.cofs,c)
	l.isMin = s
}

func (l *LinearFunction)ReadFunctionFromConsole(){
	l.ReadFunctionFromBuffer(bufio.NewReader(os.Stdin))
}

func (l *LinearFunction)ReadFunctionFromBuffer(consoleReader *bufio.Reader){
	initNEXTLINE()

	var amountOfCosf int 

	for true{
		fmt.Println("Введите количество переменных в функции")

		readedString,err := consoleReader.ReadString('\n')
		if err != nil{
			fmt.Println("ошибка при чтении:",err.Error())
			os.Exit(1)
		}

		readedString = subStringBefore(readedString,NEXTLINE)
		//fmt.Println(readedString)

		amountOfCosf,err = strconv.Atoi(readedString)
		if err != nil {
			fmt.Println("Вам необходимо ввести целое число\nПопробуйте еще раз")
			continue
		}

		if amountOfCosf <= 0{
			fmt.Println("КОличество коэффицентов может быть только строго положительным\nПопробуйте еще раз")
			continue
		}
		break
	}

	l.cofs = make([]float64,amountOfCosf+1)

	for cofNumber := 0; cofNumber < len(l.cofs)-1;cofNumber++{
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

	for true{
		fmt.Println("Введите свободный член")

		readedString,err := consoleReader.ReadString('\n')
		if err != nil{
			fmt.Println("ошибка при чтении:",err.Error())
			os.Exit(1)
		}

		readedString = subStringBefore(readedString,NEXTLINE)

		l.cofs[len(l.cofs)-1],err = strconv.ParseFloat(readedString,64)
		if err != nil {
			fmt.Println("Вам необходимо ввести число\nПопробуйте еще раз")
			continue
		}
		break
	}

	for true{
		fmt.Println("Тип оптимизационной задачи: если задача на минимум-введите min, если задача на максимум введит max...")
		readedString,err := consoleReader.ReadString('\n')
		if err != nil{
			if err != io.EOF{
				fmt.Println("ошибка при чтении:",err.Error())
				os.Exit(1)
			}
		}

		readedString = subStringBefore(readedString,NEXTLINE)

		switch readedString{
		case MIN:
			l.isMin = true
		case MAX:
			l.isMin = false
		default:
			fmt.Println("Допустимые значения для ввода только min и max\nПопробуйте еще раз")
			continue
		}

		break
	}
}

func (l *LinearFunction)Cofs()[]float64{
	ret := make([]float64,len(l.cofs))
	copy(ret,l.cofs)
	return ret
}

func (l *LinearFunction)Value()float64{
	return l.cofs[len(l.cofs)-1]
}

func (l *LinearFunction)AmountOfCofs()int{
	return len(l.cofs)-1
}

func (l *LinearFunction)CalculationAmount()int{
	return l.counted
}

func (l *LinearFunction)IsMin()bool{
	return l.isMin
}

func (l *LinearFunction)Count(variabales *Vector)(float64,error){
	switch{
	case len(*variabales) > (len(l.cofs)-1):
		return 0.0,errors.New("Too many variables for counting function")
	case len(*variabales) < (len(l.cofs)-1):
		return 0.0,errors.New("Not enough variables for counting function")
	}

	l.counted++

	result := l.cofs[len(l.cofs)-1]
	for num,cof := range (*variabales){
		result += cof*l.cofs[num]
	}

	return result,nil
}

func (l *LinearFunction)Copy()LinearFunction{
	ret := NewLinearFunction()

	ret.Init(l.cofs,l.isMin)
	return ret
}