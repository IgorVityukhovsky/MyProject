//ForFirst use: go get -u github.com/chromedp/chromedp

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/chromedp"
)

func main() {
	//l := os.Getenv("zzz")
	//p := os.Getenv("zz")
	//var сhange_number string = "ИЗМ-000042185"
	var service string = "work"

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
		chromedp.Navigate(`https://www.w3schools.com/TAGS/tryit.asp?filename=tryhtml5_input_type_button`),
		chromedp.Click(`input[@value='Click me']`, chromedp.NodeVisible),
		chromedp.Sleep(5*time.Second),
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println(service)
	input.DispatchKeyEvent(input.KeyDown).WithKey("A").WithModifiers(input.ModifierCtrl)
	input.DispatchKeyEvent(input.KeyDown).WithKey("A").WithModifiers(input.ModifierCtrl)
	//input.DispatchKeyEvent(input.KeyDown).WithKey(kb.End)
	time.Sleep(time.Hour)
}
