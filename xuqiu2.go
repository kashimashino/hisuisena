package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://blog.lenconda.top/")
	if err != nil {
		panic("resp error!")
	}
	defer resp.Body.Close()

	f, err := os.Create("D:/learn.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	dom.Find("p").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("time").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("h1").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("span").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})



		resp, err = http.Get("https://blog.lenconda.top/page/2/")
		if err != nil {
			panic("resp error!")
		}
		defer resp.Body.Close()

		dom, err = goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			panic(err)
		}
		dom.Find("p").Each(func(i int, selection1 *goquery.Selection) {
			f.WriteString(selection1.Text()+"\n")
		})
		dom.Find("time").Each(func(i int, selection1 *goquery.Selection) {
			f.WriteString(selection1.Text()+"\n")
		})
		dom.Find("h1").Each(func(i int, selection1 *goquery.Selection) {
			f.WriteString(selection1.Text()+"\n")
		})
		dom.Find("span").Each(func(i int, selection1 *goquery.Selection) {
			f.WriteString(selection1.Text()+"\n")
		})

	resp, err = http.Get("https://blog.lenconda.top/page/3/")
	if err != nil {
		panic("resp error!")
	}
	defer resp.Body.Close()

	dom, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	dom.Find("p").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("time").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("h1").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("span").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})

	resp, err = http.Get("https://blog.lenconda.top/page/4/")
	if err != nil {
		panic("resp error!")
	}
	defer resp.Body.Close()

	dom, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	dom.Find("p").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("time").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("h1").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("span").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})

	resp, err = http.Get("https://blog.lenconda.top/page/5/")
	if err != nil {
		panic("resp error!")
	}
	defer resp.Body.Close()

	dom, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	dom.Find("p").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("time").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("h1").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("span").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})

	resp, err = http.Get("https://blog.lenconda.top/page/6/")
	if err != nil {
		panic("resp error!")
	}
	defer resp.Body.Close()

	dom, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	dom.Find("p").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("time").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("h1").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("span").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})

	resp, err = http.Get("https://blog.lenconda.top/page/7/")
	if err != nil {
		panic("resp error!")
	}
	defer resp.Body.Close()

	dom, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	dom.Find("p").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("time").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("h1").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
	dom.Find("span").Each(func(i int, selection1 *goquery.Selection) {
		f.WriteString(selection1.Text()+"\n")
	})
}