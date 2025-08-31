package ourcode

import (
	"net/http"
	"text/template"
)

// HomeHandler handles GET requests to the root ("/") route.
// It loads and renders the index.html template.
// If the method is not GET or path is not "/", it returns appropriate error messages.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		RenderWithError(w, "Method not allowed", 405)
		return
	}

	if r.URL.Path != "/" {
		RenderWithError(w, "Page not found", 404)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderWithError(w, "Template not found", 404)
		return
	}

	data := PageData{}
	if err := tmpl.Execute(w, data); err != nil {
		RenderWithError(w, "Internal server error", 500)
		return
	}
}

// AsciiArtHandler handles POST requests to generate ASCII art.
// It parses form data, validates input, and renders the result in the template.
// If any error occurs, it returns an appropriate error response.
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		RenderWithError(w, "Method not allowed", 405)
		return
	}

	err := r.ParseForm()
	if err != nil {
		RenderWithError(w, "Bad request", 400)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || len(text) > 1000 {
		RenderWithError(w, "Bad request", 400)
		return
	}

	if banner == "" {
		RenderWithError(w, "Template not found", 404)
		return
	}

	validBanners := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}

	if !validBanners[banner] {
		RenderWithError(w, "Template not found", 404)
		return
	}

	result, err := GenerateASCIIArt(text, banner)
	if err != nil {
		RenderWithError(w, "Error generating ASCII art:", 500)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderWithError(w, "Template not found", 404)
		return
	}

	data := PageData{
		Input:  text,
		Banner: banner,
		Result: result,
	}

	if err := tmpl.Execute(w, data); err != nil {
		RenderWithError(w, "Internal server error", 500)
		return
	}
}

// CssHandler serves CSS files from the /css/ directory.
// If the request path is "/css/" exactly, it returns a 404 error.
func CssHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/css/" {
		RenderWithError(w, "Page not found", 404)
		return
	}

	http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))).ServeHTTP(w, r)
}
