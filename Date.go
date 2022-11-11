package main

import (
	"fmt"
	"time"
)

func main() {
	//var n int = 3
	// weekday := time.Now().Weekday()

	// fmt.Println(weekday)      // "Tuesday"
	// fmt.Println(int(weekday)) // "2"

	// l, err := time.LoadLocation("Europe/Vienna")
	// if err != nil {
	// 	panic(err)
	// }
	t := (time.Now()) //.Format("02.01.06")
	fmt.Printf("Дата по умолчанию: %v\n", t)

	// Добавить 3 дня
	//r1 := t.Add(72 * time.Hour) работает
	// r1 := t.AddDate(0, 0, n)
	// fmt.Printf("Добавили +3 дня: %v\n", r1.Format("02.01.06"))
	// fmt.Printf("День недели: ", r1.Weekday())
	date_slice := []time.Time{}
	holidays_slice := []string{"16.11.22", "17.11.22"}

	for n := 1; (len(date_slice)) <= 8; n++ {

		date := (time.Now())
		next_date := date.AddDate(0, 0, n)
		if (int(next_date.Weekday())) != 6 && (int(next_date.Weekday())) != 0 {

			fmt.Println("Будний день ", next_date.Format("02.01.06"))
			date_slice = append(date_slice, next_date)

		} else {
			fmt.Println("Выходной день ", next_date.Format("02.01.06"))
		}
	}
	fmt.Println(date_slice)

}
