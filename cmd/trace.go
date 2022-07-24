/*
Copyright © 2022 NAME HERE <akumar00029@gmail.com>

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
	Short: "Trace the IP",
	Long:  `Trace IP.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide IP to trace")
			return
		}
		for _, ip := range args {
			showData(ip)
			fmt.Println("------------------------------------------------")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type IP struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	Anycast  bool   `json:"anycast"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"

	responseBytes := getData(url)

	data := &IP{}

	err := json.Unmarshal(responseBytes, &data)

	if err != nil {
		log.Println("Unable to unmarshal response")
		return
	}

	c := color.New(color.FgRed).Add(color.Bold)

	c.Println(":::::::::::::::::::::::::::")
	c.Println(":::     DATA Found      :::")
	c.Println(":::::::::::::::::::::::::::")

	fmt.Println("IP       : " + data.IP)
	fmt.Println("Hostname : " + data.Hostname)
	fmt.Println("City     : " + data.City)
	fmt.Println("Region   : " + data.Region)
	fmt.Println("Country  : " + data.Country)
	fmt.Println("Loc      : " + data.Loc)
	fmt.Println("Org      : " + data.Org)
	fmt.Println("Postal   : " + data.Postal)
	fmt.Println("Timezone : " + data.Timezone)
	fmt.Println("Readme   : " + data.Readme)
}

func getData(url string) []byte {

	response, err := http.Get(url)

	if err != nil {
		log.Println("Unable to get response")
		return nil
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println("Unable to read response")
		return nil
	}

	return responseBytes

}
