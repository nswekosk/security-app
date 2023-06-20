package api

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
)

var jwtKey = []byte("your_secret_key")

// Credentials is a struct to read the username and password from the request body
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Server serves the directory browser API and webapp.
type Server struct {
	handler http.Handler
}

// NewServer creates a directory browser server.
// It serves webassets from the provided filesystem.
func NewServer(webassets fs.FS) (*Server, error) {
	mux := http.NewServeMux()
	s := &Server{handler: CORS(mux)}

	// API routes
	mux.Handle("/api/hello", http.HandlerFunc(s.hello))

	// web assets
	hfs := http.FS(webassets)
	files := CORS(http.FileServer(hfs))
	mux.Handle("/static/", files)
	mux.Handle("/favicon.ico", files)
	mux.Handle("/files", authMiddleware(http.HandlerFunc(fileHandler)))
	mux.Handle("/login", http.HandlerFunc(loginHandler))

	// fall back to index.html for all unknown routes
	index, err := extractIndexHTML(hfs)
	if err != nil {
		return nil, err
	}
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write(index); err != nil {
			log.Println("failed to serve index.html", err)
		}
	}))

	return s, nil
}

func CORS(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			w.Header().Set("VARY", "Origin")
			w.Header().Set("VARY", "Access-Control-Allow-Method")
			w.Header().Set("VARY", "Access-Control-Allow-Headers")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.WriteHeader(http.StatusOK)
		}
		h.ServeHTTP(w, r)
	})
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.handler)
}

// hello is an example API endpoint
func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

func extractIndexHTML(fs http.FileSystem) ([]byte, error) {
	f, err := fs.Open("index.html")
	if err != nil {
		return nil, fmt.Errorf("could not open index.html: %w", err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("could not read index.html: %w", err)
	}

	return b, nil
}
