/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

var (
	year  int
	month int
)

var weekdays = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

func showCalendar() {
	for _, weekday := range weekdays {
		fmt.Printf("%5s", weekday)
	}
	fmt.Println()

	now := time.Now()
	showMonth := time.Month(month)
	showTime := time.Date(year, showMonth, 1, 0, 0, 0, 0, now.Location())
	for {
		// Sunday -> Saturday  0 -> 6
		startWd := showTime.Weekday()
		fmt.Printf("%s", strings.Repeat(" ", int(startWd)*5))
		for ; startWd <= time.Saturday; startWd++ {
			fmt.Printf("%5d", showTime.Day())
			showTime = showTime.Add(time.Hour * 24)
			if showTime.Month() != showMonth {
				return
			}
		}
		fmt.Println()
	}
}

// dateCmd represents the date command
var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if year < 1000 || year > 9999 {
			fmt.Fprintf(os.Stderr, "invalid year should in [1000, 9999], actual: %d", year)
			os.Exit(1)
		}
		if month < 1 || month > 12 {
			fmt.Fprintf(os.Stderr, "invalid month should in [1, 12], actual: %d", month)
			os.Exit(1)
		}

		// 显示指定月份的日历
		showCalendar()
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	dateCmd.PersistentFlags().IntVarP(
		&year, "year", "y", time.Now().Year(), "year to show(should in [1000, 9999])",
	)
	dateCmd.PersistentFlags().IntVarP(
		&month, "month", "m", int(time.Now().Month()), "month to show(should in [1, 12])",
	)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
