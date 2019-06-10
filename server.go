/**
 * Copyright 2019 Shawn Anastasio
 *
 * This file is part of upaste.
 *
 * upaste is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * upaste is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with upaste.  If not, see <https://www.gnu.org/licenses/>.
 */

/**
 * This file contains the HTTP server.
 */

package main

import (
    "log"
    "fmt"
    "os"
    "strings"
    "math/rand"
    "io/ioutil"
    "net/http"
    textTemplate "text/template"
    htmlTemplate "html/template"
)

// CONSTANTS
const idLength = 6

// GLOBALS
var globalConfig Config

// HELPERS
type PasteNotFoundError string
func (e PasteNotFoundError) Error() string {
    return fmt.Sprintf("Unknown paste id: %s", string(e))
}

func loadPaste(id string) ([]byte, error) {
    filename := fmt.Sprintf("%s/%s", globalConfig.FileStorePath, id)
    contents, err := ioutil.ReadFile(filename)
    if err != nil {
        return []byte{}, err
    }

    return contents, nil
}

func randomString(length int) string {
    const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
    str := make([]byte, length)
    for i := 0; i < length; i++ {
        str[i] = byte(characters[rand.Intn(len(characters))])
    }
    return string(str)
}

func newPasteID() string {
    for {
        str := randomString(idLength) 
        path := fmt.Sprintf("%s/%s", globalConfig.FileStorePath, str)

        // Make sure this str isn't taken
        if _, err := os.Stat(path); os.IsNotExist(err) {
            return str
        }
    }
}

func isTerminal(r *http.Request) bool {
    ua := r.UserAgent()

    if strings.HasPrefix(ua, "curl") || strings.HasPrefix(ua, "Wget") {
        return true;
    }

    return false;
}

// HTTP HANDLERS
func indexHandler(w http.ResponseWriter, r *http.Request) {
    // Depending on whether or not the client is a terminal,
    // return either a pure text index or an HTML one
    var err error
    if isTerminal(r) {
        t := textTemplate.New("index")
        t, err = t.Parse(IndexTemplateText)
        if err != nil {
            log.Printf("Couldn't parse template: %v\n", err)
            errorHandler(w, r, http.StatusInternalServerError)
            return
        }

        t.Execute(w, globalConfig)
    } else {
        t := htmlTemplate.New("index")
        t, err = t.Parse(IndexTemplateHTML)
        if err != nil {
            log.Printf("Couldn't parse template: %v\n", err)
            errorHandler(w, r, http.StatusInternalServerError)
            return
        }

        t.Execute(w, globalConfig)
    }
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    var err error
    t := htmlTemplate.New("upload")
    t, err = t.Parse(UploadTemplateHTML)
    if err != nil {
        log.Printf("Couldn't parse template: %v\n", err)
        errorHandler(w, r, http.StatusInternalServerError)
        return
    }

    t.Execute(w, globalConfig)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    pasteID := r.URL.Path[1:]

    data, err := loadPaste(pasteID)
    if err != nil {
        log.Printf("Couldn't handle %v: %v\n", pasteID, err)
        errorHandler(w, r, http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "%s", data)
}

func upasteHandler(w http.ResponseWriter, r *http.Request) {
    data := r.FormValue("upaste")
    if len(data) == 0 || len(data) > globalConfig.MaxPasteSizeBytes {
        errorHandler(w, r, http.StatusBadRequest)
        return
    }

    pasteID := newPasteID()
    path := fmt.Sprintf("%s/%s", globalConfig.FileStorePath, pasteID)
    err := ioutil.WriteFile(path, []byte(data), 0660)
    if err != nil {
        log.Printf("Couldn't write data: %v", err)
        errorHandler(w, r, http.StatusInternalServerError)
        return
    }

    // Return the URL
    fmt.Fprintf(w, "%s/%s\n", globalConfig.ServerURL, pasteID)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    switch status {
        case http.StatusNotFound:
            fmt.Fprint(w, "Paste not found.\n")
        case http.StatusBadRequest:
            fmt.Fprint(w, "Bad request.\n")
        default:
            fmt.Fprint(w, "Unknown error occurred.\n")
    }
}

// ENTRY
func StartServer(c Config) {
    // Initialize things from config
    globalConfig = c

    // Start cleaner thread
    go StartCleaner(c)

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            switch r.URL.Path {
                case "/":
                    indexHandler(w, r)
                case "/upload":
                    uploadHandler(w, r)
                case "/favicon.ico":
                    // TODO
                    errorHandler(w, r, http.StatusNotFound)
                default:
                    defaultHandler(w, r)
            }
        } else if r.Method == "POST" {
            upasteHandler(w, r)
        } else {
            errorHandler(w, r, http.StatusNotImplemented)
        }
    })

    listen := fmt.Sprintf("%s:%d", c.ListenAddress, c.ListenPort)
    log.Printf("Starting upaste on %s\n", listen)
    log.Fatal(http.ListenAndServe(listen, nil))
}
