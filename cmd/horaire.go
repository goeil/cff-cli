/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cff/api"
	"cff/tui"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// horaireCmd represents the horaire command
var horaireCmd = &cobra.Command{
	Use:     "horaire",
	Short:   "A brief description of your command",
	Version: "1.2, C. Jung, 11.07.2024",

	Run: func(cmd *cobra.Command, args []string) {

		params := api.RequestParameters{}

		switch len(args) {
		case 0, 1:
			os.Exit(1)

		case 2:

			params.Start = args[0]
			params.End = args[1]

		case 3:
			params.Start = args[0]
			params.End = args[1]
			params.Time = args[2]

		case 4:
			params.Start = args[0]
			params.End = args[1]
			params.Time = args[2]
			params.Date = args[3]
		}
		loader := tui.InitialLoader()
		p := tea.NewProgram(&loader)

		//go func() { // GOLD mettre ça dans une routine asynchrone permet d'envoyer des Msg !
		go func() {

			loader.GetClient(params)
			loader.Unmarshal()
			p.Send(tui.FetchedApi(""))
		}()

		if _, err := p.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		horaires := loader.GetHoraires()

		list := tui.New(horaires)

		/*/pour un bubble interactif :
		p := tea.NewProgram(list)
		if _, err := p.Run(); err != nil {
			fmt.Println(err)
		}*/

		fmt.Println(list.ToString())

	},
}

func init() {
	rootCmd.AddCommand(horaireCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// horaireCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// horaireCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
