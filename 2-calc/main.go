package main

import (
	"fmt"
	"strings"
	"strconv"
)
var mapMenu = map[string]func([]float64)(float64, string){
	"1":amg,
	"2":sum,
	"3":med,
}
func main(){
	var varmenu = []string{
      	"Выберите одно из действий:",
		"1: Средне",
		"2: Сумма",
		"3: Медиана",
		"Ваш выбор",
   	}
	menu := promtData(varmenu...)
	outpint := outpfloat()
	a := mapMenu[menu]
	res , str := a(outpint)

	fmt.Println(str , res)
	
	
	
	


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

func amg(fs []float64)(float64 , string){
	var sum , nomer float64
	nomer = 0
	for _, value := range fs{
		sum += value
		nomer ++

	}
	result := sum / nomer
	return result , "Среедне = "
}
func  sum(fs []float64)(float64, string){
	var sum float64
	for _ , value := range fs{
		sum +=value
	}
	return sum , "Сумма :"
}
func med(fs []float64)(float64 , string){
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
		return fs[n/2] , ""
	}
	return (fs[n/2 -1]+ fs[n/2]) /2 , "Медиана ="
}
func promtData(promt ...string)string{
   var res string
   for index , data := range promt{
      if index == len(promt) -1 {
         fmt.Printf("%v: " , data)
      }else{
         fmt.Println(data)
      }
   }
   fmt.Scanln(&res)
   return res
}