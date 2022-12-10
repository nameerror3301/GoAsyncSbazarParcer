package app

import (
	"GoAsyncSbazarPrcer/internal/models"
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
)

func Run() {
	for i := 1; i <= 50; i++ {
		var (
			urlElectronic = fmt.Sprintf("https://www.sbazar.cz/30-elektro-pocitace/cela-cr/cena-neomezena/nejnovejsi/%s", strconv.Itoa(i))
		)

		if err := models.FindLinks(urlElectronic); err != nil {
			logrus.Println(err)
		}
	}
}
