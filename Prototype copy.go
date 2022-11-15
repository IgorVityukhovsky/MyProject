//ForFirst use: go get -u github.com/chromedp/chromedp \

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func main() {
	l := os.Getenv("zzz")
	p := os.Getenv("zz")
	var сhange_number string = "ИЗМ-000042663"
	var service string
	//var coordinator string = "Витюховский Игорь (Igor.Vityukhovsky)"

	//fmt.Println("Введите номер ИЗМ: ")
	//fmt.Scanf("%s\n", &сhange_number)

	//today := (time.Now()).Format("02.01.06")
	date_slice := []string{}
	var holidays_string string = ("02.01.23 03.01.23 04.01.23 05.01.23 06.01.23")
	//start_time := " 09:00:00"

	for n := 1; (len(date_slice)) <= 8; n++ {

		date := (time.Now())
		next_date := date.AddDate(0, 0, n)
		if (int(next_date.Weekday())) != 6 && (int(next_date.Weekday())) != 0 {

			lookFor := next_date.Format("02.01.06")
			contain := strings.Contains(holidays_string, lookFor)

			if contain == false {
				date_slice = append(date_slice, lookFor)
			}

		}
	}

	//plan_date := today + start_time
	//src_info_date := (date_slice[0]) + start_time
	//multi_date := (date_slice[1]) + start_time
	//delete_vm_date := (date_slice[2]) + start_time
	//create_vm_date := (date_slice[3]) + start_time
	//setup_os_win_date := (date_slice[4]) + start_time
	//setup_os_nx_date := (date_slice[5]) + start_time
	//db_monitoring_date := (date_slice[6]) + start_time
	//setup_src_date := (date_slice[7]) + start_time

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
		//chromedp.Click(`#X3Button`),
		//chromedp.Click(`aria-label='Сохранить запись и выйти. (Ctrl+Shift+F2)'`, chromedp.NodeVisible),
		chromedp.Click(`/html/body/div[3]/div[2]/div/div[2]/div[1]/div/div[5]/div/div[1]/div/table/tbody/tr/td[1]/table/tbody/tr/td[2]/table/tbody/tr[2]/td[2]/em/button`, chromedp.NodeVisible),
	); err != nil {
		log.Fatal(err)
	}

}
