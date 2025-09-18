package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

// Estamos diciendo que el titulo de nuestro archivo debe llamarse asi
// Ademas puede contener la siguiente palabra cualquier letra o numero con mas de un caracter
var validPath = regexp.MustCompile("^/(edit|view|save)/([a-zA-Z0-9]+)$")

// Cache para guardar las templates (Must provoca un panic si hay un error cargando templates)
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	fileName := p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600) // Archivo con permisos de lectura y escritura
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

// makeHandler es una función de utilidad que toma como argumento un handler
// personalizado (fn) que necesita un título extraído de la URL.
//
// En Go, los handlers que se registran con http.HandleFunc deben tener
// exactamente esta firma:
//     func(http.ResponseWriter, *http.Request)
//
// Pero si queremos que nuestras funciones también reciban un "title" como
// tercer parámetro (por ejemplo: /view/TituloPagina), no podemos usarlas
// directamente. Aquí es donde makeHandler entra en acción.
//
// Esta función devuelve un nuevo handler (función anónima) compatible con
// http.HandleFunc, que:
//   1. Valida que la URL cumpla con el formato esperado usando una expresión regular.
//   2. Extrae el título de la URL si es válido.
//   3. Llama a la función original (fn) pasándole el ResponseWriter, la Request
//      y el título extraído.
//
// Si la URL no es válida, se responde con un error 404.
//
// Ejemplo de uso:
//     http.HandleFunc("/view/", makeHandler(viewHandler))
//
// En ese caso, cuando el usuario acceda a /view/Nombre,
// la función viewHandler será llamada como:
//     viewHandler(w, r, "Nombre")
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		matchers := validPath.FindStringSubmatch(r.URL.Path) // Busca una coincidencia con la regex en el path
		if matchers == nil {
			http.NotFound(w, r) // Si no encuentra nada que coincida devuelve un error
			return
		}
		fn(w, r, matchers[2])
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body") // Capturamos el valor de un formulario
	page := &Page{Title: title, Body: []byte(body)}
	err := page.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusPermanentRedirect)

}

// Controladores o Handlers
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusPermanentRedirect)
	}

	renderTemplate(w, page, "view")
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}

	renderTemplate(w, page, "edit")
}

func renderTemplate(w http.ResponseWriter, page *Page, templateName string) {
	err := templates.ExecuteTemplate(w, templateName+".html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
