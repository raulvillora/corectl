package cmd

import (
	"fmt"
	"github.com/qlik-oss/corectl/pkg/urtag"

	"github.com/qlik-oss/corectl/internal"
	"github.com/qlik-oss/corectl/internal/log"
	"github.com/spf13/cobra"
)

var setScriptCmd = withLocalFlags(&cobra.Command{
	Use:     "set <path-to-script-file.qvs>",
	Args:    cobra.ExactArgs(1),
	Short:   "Set the script in the current app",
	Long:    "Set the script in the current app",
	Example: "corectl script set ./my-script-file.qvs",

	Run: func(ccmd *cobra.Command, args []string) {

		ctx, _, doc, params := urtag.NewCommunicator(ccmd).OpenAppSocket(true)
		scriptFile := args[0]
		if scriptFile != "" {
			internal.SetScript(ctx, doc, scriptFile)
		} else {
			log.Fatalln("no loadscript (.qvs) file specified.")
		}
		if !params.NoSave() {
			internal.Save(ctx, doc, params.NoData())
		}
	},
}, "no-save")

var getScriptCmd = &cobra.Command{
	Use:     "get",
	Args:    cobra.ExactArgs(0),
	Short:   "Print the reload script",
	Long:    "Print the reload script currently set in the app",
	Example: "corectl script get",

	Run: func(ccmd *cobra.Command, args []string) {
		ctx, _, doc, _ := urtag.NewCommunicator(ccmd).OpenAppSocket(false)
		script, err := doc.GetScript(ctx)
		if err != nil {
			log.Fatalf("could not retrieve script: %s\n", err)
		}
		if len(script) == 0 { // This happens if the script is set to an empty file
			fmt.Println("The loadscript is empty")
		} else {
			fmt.Println(script)
		}
	},
}

var scriptCmd = &cobra.Command{
	Use:   "script",
	Short: "Explore and manage the script",
	Long:  "Explore and manage the script",
	Annotations: map[string]string{
		"command_category": "sub",
	},
}

func init() {
	scriptCmd.AddCommand(setScriptCmd, getScriptCmd)
}
