# Final Project Worklog

This file contains the worklog for the final project.

## Project Preparation

### Create the GitHub Repository

Go to GitHub and create a new repository:
- Repository name: `ChineseZodiacGo`
- Description: A Go program that generates the Chinese Zodiac cycle and provides information about each year.
- Make it public
- Do NOT initialize with a README (simpler for the first push)

After the repo is created, here is what is shown on the page:
```bash
# …or create a new repository on the command line
echo "# ChineseZodiacGo" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/hgbbus/ChineseZodiacGo.git
git push -u origin main

# …or push an existing repository from the command line
git remote add origin https://github.com/hgbbus/ChineseZodiacGo.git
git branch -M main
git push -u origin main
```

We are not going to do anything about the above instructions yet.

### Create the Go Project Locally

In the Final Project folder, initialize a new Go module:
```bash
% go mod init github.com/hgbbus/ChineseZodiacGo
```

This is important because:
- It becomes your module path
- Future imports will use this
- It avoids renaming later

You should see a `go.mod` file created with the following content:
```go
module github.com/hgbbus/ChineseZodiacGo

go 1.25.6
```

### Create the Main Go File and Test It

Create a new file named `main.go` in the project folder. This will be the entry point of your Go program.

The content of `main.go` is as follows for now:
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Chinese Zodiac Program!")
	// Your code to implement the Chinese Zodiac logic goes here
}
```

Test it:
```bash
% go run .
Welcome to the Chinese Zodiac Program!
```

Or build and run it:
```bash
% go build   

% ls
ChineseZodiac.md   ChineseZodiacGo*   go.mod             main.go            worklog.md

% ./ChineseZodiacGo 
Welcome to the Chinese Zodiac Program!
```

### Add the README File

Create a `README.md` file in the project folder with the following content:
````markdown
# Chinese Zodiac Go Program

This Go program generates the Chinese Zodiac cycle and provides information about each year.

## Run the Program

To run the program, use the following command:
```bash
go run .
```

## Build the Program

To build the program, use the following command:
```bash
go build
```

After building, you can run the executable:
```bash
./ChineseZodiacGo
```
````

### Initialize Git and Commit

Inside the project folder:
```bash
% git init
Initialized empty Git repository in .../Final Project/.git/
```

Create a `.gitignore` file to ignore the executable and any other files you don't want to track:
```
# Binaries
ChineseZodiacGo
*.ext
*.out

# IDE files
.vscode/
.idea/

# PDF files
*.pdf
```

Now, configure your Git `user.email` for the repository to make your real email address private:
```bash
% git config user.email "your-noreply-github-email-address"
```

Now, add all files and make the first commit:
```bash
% git add .
% git commit -m "Initial commit with main.go and README.md"
```

### Connect to GitHub and Push

Add the remote `origin` pointing to your GitHub repository:
```bash
% git remote add origin https://github.com/hgbbus/ChineseZodiacGo.git
```

Then push the commit to GitHub:
```bash
# git branch -M main (not necessary if you already have a main branch)
% git push -u origin main
```

Commit the worklog file changes and push again:
```bash
% git add worklog.md
% git commit -m "Add connecting to github and push steps to worklog"
% git push
```

### Add Test File

Create a new file named `main_test.go` in the project folder with the following content:
```go
package main

import "testing"

func TestMain(t *testing.T) {
    // This is a placeholder test. You can add actual tests later.
}
```

Run the test(s):
```bash
% go test
PASS
ok      github.com/hgbbus/ChineseZodiacGo       0.230s
```

Or run with verbose output:
```bash
% go test -v
=== RUN   TestMain
    main_test.go:8: Starting main function test...
    main_test.go:12: Main function executed successfully.
--- PASS: TestMain (0.00s)
PASS
ok      github.com/hgbbus/ChineseZodiacGo       0.200s
```

Commit the test file and push:
```bash
% git add .
% git commit -m "Add main_test.go with a placeholder test"
% git push
```

The common practice for the test code is to separate it from the main code, but for simplicity, we are putting it in the same package and folder for now. However, we still want to use a separate package name for the test file. Go allows us to do this by using the `_test` suffix in the package name. So, we can change the package declaration in `main_test.go` to:
```go
package main_test
...
```

Now, commit the change and push again:
```bash
% git add .
% git commit -m "Change test package to main_test"
% git push
```

## Solution Implementation

### Add a Menu-Driven Interface

Add the following code to the main function in the `main.go` to create a simple menu-driven interface:
```go
...
	for {
		fmt.Println()
		fmt.Println("Please select an option:")
		fmt.Println()
		fmt.Println(" 1. Generate and display the full 60-year cycle starting from a specified year.")
		fmt.Println(" 2. Look up the zodiac information for a specific Gregorian year.")
		fmt.Println(" 0. Exit")
		fmt.Println()
		fmt.Print("Enter your choice (1, 2, or 0): ")

		var choice int
		fmt.Scanln(&choice)
		fmt.Println()

		switch choice {
		case 1:
			fmt.Print("Enter the starting Gregorian year: ")
			var startYear int
			fmt.Scanln(&startYear)
			// generateCycle(startYear)
		case 2:
			fmt.Print("Enter the Gregorian year to look up: ")
			var year int
			fmt.Scanln(&year)
			// lookupYear(year)
		case 0:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 0.")
		}
	}
```

Run it to see the menu and try entering different options:
```bash
% go run .
Welcome to the Chinese Zodiac Program!

Please select an option:

 1. Generate and display the full 60-year cycle starting from a specified year.
 2. Look up the zodiac information for a specific Gregorian year.
 0. Exit

Enter your choice (1, 2, or 0): 

...

```

Now, write the test code for the menu interface in `main_test.go`. Since the `main()` function is not exported, we cannot directly test it. We have two optins: (1) We change the package name back to the `main` package; (2) We write a helper function such as `Main()` that calls the `main()` function. We go with the second option.

The details of the test code are in `main_test.go`. For brevity, we omit the code snippet here.

Run the test code to ensure it works:
```bash
% go test .
ok      github.com/hgbbus/ChineseZodiacGo       (cached)
```

Now, commit the changes and push:
```bash
% git add .
% git commit -m "Add menu-driven interface and corresponding test code"
% git push
```

### Implement Yin-Yang and Five Element Logic

Add the following code to implement the logic for determining the element and Yin/Yang based on the stem index:
```go
func getElement(stemIndex int) string {
	elements := []string{"Wood", "Fire", "Earth", "Metal", "Water"}
	return elements[stemIndex/2]
}

func getYinYang(stemIndex int) string {
	if stemIndex%2 == 0 {
		return "Yang"
	}
	return "Yin"
}
```

Now, write test cases for these functions in `main_test.go`. Detailed code snippets are in the test file. Omited here for brevity.

Run the tests to ensure they work:
```bash
% go test .    
ok      github.com/hgbbus/ChineseZodiacGo       0.202s
```

Now, commit the changes and push:
```bash
% git add .
% git commit -m "Implement Yin-Yang and Five Element logic with tests"
% git push
```
