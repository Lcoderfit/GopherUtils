package main

import (
	"flag"
	"time"
)

/*
var (
    period time.Duration
)

func init() {
    flag.DurationVar(&period, "period", 1*time.Second, "sleep period")
}

func main() {
    flag.Parse()
	fmt.Printf("Sleeping for %v...", period)
	time.Sleep(period)
	fmt.Println()
}
*/

var (
	period time.Duration
)

func init() {
	flag.DurationVar(&period, "period", 1*time.Second, "sleep period")
}

func main() {

}
