package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/olekukonko/tablewriter"
)

type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Org      string `json:"org"`
	Country  string `json:"country"`
	Region   string `json:"region"`
	Zipcode  string `json:"postal"`
	Timezone string `json:"timezone"`
}

func main() {
	pipeFlag := flag.Bool("p", false, "print with a pipe delimiter instead of a table")
	flag.Parse()

	if len(flag.Args()) == 0 {
		s := `Get IP address metadata from IPInfo.io
Requires at least one IP address to lookup

Usage:
	whoip {-p} {ip-address} {ip-address}..

Options:
	-p	print with a pipe delimiter instead of a table`
		fmt.Println(s)
		return
	}

	var wg sync.WaitGroup
	ipStream := make(chan IPInfo)
	ipRegex := regexp.MustCompile(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`)

	for _, arg := range flag.Args() {
		ipResult := ipRegex.FindString(arg)
		if ipResult == "" {
			fmt.Printf("Skipping invalid IP address: '%s' \n", arg)
			continue
		}
		wg.Add(1)
		go getIPInfo(arg, ipStream, &wg)
	}

	// stream closer
	go func() {
		wg.Wait()
		close(ipStream)
	}()

	var rows []IPInfo
	for result := range ipStream {
		rows = append(rows, result)
	}

	if *pipeFlag {
		printInfo(rows)
	} else {
		printTableInfo(rows)
	}
}

// printInfo prints the IP address information with a pipe '|' delimiter
func printInfo(ipInfo []IPInfo) {
	for _, ip := range ipInfo {
		s := []string{
			ip.IP,
			ip.Country,
			ip.Region,
			ip.City,
			ip.Org,
			ip.Timezone,
			ip.Zipcode,
		}
		fmt.Println(strings.Join(s, "|"))
	}
}

// printTableInfo prints the IP address information in a formatted table
func printTableInfo(ipInfo []IPInfo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"IP", "COUNTRY", "REGION", "CITY", "ORGANIZATION", "TIMEZONE", "POSTAL"})
	for _, ip := range ipInfo {
		table.Append([]string{
			ip.IP,
			ip.Country,
			ip.Region,
			ip.City,
			ip.Org,
			ip.Timezone,
			ip.Zipcode,
		})
	}
	table.SetAlignment(1)
	table.Render()
}

// getIPInfo gets the IP address metadata from ipinfo.io
func getIPInfo(ip string, ch chan<- IPInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	var info IPInfo
	url := fmt.Sprintf("https://ipinfo.io/%s/json", ip)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Received error from ipinfo.io: %s \n", err.Error())
		return
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		fmt.Printf("Error decoding ipinfo.io response: %s \n", err.Error())
		return
	}

	ch <- info
}
