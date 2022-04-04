/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the given IP",
	Long:  `This command trace the given IP address and return it basic details`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("Please enter a valid IP address")
		}
	},
}

type IpDetail struct {
	IP           string `json:"ip"`
	City         string `json:"city"`
	Region       string `json:"region"`
	Country      string `json:"country"`
	Loc          string `json:"loc"`
	Postal       string `json:"postal"`
	Timezone     string `json:"timezone"`
	Organization string `json:"organization"`
}

func getData(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get response from the given IP")
	}
	readResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read the response")
	}
	return readResponse
}

func showData(ip string) {
	url := "https://ipinfo.io/" + ip + "/geo"
	readResponse := getData(url)

	data := IpDetail{}

	err := json.Unmarshal(readResponse, &data)
	if err != nil {
		log.Println("Unable to unmashall the response")
	}

	d := color.New(color.FgCyan, color.Bold).Add(color.Underline)
	d.Println("IP Found :")
	fmt.Printf("IP :%s\nCITY :%s\nREGION :%s\nCOUNTRY :%s\nLOCATION :%s\nPOSTAL :%s\nTIMEZONE :%s\n", data.IP, data.City, data.Region, data.Country, data.Loc, data.Postal, data.Timezone)
}

func init() {
	rootCmd.AddCommand(traceCmd)
}
