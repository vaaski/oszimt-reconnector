package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	cp "github.com/otiai10/copy"
)

const DIST_FOLDER = "./dist"
const PROGRAM_NAME = "oszimt-reconnector"

func main() {
	fmt.Println("creating mac bundle(s)...")

	distContents, err := os.ReadDir(DIST_FOLDER)
	if err != nil {
		panic(err)
	}

	darwinArchFolders := []string{}

	for _, file := range distContents {
		if file.IsDir() && strings.Contains(file.Name(), "darwin") {
			darwinArchFolders = append(darwinArchFolders, file.Name())
		}
	}

	for _, currentArch := range darwinArchFolders {
		fmt.Println("creating bundle for", currentArch)

		appFolder := path.Join(DIST_FOLDER, currentArch, PROGRAM_NAME+".app")
		currentExecutable := path.Join(DIST_FOLDER, currentArch, PROGRAM_NAME)

		cp.Copy("./mac-bundle/Contents", path.Join(appFolder, "Contents"))
		cp.Copy(currentExecutable, path.Join(appFolder, "Contents/MacOS/"+PROGRAM_NAME))
	}

}
