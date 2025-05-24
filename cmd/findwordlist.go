/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

var index int
var source string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "findwordlist <search-pattern>",
  Short: "findwordlist: A small tool that simplifies working with wordlists in the shell.",
  Long: `findwordlist: A small tool that simplifies working with wordlists in the shell.`,
	Run: func(cmd *cobra.Command, args []string) { 

    //is a searchstring given?
    searchstring := ""
    if(len(args) > 0){
      searchstring = args[0]
    }

    //using the given directory or the default directory in the user's $HOME
    directory := ""
    if (source != ""){
       directory = source
    } else {
      home, err := getHomeDirectory()
      cobra.CheckErr(err)

      directory = path.Join(home, "opt/wordlists")
    }

    //getting all files in the directory
    wordlists, err := searchAllWordlists(directory, searchstring)
    cobra.CheckErr(err)

    if (index > -1) {
      //print a Specific wordlist by ID
      for _, wordlist := range wordlists {
        if(wordlist.Id == index) {
          fmt.Printf("%s\n", wordlist.Path)
        }
      }
    } else {
      //Print all results
      for _, wordlist := range wordlists {
        lineCount, err := lineCount(wordlist.Path)
        cobra.CheckErr(err)
        //fmt.Printf("%d - %s - (%d lines)\n", wordlist.Id, wordlist.Path, lineCount)
        fmt.Printf("%s [#%d] [%d lines]\n", wordlist.Path, wordlist.Id, lineCount)
      }
    }
   },
}

func getHomeDirectory() (string, error) {
  user, err := user.Current()  
  if err != nil {
    return "", err
  }
  return user.HomeDir, nil
}

func searchAllWordlists (directory, searchstring string) ([]Wordlist, error)  {
  var wordlists []Wordlist
  i := 0;
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Return if there's an error accessing a file
		}

		// Check if it's a file, not a directory
    if(!info.IsDir()){
      i++
		  if (strings.Contains(strings.ToLower(path), strings.ToLower(searchstring))){

        wordlist := Wordlist {
          Id: i,
          Path: path,
        }
			  wordlists = append(wordlists, wordlist)
		  }
    }
		return nil
	})

  return wordlists, err;
}

func lineCount(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	count := 0
	for {
		_, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		count++
	}
	return count, nil
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
	rootCmd.Flags().IntVarP(&index, "index", "i", -1, "Specifies the index of a single wordlist")
}


