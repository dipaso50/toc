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
	"tocV2/toc/application/organizebymodification"
	application "tocV2/toc/infraestructure/organizebymodification"

	"github.com/spf13/cobra"
)

var (
	configFilemod string
)

// organizeByModificationCmd represents the organizeByModification command
var organizeByModificationCmd = &cobra.Command{
	Use:   "organizeByModification",
	Short: "Organiza carpetas en base a la fecha de modificación de ficheros",
	Long:  `Organiza los ficheros de una carpeta en base a su fecha de modificación`,
	Run: func(cmd *cobra.Command, args []string) {
		organizer := application.OsOrganizer{}
		organizeService := organizebymodification.NewService(organizer)

		organizeService.Execute(configFilemod)
	},
}

func init() {
	rootCmd.AddCommand(organizeByModificationCmd)
	organizeByModificationCmd.Flags().StringVarP(&configFilemod, "folder", "f", "", "Carpeta que vamos a organizar")
	organizeByModificationCmd.MarkFlagRequired("folder")
}
