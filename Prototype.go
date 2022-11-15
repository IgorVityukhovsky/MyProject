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
	var сhange_number string = "ИЗМ-000042688"
	var service string
	var coordinator string = "Витюховский Игорь (Igor.Vityukhovsky)"
	var many_delete = "\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f\x7f"
	var save_and_exit = `/html/body/div[3]/div[2]/div/div[2]/div[1]/div/div[5]/div/div[1]/div/table/tbody/tr/td[1]/table/tbody/tr/td[2]/table/tbody/tr[2]/td[2]/em/button`

	//fmt.Println("Введите номер ИЗМ: ")
	//fmt.Scanf("%s\n", &сhange_number)

	today := (time.Now()).Format("02.01.06")
	date_slice := []string{}
	var holidays_string string = ("02.01.23 03.01.23 04.01.23 05.01.23 06.01.23")
	start_time := " 09:00:00"

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

	plan_date := today + start_time
	src_info_date := (date_slice[0]) + start_time
	multi_date := (date_slice[1]) + start_time
	delete_vm_date := (date_slice[2]) + start_time
	create_vm_date := (date_slice[3]) + start_time
	setup_os_win_date := (date_slice[4]) + start_time
	setup_os_nx_date := (date_slice[5]) + start_time
	db_monitoring_date := (date_slice[6]) + start_time
	setup_src_date := (date_slice[7]) + start_time

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
			chromedp.SendKeys(`#X12`, kb.Delete+many_delete),
			chromedp.KeyEvent(coordinator), //chromedp.SetValue(`#X12`, coordinator)
			chromedp.Click(`#X26`),         //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(plan_date), //вводим дату для ЗНР планирование
			chromedp.KeyEvent(kb.Tab+"\t\t\t\t\t\t\t\t\t\t"),
			chromedp.Click(`#X55`), //результат выполнения
			chromedp.SendKeys(`#X55`, kb.Home),
			chromedp.SendKeys(`#X55`, kb.Delete+many_delete),
			chromedp.KeyEvent("Успешно"),
			chromedp.Click(save_and_exit),

			// ЗНР 2

			chromedp.WaitVisible(`#X176_2Border`),
			chromedp.Click(`#X176_2Border`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(src_info_date), //вводим дату для ЗНР 2
			chromedp.Click(save_and_exit),

			// ЗНР 3

			chromedp.WaitVisible(`#X176_3`),
			chromedp.Click(`#X176_3`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(multi_date), //вводим дату для ЗНР 3
			chromedp.Click(save_and_exit),

			// ЗНР 4

			chromedp.WaitVisible(`#X176_4`),
			chromedp.Click(`#X176_4`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(multi_date), //вводим дату для ЗНР 4
			chromedp.Click(save_and_exit),

			// ЗНР 5

			chromedp.WaitVisible(`#X176_5`),
			chromedp.Click(`#X176_5`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(multi_date), //вводим дату для ЗНР 5
			chromedp.Click(save_and_exit),

			// ЗНР 6

			chromedp.WaitVisible(`#X176_6`),
			chromedp.Click(`#X176_6`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(delete_vm_date), //вводим дату для ЗНР 6
			chromedp.Click(save_and_exit),

			// ЗНР 7

			chromedp.WaitVisible(`#X176_7`),
			chromedp.Click(`#X176_7`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(create_vm_date), //вводим дату для ЗНР 7
			chromedp.Click(save_and_exit),

			// ЗНР 8

			chromedp.WaitVisible(`#X176_8`),
			chromedp.Click(`#X176_8`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(setup_os_win_date), //вводим дату для ЗНР 8
			chromedp.Click(save_and_exit),

			// ЗНР 9

			chromedp.WaitVisible(`#X176_9`),
			chromedp.Click(`#X176_9`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(setup_os_nx_date), //вводим дату для ЗНР 9
			chromedp.Click(save_and_exit),

			// ЗНР 10

			chromedp.WaitVisible(`#X176_10`),
			chromedp.Click(`#X176_10`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(db_monitoring_date), //вводим дату для ЗНР 10
			chromedp.Click(save_and_exit),

			// ЗНР 11

			chromedp.WaitVisible(`#X176_11`),
			chromedp.Click(`#X176_11`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(setup_src_date), //вводим дату для ЗНР 11
			chromedp.Click(save_and_exit),
		); err != nil {
			log.Fatal(err)
		}
	}
	if service == "РБТ-Перенос ВМ в DMZ-" {
		if err := chromedp.Run(ctx,
			chromedp.WaitVisible(`#X167_t`), //план работ
			chromedp.Click(`#X167_t`),
			chromedp.WaitVisible(`#X176_1`), //ЗНР планирование
			chromedp.Click(`#X176_1`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.WaitVisible(`#X12`),     //исполнитель
			chromedp.Click(`#X12`),
			chromedp.SendKeys(`#X12`, kb.Home),
			chromedp.SendKeys(`#X12`, kb.Delete+many_delete),
			chromedp.KeyEvent(coordinator), //chromedp.SetValue(`#X12`, coordinator)
			chromedp.Click(`#X26`),         //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(plan_date), //вводим дату для ЗНР планирование
			chromedp.KeyEvent(kb.Tab+"\t\t\t\t\t\t\t\t\t\t"),
			chromedp.Click(`#X55`), //результат выполнения
			chromedp.SendKeys(`#X55`, kb.Home),
			chromedp.SendKeys(`#X55`, kb.Delete+many_delete),
			chromedp.KeyEvent("Успешно"),
			chromedp.Click(save_and_exit),

			// ЗНР 2

			chromedp.WaitVisible(`#X176_2Border`),
			chromedp.Click(`#X176_2Border`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(src_info_date), //вводим дату для ЗНР 2
			chromedp.Click(save_and_exit),

			// ЗНР 3

			chromedp.WaitVisible(`#X176_3`),
			chromedp.Click(`#X176_3`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(multi_date), //вводим дату для ЗНР 3
			chromedp.Click(save_and_exit),

			// ЗНР 4

			chromedp.WaitVisible(`#X176_4`),
			chromedp.Click(`#X176_4`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(delete_vm_date), //вводим дату для ЗНР 4
			chromedp.Click(save_and_exit),

			// ЗНР 5

			chromedp.WaitVisible(`#X176_5`),
			chromedp.Click(`#X176_5`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(create_vm_date), //вводим дату для ЗНР 5
			chromedp.Click(save_and_exit),

			// ЗНР 6

			chromedp.WaitVisible(`#X176_6`),
			chromedp.Click(`#X176_6`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(setup_os_win_date), //вводим дату для ЗНР 6
			chromedp.Click(save_and_exit),

			// ЗНР 7

			chromedp.WaitVisible(`#X176_7`),
			chromedp.Click(`#X176_7`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(setup_os_nx_date), //вводим дату для ЗНР 7
			chromedp.Click(save_and_exit),

			// ЗНР 8

			chromedp.WaitVisible(`#X176_8`),
			chromedp.Click(`#X176_8`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(db_monitoring_date), //вводим дату для ЗНР 8
			chromedp.Click(save_and_exit),

			// ЗНР 9

			chromedp.WaitVisible(`#X176_9`),
			chromedp.Click(`#X176_9`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(setup_src_date), //вводим дату для ЗНР 9
			chromedp.Click(save_and_exit),
		); err != nil {
			log.Fatal(err)
		}
	}
	if service == "РБТ-Снятие ресурса с внешней публикации-" || service == "РБТ-Публикация приложения на WAF-" {
		if err := chromedp.Run(ctx,
			chromedp.WaitVisible(`#X167_t`), //план работ
			chromedp.Click(`#X167_t`),
			chromedp.WaitVisible(`#X176_1`), //ЗНР планирование
			chromedp.Click(`#X176_1`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.WaitVisible(`#X12`),     //исполнитель
			chromedp.Click(`#X12`),
			chromedp.SendKeys(`#X12`, kb.Home),
			chromedp.SendKeys(`#X12`, kb.Delete+many_delete),
			chromedp.KeyEvent(coordinator), //chromedp.SetValue(`#X12`, coordinator)
			chromedp.Click(`#X26`),         //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(plan_date), //вводим дату для ЗНР планирование
			chromedp.KeyEvent(kb.Tab+"\t\t\t\t\t\t\t\t\t\t"),
			chromedp.Click(`#X55`), //результат выполнения
			chromedp.SendKeys(`#X55`, kb.Home),
			chromedp.SendKeys(`#X55`, kb.Delete+many_delete),
			chromedp.KeyEvent("Успешно"),
			chromedp.Click(save_and_exit),

			// ЗНР 2

			chromedp.WaitVisible(`#X176_2Border`),
			chromedp.Click(`#X176_2Border`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(src_info_date), //вводим дату для ЗНР 2
			chromedp.Click(save_and_exit),

			// ЗНР 3

			chromedp.WaitVisible(`#X176_3`),
			chromedp.Click(`#X176_3`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(multi_date), //вводим дату для ЗНР 3
			chromedp.Click(save_and_exit),

			// ЗНР 4

			chromedp.WaitVisible(`#X176_4`),
			chromedp.Click(`#X176_4`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(delete_vm_date), //вводим дату для ЗНР 4
			chromedp.Click(save_and_exit),
		); err != nil {
			log.Fatal(err)
		}
	}
	if service == "РБТ-Изменение политики безопасности WAF-" || service == "РБТ-Удаление DNS записи в публичном домене-" || service == "РБТ-Добавление DNS записи в публичном домене-" {
		if err := chromedp.Run(ctx,
			chromedp.WaitVisible(`#X167_t`), //план работ
			chromedp.Click(`#X167_t`),
			chromedp.WaitVisible(`#X176_1`), //ЗНР планирование
			chromedp.Click(`#X176_1`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.WaitVisible(`#X12`),     //исполнитель
			chromedp.Click(`#X12`),
			chromedp.SendKeys(`#X12`, kb.Home),
			chromedp.SendKeys(`#X12`, kb.Delete+many_delete),
			chromedp.KeyEvent(coordinator), //chromedp.SetValue(`#X12`, coordinator)
			chromedp.Click(`#X26`),         //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(plan_date), //вводим дату для ЗНР планирование
			chromedp.KeyEvent(kb.Tab+"\t\t\t\t\t\t\t\t\t\t"),
			chromedp.Click(`#X55`), //результат выполнения
			chromedp.SendKeys(`#X55`, kb.Home),
			chromedp.SendKeys(`#X55`, kb.Delete+many_delete),
			chromedp.KeyEvent("Успешно"),
			chromedp.Click(save_and_exit),

			// ЗНР 2

			chromedp.WaitVisible(`#X176_2Border`),
			chromedp.Click(`#X176_2Border`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(src_info_date), //вводим дату для ЗНР 2
			chromedp.Click(save_and_exit),

			// ЗНР 3

			chromedp.WaitVisible(`#X176_3`),
			chromedp.Click(`#X176_3`),
			chromedp.WaitVisible(`#X44Icon`), //ждать "к исполнению" (гарантия загрузки нужной страницы)
			chromedp.Click(`#X26`),           //дата
			chromedp.SendKeys(`#X26`, kb.Home),
			chromedp.SendKeys(`#X26`, kb.Delete+many_delete),
			chromedp.KeyEvent(multi_date), //вводим дату для ЗНР 3
			chromedp.Click(save_and_exit),
		); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Работы успешно распланированы")
	time.Sleep(time.Hour)
}
