package model

type Coffee_recipe struct {
	// gorm.Model
	Id       string
	Brew_day string
	Device   string
	Grind    string
	Beans    string
	Shop     string
	Review   string
	Other    string
}
