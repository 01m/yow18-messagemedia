// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/messagemedia/messages-go-sdk"
	"github.com/messagemedia/messages-go-sdk/messages_pkg"
	"github.com/messagemedia/messages-go-sdk/models_pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var message string
var to string

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("send called")

		if "" == message {
			fmt.Println("Error: message is empty")
			return
		}

		if "" == to {
			fmt.Println("Error: mobile number is empty")
			return
		}

		fmt.Println("Sending " + message + " to " + to)

		messagemedia_messages_sdk.Config.BasicAuthUserName = viper.GetString("api_key")
		messagemedia_messages_sdk.Config.BasicAuthPassword = viper.GetString("api_secret")

		messages := messages_pkg.NewMESSAGES()
		bodyValue := []byte(fmt.Sprintf(`{
			"messages": [{
				"content": "%s",
				"destination_number": "%s",
				"format": "SMS"
			}]
		}`, message, to))

		var body *models_pkg.SendMessagesRequest
		json.Unmarshal(bodyValue, &body)

		result, err := messages.CreateSendMessages(body)
		fmt.Println(result, err)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")
	sendCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "message")
	sendCmd.PersistentFlags().StringVarP(&to, "to", "t", "", "mobile number")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
