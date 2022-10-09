package controllers

import (
	"net/http"

	tmpl "html/template"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/the-science-of-deduction/utils"
)

// Template is a struct for go html/template
type Template interface {
	Index(c *gin.Context)
	CaseFiles(c *gin.Context)
	CaseFileView(c *gin.Context)
	Forum(c *gin.Context)
	HiddenMessages(c *gin.Context)
	NotFound(c *gin.Context)
}

type template struct {
}

type casesFiles struct {
	URL       string
	Title     string
	IsDeleted bool
}

// Ongoing Case Files
var (
	casefiles []casesFiles = []casesFiles{
		{
			URL:       "#",
			Title:     "Analysis of Tobacco Ash: DELETED!!",
			IsDeleted: true,
		},
		{
			URL:   "#",
			Title: "The Aluminium Crutch: Ongoing case",
		},
		{
			URL:   "#",
			Title: "Analysis of Perfumes: Ongoing case",
		},
		{
			URL:   "#",
			Title: "Hidden Message #3 - I need your help",
		},
		{
			URL:   "#",
			Title: "Hidden Message #2 - Answer revealed",
		},
		{
			URL:   "#",
			Title: "Hidden Message #1 - Answer revealed",
		},
		{
			URL:   "the-green-ladder.html",
			Title: "The Green Ladder: A cast iron alibi?",
		},
	}

	// Archived Case Files
	archivedcasefiles []casesFiles = []casesFiles{
		{
			URL:   "#",
			Title: "The Confusion of Isadora Persano",
		},
		{
			URL:   "#",
			Title: "The Abernetty Family",
		},
		{
			URL:   "#",
			Title: "The Crooked House",
		},
		{
			URL:   "#",
			Title: "The Man With Four Legs",
		},
		{
			URL:   "#",
			Title: "The Killer Cats of Greenwich",
		},
		{
			URL:   "#",
			Title: "The Kirkcudbright Killer",
		},
		{
			URL:   "#",
			Title: "The Ghost of St Bartholomew's",
		},
		{
			URL:   "#",
			Title: "The Purple Woman",
		},
		{
			URL:   "#",
			Title: "The Laughing Pilot",
		},
		{
			URL:   "#",
			Title: "The Missing Jars",
		},
		{
			URL:   "#",
			Title: "The Invisible Porter",
		},
		{
			URL:   "#",
			Title: "The Subdivided Crooner",
		},
		{
			URL:   "#",
			Title: "The Pale Man",
		},
		{
			URL:   "#",
			Title: "The Iron Football",
		},
	}
)

// Recent Activities
var (
	recentPostURL   string = "#"
	recentPostTitle string = "If you want my help, write to me at 221B Baker Street, London or contact me through John's blog. - SH"

	// Copyright
	copyright string = "Copyright © 2022 mrinjamul. All rights reserved."
)

// Index is a function for index page
func (t *template) Index(c *gin.Context) {
	// No of Notes
	noOfNotes := 336
	// Last Update Time for posts
	lastUpdateTime := "May 9, 17:35"

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":             "The Science of Deduction — Home",
		"IsHome":            "active",
		"NoN":               noOfNotes,
		"Time":              lastUpdateTime,
		"caseFiles":         casefiles,
		"archivedCaseFiles": archivedcasefiles,
		"recentPostURL":     recentPostURL,
		"recentPostTitle":   recentPostTitle,
		"Copyright":         copyright,
	})
}

// CaseFiles is a function for case files page
func (t *template) CaseFiles(c *gin.Context) {
	c.HTML(http.StatusNotFound, "case-files.html", gin.H{
		"Title":             "The Science of Deduction — Case Files",
		"IsCase":            "active",
		"caseFiles":         casefiles,
		"archivedCaseFiles": archivedcasefiles,
		"recentPostURL":     recentPostURL,
		"recentPostTitle":   recentPostTitle,
		"Copyright":         copyright,
	})
}

// CaseFileView is a function for case file view page
func (t *template) CaseFileView(c *gin.Context) {

	content := "# The Green Ladder: A cast iron alibi? \n A PC Jane Downing has come to ask for my help. Her husband died. Everyone thinks it's a tragic accident. She's convinced his brother did it. Brother has a cast iron alibi. I love cast iron alibis so I'm taking on the case."
	htmlContent := tmpl.HTML(utils.MarkdownToHTML(content))

	c.HTML(http.StatusNotFound, "case-files-view.html", gin.H{
		"Title":             "The Science of Deduction — Case Files",
		"IsCase":            "active",
		"CaseTitle":         "The Green Ladder: A cast iron alibi?",
		"IsClosed":          true,
		"CaseContent":       htmlContent,
		"caseFiles":         casefiles,
		"archivedCaseFiles": archivedcasefiles,
		"recentPostURL":     recentPostURL,
		"recentPostTitle":   recentPostTitle,
		"Copyright":         copyright,
	})
}

// Forum
func (t *template) Forum(c *gin.Context) {

	c.HTML(http.StatusNotFound, "hidden-messages.html", gin.H{
		"Title":             "The Science of Deduction — Forum",
		"IsHidden":          "active",
		"caseFiles":         casefiles,
		"archivedCaseFiles": archivedcasefiles,
		"recentPostURL":     recentPostURL,
		"recentPostTitle":   recentPostTitle,
		"Copyright":         copyright,
	})
}

// HiddenMessages is a function for hidden messages page
func (t *template) HiddenMessages(c *gin.Context) {
	c.HTML(http.StatusNotFound, "hidden-messages.html", gin.H{
		"Title":             "The Science of Deduction — Hidden Messages",
		"IsHidden":          "active",
		"caseFiles":         casefiles,
		"archivedCaseFiles": archivedcasefiles,
		"recentPostURL":     recentPostURL,
		"recentPostTitle":   recentPostTitle,
		"Copyright":         copyright,
	})
}

// NotFound is a function for not found page
func (t *template) NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"Title":     "404 Not Found",
		"Copyright": "Copyright © 2022 mrinjamul. All rights reserved.",
	})
}

// NewTemplate is a function for new template
func NewTemplate() Template {
	return &template{}
}
