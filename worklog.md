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

### Go Test Notes

#### Run all tests:
```bash
% go test ./...
```

#### Run a specific test function:
```bash
% go test -run TestMain
```

#### Run with coverage:
```bash
% go test -cover
```

#### Generate a coverage report:
```bash
% go test -coverprofile=coverage.out
% go tool cover -html=coverage.out
```

The `go tool cover -html=coverage.out` command will interpret the coverage profile (`coverage.out`) and open a web browser showing the coverage report in a visual format, highlighting which lines of code were covered by the tests and which were not.

If needed, you can provide an option to specify the output file for the HTML report:
```bash
% go tool cover -html=coverage.out -o coverage.html
```

This will save the HTML report as `coverage.html` in the current directory, which you can then open with a web browser to view the coverage report.

Commit the test notes to the worklog and push:
```bash
% git add worklog.md
% git commit -m "Add Go test notes to worklog"
% git push
```

### Implement the Lookup Function and GitHub Actions

Add the following code to implement the `getZodiacInfo(year int)` function that calculates the zodiac information for a given year:
```go

var (
	Stems = [...]string{"Jia", "Yi", "Bing", "Ding", "Wu", "Ji", "Geng", "Xin", "Ren", "Gui"}
	Elements = [...]string{"Wood", "Fire", "Earth", "Metal", "Water"}
	Branches = [...]string{"Zi", "Chou", "Yin", "Mao", "Chen", "Si", "Wu", "Wei", "Shen", "You", "Xu", "Hai"}
	Animals = [...]string{"Rat", "Ox", "Tiger", "Rabbit", "Dragon", "Snake", "Horse", "Goat", "Monkey", "Rooster", "Dog", "Pig"}
)

func mod(a, b int) int {
	return (a % b + b) % b
}

func getZodiacInfo(year int) (string, string, string, string, string) {
	stemIndex := mod(year - 1984, 10)
	branchIndex := mod(year - 1984, 12)
	
	stem := Stems[stemIndex]
	branch := Branches[branchIndex]
	element := GetElement(stemIndex)
	yinYang := GetYinYang(stemIndex)
	animal := Animals[branchIndex]

	return stem, branch, element, yinYang, animal
}
```

Now, write test cases for the `getZodiacInfo` function in `main_test.go`. Detailed code snippets are in the test file. Omited here for brevity.

Run the tests to ensure they work:
```bash
% go test .
ok      github.com/hgbbus/ChineseZodiacGo       0.230s
```

Now, we could commit the changes and push. However, let's try GitHub Actions to run the tests automatically on push. We will set up a simple workflow for this.

#### Create GitHub Actions Workflow

In the project folder, create a new directory named `.github/workflows`. Inside this directory, create a new file named `go.yml` with the following content:

```yaml
name: Test and Coverage             # Just a name for the workflow

on:                                 # When to run the workflow
  push:                             #   Run on push to main branch
    branches: [ main ]
  pull_request:                     #   Or run on pull request to main branch 
    branches: [ main ]

jobs:                               # Define the jobs to run (a workflow can contain multiple jobs)
  test:                             # A job named "test" (just a name)
    runs-on: ubuntu-latest          # Tells GitHub to run this job on the latest version of Ubuntu

    steps:                          # Steps are the individual tasks that make up the job
      - name: Checkout code         # Step 1: Check out the code from the repository
        uses: actions/checkout@v5   #     Use a pre-built GitHub action that checks out your code to the runner
                                    #     (v5 means major version 5 of the prebuilt action on GitHub)

      - name: Set up Go             # Step 2: Set up the Go programming environment
        uses: actions/setup-go@v5   #     Use a pre-built GitHub action that sets up Go on the runner
                                    #     (v5 means major version 5 of the prebuilt action on GitHub)
        with:
          go-version: '1.25'        #     Specify the version of Go to use (1.20 in this case)

      - name: Run tests with coverage  # Step 3: Run the tests and collect coverage information
        run: go test ./... -coverprofile=coverage.out -covermode=atomic

      - name: Display coverage in console  # Step 4: Display the coverage report in the console
        run: go tool cover -func=coverage.out
```

This workflow will run on every push and pull request to the `main` branch. It checks out the code, sets up Go, builds the project, and runs the tests.

#### Commit the Workflow and Push

Now, commit the implementation, test code, and workflow file and push:
```bash
% git add .
% git commit -m "Implement getZodiacInfo function and set up GitHub Actions workflow"
% git push
```

That push was successful, and you should see the workflow running in the "Actions" tab of your GitHub repository. To see the details of the workflow run, click on the workflow name (e.g., "Test and Coverage"); then click on the specific run to see the logs and results of each step; then click on the specific job (e.g., "test") to see the logs from each step for that job, including the output of the tests and coverage report.

However, there is a warning from `actions/setup-go` about not seeing a `go.sum` file, which is used for dependency management. Normally, to fix this warning, we can run the following command in the terminal:
```bash
% go mod tidy
```

But this doesn't generate a `go.sum` file because we don't have any external dependencies in our project. We will therefore have to ignore this warning for now.

