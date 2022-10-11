package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	tmpl "html/template"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/the-science-of-deduction/models"
	"github.com/mrinjamul/the-science-of-deduction/repository"
	"github.com/mrinjamul/the-science-of-deduction/utils"
)

// Template is a struct for go html/template
type Template interface {
	Index(c *gin.Context)
	CaseFiles(c *gin.Context)
	CaseFileNew(c *gin.Context)
	CreateCaseFile(c *gin.Context)
	CaseFileEdit(c *gin.Context)
	UpdateCaseFile(c *gin.Context)
	CaseFileView(c *gin.Context)
	Forum(c *gin.Context)
	HiddenMessages(c *gin.Context)
	NotFound(c *gin.Context)
}

type template struct {
	postRepo repository.CaseFiles
}

// Recent Activities
var (
	recentPostURL   string = "mailto:mrinjamul@pm.me"
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

	archivedfiles, err := t.postRepo.GetArchivedCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}
	casefiles, err := t.postRepo.GetActiveCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":             "The Science of Deduction — Home",
		"IsHome":            "active",
		"NoN":               noOfNotes,
		"Time":              lastUpdateTime,
		"caseFiles":         casefiles,
		"archivedCaseFiles": archivedfiles,
		"recentPostURL":     recentPostURL,
		"recentPostTitle":   recentPostTitle,
		"Copyright":         copyright,
	})
}

// CaseFiles is a function for case files page
func (t *template) CaseFiles(c *gin.Context) {
	archivedfiles, err := t.postRepo.GetArchivedCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}
	caseFiles, err := t.postRepo.GetActiveCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}
	ongoingCases, err := t.postRepo.GetOnGoingCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}
	recentCases, err := t.postRepo.GetRecentCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}

	c.HTML(http.StatusNotFound, "case-files.html", gin.H{
		"Title":             "The Science of Deduction — Case Files",
		"IsCase":            "active",
		"recentCases":       recentCases,
		"ongoingCases":      ongoingCases,
		"caseFiles":         caseFiles,
		"archivedCaseFiles": archivedfiles,
		"recentPostURL":     recentPostURL,
		"recentPostTitle":   recentPostTitle,
		"Copyright":         copyright,
	})
}

// CaseFileNew is a function for new case file page
func (t *template) CaseFileNew(c *gin.Context) {

	archivedfiles, err := t.postRepo.GetArchivedCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}
	casefiles, err := t.postRepo.GetActiveCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}

	c.HTML(http.StatusNotFound, "case-file-new.html", gin.H{
		"Title":             "The Science of Deduction — Case Files",
		"IsCase":            "active",
		"caseFiles":         casefiles,
		"archivedCaseFiles": archivedfiles,
		"recentPostURL":     recentPostURL,
		"recentPostTitle":   recentPostTitle,
		"Copyright":         copyright,
	})
}

// CreateCaseFile is a function for creating a new post
func (t *template) CreateCaseFile(c *gin.Context) {
	// Get the post title from the form
	title := c.PostForm("title")
	title = strings.TrimSpace(title)
	// Get the post description from the form
	description := c.PostForm("description")
	description = strings.TrimSpace(description)
	// Get the post author from the form
	author := c.PostForm("author")
	author = strings.TrimSpace(author)
	// Get the post content from the form
	content := c.PostForm("contents")
	// Get the post tags from the form
	// tags := c.PostForm("tags")
	// Get isClosed from the form
	isClosed := c.PostForm("isClosed")
	// Get isArchived from the form
	isArchived := c.PostForm("isArchived")
	// Get isDeleted from the form
	isDeleted := c.PostForm("isDeleted")
	// Get the post date from the form
	date := time.Now()
	// url for the post
	url := "/case-files/" + strings.ReplaceAll(title, " ", "-") + ".html"
	// check if feilds are empty
	if title == "" { // || author == "" || content == "" {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title":        "Error 400",
			"ErrorMessage": "Please fill all the feilds",
			"Copyright":    copyright,
		})
		return
	}

	if author == "" {
		author = "Anonymous"
	}

	// Create a new post
	post := models.CaseFiles{
		Title:       title,
		Description: description,
		Author:      author,
		Content:     content,
		URL:         url,
		UpdatedAt:   date,
	}

	if isArchived == "on" {
		post.IsArchived = true
	}
	if isClosed == "on" {
		post.IsClosed = true
	}
	if isDeleted == "on" {
		post.IsDeleted = true
	}

	// add case file to db
	err := t.postRepo.CreateCaseFile(c, &post)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title":        "404 Not Found",
			"ErrorMessage": "Something went wrong\n" + err.Error(),
			"Copyright":    copyright,
		})
	}

	// Redirect to the post page
	c.Redirect(http.StatusMovedPermanently, url)
}

