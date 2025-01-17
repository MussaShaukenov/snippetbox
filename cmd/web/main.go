package main

import (
	"flag"
<<<<<<< HEAD
	"github.com/MussaShaukenov/snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql" // New import”
<<<<<<< HEAD
	"html/template"
=======
>>>>>>> parent of ab1feef (chapter 5: database-driven responses)
=======
>>>>>>> parent of 7edf881 (chapter 5: dynamic html templates)
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom loggers, but
// we'll add more to it as the build progresses.
type application struct {
<<<<<<< HEAD
<<<<<<< HEAD
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
=======
	errorLog *log.Logger
	infoLog  *log.Logger
>>>>>>> parent of ab1feef (chapter 5: database-driven responses)
=======
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
>>>>>>> parent of 7edf881 (chapter 5: dynamic html templates)
}

func main() {
	// Define a new command-line flag with the name 'addr', a default value of ":4000"
	// and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are
	// encountered during parsing the application will be terminated.
	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// Create a logger for writing error messages in the same way, but use stderr as
	// the destination and use the log.Lshortfile flag to include the relevant
	// file name and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

<<<<<<< HEAD
	// To keep the main() function tidy I've put the code for creating a connection
	// pool into the separate openDB() function below. We pass openDB() the DSN
	// from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()

	app := &application{
<<<<<<< HEAD
		infoLog:       infoLog,
		errorLog:      errorLog,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
=======
	// Initialize a new instance of our application struct, containing the
	// dependencies.
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
>>>>>>> parent of ab1feef (chapter 5: database-driven responses)
=======
		infoLog:  infoLog,
		errorLog: errorLog,
		snippets: &models.SnippetModel{DB: db},
>>>>>>> parent of 7edf881 (chapter 5: dynamic html templates)
	}

	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to dereference the pointer (i.e.
	// prefix it with the * symbol) before using it. Note that we're using the
	// log.Printf() function to interpolate the address with the log message.
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
