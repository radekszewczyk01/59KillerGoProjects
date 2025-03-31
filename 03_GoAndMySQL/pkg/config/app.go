package config

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
)

var (
	db *gorm.DB
)

type Config struct {
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBName     string `json:"db_name"`
}

func LoadConfig() (*Config, error) {
	file, err := ioutil.ReadFile("C:/Users/ahamed/GolandProjects/59KillerGOProjects/03_GoAndMySQL/pkg/config/config.json") // wczytanie pliku JSON
	if err != nil {
		return nil, fmt.Errorf("błąd przy wczytywaniu pliku konfiguracyjnego: %v", err)
	}

	var config Config
	err = json.Unmarshal(file, &config) // przetwarzanie JSON na strukturę
	if err != nil {
		return nil, fmt.Errorf("błąd przy przetwarzaniu JSON: %v", err)
	}

	return &config, nil
}

func Connect() {
	// Wczytaj dane z pliku konfiguracyjnego
	config, err := LoadConfig()
	fmt.Println(config.DBHost, config.DBPort, config.DBUser, config.DBPassword)
	if err != nil {
		log.Fatal("Nie udało się załadować konfiguracji:", err)
	}

	// Stwórz połączenie z bazą danych
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Nie udało się połączyć z bazą danych:", err)
	}
	db = d
	fmt.Println("✅ Połączenie z bazą danych powiodło się!")
}

func GetDB() *gorm.DB {
	return db
}
