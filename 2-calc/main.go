package main

import (
"fmt"
"strings"
"strconv"
)

func main(){
	
	operation := outpstring()
	outpint := outpfloat()

	result , a := raspredelenie(operation , outpint)
	fmt.Println("Ваш результат ",  a, result)
	
	
	


}

func outpstring() (int){
	var vibor string

	fmt.Print("Введит еопирацию(AVG - среднее(1), SUM - сумму(2), MED - медиану(3)) Либо номер:")
	fmt.Scan(&vibor)
	
	if vibor == "AVG" ||  vibor == "1"{

		fmt.Println("Вы выбрали операцию AVG - среднее")
		return 1

	}else if vibor == "SUM" || vibor == "2" {

		fmt.Println("Вы выбрали операцию SUM - сумму")
		return 2
		
	}else if vibor == "MED" || vibor == "3"{

		fmt.Println("Вы выбрали операцию MED - медиану")
		return 3

	}
	return 0
	

} 
func outpfloat()[]float64{
	var outstring string
	resultnum  := []float64{}
	fmt.Print("Введите число через запитую:")
	fmt.Scan(&outstring)
	result := strings.Split(outstring , ",")
	for _ , value := range result{
		f, err:= strconv.ParseFloat(value , 64)
		if err != nil {
			continue
		}
		resultnum = append(resultnum, f)
	}
	return resultnum
}
func raspredelenie(vibor int, fs[]float64)(float64 , string){

	switch{
		case vibor == 1:
			result := amg(fs)
			a := "среднее = "
			return result , a
		
		case vibor == 2:
			result := sum(fs)
			a := "сумма =  "
			return  result, a

		
		case vibor == 3:
			result := med(fs)
			a := "медиана = "
			return result, a
		
	}
	return 0 ,""
}
func amg(fs []float64)float64{
	var sum , nomer float64
	for index, value := range fs{
		sum += value
		nomer = float64(index)
		
		
	}
	result := sum / nomer
	return result
}
func  sum(fs []float64)float64{
	var sum float64
	for _ , value := range fs{
		sum +=value
	}
	return sum
}
func med(fs []float64)float64{
	n := len(fs)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if fs[j] > fs[j+1] {
				fs[j], fs[j+1] = fs[j+1], fs[j]
			}
		}
	}
	fmt.Println(fs)
	if n%2 ==1 {
		return fs[n/2]
	}
	return (fs[n/2 -1]+ fs[n/2]) /2 
}
