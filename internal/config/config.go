package config

import (
	"sync"
	"testingwork-kafka/pkg/logging"

	"github.com/ilyakaznacheev/cleanenv"
)

// создаём структуру которая соответствует конфигу в yml формате
type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// объявляем две переменные инстанс используем для записи туда конфига, а ванс
// для того что бы убедиться что наш код сработает только один раз
var instance *Config
var once sync.Once

func GetConfig() *Config {
	//ванс говорит о том что код выполнится только один раз, а остальные запуски
	//проигнорируются
	once.Do(func() {
		//передаём логер
		logger := logging.GetLogger()
		logger.Info("чтение конфига")
		instance = &Config{}
		//записываем в инстанс конфиг из файла, проверяем на ошибки, возвращаем конфиг выше
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info("ошибка чтения конфига: ", help)
			logger.Fatal(err)
		}
	})
	return instance
}
