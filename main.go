package main

import (
	"fmt"
	"github.com/JoelAOgden/Deliveroo/cron"
	"os"
	"text/tabwriter"
)

func main() {

	if len(os.Args) != 7 {
		panic(fmt.Sprintf("incorrect number of args, requires 6, given %v", len(os.Args)))
	}

	cronArgs := os.Args[1:6]

	service := cron.Service{}

	cronInfo, err := service.Parse(cronArgs)
	if err != nil {
		panic(err)
	}

	printCronInfo(*cronInfo, os.Args[6])
}

func printCronInfo(info cron.Info, cmd string) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Fprintln(w, fmt.Sprintf("Minute\t%v", info.Minute))
	fmt.Fprintln(w, fmt.Sprintf("Hour\t%v", info.Hour))
	fmt.Fprintln(w, fmt.Sprintf("Day Of The Month\t%v", info.DayOfTheMonth))
	fmt.Fprintln(w, fmt.Sprintf("Month\t%v", info.Month))
	fmt.Fprintln(w, fmt.Sprintf("Day Of The Week\t%v", info.DayOfTheWeek))
	fmt.Fprintln(w, fmt.Sprintf("Command\t%s", cmd))
	w.Flush()

}
