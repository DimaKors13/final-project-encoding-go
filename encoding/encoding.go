package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {

	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		return err
	}
	j.DockerCompose = &models.DockerCompose{}
	err = json.Unmarshal(jsonFile, j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при десериализации: %s", err.Error())
		return err
	}

	yamlData, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при сериализации: %s", err.Error())
		return err
	}

	file, err := os.Create(j.FileOutput)
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s", err.Error())
		return err
	}

	defer file.Close()

	_, err = file.Write(yamlData)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {

	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		return err
	}

	y.DockerCompose = &models.DockerCompose{}
	err = yaml.Unmarshal(yamlFile, y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при десериализации: %s", err.Error())
		return err
	}

	jsonData, err := json.Marshal(y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при сериализации: %s", err.Error())
		return err
	}

	file, err := os.Create(y.FileOutput)
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s", err.Error())
		return err
	}
	// когда программа завершится, надо закрыть дескриптор файла
	defer file.Close()

	// записываем слайс байт в файл
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
		return err
	}

	return nil
}
