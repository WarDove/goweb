package main

import (
	"crypto/sha1"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		sID, err := uuid.NewV4()
		checkError(err)
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	//c.MaxAge = 3600
	http.SetCookie(w, c)
	return c
}

func appendCookie(w http.ResponseWriter, c *http.Cookie, imgNames ...string) *http.Cookie {

	s := c.Value

	for _, v := range imgNames {
		if !strings.Contains(s, v) {
			s += "|" + v
		}
	}

	c.Value = s
	http.SetCookie(w, c)
	return c
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	var imgNames []string
	as := req.Header.Get("Host")

	if req.Method == http.MethodPost {
		// IF if was not multiple choice for file , we would use r.FormFile - which auto parses one file
		// and gives the File itself and its header + error  ( req.FormFile ("name of the input from html") )
		err := req.ParseMultipartForm(32 << 20)

		if err != nil {
			log.Println(err)
			//w.WriteHeader(http.StatusInternalServerError) or
			http.Error(w, "Wrong file type", http.StatusBadRequest)
			return
		}

		fhs := req.MultipartForm.File["images"]

		for _, fh := range fhs {

			ext := strings.Split(fh.Filename, ".")[1]

			h := sha1.New()

			mf, err := fh.Open()
			checkError(err)
			defer mf.Close()

			io.Copy(h, mf)

			fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

			wd, err := os.Getwd()
			checkError(err)

			path := filepath.Join(wd, "public", "pics", fname)
			f, err := os.Create(path)
			checkError(err)
			mf.Seek(0, 0)
			io.Copy(f, mf)

			imgNames = append(imgNames, fname)

		}

	}
	c = appendCookie(w, c, imgNames...)

	data := strings.Split(c.Value, "|")

	tpl.ExecuteTemplate(w, "index.gohtml", data[1:])

}

func clnCookie(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/clean", clnCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":8080", nil)

}
