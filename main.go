package main

import (
	"html/template"
	"net/http"
)

// User struct
type User struct {
	Firstname       string
	Email           string
	Password        string
	ConfirmPassword string
	Errors          map[string]string
}

// Job struct
type Job struct {
	Position       string
	CompanyName    string
	CompanyLogo    string
	Location       string
	JobDescription template.HTML
	URL            string
}

// Jobs struct
type Jobs struct {
	Jobs []Job
}

// PostJob struct
type PostJob struct {
	JobTitle       string
	JobLocation    string
	CompanyPhoto   string
	CompanyName    string
	JobDescription string
	CompanyWebsite string
	ApplyURL       string
	Email          string
	Errors         map[string]string
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/signin", signinHandler)
	http.HandleFunc("/signin-errors", signinErrorsHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/signup-errors", signupErrorsHandler)
	http.HandleFunc("/forgotten-password", forgottenPasswordHandler)
	http.HandleFunc("/post-job", postJobHandler)
	http.HandleFunc("/post-job-errors", postJobErrorsHandler)
	http.HandleFunc("/jobs", joblistHandler)
	http.HandleFunc("/dashboard", dashBoardHandler)
	http.HandleFunc("/not-found", pageNotFoundHandler)
	http.HandleFunc("/internal-error", internalErrorHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/signin.html",
	}
	tp := template.Must(template.ParseFiles(files...))
	tp.ExecuteTemplate(w, "layout", nil)
}

func signinErrorsHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/signin.html",
	}
	funcMap := template.FuncMap{
		"errclass": errclass,
	}
	tp := template.Must(template.New("templates/layout.html").Funcs(funcMap).ParseFiles(files...))
	u := &User{
		Errors: map[string]string{
			"Form":     "Something is wrong with sign in",
			"Email":    "Invalid email",
			"Password": "Incorrect password",
		},
	}

	tp.ExecuteTemplate(w, "layout", u)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/signup.html",
	}
	tp := template.Must(template.ParseFiles(files...))
	tp.ExecuteTemplate(w, "layout", nil)
}

func signupErrorsHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/signup.html",
	}
	funcMap := template.FuncMap{
		"errclass": errclass,
	}
	tp := template.Must(template.New("templates/layout.html").Funcs(funcMap).ParseFiles(files...))
	u := &User{
		Firstname: "John",
		Errors: map[string]string{
			"Form":            "This is a general error for this form",
			"Firstname":       "Incorrrect name",
			"Email":           "Invalid email",
			"Password":        "Incorrect password",
			"ConfirmPassword": "Type password again",
		},
	}

	tp.ExecuteTemplate(w, "layout", u)
}

func forgottenPasswordHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/forgotten-password.html",
	}
	tp := template.Must(template.ParseFiles(files...))
	tp.ExecuteTemplate(w, "layout", nil)
}

func postJobHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/post-job.html",
	}
	tp := template.Must(template.ParseFiles(files...))
	tp.ExecuteTemplate(w, "layout", nil)
}

func postJobErrorsHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/post-job.html",
	}

	funcMap := template.FuncMap{
		"errclass": errclass,
	}

	tp := template.Must(template.New("templates/layout.html").Funcs(funcMap).ParseFiles(files...))

	p := &PostJob{
		Errors: map[string]string{
			"Form":           "This is a general error for this form",
			"JobTitle":       "Job Title is required",
			"JobLocation":    "Specify at least one location",
			"CompanyPhoto":   "Company photo is required",
			"CompanyName":    "Company name is required",
			"JobDescription": "Job Description is required",
			"CompanyWebsite": "Company Website is required",
			"ApplyURL":       "Apply URL is required",
			"Email":          "Email is required",
		},
	}
	tp.ExecuteTemplate(w, "layout", p)
}

func joblistHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/job-list.html",
	}
	tp := template.Must(template.ParseFiles(files...))

	j := &Jobs{
		Jobs: []Job{
			{
				Position:    "Go Developer",
				CompanyName: "Stripe",
				CompanyLogo: "",
				Location:    "UK Remote",
				JobDescription: `<h1>Go Developer Job Description</h1>
				<h2>JD H2 Tag</h2>
				<p>This is a job description</p>`,
				URL: "https://www.stripe.com",
			},
			{
				Position:    "Frontend Developer",
				CompanyName: "Uber",
				CompanyLogo: "",
				Location:    "US Remote",
				JobDescription: `<h1>Frontend Developer</h1>
				<h4>Requirements</h4>
				<ul>
					<li>3 + Years experience in HTML/CSS/JS</li>
					<li>VueJS Experience</li>
				</ul>
				`,
				URL: "https://www.google.com",
			},
		},
	}

	tp.ExecuteTemplate(w, "layout", j)
}

func dashBoardHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/admincp/layout.html",
		"templates/admincp/dashboard.html",
	}
	tp := template.Must(template.ParseFiles(files...))
	tp.ExecuteTemplate(w, "layout", nil)
}

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/errors/layout.html",
		"templates/errors/error404.html",
	}
	tp := template.Must(template.ParseFiles(files...))
	tp.ExecuteTemplate(w, "layout", nil)
}

func internalErrorHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/errors/layout.html",
		"templates/errors/error500.html",
	}
	tp := template.Must(template.ParseFiles(files...))
	tp.ExecuteTemplate(w, "layout", nil)
}

func errclass(s *string) string {
	if s == nil {
		return ""
	}
	if len(*s) > 0 {
		return "input-error"
	}
	return ""
}
