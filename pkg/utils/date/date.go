package date

import (
	"fmt"
	"time"
)

var longMounthNamesFr = []string{
	"Janvier",
	"Février",
	"Mars",
	"Avril",
	"Mai",
	"Juin",
	"Juillet",
	"Août",
	"Septembre",
	"Octobre",
	"Novembre",
	"Décembre",
}

func FormatDateFr(d *time.Time) string {
	return fmt.Sprintf("%d %s %d", d.Day(), longMounthNamesFr[d.Month()-1], d.Year())
}
