package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hamzajer/projy/utils" // Import our new utils package
)

func main() {
	log.SetFlags(0) // Remove timestamp from log messages

	// 1. Check Arguments
	if len(os.Args) != 3 {
		log.Fatal("Usage: projy <language> <path>")
	}
	language := os.Args[1]
	repoPath := os.Args[2]
	repoName := filepath.Base(repoPath) // Extract repo name from path

	// 2. Check if path already exists
	if _, err := os.Stat(repoPath); !os.IsNotExist(err) {
		log.Fatalf("Error: Path '%s' already exists.", repoPath)
	}

	// 3. Main creation logic
	var err error
	supportedLangs := []string{"go", "python", "node-api", "react", "java", "csharp", "rust", "c", "cpp"}

	switch language {
	case "go":
		err = utils.CreateGoRepo(repoPath, repoName)
	case "python":
		err = utils.CreatePythonRepo(repoPath, repoName)
	case "node-api":
		err = utils.CreateNodeApiRepo(repoPath, repoName)
	case "react":
		err = utils.CreateReactRepo(repoPath, repoName)
	case "java":
		err = utils.CreateJavaRepo(repoPath, repoName)
	case "csharp":
		err = utils.CreateCSharpRepo(repoPath, repoName)
	case "rust":
		err = utils.CreateRustRepo(repoPath, repoName)
	case "c":
		err = utils.CreateCRepo(repoPath, repoName)
	case "cpp":
		err = utils.CreateCppRepo(repoPath, repoName)
	default:
		log.Fatalf("Error: Unsupported language '%s'. Supported: [%s]", language, strings.Join(supportedLangs, ", "))
	}

	// 4. Final error check
	if err != nil {
		// If creation failed, try to clean up the partially created directory
		os.RemoveAll(repoPath)
		log.Fatalf("Error creating repo: %s", err)
	}

	fmt.Printf("Successfully created '%s' project at: %s\n", language, repoPath)
	fmt.Println("Happy coding!")
}
