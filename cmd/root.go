// Copyright © 2017 Joseph Schneider <https://github.com/astropuffin>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var lastCheckin = time.Now()
var currentlyHealthy = true
var heartbeat int
var grace int


// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "webhealth",
	Short: "A simple healtchcheck webserver",
	Long: `A simple healthcheck webserver. It expects an inbound ping within the heartbeat interval.`,
	Run: func(cmd *cobra.Command, args []string) {
        do()
    },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
    RootCmd.Flags().IntVar(&heartbeat, "heartbeat", 10, "heartbeat interval in seconds")
    RootCmd.Flags().IntVar(&grace, "grace", 3, "number of intervals that can be missed before considered unhealthy")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
    if currentlyHealthy {
		io.WriteString(w, "ok")
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		io.WriteString(w, "not ok")
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	lastCheckin = time.Now()
	io.WriteString(w, "pong")
}

func updateStatus() {
	for {
	    if time.Now().Before(lastCheckin.Add(time.Duration(heartbeat * grace) * time.Second)) {
		    currentlyHealthy = true
			fmt.Println("healthy ", time.Now())
		} else {
		    currentlyHealthy = false
			fmt.Println("not healthy ", time.Now())
		}
		time.Sleep(time.Duration(heartbeat) * time.Second)
	}
}

func do() {
	fmt.Println("starting: ", time.Now())
	go updateStatus()
	http.HandleFunc("/health", healthcheck)
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(":8000", nil)
}
