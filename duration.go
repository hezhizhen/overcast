package overcast

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	utilz "github.com/hezhizhen/tiny-tools"
)

// Duration counts the duration of all episodes of the podcast
func Duration(url string) time.Duration {
	podcastPage, err := goquery.NewDocument(url)
	utilz.Check(err)
	durationInMinute := 0 // TODO: get more precise duration
	list := podcastPage.Find("body > div.container.pure-g > div > div > div.pure-u-1.pure-u-sm-3-5 > a")
	list.Each(func(i int, s *goquery.Selection) {
		text := s.Find("div.cellcontent > div.titlestack > div.caption2.singleline").Text()
		text = strings.TrimSpace(text)
		dateAndDuration := strings.Split(text, " • ")
		durationWithUnit := dateAndDuration[1]
		durationStr := strings.Split(durationWithUnit, " min")[0]
		duration, err := strconv.Atoi(durationStr)
		utilz.Check(err)
		durationInMinute += duration
	})
	d, err := time.ParseDuration(fmt.Sprintf("%dm", durationInMinute))
	utilz.Check(err)
	return d
}
