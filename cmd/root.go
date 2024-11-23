/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Venukishore-R/caching-proxy/internal/app"
	"github.com/Venukishore-R/caching-proxy/internal/proxy-server/proxy"
	"github.com/spf13/cobra"
)

var (
	Port       string
	Origin     string
	ClearCache bool
	Proxy      *proxy.Proxy
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "caching-proxy",
	Short: "Caching server that caches responses from other server",
	Long: `Caching proxy server that forwards requests to the actual server 
and caches the responses. If the same request is made again, 
it will return the cached response instead of forwarding the request to the server.`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: "1.0.0",

	// Uncomment the following line if your application has a default action associated with it:

	Run: func(cmd *cobra.Command, args []string) {
		var port int
		var err error

		if Port != "" {
			port, err = strconv.Atoi(Port)
			if err != nil {
				fmt.Printf("Invalid port number: %v\n", err)
				return
			}
		}

		Proxy = proxy.NewProxy(Origin, ClearCache)

		server := &app.Server{
			Port:       port,
			Origin:     Origin,
			ClearCache: ClearCache,
			Proxy:      Proxy,
		}

		server.StartServer()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.caching-proxy.yaml)")

	rootCmd.PersistentFlags().BoolVarP(&ClearCache, "clear-cache", "c", false, "Clear cache of proxy server")
	rootCmd.PersistentFlags().StringVarP(&Port, "port", "p", "", "Port number to start cache server")
	rootCmd.PersistentFlags().StringVarP(&Origin, "origin", "o", "", "Origin URL of the server to cache the response")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
