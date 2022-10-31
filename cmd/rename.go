package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename file/folder",
	Run: func(cmd *cobra.Command, args []string) {
		oldName := selectFilePrompt()
		newName := stringPrompt(oldName)
		terminal := exec.Command("mv", oldName, newName)
		err := terminal.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Succesfully rename %v to %v \n", oldName, newName)
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func stringPrompt(s string) string {
	prompt := promptui.Prompt{
		Label: fmt.Sprintf("Rename %v to ", s),
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func selectFilePrompt() string {
	f, _ := os.Open(".")
	files, _ := f.ReadDir(0)
	fileNames := make([]string, 0, len(files))
	for _, v := range files {
		var vName string
		if v.IsDir() {
			vName = fmt.Sprintf("%v/", v.Name())
		} else {
			vName = v.Name()
		}
		fileNames = append(fileNames, vName)
	}
	prompt := promptui.Select{
		Label: "Select File/Folder",
		Items: fileNames,
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	return result
}
