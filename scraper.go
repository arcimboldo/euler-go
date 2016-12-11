// get problem number as input
// get the page and parse it and print the problem text

package main
import (
	"fmt"
	"os"
	"strconv"
	"log"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const eulerurl = "https://projecteuler.net/problem="

func createProblemProgram(pn string, problem string){
	// Find the current working directory.
	// Check if 'args[1]' directory exists, in case create it
	// create file args[1]/main.go with the problem as comment plus standard stuff
	cwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	path := cwd + "/problems/" + pn
	stat, err := os.Stat(path)
	if err == nil && ! stat.IsDir() {
		log.Fatalf("path %s should be a dir but it's not", path)
	}
	os.MkdirAll(path, 0755)

	file, err := os.OpenFile(path + "/main.go", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Unable to create file %s: %s", path, err)
	}
	defer file.Close()
	
	sproblem := strings.Split(problem, "\n")
	for i:= range sproblem {
		file.WriteString("// " + sproblem[i] + "\n")
	}
	file.WriteString("\n")
	file.WriteString("package main\n\n")
	file.WriteString("import (\n    \"fmt\"\n)\n\n")
	file.WriteString("func main(){\n\n}\n")

	tfile, err := os.OpenFile(path + "/main_test.go", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Unable to create file %s: %s", path, err)
	}
	defer tfile.Close()
	tfile.WriteString("package main\n\nimport (\n    \"testing\"\n)\n\n")
	
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: %s N\n", args[0])
		os.Exit(1)
	} else if len(args) > 2 {
		fmt.Printf("Too many arguments. Usage: %s N\n", args[0])
		os.Exit(1)
	} else  if _, err :=  strconv.Atoi(args[1]); err != nil {
		fmt.Printf("Argument '%s' is not an integer\n", args[1])
		os.Exit(1)
	}

	data, err := os.Open(args[1] + ".txt")
	var doc *goquery.Document
	if err == nil {
		doc, err = goquery.NewDocumentFromReader(data)
	} else {
		url := eulerurl + args[1]
		doc, err = goquery.NewDocument(url)
		if err != nil {
			log.Fatal(err)
		}
	}
	var problem string
	doc.Find("div[role=problem] p").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		problem += s.Text() + "\n"
	})
	fmt.Printf("%s", problem)

	createProblemProgram(args[1], problem)
}
