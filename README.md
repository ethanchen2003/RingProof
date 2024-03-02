# RingProof
Application for Caller ID verification.
This application provides spam number identification and will rely on crowd-sourcing information. It provides basic functionalities such as user registration, authentication, phone number submission, identification, and spam filtering.

## Features
- **User Registration and Authentication**: Securely register and authenticate users.
- **Phone Number Submission**: Users can submit phone numbers and mark them as spam or not.
- **Spam Filtering**: Simple algorithm to identify numbers as spam based on user submissions.
- **Phone Number Query**: API to query the spam status and identification of phone numbers.
- **Basic Web Interface**: A simple frontend to interact with the backend services.

## Tech Stack
- **Frontend**: HTML, CSS, JavaScript (optional: React/Vue.js for more dynamic interfaces)
- **Backend**: Go (Gin/Gorilla Mux for routing, GORM/SQL driver for database interaction)
- **Database**: MySQL
- **Authentication**:
- **Testing**: Go's built-in testing package
- **Version Control**: Git


## Project Structure

- `/cmd`: Main applications for this project.
- `/internal`: Private application and library code.
- `/pkg`: Library code that's ok to use by external applications.
- `/api`: API controllers and routing information.
- `/models`: Data models and database interaction.
- `/middleware`: HTTP middleware functions.
- `/auth`: Authentication and authorization logic.

## Acknowledgments

- Go community for invaluable resources and libraries.
- Gin/Gorilla Mux for routing.
- GORM for object-relational mapping.
- All contributors and users of this project.

## Instructions

1. set envrionment variable
GOPATH and GOROOT

For best pratice, set
export GOPATH=$HOME/go
export GOROOT="$(/usr/local/bin/brew --prefix golang)/libexec"

2. under GOPATH, create folder if it doesn't exist. (if you install golang correct, the folder should be there)
mkdir src

3. under src, run command:
 mkdir github.com
 cd github.com
 
4. move phoneapp folder to github.com

5. under phoneapp/web folder, run

go run main.go

6. Open a browser, enter http://localhost:8080, the login page is up.

7. enter username as user1, password as password1


The structure of the code:
	it has three packages: main, security, and service

	both of security and service are local packages and cannot download from github.com
	GoLang may complain about finding the package. Run following command can solve this:
	
	go mod edit -replace github.com/phoneapp/security=../security
	go mod edit -replace github.com/phoneapp/service=../service
	go mod tidy
	
	

 web/main.go is the main program, it starts the server and register the http handler.
 security/login.go is security package, it contains the functions about login, logout, create security session.
 service/service.go will implement two functions: SubmitSpamPhone and CheckPhone. Both of them will insert/update to database.



