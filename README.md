# My Go Learning Repository

Welcome to my Go learning repository! Here, I document my journey of learning Go by adding codes, notes, and resources for various projects and applications. My goal is to build small projects and applications to apply my knowledge and skills in areas such as server development, CLI tools, microservices, and embedded systems.

## Table of Contents

- [Introduction](#introduction)
- [Projects](#projects)
- [Notes](#notes)
- [Important Links](#important-links)
- [Conclusion](#conclusion)

## Introduction

In this repository, you'll find a collection of projects and code snippets written in Go. Each project focuses on a specific aspect of Go programming, ranging from server development to building CLI tools and microservices. The goal is to gain practical experience and deepen my understanding of the language and its applications.

## Projects

Here's a list of projects I've worked on or am planning to build:

1. **Video Conference App**  
   A web application for video conferencing with real-time chat functionality.  
   **Technologies**: Go, WebRTC, HTML/CSS, JavaScript

2. **Basic WebSocket App (Live Notes Sharing)**  
   A Next.js-based frontend application that enables real-time notes sharing through WebSocket connections.  
   **Technologies**: Go, WebSocket, Next.js

3. **Terminal-Based Games (UNO)**  
   A CLI game version of the popular card game UNO, playable directly from the terminal.  
   **Technologies**: Go, CLI

4. **Online Chess Multiplayer App**  
   A web-based multiplayer chess application with real-time game syncing and move tracking using HTMX.  
   **Technologies**: Go, HTMX, HTML/CSS, JavaScript

## Important Links

Here are some important links and resources that have been helpful in my Go learning journey:

- [Official Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Gophercises](https://gophercises.com/)
- [A Tour of Go](https://tour.golang.org/)
- [Hitesh Choudhary's Golang Course](https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa)

<!-- Add more links as needed -->

## Conclusion

This repository serves as a comprehensive documentation of my journey of learning and mastering Go programming language. By building various projects and exploring different aspects of Go, I aim to solidify my understanding and expertise in this versatile language. Feel free to explore the code, projects, and notes, and I hope you find them helpful in your own Go learning journey.

## Notes

This section is dedicated to my personal notes on the code, concepts, and lessons learned while working on various projects. It serves as a reference for future learning and troubleshooting.

# 02. Before you start with golang

- Go enthusiasts are called "Gophers"
- Go is termed as the Next-Gen programming language
  - A lot of people go ahead and say that if C was built in present day, it would be have been built like Go
- Go utilizes the true power of Cloud infrastructure

## Features

- Golang is compiled
- Runtime is built into the final executable. So it can run directly without a VM
- Executables are different for different OS
- Golang can be used to build from system apps to web apps
- Go code is already in production
- Object oriented - Both yes and no
  - We don't have classes in Golang (we have structs)
  - There's no overloading
- Lexer does a lot of work

# 03. Golang installation and hello world

- [GoLang](https://go.dev/)

## Hello World

- After install Go and VS Code extensions, we'll need to write our Hello World program
- Init (Something like `npm init`)

```sh
go mod init hello
```

- Go has a single entry point like in C/C++. main is a reserved keyword
- Automatic imports are done by VSCode
### main.go

```golang
package main

import "fmt"

func main()  {
	fmt.Println("Hello, World!")
}
```
- To run the file:
```
go run main.go
```