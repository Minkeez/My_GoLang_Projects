# Hello, World! CLI App in Go
A simple command-line application written in Go (Golang) that prints **"Hello, World!"** when executed. This project also demonstrates different text formats (uppercase, lowercase, and colored output) and accepts user input via command-line arguments.

### Features
- Prints `"Hello, World!"`
- Supports different output formats (uppercase, lowercase, colored text)
- Accepts custom messages via command-line arguments
- Fully compiled executable for easy execution

---

### Installation
1. **Clone the Repository**
  ```
  git clone https://github.com/Minkeez/My_GoLang_Projects.git
  ```
2. **Initialize Go Modules (if not already done)**
  ```
  go mod init 01_Hello_World_CLI
  cd 01_Hello_World_CLI
  ```
3. **Install Dependencies (for colored output)**
  ```
  go get github.com/fatih/color
  ```

---

### Usage
#### Run Directly
Execute the program using:
```
go run main.go
```
Expected output:
```
Hello, World!
```
#### Run with Different Formats
UPPERCASE flag
```
go run main.go --uppercase
```
Expected output:
```
HELLO, WORLD!
```
lowercase flag
```
go run main.go --lowercase
```
Expected output:
```
hello, world!
```
#### Run with Custom Message
Let's say I want to print my name "HourCode"
```
go run main.go "HourCode"
```
Expected output:
```
Custom Message: HourCode
```
#### Run with Colored Output
This is show you RGB text colors
```
go run main.go --color
```
Expected output (with color):
```
Hello, World! | Red Color
Hello, World! | Green Color
Hello, World! | Blue Color
```
#### Combine Flag and Custom Message
I want to print my name in RGB colors
```
go run main.go --color "HourCode"
```
Expected output (with color):
```
HourCode | Red Color
HourCode | Green Color
HourCode | Blue Color
```
#### Compile and Run the Executable
You also compile it once, and then run it
```
go build -o main.exe
```
and when you want to use it. just double-click or print it through cmd:
```
main.exe
```