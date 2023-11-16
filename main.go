package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	directory          = "."
	extensionList      string
	excludedExtensions string
	showHelpFlag       bool
)

func init() {
	flag.BoolVar(&showHelpFlag, "h", false, "Show help message")
	flag.BoolVar(&showHelpFlag, "help", false, "Show help message")
	flag.StringVar(&extensionList, "e", "", "List of extensions to include (comma-separated)")
	flag.StringVar(&excludedExtensions, "x", "", "List of extensions to include (comma-separated)")
	flag.Usage = showHelp
}

func main() {
	flag.Parse()

	if flag.NArg() > 0 || (extensionList != "" && excludedExtensions != "") {
		log.Fatal("Error: Invalid usage.")
	}

	if showHelpFlag {
		flag.Usage()
		return
	}

	if extensionList == "" && excludedExtensions == "" {
		countLinesAllFiles()
	} else if extensionList != "" {
		countLinesWithExtension()
	} else if excludedExtensions != "" {
		countLinesWithExcludedExtension()
	}
}

func showHelp() {
	fmt.Println("Welcome to catline!")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -e <extensions>   List of extensions to include (comma-separated)")
	fmt.Println("  -x <extensions>   List of extensions to exclude (comma-separated)")
	fmt.Println("  -h, -help         Show this help message")
	fmt.Println()
	fmt.Println("Arguments must be used separately. Combinations like '-e txt -x php'")
	fmt.Println("are not allowed.")
	fmt.Println()
	fmt.Println("To read all files within a directory, simply run the program without")
	fmt.Println("including any additional commands.")
}

func countLines(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var lineCount int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineCount++
	}

	return lineCount, nil
}

func countLinesAllFiles() {
	foundFiles := false

	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Dir(path) == directory {

			lineCount, err := countLines(path)
			if err != nil {
				log.Fatal("Error:", err)
			}

			if lineCount > 0 {
				foundFiles = true
				fmt.Printf("%-15s %d\n", path, lineCount)
			}
		}
		return nil
	})
	if !foundFiles {
		fmt.Println("No files found in the specified directory.")
	}
}

func countLinesWithExcludedExtension() {
	foundFiles := false
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Dir(path) == directory {

			if err != nil {
				log.Fatal("Error:", err)
			}

			extension := strings.ToLower(filepath.Ext(path))
			extension = strings.Replace(extension, ".", "", -1)

			if !strings.Contains(strings.ToLower(excludedExtensions), extension) {
				lineCount, err := countLines(path)
				if err != nil {
					log.Fatal("Error:", err)
				}

				if lineCount > 0 {
					foundFiles = true
					fmt.Printf("%-15s %d\n", path, lineCount)
				}
			}
		}
		return nil
	})

	if !foundFiles {
		fmt.Println("No files found in the specified directory.")
	}
}

func countLinesWithExtension() {
	foundFiles := false

	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Dir(path) == directory {
			extension := strings.ToLower(filepath.Ext(path))
			extension = strings.Replace(extension, ".", "", -1)

			if strings.Contains(strings.ToLower(extensionList), extension) {
				lineCount, err := countLines(path)

				if err != nil {
					log.Fatal("Error:", err)
				}

				if lineCount > 0 {
					foundFiles = true
					fmt.Printf("%-15s %d\n", path, lineCount)
				}
			}
		}
		return nil
	})

	if !foundFiles {
		fmt.Println("No files found in the specified directory.")
	}
}
