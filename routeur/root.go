package routeur

import (
	"boutique/controller"
	"fmt"
	"log"
	"net/http"
)

func NotFoundHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Créez un ResponseWriter personnalisé pour capturer le statut
		customWriter := &statusResponseWriter{ResponseWriter: w, status: http.StatusOK}

		// Appelez le gestionnaire suivant avec notre ResponseWriter personnalisé
		next.ServeHTTP(customWriter, r)

		// Vérifiez si le statut 404 a été capturé
		if customWriter.status == http.StatusNotFound {
			// Redirigez vers la page d'accueil
			http.Redirect(w, r, "/", http.StatusFound)
		}
	})
}

// statusResponseWriter est une enveloppe autour de http.ResponseWriter qui nous permet de capturer le code de statut HTTP
type statusResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func InitServe() {
	FileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", FileServer))

	// Définition d'une route par défaut pour les chemins non trouvés
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			controller.NotFoundPageHandler(w, r)
		} else {
			controller.IndexHandler(w, r)
		}
	})

	if err := http.ListenAndServe(controller.Port, nil); err != nil {
		fmt.Printf("ERREUR LORS DE L'INITIATION DES ROUTES %v \n", err)
		log.Fatal(err)
	}
}
