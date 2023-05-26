package test

import (
	"bytes"
	"log"
	"os/exec"
	"sync"
	"testing"

	"io"
	"os"
	"path/filepath"

	"github.com/google/go-cmdtest"
)

var once sync.Once
var absPath string

var testModFileNames []string


func setup() {

	testModFileNames = make([]string, 0)
	abs, err := filepath.Abs("")
	
	if err != nil {
		log.Fatalf("Error navigating to makefile: %v",err)
	}

	absPath = abs

	makeFilePath := filepath.Join(abs,"../Makefile")
	cmd := exec.Command("make", "build", "-f", makeFilePath)

	var stdout bytes.Buffer
	var sterr bytes.Buffer
    cmd.Stdout = &stdout
	cmd.Stderr = &sterr

	if err := cmd.Run(); err != nil {
		log.Println(stdout.String())
		log.Println(sterr.String())
		log.Fatalf("Error building gen-attributions: %v", err)
	}

	testModFileNames = append(testModFileNames, "top_level.go.mod" )
	testModFileNames = append(testModFileNames, "ignore_dir.go.mod" )
	testModFileNames = append(testModFileNames, "skip_nonlicense.go.mod" )
	testModFileNames = append(testModFileNames, "sub_dir.go.mod" )
	testModFileNames = append(testModFileNames, "indirect_dep.go.mod" )
}

func TestCLI(t *testing.T) {
	once.Do(setup)
	testDataPath := filepath.Join(absPath, "../testdata")
	ts, err := cmdtest.Read(testDataPath)
	if err != nil {
        t.Fatal(err)
    }
	
	ts.Setup = func(rootDir string) error {
		for _, modFile := range testModFileNames{

			err := copyFile(filepath.Join(testDataPath,modFile),filepath.Join(rootDir,modFile))

			if err != nil{
				return err
			}
		}
		return nil
	}
	ts.Commands["gen-attributions"] = cmdtest.Program(filepath.Join(testDataPath,"../bin/gen-attributions"))
	ts.Run(t, false)
}

func copyFile(sourceFile, destinationFile string) error {
	// Open the source file

	src, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer src.Close()

	// Create the destination file
	dest, err := os.Create(destinationFile)
	if err != nil {
		return err
	}
	defer dest.Close()

	// Copy the contents from source to destination
	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}

	return nil
}