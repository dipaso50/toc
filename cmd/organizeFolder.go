package cmd

import (
	"tocV2/toc/application/organizebyrule"
	application "tocV2/toc/infraestructure/organizebyrule"

	"github.com/spf13/cobra"
)

var (
	configFile string
	folder     string
)

// organizeFolderCmd represents the organizeFolder command
var organizeFolderCmd = &cobra.Command{
	Use:   "byRule",
	Short: "Organiza carpetas en base a reglas",
	Long:  `Organiza carpetas en base a reglas definidas en el fichero json pasado como parámetro`,
	Run: func(cmd *cobra.Command, args []string) {

		rules := application.ReadConfig(configFile)

		ruleManager := organizebyrule.NewRuleManager()

		for _, r := range rules.Rules {
			ruleManager.AddRule(r.Expresion, r.Folder)
		}

		organizer := application.OsOrganizer{}
		organizeService := organizebyrule.NewService(organizer, ruleManager)

		organizeService.Execute(folder)
	},
}

func init() {
	rootCmd.AddCommand(organizeFolderCmd)
	organizeFolderCmd.Flags().StringVarP(&folder, "folder", "f", "", "Carpeta que vamos a organizar")
	organizeFolderCmd.MarkFlagRequired("folder")

	organizeFolderCmd.Flags().StringVarP(&configFile, "config", "c", "", "Fichero de configuración de reglas json")
	organizeFolderCmd.MarkFlagRequired("config")
}
