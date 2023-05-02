package web

import (
	"embed"
	"github.com/fsnotify/fsnotify"
	"html/template"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

//go:embed *.gohtml
var embedFS embed.FS

var externalFs fs.FS

var embedTemplate *template.Template

var externalTemplate *template.Template

var funcMap = map[string]any{
	"contains":  strings.Contains,
	"hasPrefix": strings.HasPrefix,
	"hasSuffix": strings.HasSuffix,
	"loop": func(n int) []struct{} {
		return make([]struct{}, n)
	},
	"add": func(i int, delta int) int {
		return i + delta
	},
}

func init() {
	tmpl, err := template.New("embed").Funcs(funcMap).ParseFS(embedFS, "*.gohtml")
	if err != nil {
		log.Fatal("embed template parse error:", err)
	}
	embedTemplate = tmpl
}

func ConfigTemplateDir(templateDir string) (*fsnotify.Watcher, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	externalTemplateDir := dir + string(os.PathSeparator) + templateDir
	//dir exists or not
	stat, err := os.Stat(externalTemplateDir)
	if err != nil || !stat.IsDir() {
		return nil, err
	}

	externalFs = os.DirFS(externalTemplateDir)
	externalTemplate, err = template.New("external").Funcs(funcMap).ParseFS(externalFs, "*.gohtml")
	if err != nil {
		log.Println("external template3", externalTemplateDir, templateDir, dir, err)
		return nil, err
	}
	log.Println("external template")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)

					externalTemplate, err = template.New("external").Funcs(funcMap).ParseFS(externalFs, "*.gohtml")
					if err != nil {
						log.Println(err)
					}
				}
				if event.Has(fsnotify.Remove) {
					log.Println("remove file:", event.Name, externalTemplateDir)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	err = watcher.Add(externalTemplateDir)
	if err != nil {
		if watcher != nil {
			watcher.Close()
		}
		return nil, err
	}
	return watcher, nil
}

func ExecuteTemplate(wr io.Writer, name string, data any) {
	var err error
	if externalTemplate != nil {
		err = externalTemplate.ExecuteTemplate(wr, name, data)
	}
	if err != nil || externalTemplate == nil {
		err = embedTemplate.ExecuteTemplate(wr, name, data)
		if err != nil {
			log.Println(err)
		}
	}
}
