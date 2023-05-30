package test

import (
        "bytes"
        "io"
        "log"
        "os"
        "os/exec"
        "path/filepath"
        "sync"
        "testing"

        "github.com/google/go-cmdtest"
)

var once sync.Once
var absPath string

var testModFileNames []string

func setup() {

        abs, err := filepath.Abs("")

        if err != nil {
            log.Fatalf("Error navigating to makefile: %v", err)
        }

        absPath = abs

		// Building CLI Binary for the tests to use
        makeFilePath := filepath.Join(abs, "../Makefile")
        cmd := exec.Command("make", "build", "-f", makeFilePath)

        var stdout bytes.Buffer
        var stderr bytes.Buffer
        cmd.Stdout = &stdout
        cmd.Stderr = &stderr

        if err := cmd.Run(); err != nil {
                log.Println(stdout.String())
                log.Println(stderr.String())
                log.Fatalf("Error building attribution-gen: %v", err.Error())
        }

        // Go Mod files used to test generator
        testModFileNames = []string{
                "top_level.go.mod",
                "ignore_dir.go.mod",
                "skip_nonlicense.go.mod",
                "sub_dir.go.mod",
                "indirect_dep.go.mod",
        }
}

func TestCLI(t *testing.T) {
        once.Do(setup)
        testDataPath := filepath.Join(absPath, "../testdata")
        ts, err := cmdtest.Read(testDataPath)
        if err != nil {
                t.Fatal(err)
        }

        ts.Setup = func(rootDir string) error {
                for _, modFile := range testModFileNames {
                        err := copyFile(filepath.Join(testDataPath, modFile), filepath.Join(rootDir, modFile))
                        if err != nil {
                                return err
                        }
                }
                return nil
        }
        ts.Commands["attribution-gen"] = cmdtest.Program(filepath.Join(testDataPath, "../bin/attribution-gen"))
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