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
	var сhange_number string = "ИЗМ-000042067"
	var service string
	var coordinator string = "Координатор"

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
		chromedp.Click(`#ext-gen-top356`),                //кнопка поиска
		chromedp.WaitVisible(`#ext-gen-top408`),          //ждём кнопку печати
		chromedp.Click(`#X3Button`),                      //стрелка вниз (развернуть)
		chromedp.Click(`#X3Popup_6`),                     //изменение
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
	if service == "РБТ-Переустановка/ смена ОС на ВМ-" {
		if err := chromedp.Run(ctx,
			chromedp.WaitVisible(`#X167_t`), //план работ
			chromedp.Click(`#X167_t`),
			chromedp.WaitVisible(`#X176_1`), //ЗНР планирование
			chromedp.Click(`#X176_1`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.WaitVisible(`#X12`),     //исполнитель
			chromedp.Click(`#X12`),
			chromedp.SendKeys(`#X12`, kb.Home),
			chromedp.SendKeys(`#X12`, kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete),
			chromedp.KeyEvent(coordinator), //chromedp.SetValue(`#X12`, coordinator), #X26
			chromedp.Click(`#X26`),         //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete),
			chromedp.KeyEvent("01.11.22 09:00:00"),
			chromedp.KeyEvent(kb.Tab+kb.Tab+kb.Tab+kb.Tab+kb.Tab+kb.Tab+kb.Tab+kb.Tab+kb.Tab+kb.Tab+kb.Tab),
			chromedp.Click(`#X55`), //результат выполнения
			chromedp.SendKeys(`#X55`, kb.Home),
			chromedp.SendKeys(`#X55`, kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete+kb.Delete),
			chromedp.KeyEvent("Test"),
			chromedp.Click(`#ext-gen-top474`), //сохранить и выйти
		); err != nil {
			log.Fatal(err)
		}
	}
	time.Sleep(time.Hour)
}
