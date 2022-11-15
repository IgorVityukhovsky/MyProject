package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

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
				//fmt.Println("Будний день ", lookFor)
				date_slice = append(date_slice, lookFor)
			}

		} //else {
		//	fmt.Println("Выходной день ", next_date.Format("02.01.06"))
		//}
	}
	fmt.Println(date_slice)

	plan_date := today + start_time
	src_info_date := (date_slice[0]) + start_time
	multi_date := (date_slice[1]) + start_time
	delete_vm_date := (date_slice[2]) + start_time
	create_vm_date := (date_slice[3]) + start_time
	setup_os_win_date := (date_slice[4]) + start_time
	setup_os_nx_date := (date_slice[5]) + start_time
	db_monitoring_date := (date_slice[6]) + start_time
	setup_src_date := (date_slice[7]) + start_time
	fmt.Println(plan_date)
	fmt.Println(src_info_date)
	fmt.Println(multi_date)
	fmt.Println(delete_vm_date)
	fmt.Println(create_vm_date)
	fmt.Println(setup_os_win_date)
	fmt.Println(setup_os_nx_date)
	fmt.Println(db_monitoring_date)
	fmt.Println(setup_src_date)

}
