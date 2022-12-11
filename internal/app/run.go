package app

import (
	"GoAsyncSbazarParcer/internal/models"
	"fmt"
	"strconv"
	"sync"

	gojson "github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

// Константные имена категорий
const (
	Electronic = "Elec"
	Сlothing   = "Сlothing"
	Hobby      = "Hobby"
	BabyMoM    = "BabyMoM"
	Sport      = "Sport"
)

func Run() {
	var wg sync.WaitGroup

	// i == Кол-во страниц которые будут собраны
	for i := 1; i <= 15; i++ {
		wg.Add(5)
		var (
			urlElectronic = fmt.Sprintf("https://www.sbazar.cz/30-elektro-pocitace/cela-cr/cena-neomezena/nejnovejsi/%s", strconv.Itoa(i))
			urlСlothing   = fmt.Sprintf("https://www.sbazar.cz/15-obleceni-obuv-doplnky/cela-cr/cena-neomezena/nejnovejsi/%s", strconv.Itoa(i))
			urlHobby      = fmt.Sprintf("https://www.sbazar.cz/33-starozitnosti-hobby-umeni/cela-cr/cena-neomezena/nejnovejsi/%s", strconv.Itoa(i))
			urlBabyMoM    = fmt.Sprintf("https://www.sbazar.cz/29-detsky-bazar/cela-cr/cena-neomezena/nejnovejsi/%s", strconv.Itoa(i))
			urlSport      = fmt.Sprintf("https://www.sbazar.cz/27-sport/cela-cr/cena-neomezena/nejnovejsi/%s", strconv.Itoa(i))
		)

		go func(urlElectronic string) {
			if err := models.FindProduct(urlElectronic, Electronic); err != nil {
				logrus.Println(err)
			}
			wg.Done()
		}(urlElectronic)

		go func(urlСlothing string) {
			if err := models.FindProduct(urlСlothing, Сlothing); err != nil {
				logrus.Println(err)
			}
			wg.Done()
		}(urlСlothing)

		go func(urlHobby string) {
			if err := models.FindProduct(urlHobby, Hobby); err != nil {
				logrus.Println(err)
			}
			wg.Done()
		}(urlHobby)

		go func(urlBabyMoM string) {
			if err := models.FindProduct(urlBabyMoM, BabyMoM); err != nil {
				logrus.Println(err)
			}
			wg.Done()
		}(urlBabyMoM)

		go func(urlSport string) {
			if err := models.FindProduct(urlSport, Sport); err != nil {
				logrus.Println(err)
			}
			wg.Done()
		}(urlSport)
		wg.Wait()
	}

	/*
		Отправка данных в другой микросервис
	*/
	fmt.Println(string(MarshalData(models.Elec)))
	fmt.Println()
	fmt.Println()
	fmt.Println(string(MarshalData(models.Сlothing)))
	fmt.Println()
	fmt.Println()
	fmt.Println(string(MarshalData(models.Hobby)))
	fmt.Println()
	fmt.Println()
	fmt.Println(string(MarshalData(models.BabyMoM)))
	fmt.Println()
	fmt.Println()
	fmt.Println(string(MarshalData(models.Sport)))
}

func MarshalData(data interface{}) []byte {
	out, err := gojson.Marshal(data)
	if err != nil {
		logrus.Errorf("Err marshal data in struct - %s", err)
	}
	return out
}
