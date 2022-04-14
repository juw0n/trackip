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
	IP              string  `json:"ip"`
	Type            string  `json:"type"`
	Continent       string  `json:"continent"`
	Country         string  `json:"country"`
	Country_Code    string  `json:"country_code"`
	Country_Capital string  `json:"country_capital"`
	Country_Phone   string  `json:"country_phone"`
	Region          string  `json:"region"`
	City            string  `json:"city"`
	Latitude        float32 `json:"latitude"`
	Longitude       float32 `json:"longitude"`
	Timezone        string  `json:"timezone"`
	TimezoneName    string  `json:"timezone_name"`
	Currency        string  `json:"currency"`
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
	url := "http://ipwhois.app/json/" + ip
	readResponse := getData(url)

	data := IpDetail{}

	err := json.Unmarshal(readResponse, &data)
	if err != nil {
		log.Println("Unable to unmashall the response")
	}

	d := color.New(color.FgCyan, color.Bold).Add(color.Underline)
	d.Println("IP Found :")
	fmt.Printf("IP :%s\nTYPE :%s\nCONTINENT :%s\nCOUNTRY :%s\nCOUNTRY_CODE :%s\nCOUNTRY_CAPITAL :%s\nCOUNTRY_PHONE :%s\nREGION :%s\nCITY :%s\nLATITUDE :%v\nLONGITUDE :%v\nTIMEZONE :%s\nTIMEZONE_NAME :%s\nCURRENCY :%s\n", data.IP, data.Type, data.Continent, data.Country, data.Country_Code, data.Country_Capital, data.Country_Phone, data.Region, data.City, data.Latitude, data.Longitude, data.Timezone, data.TimezoneName, data.Currency)
}

func init() {
	rootCmd.AddCommand(traceCmd)
}
