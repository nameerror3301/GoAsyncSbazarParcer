package models

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

type RequesLast struct {
	User struct {
		Name        string `json:"user_name"`
		DateRegistr string `json:"user_date_reg"`
		PhoneNumber string `json:"user_phone"`
	} `json:"user_data"`
	Products struct {
		ProdName    string `json:"product_name"`
		PhotoUrl    string `json:"photo_url"`
		Price       string `json:"price"`
		Description string `json:"description"`
		Url         string `json:"url"`
	} `json:"product_data"`
}

var Elec []RequesLast
var Сlothing []RequesLast
var Hobby []RequesLast
var BabyMoM []RequesLast
var Sport []RequesLast

// Универсальная функция для сбора данных со всех категорий
func FindProduct(url string, category string) error {
	idx := 1
	data, err := Request(url)
	if err != nil {
		logrus.Println(err)
	}
	defer data.Close()

	// Загружаем ответ от ресурса
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		logrus.Errorf("Err load data - %s", err)
	}
	defer data.Close()

	doc.Find("a[class=c-item__link]").Each(func(_ int, s *goquery.Selection) {
		val, _ := s.Attr("href")
		c := colly.NewCollector()

		c.OnHTML("div[class=p-uw-item__content]", func(e *colly.HTMLElement) {
			productName := e.DOM.Find("h1[class=p-uw-item__header]").Text()
			photoUrl, _ := e.DOM.Find("div[class=ob-c-carousel__item-content] img").Attr("src")
			description := e.DOM.Find("p[class=p-uw-item__description]").Text()
			price := e.DOM.Find("div[class=p-uw-item__first-line] b[class=c-price__price]").Text()
			currency := e.DOM.Find("span[class=c-price__currency]").Text()
			userName := e.DOM.Find("div[class=c-seller-info__name-wrapper]").Text()
			dateCreate := e.DOM.Find("div[class=c-seller-info__date]").Text()
			phoneNumber := e.DOM.Find("span[itemprop=telephone]").Text()

			switch category {
			case "Elec":
				AppendData(&Elec, userName, dateCreate, phoneNumber, productName, fmt.Sprintf("https:%s", photoUrl), price, currency, description, val)
			case "Сlothing":
				AppendData(&Сlothing, userName, dateCreate, phoneNumber, productName, fmt.Sprintf("https:%s", photoUrl), price, currency, description, val)
			case "Hobby":
				AppendData(&Hobby, userName, dateCreate, phoneNumber, productName, fmt.Sprintf("https:%s", photoUrl), price, currency, description, val)
			case "BabyMoM":
				AppendData(&BabyMoM, userName, dateCreate, phoneNumber, productName, fmt.Sprintf("https:%s", photoUrl), price, currency, description, val)
			case "Sport":
				AppendData(&Sport, userName, dateCreate, phoneNumber, productName, fmt.Sprintf("https:%s", photoUrl), price, currency, description, val)
			}

			logrus.Infof("%d -> %s [%s]\n", idx, val, category)
		})

		c.OnError(func(r *colly.Response, err error) {
			logrus.Errorf("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
			time.Sleep(time.Second * 5)
		})

		c.Visit(val)
		idx++
	})
	return nil
}

func AppendData(data *[]RequesLast, respData ...string) {
	*data = append(*data, RequesLast{
		User: struct {
			Name        string "json:\"user_name\""
			DateRegistr string "json:\"user_date_reg\""
			PhoneNumber string "json:\"user_phone\""
		}{
			Name:        respData[0],
			DateRegistr: respData[1],
			PhoneNumber: respData[2],
		},
		Products: struct {
			ProdName    string "json:\"product_name\""
			PhotoUrl    string "json:\"photo_url\""
			Price       string "json:\"price\""
			Description string "json:\"description\""
			Url         string "json:\"url\""
		}{
			ProdName:    respData[3],
			PhotoUrl:    respData[4],
			Price:       fmt.Sprintf("%s %s", respData[5], respData[6]),
			Description: respData[7],
			Url:         respData[8],
		},
	})
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

	fmt.Println(resp.StatusCode)
	return resp.Body, nil
}
