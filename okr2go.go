package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/browser"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kyokomi/emoji"
	"github.com/oxisto/go-httputil"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	file, err := os.OpenFile("okr2go.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	}

	log.SetLevel(logrus.DebugLevel)

	emoji.Printf("Welcome to okr2go! Your :books:tracker is ready at http://localhost:4300.\n")

	browser.OpenURL("http://localhost:4300")

	router := handlers.LoggingHandler(&httputil.LogWriter{Logger: log, Level: logrus.InfoLevel, Component: "http"}, NewRouter())
	err = http.ListenAndServe("0.0.0.0:4300", router)

	log.Errorf("An error occurred: %v", err)
}

type Objective struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	KeyResults  []KeyResult `json:"keyResults"`
}

type KeyResult struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Current      int64    `json:"current"`
	Target       int64    `json:"target"`
	Contributors []string `json:"contributors"`
	Comments     []string `json:"comments"`
}

// NewRouter returns a configured mux router containing all REST endpoints
func NewRouter() *mux.Router {
	// pack angular ui
	box := packr.NewBox("./okr2go-ui/dist/okr2go-ui")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/objectives", GetObjectives)
	router.PathPrefix("/").Handler(http.FileServer(box))

	return router
}

func GetObjectives(w http.ResponseWriter, r *http.Request) {
	var err error

	objectives, err := ParseMarkdown("example.md")

	if err != nil {
		httputil.JSONResponse(w, r, nil, err)
		return
	}

	httputil.JSONResponse(w, r, objectives, err)
}

func ParseMarkdown(path string) ([]*Objective, error) {
	data, err := ioutil.ReadFile("example.md")
	if err != nil {
		return nil, err
	}

	// strip window-style \r from the data, otherwise the Markdown library is doing weird things
	data = []byte(strings.ReplaceAll(string(data), "\r", ""))

	md := blackfriday.New()
	ast := md.Parse(data)

	walker := &ObjectiveWalker{}
	ast.Walk(walker.Walk)

	objectives, err := walker.Result()

	return objectives, err
}

type ObjectiveWalker struct {
	err        error
	objective  *Objective
	objectives []*Objective
}

func (o ObjectiveWalker) Result() ([]*Objective, error) {
	if o.err != nil {
		return nil, o.err
	}

	return o.objectives, nil
}

func Dump(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if entering {
		fmt.Printf("%+v\n", node)
	}

	return blackfriday.GoToNext
}

func (o *ObjectiveWalker) Walk(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	var err error

	if !entering {
		return blackfriday.GoToNext
	}

	if node.Type == blackfriday.Heading && node.HeadingData.Level == 1 {
		// if there is a new high level heading, we have a new objective
		o.objective = &Objective{}

		// add a reference to it to our objective list
		o.objectives = append(o.objectives, o.objective)

		log.Debug("Started new objective\n")

		return blackfriday.GoToNext
	}

	// take the first available text as the objective name
	if node.Type == blackfriday.Text && o.objective.Name == "" {
		o.objective.Name = string(node.Literal)

		log.Debugf("Setting objective text to %s\n", o.objective.Name)

		return blackfriday.GoToNext
	}

	if node.Type == blackfriday.Text {
		// yay, more text. could either be part of the description or the objective table
		text := string(node.Literal)

		// not really a good way to parse this
		if strings.Contains(text, "| -- |") {
			// this is the table heading, just ignore it

			return blackfriday.SkipChildren
		}

		// its probably the objectives table
		if strings.HasPrefix(text, "\n|") {
			o.objective.KeyResults = []KeyResult{}

			// split per line first
			lines := strings.Split(text, "\n")

			for _, line := range lines {
				// skip empty lines
				if line == "" {
					continue
				}

				keyResult := KeyResult{}

				// split fields
				fields := strings.Split(line, "|")

				log.Debugf("Parsing objective content %+v...\n", fields)

				// ID
				keyResult.ID = strings.TrimSpace(fields[1])

				// Name
				keyResult.Name = strings.TrimSpace(fields[2])

				// Current (parse to int/float?)
				keyResult.Current, _ = strconv.ParseInt(strings.TrimSpace(fields[3]), 10, 64)

				// Target (parse to int/float?)
				keyResult.Target, _ = strconv.ParseInt(strings.TrimSpace(fields[4]), 10, 64)

				// Contributors (split string)
				keyResult.Contributors = trimAndSplit(fields[5])

				// Comments (split string)
				keyResult.Comments = trimAndSplit(fields[6])

				// add the result
				o.objective.KeyResults = append(o.objective.KeyResults, keyResult)
			}
		} else {
			log.Debugf("Appending '%s' to description", node.Literal)

			// otherwise, just append it to the description
			if o.objective.Description == "" {
				o.objective.Description = text
			} else {
				o.objective.Description += " " + text
			}
		}

		return blackfriday.GoToNext
	}
	// do not continue if there is an error and store the error
	if err != nil {
		o.err = err
		return blackfriday.Terminate
	}

	return blackfriday.GoToNext
}

func trimAndSplit(field string) []string {
	// make sure the array is at least emtpy and not null
	r := []string{}

	// split string
	elements := strings.Split(strings.TrimSpace(field), ",")

	// loop through elements
	for _, element := range elements {
		element = strings.TrimSpace(element)

		// don't add empty elements
		if element == "" {
			continue
		}

		r = append(r, element)
	}

	return r
}

func ParseHeading(node *blackfriday.Node) (*Objective, error) {
	if node.Type != blackfriday.Heading {
		return nil, errors.New("expected heading node")
	}

	// first child should be text, otherwise we cannot parse the name
	if node.FirstChild == nil || node.FirstChild.Type != blackfriday.Text {
		return nil, errors.New("heading node does not have text")
	}

	objective := Objective{
		Name: string(node.FirstChild.Literal),
	}

	return &objective, nil
}
