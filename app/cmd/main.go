package main

import (
	"GoAsyncSbazarParcer/internal/app"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})
}

/*
	<div class="p-uw-item__content" - Элемент в котором находятся все данные

	<h1 class="p-uw-item__header" > font style="vertical-align: inherit;" - Название товара (TEXT)

	<b class="c-price__price" > font style="vertical-align: inherit;" - Цена товара (TEXT)

	<p class="p-uw-item__description" > font style="vertical-align: inherit;" - Описание товара (TEXT)

	<a class="c-seller-info__name c-seller-info__name--link" > font class="vertical-align: inherit;" - Имя Пользователя (TEXT)

	<div class="c-seller-info__date" > font class="vertical-align: inherit;" - Дата регистрации (TEXT)
*/

func main() {
	t := time.Now()
	app.Run()

	fmt.Println(time.Since(t))

}