// CaseFileEdit is a function for editing a case file
func (t *template) CaseFileEdit(c *gin.Context) {

	// get id from http request
	id := c.Param("id")
	// parse title from id
	title := strings.ReplaceAll(id, "-", " ")
	// title = strings.ReplaceAll(title, "%20", " ")
	title = strings.TrimSuffix(title, ".html")
	title = strings.TrimSpace(title)
	// fetch case file from db
	post, err := t.postRepo.GetCaseFileByTitle(c, title)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title":        "404 Not Found",
			"IsCase":       "active",
			"ErrorMessage": "Something went wrong\n" + err.Error(),
			"Copright":     copyright,
		})
		return
	}
	var (
		isClosed   string
		isArchived string
		isDeleted  string
	)
	if post.IsClosed {
		isClosed = "checked"
	} else {
		isClosed = ""
	}
	if post.IsArchived {
		isArchived = "checked"
	} else {
		isArchived = ""
	}
	if post.IsDeleted {
		isDeleted = "checked"
	} else {
		isDeleted = ""
	}

	archivedfiles, err := t.postRepo.GetArchivedCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}
	casefiles, err := t.postRepo.GetActiveCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}

	c.HTML(http.StatusNotFound, "case-file-new.html", gin.H{
		"Title":             "The Science of Deduction — Case Files",
		"IsCase":            "active",
		"postId":            post.Id,
		"postTitle":         post.Title,
		"postDescription":   post.Description,
		"postAuthor":        post.Author,
		"postIsClosed":      isClosed,
		"postIsArchived":    isArchived,
		"postIsDeleted":     isDeleted,
		"postContents":      post.Content,
		"caseFiles":         casefiles,
		"archivedCaseFiles": archivedfiles,
		"recentPostURL":     recentPostURL,
		"recentPostTitle":   recentPostTitle,
		"Copyright":         copyright,
	})
}

// UpdateCaseFile is a function for updating a case file
func (t *template) UpdateCaseFile(c *gin.Context) {
	// Get id from the form
	caseid := c.PostForm("caseid")
	id, err := strconv.Atoi(caseid)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title":        "404 Not Found",
			"ErrorMessage": "Something went wrong\n" + err.Error(),
			"Copyright":    copyright,
		})
	}
	// Get the post title from the form
	title := c.PostForm("title")
	title = strings.TrimSpace(title)
	// Get the post description from the form
	description := c.PostForm("description")
	description = strings.TrimSpace(description)
	// Get the post author from the form
	author := c.PostForm("author")
	author = strings.TrimSpace(author)
	// Get the post content from the form
	content := c.PostForm("contents")
	// Get the post tags from the form
	// tags := c.PostForm("tags")
	// Get isClosed from the form
	isClosed := c.PostForm("isClosed")
	// Get isArchived from the form
	isArchived := c.PostForm("isArchived")
	// Get isDeleted from the form
	isDeleted := c.PostForm("isDeleted")
	// Get the post date from the form
	date := time.Now()
	// url for the post
	url := "/case-files/" + strings.ReplaceAll(title, " ", "-") + ".html"
	// check if feilds are empty
	if title == "" { // || author == "" || content == "" {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title":        "Error 400",
			"ErrorMessage": "Please fill all the feilds",
			"Copyright":    copyright,
		})
		return
	}

	if author == "" {
		author = "Anonymous"
	}

	// Update a new post
	post := models.CaseFiles{
		Id:          uint(id),
		Title:       title,
		Description: description,
		Author:      author,
		Content:     content,
		URL:         url,
		UpdatedAt:   date,
	}

	if isArchived == "on" {
		post.IsArchived = true
	} else {
		post.IsArchived = false
	}
	if isClosed == "on" {
		post.IsClosed = true
	} else {
		post.IsClosed = false
	}
	if isDeleted == "on" {
		post.IsDeleted = true
	} else {
		post.IsDeleted = false
	}

	// add case file to db
	err = t.postRepo.UpdateCaseFile(c, &post)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title":        "404 Not Found",
			"ErrorMessage": "Something went wrong\n" + err.Error(),
			"Copyright":    copyright,
		})
	}

	// Redirect to the post page
	c.Redirect(http.StatusMovedPermanently, url)
}

