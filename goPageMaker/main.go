package main

import (
	// "bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"reflect"
	"strings"
	// "reflect"
)

// tell program page or have it have a csv of pages
// use csv and have like 10 max per page
// have it then use assets to place into page

func main() {

	print("Running")
	// reads what files are in the assets folder
	assets := readAssetsDirectory()
	count := len(assets)
	// isFolder := checkIsFolder(count, assets)

	/*
		currently, assets is everything that is in the folder assets
		I want it to save a list of all assets in the folder in a file locally

		to do this i have to read the assets csv file and from it, compare against the assets slice
	*/

	// load csv and check which assets are new
	isNew := make([]bool, count)
	assetFile := readAssets("assets.csv")

	print(assetFile)
	preExistingAssets := assetFile.determineAssets()
	checkNewAssets(preExistingAssets, assets, isNew)

	writePagePreffix("file.md", 0)

	// if there are new files (files not present in the CSV file)

	createNewPage()

	// TODO reading: file, err := os.Open("../assets/file.txt") // For read access.

}

// func checkIsFolder(count int, assets []fs.DirEntry) (isFolder []bool) {
// 	isFolder = make([]bool, count)

// 	for i, asset := range assets {
// 		isFolder[i] = asset.IsDir()
// 	}

// 	return
// }

func checkNewAssets(preExistingAssets []string, assets []fs.DirEntry, newAssets []bool) {
	// loop through assets, loop through

	print(assets, preExistingAssets)

	for _, v := range assets {
		// loopts through assets
		// print("a", assets[0].IsDir())
		// print(i, v)

		if itemInArray(v.Name(), preExistingAssets) {

			print("inArray", v.Name())
		}

	}

	// for i, v := range preExistingAssets {

	// }
}

func itemInArray(item string, arr []string) (inArray bool) {
	inArray = false
	for _, v := range arr {
		if reflect.DeepEqual(v, item) {
			inArray = true
		}
	}
	return
}

func readAssets(filename string) (a AssetFile) {
	return AssetFile{CSVFile: readCSV(filename)}
}

// returns an array of headings and a 2d array of
func readCSV(fileName string) (csv CSVFile) {
	// read fileName into CSVFile
	fileContents := readFile(fileName)

	// go through each line in CSV and
	for i, csvCell := range fileContents {
		// print("csv:", csvCell)
		if i == 0 { // adds headings to headings attribute
			csv.headings = strings.Split(csvCell, ",")
		} else { // ads csv items to contents attribute
			csv.contents = append(csv.contents, strings.Split(csvCell, ","))
		}
	}

	return
}

func readFile(fileName string) (lines []string) {
	data, err := os.ReadFile(fileName) // For read access.
	checkError(err)

	oneLine := strings.ReplaceAll(string(data), "\r", "")
	// print(oneLine)
	// print()

	// for _,letter := range oneLine {
	// 	print(letter, string(letter))
	// }

	// var fileContents []string = strings.Split(oneLine, "\n")
	lines = strings.Split(oneLine, "\n")

	return
}

func readLine(file fs.File, b []byte) (lengthOfLine int, err error) {
	// read contents of a file
	lengthOfLine, err = file.Read(b)
	if err == io.EOF {
		print("End of file")
	} else if err != nil {
		// print(err)
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", lengthOfLine, b[:lengthOfLine])
	return
}

func writeFile(fileName, text string) {
	if err := os.WriteFile(fileName, []byte(text), 0666); err != nil {
		log.Fatal(err)
	}
}

func constructMarkdownLink(embed bool, displayText, path string) string {
	if embed {
		return fmt.Sprintf("![%s](%s)", displayText, path)
	}
	return fmt.Sprintf("[%s](%s)", displayText, path)
}

func appendToFile(file os.File) {
	// is supposed to add a
}

func createNewPage() {}

func writePagePreffix(fileName string, pageNumber int) error {
	// write to file:
	// Page #
	// prev next
	writeFile(fileName, fmt.Sprintf("Page %d", pageNumber))
	err := writePrevNextPage(fileName, pageNumber)

	return err
}

func writePrevNextPage(fileName string, pageNumber int) error {
	path := "../pages/"
	links := ""

	if pageNumber > 1 {

		links += constructMarkdownLink(false, "Page 1", (path + fmt.Sprintf("Page%d.md", (pageNumber-1))))
	}

	writeFile(fileName, links)

	return nil
}

type CSVFile struct {
	headings []string
	contents [][]string
}

func (c *CSVFile) getIndexOfColumn(header string) (index int, err error) {
	for i, heading := range c.headings {
		if reflect.DeepEqual(heading, header) {
			index = i
		}
	}

	return
}

type AssetFile struct {
	CSVFile
}

func (a *AssetFile) determineAssets() (assets []string) {
	// splits out the column in CSV file that refers to assets

	// determines column of asset column
	iOfAssets, _ := a.getIndexOfColumn("assetName")
	iOfFileType, _ := a.getIndexOfColumn("fileType")

	for i, row := range a.contents {

		if i > 0 {
			item := row[iOfAssets]

			if reflect.DeepEqual(row[iOfFileType], "folder") {
			} else {
				item += "." + row[iOfFileType]
			}
			assets = append(assets, item)
		}
	}

	print("ass", assets)

	return
}

func assetsCSVPath() string {
	return "assets.csv"
}