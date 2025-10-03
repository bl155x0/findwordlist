/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type Wordlist struct {
  Id int
  Path string
}

var versionFlag bool
var source string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "findwordlist <search-pattern>",
  Short: "findwordlist: A small tool that simplifies working with wordlists in the shell.",
  Long: `findwordlist: A small tool that simplifies working with wordlists in the shell.`,
	Run: func(cmd *cobra.Command, args []string) { 

		//Just print the version and exit
		if(versionFlag) {
			fmt.Println("findwordlist version 0.3")
			return
		}

		//Determine the wordlist directory
    directory,err := getWordlistDirectory(source)
    cobra.CheckErr(err)

		//Search  wordlists
		wordlists, err := searchAllWordlists(directory)
    cobra.CheckErr(err)

		//call fzf and let the user select which wordlist to use
		selection, err := fzf(wordlists)
    cobra.CheckErr(err)

		//print the selection ready to be evaluated by a shell
		fmt.Printf("export W=%q\n", selection)
   },
}

func getHomeDirectory() (string, error) {
  user, err := user.Current()  
  if err != nil {
    return "", err
  }
  return user.HomeDir, nil
}

func searchAllWordlists (directory string) ([]Wordlist, error)  {
  var wordlists []Wordlist
  i := 0;
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Return if there's an error accessing a file
		}

		// Check if it's a file, not a directory
    if(!info.IsDir()){
			i++
			wordlist := Wordlist {
				Id: i,
				Path: path,
			}
			wordlists = append(wordlists, wordlist)
    }
		return nil
	})

  return wordlists, err;
}


// getWordlistDirectory determines the directory of wordlists to 
// It just returns the givenDirectory is not empty,
// otherwise it reads the WORDLISTS env variable and returns it's value 
// if the variable is not set, the directory "~/opt/wordlists/" is returned as fallback
func getWordlistDirectory(givenDirectory string) (string,error)  {
	//Use the given directory as highes priority
	if(givenDirectory != ""){
		return givenDirectory, nil
	}

	//Next check if the WORDLISTS env variable is set
	envWordlists := os.Getenv("WORDLISTS")
	if (envWordlists != "") {
		return envWordlists,nil
	}

	//If nothing fall back to ~/opt/wordlists
	home, err := getHomeDirectory()
	if err != nil {
		return "", err
	}
	return path.Join(home, "opt/wordlists"), nil
}

// call fzf and pass a list of wordlist
func fzf(wordlists []Wordlist) (string, error) {
	//prozess the given wordlis
	var lists []string
	for _, list := range wordlists {
		lists = append(lists, list.Path)
	}

	//call fzf
	cmd := exec.Command("fzf")
	cmd.Stdin = strings.NewReader(strings.Join(lists, "\n"))
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	selection := strings.TrimSpace(string(out))
	return selection, nil
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
	rootCmd.Flags().StringVarP(&source,"source", "s", "", "Specifies the directory to search for wordlist files")
	rootCmd.Flags().BoolVar(&versionFlag,"version", false, "Prints version and exit")
}