// CaseFileView is a function for case file view page
func (t *template) CaseFileView(c *gin.Context) {

	// get id from http request
	id := c.Param("id")
	// parse title from id
	title := strings.ReplaceAll(id, "-", " ")
	title = strings.TrimSuffix(title, ".html")
	title = strings.TrimSpace(title)
	// fetch case file from db
	post, err := t.postRepo.GetCaseFileByTitle(c, title)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title":        "404 Not Found",
			"IsCase":       "active",
			"ErrorMessage": "Something went wrong\n" + err.Error(),
			"Copright":     copyright,
		})
		return
	}
	// check if post is empty
	if post.Id != 0 {

		htmlContent := tmpl.HTML(utils.MarkdownToHTML(post.Content))
		archivedfiles, err := t.postRepo.GetArchivedCaseFiles(c)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "404.html", gin.H{
				"Title":        "The Science of Deduction — 404",
				"ErrorMessage": "Internal Server Error",
				"Copright":     copyright,
			})
			return
		}
		casefiles, err := t.postRepo.GetActiveCaseFiles(c)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "404.html", gin.H{
				"Title":        "The Science of Deduction — 404",
				"ErrorMessage": "Internal Server Error",
				"Copright":     copyright,
			})
			return
		}

		url := strings.TrimSuffix(post.URL, ".html")

		c.HTML(http.StatusNotFound, "case-files-view.html", gin.H{
			"Title":             "The Science of Deduction — Case Files",
			"IsCase":            "active",
			"url":               url,
			"CaseTitle":         post.Title,
			"CaseAuthor":        post.Author,
			"IsClosed":          false,
			"CaseContent":       htmlContent,
			"caseFiles":         casefiles,
			"archivedCaseFiles": archivedfiles,
			"recentPostURL":     recentPostURL,
			"recentPostTitle":   recentPostTitle,
			"Copyright":         copyright,
		})
		return
	}

	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"Title":        "404 Not Found",
		"IsCase":       "active",
		"ErrorMessage": "Lost in the Mind Palace !!",
		"Copyright":    copyright,
	})
}

// Forum
func (t *template) Forum(c *gin.Context) {

	archivedfiles, err := t.postRepo.GetArchivedCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}
	casefiles, err := t.postRepo.GetActiveCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}

	c.HTML(http.StatusNotFound, "forum.html", gin.H{
		"Title":             "The Science of Deduction — Forum",
		"IsForum":           "active",
		"caseFiles":         casefiles,
		"archivedCaseFiles": archivedfiles,
		"recentPostURL":     recentPostURL,
		"recentPostTitle":   recentPostTitle,
		"Copyright":         copyright,
	})
}

// HiddenMessages is a function for hidden messages page
func (t *template) HiddenMessages(c *gin.Context) {

	archivedfiles, err := t.postRepo.GetArchivedCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}
	casefiles, err := t.postRepo.GetActiveCaseFiles(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "404.html", gin.H{
			"Title":        "The Science of Deduction — 404",
			"ErrorMessage": "Internal Server Error",
			"Copright":     copyright,
		})
		return
	}

	c.HTML(http.StatusNotFound, "hidden-messages.html", gin.H{
		"Title":             "The Science of Deduction — Hidden Messages",
		"IsHidden":          "active",
		"caseFiles":         casefiles,
		"archivedCaseFiles": archivedfiles,
		"recentPostURL":     recentPostURL,
		"recentPostTitle":   recentPostTitle,
		"Copyright":         copyright,
	})
}

// NotFound is a function for not found page
func (t *template) NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"Title":        "404 Not Found",
		"ErrorMessage": "Lost in the Mind Palace !!",
		"Copyright":    copyright,
	})
}

// NewTemplate is a function for new template
func NewTemplate(postRepo repository.CaseFiles) Template {
	return &template{
		postRepo: postRepo,
	}
}
