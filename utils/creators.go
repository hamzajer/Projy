package utils

import (
	"fmt"
	"strings"
)

// --- Repo Creation Logic ---
// These are Exported (uppercase) so main.go can call them.

// CreateGoRepo builds a standard Go application structure.
func CreateGoRepo(repoPath, repoName string) error {
	b := NewBuilder(repoPath)

	b.WriteFile("README.md", "# "+repoName)
	b.WriteFile(".gitignore", "*.exe\n*.out\n")
	b.WriteFile("go.mod", goModContent(repoName))
	b.WriteFile(fmt.Sprintf("cmd/%s/main.go", repoName), goMainContent())
	b.MkDir("pkg/utils") // Example "functions" folder

	return b.Err()
}

// CreatePythonRepo builds a Data Science project structure.
func CreatePythonRepo(repoPath, repoName string) error {
	b := NewBuilder(repoPath)

	b.WriteFile("README.md", "# "+repoName)
	b.WriteFile(".gitignore", pythonGitignore())
	b.WriteFile("requirements.txt", "# Add your packages here, e.g.:\npandas\nnumpy\n")
	b.MkDir("data/raw")
	b.MkDir("data/processed")
	b.MkDir("notebooks")
	b.WriteFile("notebooks/01-explore.ipynb", "") // Stub notebook
	b.MkDir("src")                                // This is your "functions" folder
	b.WriteFile("src/__init__.py", "")
	b.WriteFile("src/main.py", pythonMainContent())
	b.MkDir("tests")

	return b.Err()
}

// CreateNodeApiRepo builds a Node.js Express API structure.
func CreateNodeApiRepo(repoPath, repoName string) error {
	b := NewBuilder(repoPath)

	b.WriteFile("README.md", "# "+repoName)
	b.WriteFile(".gitignore", nodeGitignore())
	b.WriteFile("package.json", nodePackageJson(repoName))
	b.WriteFile("src/server.js", nodeServerJs()) // Main entry point
	b.MkDir("src/controllers")                   // Your "functions" folders
	b.MkDir("src/services")
	b.MkDir("src/routes")
	b.MkDir("src/models")

	return b.Err()
}

// CreateReactRepo builds a Next.js project structure.
func CreateReactRepo(repoPath, repoName string) error {
	b := NewBuilder(repoPath)

	b.WriteFile("README.md", "# "+repoName)
	b.WriteFile(".gitignore", nodeGitignore()) // Can reuse node's gitignore
	b.WriteFile("package.json", reactPackageJson(repoName))
	b.WriteFile("next.config.js", reactConfig())
	b.WriteFile("app/page.js", reactIndexPage())
	b.WriteFile("app/layout.js", reactLayout())
	b.MkDir("components") // Your "functions" folder for UI
	b.MkDir("lib")        // Your "functions" folder for logic

	return b.Err()
}

// CreateJavaRepo builds a Spring Boot project structure.
func CreateJavaRepo(repoPath, repoName string) error {
	b := NewBuilder(repoPath)
	// Sanitize repoName for Java class name
	javaClassName := strings.ReplaceAll(repoName, "-", "")
	javaClassName = strings.ReplaceAll(javaClassName, "_", "")

	mainJavaPath := fmt.Sprintf("src/main/java/com/example/%s", javaClassName)
	mainJavaFile := fmt.Sprintf("%s/%sApplication.java", mainJavaPath, strings.Title(javaClassName))

	b.WriteFile("README.md", "# "+repoName)
	b.WriteFile("pom.xml", javaPomXml(repoName))
	b.WriteFile(mainJavaFile, javaMainApp(javaClassName))
	b.WriteFile("src/main/resources/application.properties", javaAppProperties())
	b.MkDir("src/main/java/com/example/" + javaClassName + "/controller")
	b.MkDir("src/main/java/com/example/" + javaClassName + "/service") // "functions" folder
	b.MkDir("src/main/java/com/example/" + javaClassName + "/model")
	b.MkDir("src/test/java/com/example/" + javaClassName)

	return b.Err()
}

// CreateCSharpRepo builds a .NET WebAPI project structure.
func CreateCSharpRepo(repoPath, repoName string) error {
	b := NewBuilder(repoPath)

	b.WriteFile("README.md", "# "+repoName)
	b.WriteFile(repoName+".csproj", csharpProjectFile(repoName))
	b.WriteFile("Program.cs", csharpProgramCs())
	b.WriteFile("appsettings.json", csharpAppSettings())
	b.MkDir("Controllers")
	b.MkDir("Services") // Your "functions" folder
	b.MkDir("Models")
	b.MkDir("Data")

	return b.Err()
}

// CreateRustRepo builds a Rust CLI project structure.
func CreateRustRepo(repoPath, repoName string) error {
	b := NewBuilder(repoPath)

	b.WriteFile("README.md", "# "+repoName)
	b.WriteFile(".gitignore", "target/")
	b.WriteFile("Cargo.toml", rustCargoToml(repoName))
	b.WriteFile("src/main.rs", rustMainRs())
	b.MkDir("src/bin") // For multiple binaries
	b.MkDir("src/lib") // For library code ("functions" folder)

	return b.Err()
}

// CreateCRepo builds a C project structure with a Makefile.
func CreateCRepo(repoPath, repoName string) error {
	b := NewBuilder(repoPath)

	b.WriteFile("README.md", "# "+repoName)
	b.WriteFile(".gitignore", cGitignore())
	b.WriteFile("Makefile", cMakeFileContent(repoName))
	b.WriteFile("src/main.c", cMainContent())
	b.MkDir("include") // For header files
	b.MkDir("bin")     // For compiled output

	return b.Err()
}

// CreateCppRepo builds a C++ project structure with a Makefile.
func CreateCppRepo(repoPath, repoName string) error {
	b := NewBuilder(repoPath)

	b.WriteFile("README.md", "# "+repoName)
	b.WriteFile(".gitignore", cGitignore()) // Can reuse C's gitignore
	b.WriteFile("Makefile", cppMakeFileContent(repoName))
	b.WriteFile("src/main.cpp", cppMainContent())
	b.MkDir("include") // For header files
	b.MkDir("bin")     // For compiled output

	return b.Err()
}
