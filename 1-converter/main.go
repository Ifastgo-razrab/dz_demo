package main

import (
	"fmt"
	"errors"
)

const USDvEUR = 0.8436
const USDvRUB = 76.62
const EURvRUB = USDvRUB / USDvEUR
const EURvUSD = 1 / USDvEUR
const RUBvUSD = 1 /USDvRUB
const RUBvEUR = 1 / EURvRUB

func main() {
	fmt.Println("___Программа для обмена Валют____")
	
		var  namber float64
		var base_currency , target_currency string
		var err , ererr , err2  error
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
			target_currency , err2 = reverschat(base_currency)
			if err2 != nil {
				continue	
			}else{
				fmt.Println("Вы выбрали :", target_currency)
				break
			}
		}
		
		result := valute(namber , base_currency , target_currency)
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


func reverschat(base_currency string) (string ,error) {
	var target_currency string
	switch {
		case base_currency == "EUR":
			fmt.Print("Введите целивую валюту(RUB ,USD) :")
			fmt.Scan(&target_currency)
			if target_currency != "EUR" && target_currency != "USD" && target_currency != "RUB"  {
				return "",  errors.New("NO")
			}
			return target_currency , nil

		case base_currency == "RUB":
			fmt.Print("Введите целивую валюту(EUR ,USD) :")
			fmt.Scan(&target_currency)
			if target_currency != "EUR" && target_currency != "USD" && target_currency != "RUB"  {
				return "",  errors.New("NO")
			}
			return target_currency , nil

		case base_currency == "USD":
			fmt.Print("Введите целивую валюту(RUB ,EUR) :")
			fmt.Scan(&target_currency)
			if target_currency != "EUR" && target_currency != "USD" && target_currency != "RUB"  {
				return "",  errors.New("NO")
			}
			return target_currency , nil

	}
	
	return "" , nil
}
	
	





func valute(number float64, base_currency string, target_currency string) float64{
	var result float64
	switch{
		case base_currency == "USD" && target_currency == "EUR":
			result = USDvEUR * number
			return result
		case base_currency == "USD" && target_currency == "RUB":
			result = USDvRUB * number
			return result
		case base_currency == "EUR" && target_currency == "RUB":
			result = EURvRUB * number
			return result
		case base_currency == "EUR" && target_currency == "USD":
			result = EURvUSD * number
			return result
		case base_currency == "RUB" && target_currency == "USD":
			result = RUBvUSD * number
			return result
		case base_currency == "RUB" && target_currency == "EUR":
			result = RUBvEUR * number
			return result	
	}

	return 0

	
}