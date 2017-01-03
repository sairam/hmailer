package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	// file, err := os.Open("sample.html") // For read access.
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// doc, err := html.Parse(file)
	doc, err := goquery.NewDocument("https://news.ycombinator.com/item?id=13301833")
	// doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}

	tag := `tr.athing.comtr>td>table>tbody>tr>td:nth-of-type(3)>div.comment>span.c00`

	doc.Find(tag).Each(func(i int, s *goquery.Selection) {
		s.Children().Last().Remove()
		fmt.Println(s.Nodes[0].FirstChild.Data)
		s.Children().Each(func(_ int, a *goquery.Selection) {
			fmt.Println(a.Text())
		})
		fmt.Println("--------------------------------------------------------------------------------------------")

	})

}
