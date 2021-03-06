package blog

import (
	"bytes"
	"github.com/knieriem/markdown"
	"html/template"
	"io"
	"strings"
	"time"
)

var parser = markdown.NewParser(&markdown.Extensions{})

var funcMap = template.FuncMap{
	"dateTime": func(t time.Time) string {
		return t.Format("Monday, January 2, 2006, 15:04")
	},
	"markdown": func(s string) template.HTML {
		var buffer bytes.Buffer
		reader := strings.NewReader(s)
		parser.Markdown(reader, markdown.ToHTML(&buffer))
		return template.HTML(buffer.String())
	},
}

var templates map[string]*template.Template

func init() {
	templates = make(map[string]*template.Template)
	for _, tmpl := range []string{"tmpl/post_page.html", "tmpl/post_edit.html", "tmpl/post_single.html"} {
		templates[tmpl] = template.Must(
			template.New(tmpl).Funcs(funcMap).ParseFiles(
				"tmpl/_layout.html", "tmpl/_post.html", "tmpl/_pagination.html", tmpl))
	}
}

func renderPost(wr io.Writer, post Post, comments []Comment) error {
	return renderTemplate(wr, templates["tmpl/post_single.html"], "layout", map[string]interface{}{
		"baseUri":  "/",
		"Post":     &post,
		"Comments": comments,
	})
}

func renderPosts(wr io.Writer, posts []Post, page, pageCount int) error {
	var previous, next int
	if page > 1 {
		previous = page - 1
	} else {
		previous = 0
	}
	if page < pageCount {
		next = page + 1
	} else {
		next = 0
	}

	pages := make([]bool, pageCount+1)
	pages[page] = true

	return renderTemplate(wr, templates["tmpl/post_page.html"], "layout", map[string]interface{}{
		"baseUri": "/",
		"Posts":   posts,
		"Pagination": map[string]interface{}{
			"Previous": previous,
			"Next":     next,
			"Pages":    pages,
		},
	})
}

func renderEditPost(wr io.Writer, post *Post) error {
	return renderTemplate(wr, templates["tmpl/post_edit.html"], "layout", map[string]interface{}{
		"baseUri": "/",
		"Post":    post,
	})
}

func renderTemplate(wr io.Writer, t *template.Template, name string, data map[string]interface{}) error {
	// Buffer the rendered output so that potential errors don't end up mixed with the output
	var buffer bytes.Buffer
	var err error
	if err = t.ExecuteTemplate(&buffer, name, data); err != nil {
		return err
	}
	_, err = buffer.WriteTo(wr)
	return err
}
