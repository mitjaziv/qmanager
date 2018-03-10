// Handler package is used for http request handling.
//
// Disclaimer: In normal usage, I would use router to enable multiple HTTP end points.
package manager

import (
	"fmt"
	"log"
	"net/http"
	"text/tabwriter"
	"time"
)

type (
	HttpHandler interface {
		ServeHTTP(w http.ResponseWriter, r *http.Request)
		RegisterHandler()
	}

	httpHandler struct {
		manager Manager
	}
)

func NewHttpHandler(m Manager) HttpHandler {
	return &httpHandler{
		manager: m,
	}
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Header
	w.Header().Set("Content-Type", "text/plain")

	// Current time
	_, err := w.Write([]byte("Current time is: " + time.Now().String() + "\n"))
	if err != nil {
		log.Println(err)
	}

	// Write counts
	wait, progress, done := h.manager.Status()
	_, err = w.Write([]byte("Wait:\t\t" + fmt.Sprintf("%d", wait) + "\n"))
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write([]byte("Progress:\t" + fmt.Sprintf("%d", progress) + "\n"))
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write([]byte("Done:\t\t" + fmt.Sprintf("%d", done) + "\n\n"))
	if err != nil {
		log.Println(err)
	}

	// Write done list
	_, err = w.Write([]byte("Done list:\n"))
	if err != nil {
		log.Println(err)
	}

	// Create tab writer for formatted output
	tb := tabwriter.NewWriter(w, 0, 0, 2, ' ', tabwriter.TabIndent)

	// Collect done list and format it
	list := h.manager.DoneList()
	for _, val := range list {
		_, err := fmt.Fprintln(tb, val)
		if err != nil {
			log.Println(err)
		}
	}

	// Flush done list to tab writer.
	err = tb.Flush()
	if err != nil {
		log.Println(err)
	}
}

// RegisterHandler registers HTTP handler
func (h *httpHandler) RegisterHandler() {
	http.Handle("/", h)
}
