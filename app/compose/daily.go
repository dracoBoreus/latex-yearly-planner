package compose

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/components/calendar"
	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func Daily(cfg config.Config, tpls []string) ([]page.Page, error) {
	if len(tpls) != 1 {
		return nil, fmt.Errorf("exppected one tpl, got %d %v", len(tpls), tpls)
	}

	pages := make([]page.Page, 0, 366)
	day := calendar.DayTime{Time: time.Date(cfg.Year, time.January, 1, 0, 0, 0, 0, time.Local)}

	for day.Year() == cfg.Year {
		right := header.Items{}
		_, weekNum := day.ISOWeek()
		left := header.Items{
			header.NewIntItem(cfg.Year),
			header.NewTextItem("Q" + strconv.Itoa(int(math.Ceil(float64(day.Month())/3.)))),
			header.NewMonthItem(day.Month()),
			header.NewTextItem("Week " + strconv.Itoa(weekNum)),
			header.NewTimeItem(day).SetLayout("Monday, 2").Ref(),
		}

		if day.Month() != time.January || day.Day() != 1 {
			right = append(right, header.NewTimeItem(day.AddDate(0, 0, -1)).SetLayout("Mon, 2"))
		}

		if day.Month() != time.December || day.Day() != 31 {
			right = append(right, header.NewTimeItem(day.AddDate(0, 0, 1)).SetLayout("Mon, 2"))
		}

		pages = append(pages, page.Page{
			Tpl: tpls[0],
			Header: header.Header{
				Left:  left,
				Right: right,
			},
			Body: day,
		})

		day = day.AddDate(0, 0, 1)
	}

	return pages, nil
}
