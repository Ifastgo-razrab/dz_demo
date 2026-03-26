package main

import (
	"fmt"
	"errors"
)

type valutemap = map[string]map[string]float64

func main() {
	fmt.Println("___Программа для обмена Валют____")
	full_kurc := valutemap{
		"USD": {
			"EUR" : 0.8436,
			"RUB" : 76.62,
		},
		"EUR": {
			"RUB" : 90.82,
			"USD" : 1.18,
		},
		"RUB": {
			"USD" : 0.013,
			"EUR" : 0.011,
		},
	}
	fmt.Println(full_kurc["EUR"])

		var  namber float64
		var base_currency , target_currency string
		var err , ererr   error
		for {
			base_currency ,  err = outprint()
			if err != nil {
				continue
			}else{
				fmt.Println("Вы выбрали :", base_currency)
				break
			}
			
		}

		for {
			namber , ererr  = namberto()
			if ererr != nil {
				continue
			}else{
				break
			}
		}
		for{
			target_currency = reverschat(base_currency , full_kurc)
			
			fmt.Println("Вы выбрали :", target_currency)
			break
		}
		
		result := valute(namber , base_currency , target_currency, full_kurc)
		fmt.Println("Результат : ", result)
		


		
		
	

}



func outprint() (  string , error)  {
	var base_currency string
	
	
	
	fmt.Print("Введите исходную валюту (EUR , USD ,RUB): ")
	fmt.Scan(&base_currency)
	if base_currency != "EUR" && base_currency != "USD" && base_currency != "RUB"  {
			return "",  errors.New("NO")
	}
	return base_currency , nil

}



func namberto() (float64 , error){
	var  number float64
	fmt.Print("Введите количество валюты : ")
	fmt.Scan(&number)
	if number <= 0{
			return  0,  errors.New("NO")			
	}
	return number , nil

}


func reverschat(base_currency string, full_kurc valutemap) (string ) {
	var target_currency string
	viktor := []string{}
	for key , _ := range full_kurc[base_currency]{
		viktor = append(viktor, key)
	}
	fmt.Print("Выберите целевую валютную пару",viktor , ": ")
	fmt.Scan(&target_currency)
	return target_currency 
}
	
	





func valute(number float64, base_currency string, target_currency string , m valutemap) float64{
	var result float64
	a := m[base_currency][target_currency]

	result = number * a


	return result

	
}