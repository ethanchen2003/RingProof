# RingProof
Application for Caller ID verification
This full-stack application provides spam number identification that relies on crowd-sourcing information. It provides basic functionalities such as user registration, authentication, phone number submission, identification, and spam filtering.

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

