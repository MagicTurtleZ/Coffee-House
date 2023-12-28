package main

import (
	"html/template"
	"path/filepath"
	"woonbeaj/snippetbox/pkg/models"
)

type templateData struct {
    Cheque *models.Order
    Cheques []*models.Order
}

func getTemplate(dir string) ([]string, error) {
    pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
    if err != nil {
        return nil, err
    }
    return pages, nil
}

func getHtml(dir string) ([]string, error) {
    pages, err := filepath.Glob(filepath.Join(dir, "*.html"))
    if err != nil {
        return nil, err
    }
    return pages, nil
}

func parseTempl(dir string, cache *map[string]*template.Template) error {
    pages, err := getTemplate(dir)
    if err != nil {
        return err
    }
    for _, page := range pages {
        name := filepath.Base(page)
        ts, err := template.ParseFiles(page)
        if err != nil {
            return err
        }
        ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
        if err != nil {
            return err
        }
        ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
        if err != nil {
            return err
        }
        (*cache)[name] = ts
    }
    return nil
}

func parseHtml(dir string, cache *map[string]*template.Template) error {
    pages, err := getHtml(dir)
    if err != nil {
        return err
    }
    for _, page := range pages {
        name := filepath.Base(page)
        ts, _ := template.ParseFiles(page)
        (*cache)[name] = ts
    }
    return nil
}

func newCache(dir string) (map[string]*template.Template, error) {

    cache := map[string]*template.Template{}

    err := parseTempl(dir, &cache)
    if err != nil {
        return nil, err
    }
    
    err = parseHtml(dir, &cache)
    if err != nil {
        return nil, err
    }
    return cache, nil
}