Lastly, to make GitHub show a nice coverage dashboard, we can add another step to the workflow to upload the coverage report to GitHub:
```yaml
      - name: Upload coverage report  # Step 5: Upload the coverage report as an artifact (optional)
        uses: actions/upload-artifact@v4  #     Use a pre-built GitHub action that uploads artifacts
        with:
          name: coverage-report           #     Name of the artifact
          path: coverage.out              #     Path to the coverage report file generated in Step 3
```

Now, commit the change and push again:
```bash
% git add .
% git commit -m "Add step to upload coverage report to GitHub Actions workflow"
% git push
```

That upload step works, and you can see the uploaded artifact in the "Artifacts" section of the workflow run details. You can download the `coverage.out` file and use it to generate a local coverage report if needed.

A better solution for the coverage report is to use a service like Codecov or Coveralls, which can integrate with GitHub and provide a nice dashboard for coverage reports. I am not doing that for now.

### Complete the Lookup Functionality in the Menu

Add the following code to complete the lookup functionality in the menu:
```go
		case 2:
			fmt.Print("Enter the Gregorian year to look up: ")
			var year int
			fmt.Scanln(&year)
			stem, branch, element, yinYang, animal := GetZodiacInfo(year)
			fmt.Println()
			fmt.Printf("%d -> %s %s (%s %s %s)\n", year, stem, branch, yinYang, element, animal)
```

Test the menu option `2` to ensure it works correctly. Details of the test are omitted here for brevity.

Now, commit the changes and push:
```bash
% git add .
% git commit -m "Complete the lookup functionality in the menu"
% git push
```

This should trigger another run of the GitHub Actions workflow, which will run the tests and show the results in the "Actions" tab of your GitHub repository.

### Compltete the Cycle Generation Function and Menu Option

Add the following code to complete the cycle generation function and menu option:
```go
func generateCycle(startYear int) {
	// One solution using existing GetZodiacInfo function
	//for i := range 60 {
	//	year := startYear + i
	//	stem, branch, element, yinYang, animal := GetZodiacInfo(year)
	//	fmt.Printf("%d - %s %s - %s %s - %s\n", year, stem, branch, yinYang, element, animal)
	//}

	// Another solution using direct calculations
	year := startYear
	stemIndex := mod(year - 1984, 10)
	branchIndex := mod(year - 1984, 12)
	fmt.Println()
	for {
		stem := Stems[stemIndex]
		branch := Branches[branchIndex]
		element := GetElement(stemIndex)
		yinYang := GetYinYang(stemIndex)
		animal := Animals[branchIndex]

		fmt.Printf("%d - %s %s - %s %s - %s\n", year, stem, branch, yinYang, element, animal)

		year++
		stemIndex = (stemIndex + 1) % 10
		branchIndex = (branchIndex + 1) % 12

		if year == startYear + 60 {
			break
		}
	}
}
```

Change the menu option `1` to call the `generateCycle` function:
```go
        case 1:
            fmt.Print("Enter the starting Gregorian year: ")
            var startYear int
            fmt.Scanln(&startYear)
            generateCycle(startYear)
```

Test the menu option `1` to ensure it works correctly. Details of the test are omitted here for brevity.

Now, write test cases for the `generateCycle` function in `main_test.go`. Since this function prints output to the console, we can capture the output and verify it against expected results. Detailed code snippets are in the test file. Omited here for brevity.

Run the tests to ensure they work:
```bash
% go test . 
ok      github.com/hgbbus/ChineseZodiacGo       0.343s
```

Finally, commit the changes and push:
```bash
% git add .
% git commit -m "Complete the cycle generation function and menu option"
% git push
```

## Final Notes and Suggestions

For a Python implementation, the structure would be similar. Below is my suggestion for the students:

1. Start with a main function that prints a welcome message and displays a menu for the user to choose from. For each menu option, stub a function call that will implement the required functionality eventually.
2. Define the data structures for the Heavenly Stems, Earthly Branches, Elements, and Animals as lists. Don't attempt more complex data structures at this point; keep it simple.
3. Implement a function called `generate_cycle(start_year)` that generates the 60-year cycle starting from the specified year. Use modular arithmetic to wrap the indices for stems and branches.
4. Update the menu option for generating the cycle to call this function and display the results in the specified format.
5. Test the menu option for generating the cycle to ensure it works correctly.
6. Implement a function called `get_element(stem_index)` that returns the element based on the stem index. Implement another function called `get_yin_yang(stem_index)` that returns "Yin" or "Yang" based on the stem index. These are helper functions that are not necessarily needed but will be used in the next step. Optionally, these helper functions can also used by the `generate_cycle` function.
7. Implement a function called `get_zodiac_info(year)` that calculates the stem index and branch index based on the year, and returns the corresponding stem, branch, element, yin/yang, and animal information using the above helper functions.
8. Update the menu option for looking up a specific year to call the `get_zodiac_info(year)` function and display the results in the specified format.
9. Test the menu option for looking up a specific year to ensure it works correctly.
10. Done! Give yourself a pat on the back and celebrate your achievement!

(End)
