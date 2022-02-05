package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"github.com/pkg/errors"
)

var (
	letters      = regexp.MustCompile(`^[a-zA-Z]+$`)
	tags         = regexp.MustCompile(`<[^>]*>`)
	citations    = regexp.MustCompile(`\[[^\]]*\]`)
	punctuations = regexp.MustCompile(`[(),.!?";:]*`)
)

func main() {
	processFrench()
}

func loadLinks() (map[string]func(string) []string, error) {
	results := map[string]func(string) []string{}
	links, err := readLinks("large.txt")
	if err != nil {
		return nil, err
	}
	for _, l := range links {
		results[l] = parseFrenchPage
	}

	links, err = readLinks("crawled.txt")
	if err != nil {
		return nil, err
	}
	for _, l := range links {
		results[l] = parseFrenchPage
	}

	links, err = readLinks("popular.txt")
	if err != nil {
		return nil, err
	}
	for _, l := range links {
		results[l] = parseFrenchList
	}

	return results, nil
}

func readLinks(sourceFile string) ([]string, error) {
	file, err := os.Open(sourceFile)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open source file")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read lines")
	}
	return lines, nil
}

func processFrench() {
	uris, err := loadLinks()
	if err != nil {
		fmt.Printf("error retrieveing data: %+v", err)
		return
	}
	fmt.Printf("reading %v links\n", len(uris))

	wordsProcessed := []string{}
	count := 0
	for uri, function := range uris {
		fmt.Printf("processing %s...\n", uri)
		count++
		rawPage, err := getURI(uri)
		if err != nil {
			fmt.Printf("error retrieveing data: %+v", err)
			return
		}
		//output(fmt.Sprintf("output-%v-raw.txt", count), rawPage)

		wordsRaw := append(wordsProcessed, function(rawPage)...)
		wordsProcessed = processWords(wordsRaw)
		fmt.Printf("%v words processed so far in %v pages\n", len(wordsProcessed), count)
		time.Sleep(2500 * time.Millisecond)
	}

	err = output("test.txt", strings.Join(wordsProcessed, "\n"))
	if err != nil {
		fmt.Printf("error outputing data: %+v", err)
		return
	}
}

func processWords(wordsRaw []string) []string {
	unique := map[string]bool{}
	processed := []string{}
	for _, w := range wordsRaw {
		clean := cleanWord(w)
		if len(clean) > 1 && !unique[clean] {
			processed = append(processed, clean)
			unique[clean] = true
		}
	}

	return processed
}

func cleanWord(word string) string {
	if word[0] >= 65 && word[0] < 91 {
		//fmt.Printf("POTENTIAL NAME: %v\n", word)
		return ""
	}
	clean := strings.ToLower(word)
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	clean, _, _ = transform.String(t, clean)
	if !letters.MatchString(clean) {
		//fmt.Printf("NOT CLEAN: %v\n", clean)
		return ""
	}

	return clean
}

func parseFrenchPage(content string) []string {
	startIndex := strings.Index(content, "<p>")
	endIndex := strings.Index(content, "</p>")
	words := []string{}
	for startIndex >= 0 && endIndex > startIndex {
		// extract the current paragraph
		sentence := content[startIndex+3 : endIndex]

		// remove tags
		sentencec := tags.ReplaceAllString(sentence, "")
		sentencec = citations.ReplaceAllString(sentencec, "")
		sentencec = punctuations.ReplaceAllString(sentencec, "")
		words = append(words, strings.Fields(sentencec)...)
		//fmt.Printf("SENTENCE \n%s \nBECAME\n%s\n", sentence, sentencec)
		//fmt.Printf("START: %v\tEND: %v\n", startIndex, endIndex)

		// figure out the indices for the next paragraph
		content = content[endIndex+4:]
		startIndex = strings.Index(content, "<p>")
		endIndex = strings.Index(content, "</p>")

		// consider the case of errant closing tag
		if endIndex >= 0 && startIndex > endIndex {
			content = content[startIndex:]
			startIndex = 0
			endIndex = strings.Index(content, "</p>")
		}
	}

	return words
}

func parseFrenchList(content string) []string {
	// get to near the list
	startIndex := strings.Index(content, "mw-parser-output")
	content = content[startIndex:]

	// setup for first item
	startIndex = strings.Index(content, "<li>")
	endIndex := strings.Index(content, "</ul>")
	words := []string{}
	for startIndex >= 0 && startIndex < endIndex {
		content = content[startIndex:]
		linkIndex := strings.Index(content, "<a")
		content = content[linkIndex:]
		linkIndex = strings.Index(content, ">")
		closeIndex := strings.Index(content, "</a>")

		words = append(words, content[linkIndex+1:closeIndex])

		startIndex = strings.Index(content, "<li>")
		endIndex = strings.Index(content, "</ul>")
	}

	return words
}

func output(filename string, content string) error {
	err := ioutil.WriteFile(filename, []byte(content), os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "unable to store content")
	}

	return nil
}

func getURI(uri string) (string, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", errors.Wrapf(err, "unable to create request")
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := client.Do(request)
	if err != nil {
		return "", errors.Wrapf(err, "unable to get uri %s", uri)
	}
	defer resp.Body.Close()

	bodyRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "unable to get string body")
	}

	return string(bodyRaw), nil
}
