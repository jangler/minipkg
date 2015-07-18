package webui

import (
	"bytes"
	"html/template"
	"net"
	"net/http"
	"testing"
	"time"
)

const testContent template.HTML = `
<a href="/">Hyperlink</a><br>
<blockquote>Block quote</blockquote>
<button type="button">Button</button><br>
<cite>Cite</cite><br>
<code>Code</code><br>
<del>Deleted text</del><br>
<dfn>Definition</dfn><br>
<dl><dt>Term</dt><dd>Description</dd></dl>
<em>Emphasized text</em><br>
<form><fieldset><legend>Legend</legend>Fieldset</fieldset><form>
<h1>Heading 1</h1>
<h2>Heading 2</h2>
<h3>Heading 3</h3>
<h4>Heading 4</h4>
<h5>Heading 5</h5>
<h6>Heading 6</h6>
<hr>
<iframe src="/"></iframe><br>
<form>
<input type="button" value="Input button"><br>
<input type="checkbox"> Input checkbox<br>
<input type="file"><br>
Input password <input type="password"><br>
<input type="radio"> Input radio<br>
<input type="reset"><br>
<input type="submit"><br>
Input text <input type="text"><br>
</form>
<ins>Inserted text</ins><br>
<kbd>Keyboard input</kbd><br>
<label>Label</label><br>
<ol><li>Ordered list item</li></ol>
<p>Paragraph</p>
<pre>Preformatted text</pre>
<q>Quotation</q><br>
<samp>Sample output</samp><br>
<small>Small text</small><br>
<select>
<optgroup label="Option group"><option>Option</option></optgroup>
</select><br>
<strong>Strong text</strong><br>
<sub>Subscript</sub><br>
<sup>Suberscript</sup><br>
<table>
<caption>Caption</caption>
<thead><tr><th>Table header</th></tr></thead>
<tbody><tr><td>Table body</td></tr></tbody>
<tfoot><tr><td>Table footer</td></tr></tfoot>
<table>
<var>Variable</var><br>
<ul><li>Unordered list item</li></ul>
`

func TestGetOpenPort(t *testing.T) {
	if got, want := getOpenPort(), ":8080"; got != want {
		t.Errorf("getOpenPort(0) == %#v; want %#v", got, want)
	}
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatalf("%v", err)
	}
	if got, want := getOpenPort(), ":8081"; got != want {
		t.Errorf("getOpenPort(0) == %#v; want %#v", got, want)
	}
	listener.Close()
}

func TestLaunch(t *testing.T) {
	done := make(chan struct{})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		PageTemplate.Execute(w, PageData{
			Title:   "Hello, world!",
			Content: testContent,
		})
		done <- struct{}{}
	})

	http.HandleFunc("/table", func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		TableTemplate.Execute(&buf, TableData{
			Caption: "Caption",
			Header:  []string{"Header 1", "Header 2"},
			Footer:  []string{"Footer 1", "Footer 2"},
			Body: [][]string{
				[]string{"Body A1", "Body A2"},
				[]string{"Body B1", "Body B2"},
			},
		})
		PageTemplate.Execute(w, PageData{
			Title:   "Table",
			Content: template.HTML(buf.String()),
		})
		done <- struct{}{}
	})

	go Start()
	<-done
	<-done
	time.Sleep(time.Second / 10)
}
