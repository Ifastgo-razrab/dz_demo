package config



type Config struct {
	Key string
}

func NewConfig()*Config{
	//key :=  os.Getenv("KEY")
	//if key == ""{
	//	panic("Не удалось найти ключ")
	//}
	return &Config{
		Key: "$2a$10$SeGDTPkug3zqvkDEu2H6SuCtzpXbDSQGjOEhHgreRTDzl8N9kvACK",
	}
}