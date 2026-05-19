package file

import (
	"fmt"
	"os"
	"strings"
)

func Writefile(content []byte, file string) {
	data, err := os.Create(file)
	if err != nil {
		fmt.Println("Файл :" , err )
	}
	defer data.Close()
	_ , err = data.Write(content)
	if err != nil {
		fmt.Println("Файл :" , err )
	}
	fmt.Println("Файл успешно записан")
}


func ReadFile(file string)([]byte){
	data , err := os.ReadFile(file)
	if err != nil{
		fmt.Println("File:", err)
		
	}
	return data


}
func isJson(file string)bool{
	return strings.HasSuffix(file , ".json")
}
