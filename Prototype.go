//ForFirst use: go get -u github.com/chromedp/chromedp \

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func main() {
	l := os.Getenv("zzz")
	p := os.Getenv("zz")
	var сhange_number string = "ИЗМ-000042185"
	var service string

	//fmt.Println("Введите номер ИЗМ: ")
	//fmt.Scanf("%s\n", &сhange_number)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf("https://%s:%s@itsm.x5.ru/sm/index.do", l, p)),

		chromedp.Click(`#ext-gen-top356`),       //кнопка поиска
		chromedp.WaitVisible(`#ext-gen-top408`), //ждём кнопку печати

		chromedp.Click(`#X3Button`),  //стрелка вниз (развернуть)
		chromedp.Click(`#X3Popup_6`), //изменение

		chromedp.WaitVisible(`#var\/choices\/openLabel`), //ждать доступность элемента (самый нижний чек)
		chromedp.Sleep(2*time.Second),
		chromedp.SendKeys(`//*[@id="X11"]`, сhange_number), //вписываем номер изменения
		chromedp.SendKeys(`//*[@id="X11"]`, (kb.Enter)),    //сработало один раз
		chromedp.WaitVisible(`#X15Readonly`),
		chromedp.Value(`#X15Readonly`, &service, chromedp.NodeVisible), //считывает тип работ и записывает в переменную service

	); err != nil {
		log.Fatal(err)
	}
	fmt.Println(service)
	if service == "РБТ-Публикация приложения на WAF-" {
		if err := chromedp.Run(ctx,
			chromedp.Click(`#X15Readonly`)); err != nil {
			log.Fatal(err)
		}
	}
	time.Sleep(time.Hour)
}
