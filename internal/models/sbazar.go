package models

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

type AllData struct {
	Contact struct {
		Name         string
		LastActicity string
		DateRegistr  string
		PhoneNumber  int
	}
	Products struct {
		ProdName    string
		PhotoUrl    string
		Price       int
		Description string
		Url         string
	}
	IsPhone bool
}

// Сбор данных со всех категорий
func FindLinks(url string) error {

	data, err := Request(url)
	if err != nil {
		logrus.Println(err)
	}

	// Загружаем ответ от ресурса
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		logrus.Errorf("Err load data - %s", err)
	}

	// Получение ссылки
	doc.Find("a[class=c-item__link]").Each(func(_ int, s *goquery.Selection) {
		val, _ := s.Attr("href")
		/*
			Переход по полученной ссылке и сбор данных о товаре и пользователе
		*/
		fmt.Println(val)
	})

	defer data.Close()
	return nil
}

func Request(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		logrus.Fatalf("Err request to %s - %s", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Err responce - %d %s", resp.StatusCode, resp.Status)
		time.Sleep(time.Second * 5)
	}
	return resp.Body, nil
}

// func CollectingProductData(url string) ([]Electronic, error) {

// 	return nil, nil
// }
