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
