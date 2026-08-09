package api

import (
	"CLI/bins"
	"CLI/config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)
type Apiconf struct{
	Bin bins.Bin
	Config config.Config
}

func Apiconfig(config config.Config , bins bins.Bin)*Apiconf{
	return &Apiconf{
		Config: config,
		Bin: bins,
	}
}

func CreatBin(data bins.Bin){
	fmt.Println("zapusk")
	data_js, err := json.Marshal(data)
	
	if err != nil{
		fmt.Println(err)
	}
	reader := strings.NewReader(string(data_js))

	req, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b", reader )
	if err != nil{
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key" , "$2a$10$pX43ILvnHXg1nOMXXv59yeaI76B423sBVGqCzPs8gbX.ZXotgvVBy")  
	
	client := http.Client{}
	resp ,err := client.Do(req)
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(data_js))
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))
}

func Binget(Id string , Url string){
	baseurl , err := url.Parse(Url + Id)
	if err != nil{
		fmt.Println(err, "err1")
	}
	req ,err := http.NewRequest(http.MethodGet, baseurl.String(), nil)
	if err != nil{
		fmt.Println(err)
	}
	req.Header.Set("X-Master-Key" , "$2a$10$pX43ILvnHXg1nOMXXv59yeaI76B423sBVGqCzPs8gbX.ZXotgvVBy")
	req.Header.Set("X-Bin-Meta" , "false")
	

	client := &http.Client{}
	reset ,err := client.Do(req)
	if err != nil{
		fmt.Println(err)
	}
	defer reset.Body.Close()
	var bins bins.Bin
	err = json.NewDecoder(reset.Body).Decode(&bins)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bins)
	

	
}

func Delite(Id string , Url string){
	baseUrl , err := url.Parse(Url + Id)
	if err != nil {
		fmt.Println(err)
	}
	req , err := http.NewRequest(http.MethodDelete, baseUrl.String(), nil)
	if err != nil{
		fmt.Println(err)
	}
	req.Header.Set("X-Master-Key" , "$2a$10$pX43ILvnHXg1nOMXXv59yeaI76B423sBVGqCzPs8gbX.ZXotgvVBy")

	client := http.Client{}
	reset , err := client.Do(req)
	if err != nil{
		fmt.Println(err)
	}
	defer reset.Body.Close()

	body, _ := io.ReadAll(reset.Body)
	fmt.Println(string(body))

}

func Put(data bins.Bin , Id string){
	dataJs ,err := json.Marshal(data)
	
	baseUrl , err := url.Parse("https://api.jsonbin.io/v3/b/" + Id)
	fmt.Println(baseUrl)
	if err != nil {
		fmt.Println(err)
	}
	reader := strings.NewReader(string(dataJs))

	req , err := http.NewRequest(http.MethodPut , baseUrl.String(), reader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key" , "$2a$10$pX43ILvnHXg1nOMXXv59yeaI76B423sBVGqCzPs8gbX.ZXotgvVBy")

	client := http.Client{}
	reset , err := client.Do(req)
	if err!= nil{
		fmt.Println(err)
	}
	defer reset.Body.Close()

	body , _ := io.ReadAll(reset.Body)
	fmt.Println(reset.StatusCode)
	fmt.Println(string(body))
	
}