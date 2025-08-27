package main

import (
    "html/template"
    "net/http"
)

type DadosPagina struct {
    Titulo   string
    Mensagem string
    Itens    []string
}

func handler(w http.ResponseWriter, r *http.Request) {
    dados := DadosPagina{
        Titulo:   "Página com Layout",
        Mensagem: "Conteúdo dinâmico com layout reutilizável",
        Itens:    []string{"Go", "Templates", "Layouts", "Reutilização"},
    }

    tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))
    tmpl.ExecuteTemplate(w, "layout", dados)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}