/*
Copyright 2018 The nullhash Authors.
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
	"log"

	"github.com/nullhash/kcm/kcmmanager/context"

	"github.com/spf13/cobra"
)

var useContext = &cobra.Command{
	Use:   "use",
	Short: "This command is for kubeconfig from kcm",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("no context provided")
			return
		}
		if len(args) > 1 {
			log.Println("more than one context not supported")
			return
		}
		context.UseContext(args[0])
	},
}
