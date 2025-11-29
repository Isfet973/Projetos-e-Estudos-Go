package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gomarkdown/markdown"
)

func main() {
	// Servir CSS da pasta
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		css, err := os.ReadFile("./assets/style.css")
		if err != nil {
			http.Error(w, "CSS não encontrado", 404)
			return
		}
		w.Write(css)
	})

	// Servir página principal com Markdown
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Lê o arquivo Markdown
		md, err := os.ReadFile("./content/recursos.md")
		if err != nil {
			http.Error(w, "Arquivo não encontrado", 404)
			return
		}

		// Converte Markdown para HTML
		htmlContent := markdown.ToHTML(md, nil, nil)

		// Template HTML que vai receber o conteúdo
		tmpl := template.Must(template.ParseFiles("./templates/layout.html"))

		// Dados para passar ao template
		data := map[string]interface{}{
			"Title":    "Mini FMHY - Índice de Recursos",
			"Content":  template.HTML(htmlContent),
		}

		tmpl.Execute(w, data)
	})

	// Servir categorias específicas
	http.HandleFunc("/categoria/", func(w http.ResponseWriter, r *http.Request) {
		// Pega o nome da categoria da URL
		categoria := r.URL.Path[len("/categoria/"):]
		
		// Lê o arquivo correspondente
		md, err := os.ReadFile("./content/" + categoria + ".md")
		if err != nil {
			http.Error(w, "Categoria não encontrada", 404)
			return
		}

		htmlContent := markdown.ToHTML(md, nil, nil)
		tmpl := template.Must(template.ParseFiles("./templates/layout.html"))

		data := map[string]interface{}{
			"Title":   "Categoria: " + categoria,
			"Content": template.HTML(htmlContent),
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}