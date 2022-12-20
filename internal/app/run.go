package app

import (
	"GoAsyncSbazarParcer/internal/config"
	"GoAsyncSbazarParcer/internal/models"
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

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

const (
	QueryElectronic = "electronics"
	QueryСlothing   = "clothing"
	QueryHobby      = "hobby"
	QueryBabyMoM    = "babymom"
	QuerySport      = "sport"
)

func Run() {
	var wg sync.WaitGroup
	conf := config.ReadConfig()

	// i == Кол-во страниц которые будут собраны
	for i := 1; i <= 1; i++ {
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
	SendData(MarshalData(models.Elec), QueryElectronic, conf.Data.JwtToken)

	SendData(MarshalData(models.Сlothing), QueryСlothing, conf.Data.JwtToken)

	SendData(MarshalData(models.Hobby), QueryHobby, conf.Data.JwtToken)

	SendData(MarshalData(models.BabyMoM), QueryBabyMoM, conf.Data.JwtToken)

	SendData(MarshalData(models.Sport), QuerySport, conf.Data.JwtToken)
}

func SendData(data []byte, category string, token string) {
	conf := config.ReadConfig()

	url := fmt.Sprintf("%sadd?category=%s&market=sbazar", conf.Data.OutStorageAddr, category)

	reader := bytes.NewReader(data)
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		logrus.Error("Err request generation - %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		logrus.Error("Err send data - %s", err)
		time.Sleep(5 * time.Second)
		SendData(MarshalData(models.Elec), category, conf.Data.JwtToken)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		logrus.Warnf("Check jwt token - %d", http.StatusUnauthorized)
	}

	if res.StatusCode != http.StatusOK {
		logrus.Errorf("Err sending data - %s", err)
	} else {
		logrus.Info("Success sending data")
	}
}

func MarshalData(data interface{}) []byte {
	out, err := gojson.Marshal(data)
	if err != nil {
		logrus.Errorf("Err marshal data in struct - %s", err)
	}
	return out
}
