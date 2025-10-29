# **Projy 🚀**

projy (short for *Project-ify*) is a simple and fast command-line tool, written in Go, for scaffolding new developer projects. It instantly generates a clean, idiomatic directory structure and boilerplate files for various languages, so you can stop copying old projects and start coding faster.

## **Features**

* **Blazing Fast:** Written in Go, it runs instantly.  
* **Simple to Use:** projy \<language\> \<path\> is all you need.  
* **No Dependencies:** Distributed as a single binary.  
* **Wide Language Support:** Generates templates for:  
  * Go  
  * Python (Data Science)  
  * Node.js (Express API)  
  * React (Next.js)  
  * Java (Spring Boot)  
  * C\# (.NET WebAPI)  
  * Rust (CLI)  
  * C (Makefile)  
  * C++ (Makefile)

## **Installation**

### **From Source**

1. **Clone the repository:**  
   git clone https://github.com/hamzajer/Projy.git \\
   cd projy

3. **Build the binary:**  
   go build \-o projy

4. **Move the binary to your PATH:**  
   sudo mv projy /usr/local/bin/

## **Usage**

Using projy is simple. Just provide the language and the path where you want your new project to be created.  
projy \<language\> \<path\>

The \<path\> will also be used as the project's name (e.g., in package.json or go.mod).

### **Examples**

**Create a new Python data science project:**  
projy python ./my-data-science-project

**Create a new Go CLI project:**  
projy go ./my-go-cli

**Create a new React (Next.js) web app:**  
projy react ./my-web-app

**Create a C++ project:**  
projy cpp ./my-cpp-program

## **Supported Languages**

| Language | Template |
| :---- | :---- |
| go | A standard Go CLI project with cmd/ and pkg/ |
| python | A data science structure with data/, notebooks/, src/ |
| node-api | A Node.js Express API with src/, controllers/, services/ |
| react | A Next.js project with app/, components/, lib/ |
| java | A Spring Boot project with a standard Maven layout |
| csharp | A .NET 8 WebAPI project |
| rust | A Cargo project with src/main.rs and src/lib |
| c | A C project with a Makefile, src/, and include/ |
| cpp | A C++ project with a Makefile, src/, and include/ |

## **How to Contribute**

Contributions are welcome\! If you'd like to add a new language template or improve an existing one:

1. Fork the repository.  
2. Create a new branch (git checkout \-b feature/add-swift-template).  
3. Add your changes in the utils/ directory (e.g., templates.go and creators.go).  
4. Submit a Pull Request.
