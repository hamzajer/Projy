package utils

import (
	"fmt"
	"strings"
)

// --- Boilerplate Content ---
// These are un-exported (lowercase) as they are only
// used by the Create... functions in this same package.

func goMainContent() string {
	return `package main

import "fmt"

func main() {
	fmt.Println("Hello, Go Project!")
}
`
}

func goModContent(repoName string) string {
	return fmt.Sprintf(`module %s

go 1.21
`, repoName)
}

func pythonMainContent() string {
	return `import os

def main():
    print("Hello, Python Project!")

if __name__ == "__main__":
    main()
`
}

func pythonGitignore() string {
	return `
# Python
__pycache__/
venv/
*.pyc
.env
`
}

func nodePackageJson(repoName string) string {
	return fmt.Sprintf(`{
  "name": "%s",
  "version": "1.0.0",
  "description": "A new Node.js API",
  "main": "src/server.js",
  "scripts": {
    "start": "node src/server.js",
    "dev": "nodemon src/server.js"
  },
  "dependencies": {
    "express": "^4.18.2"
  },
  "devDependencies": {
    "nodemon": "^3.0.1"
  }
}
`, repoName)
}

func nodeServerJs() string {
	return `const express = require('express');
const app = express();
const port = process.env.PORT || 3000;

app.get('/', (req, res) => {
  res.send('Hello, Node.js API!');
});

app.listen(port, () => {
  console.log(\Server running on http://localhost:\${port}\);
});
`
}

func nodeGitignore() string {
	return `
# Node
node_modules/
.env
dist/
`
}

func reactPackageJson(repoName string) string {
	return fmt.Sprintf(`{
  "name": "%s",
  "version": "0.1.0",
  "private": true,
  "scripts": {
    "dev": "next dev",
    "build": "next build",
    "start": "next start",
    "lint": "next lint"
  },
  "dependencies": {
    "react": "^18",
    "react-dom": "^18",
    "next": "14.0.1"
  },
  "devDependencies": {
    "typescript": "^5",
    "@types/node": "^20",
    "@types/react": "^18",
    "@types/react-dom": "^18"
  }
}
`, repoName)
}

func reactIndexPage() string {
	return `export default function Home() {
  return (
    <main>
      <h1>Hello, Next.js!</h1>
    </main>
  )
}
`
}

func reactLayout() string {
	return `export const metadata = {
  title: 'My Next.js App',
}

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}
`
}

func reactConfig() string {
	return `/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
}

module.exports = nextConfig
`
}

func javaPomXml(repoName string) string {
	// A very minimal pom.xml for Spring Boot
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>
    <parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>3.2.0</version>
        <relativePath/>
    </parent>
    <groupId>com.example</groupId>
    <artifactId>%s</artifactId>
    <version>0.0.1-SNAPSHOT</version>
    <name>%s</name>
    <properties>
        <java.version>17</java.version>
    </properties>
    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>
    <build>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
            </plugin>
        </plugins>
    </build>
</project>
`, repoName, repoName)
}

func javaMainApp(repoName string) string {
	// Simple main application class
	return `package com.example.` + repoName + `;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@RestController
public class ` + strings.Title(repoName) + `Application {

    public static void main(String[] args) {
        SpringApplication.run(` + strings.Title(repoName) + `Application.class, args);
    }

    @GetMapping("/")
    public String hello() {
        return "Hello, Spring Boot!";
    }
}
`
}

func javaAppProperties() string {
	return `server.port=8080`
}

func csharpProjectFile(repoName string) string {
	return fmt.Sprintf(`<Project Sdk="Microsoft.NET.Sdk.Web">
  <PropertyGroup>
    <TargetFramework>net8.0</TargetFramework>
    <Nullable>enable</Nullable>
    <ImplicitUsings>enable</ImplicitUsings>
  </PropertyGroup>
  <ItemGroup>
    <PackageReference Include="Microsoft.AspNetCore.OpenApi" Version="8.0.0" />
    <PackageReference Include="Swashbuckle.AspNetCore" Version="6.4.0" />
  </ItemGroup>
</Project>
`, repoName)
}

func csharpProgramCs() string {
	return `var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();
app.UseAuthorization();
app.MapControllers();

app.MapGet("/", () => "Hello, .NET Web API!");

app.Run();
`
}

func csharpAppSettings() string {
	return `{
  "Logging": {
    "LogLevel": {
      "Default": "Information",
      "Microsoft.AspNetCore": "Warning"
    }
  },
  "AllowedHosts": "*"
}
`
}

func rustCargoToml(repoName string) string {
	return fmt.Sprintf(`[package]
name = "%s"
version = "0.1.0"
edition = "2021"

[dependencies]
`, repoName)
}

func rustMainRs() string {
	return `fn main() {
    println!("Hello, Rust!");
}
`
}

func cGitignore() string {
	return `
# C/C++
*.o
*.out
bin/
`
}

func cMainContent() string {
	return `#include <stdio.h>

int main() {
    printf("Hello, C Project!\n");
    return 0;
}
`
}

func cMakeFileContent(repoName string) string {
	return fmt.Sprintf(`CC=gcc
CFLAGS=-Iinclude -Wall
SRC=src/main.c
TARGET=bin/%s

all: $(TARGET)

$(TARGET): $(SRC)
	@mkdir -p bin
	$(CC) $(CFLAGS) -o $(TARGET) $(SRC)

clean:
	rm -f $(TARGET) *.o
`, repoName)
}

func cppMainContent() string {
	return `#include <iostream>

int main() {
    std::cout << "Hello, C++ Project!" << std::endl;
    return 0;
}
`
}

func cppMakeFileContent(repoName string) string {
	return fmt.Sprintf(`CXX=g++
CXXFLAGS=-Iinclude -Wall -std=c++17
SRC=src/main.cpp
TARGET=bin/%s

all: $(TARGET)

$(TARGET): $(SRC)
	@mkdir -p bin
	$(CXX) $(CXXFLAGS) -o $(TARGET) $(SRC)

clean:
	rm -f $(TARGET) *.o
`, repoName)
}
