# Personal Blog

A simple **Personal Blog** application built with **Go (Golang)** using the standard `net/http` package and HTML templates.

This project is part of the **roadmap.sh Backend Projects** and demonstrates how to build a server-rendered web application without using any external frameworks.

## Features

### Guest Section

* View all published articles
* Read individual articles
* Clean HTML template rendering

### Admin Section

* Login authentication using cookies
* Dashboard to manage articles
* Create new articles
* Edit existing articles
* Delete articles
* Logout functionality

## Tech Stack

* Go (Golang)
* net/http
* html/template
* JSON File Storage
* HTML5
* CSS3

## Project Structure

```
personal-blog/
│
├── articles/               # JSON files for each article
│
├── handler/                # HTTP handlers
│   ├── add_article_handler.go
│   ├── admin_handler.go
│   ├── article_handler.go
│   ├── auth.go
│   ├── delete_article_handler.go
│   ├── edit_article_handler.go
│   ├── helper.go
│   ├── home_handler.go
│   ├── login_handler.go
│   └── logout_handler.go
│
├── model/
│   └── article.go
│
├── storage/
│   └── article_storage.go
│
├── static/
│   └── style.css
│
├── templates/
│   ├── home.html
│   ├── article.html
│   └── admin/
│       ├── dashboard.html
│       ├── add_article.html
│       ├── edit_article.html
│       └── login.html
│
├── main.go
├── go.mod
└── README.md
```

## How It Works

Articles are stored as individual JSON files inside the `articles/` directory.

Example:

```
articles/
├── 1.json
├── 2.json
└── 3.json
```

Each article contains:

```json
{
  "id": 1,
  "title": "My First Article",
  "content": "Hello World!",
  "publishedAt": "2026-07-15T14:30:00Z"
}
```

## Authentication

The admin area uses a simple cookie-based authentication.

Default credentials:

| Username | Password |
| -------- | -------- |
| admin    | admin123 |

After successful login, a session cookie is created and required to access all admin pages.

## Routes

### Guest

| Method | Endpoint           | Description  |
| ------ | ------------------ | ------------ |
| GET    | `/`                | Home page    |
| GET    | `/article?id={id}` | View article |

### Authentication

| Method | Endpoint  | Description |
| ------ | --------- | ----------- |
| GET    | `/login`  | Login page  |
| POST   | `/login`  | Login       |
| GET    | `/logout` | Logout      |

### Admin

| Method | Endpoint                | Description       |
| ------ | ----------------------- | ----------------- |
| GET    | `/admin`                | Dashboard         |
| GET    | `/admin/add`            | Add article form  |
| POST   | `/admin/add`            | Create article    |
| GET    | `/admin/edit?id={id}`   | Edit article form |
| POST   | `/admin/edit?id={id}`   | Update article    |
| POST   | `/admin/delete?id={id}` | Delete article    |

## Running the Project

Clone the repository:

```bash
git clone https://github.com/<your-username>/personal-blog.git
```

Go into the project folder:

```bash
cd personal-blog
```

Run the application:

```bash
go run .
```

The server will start at:

```
http://localhost:8080
```

## Future Improvements

* Use SQLite or PostgreSQL instead of filesystem storage
* Password hashing
* Secure session management
* Markdown support
* Categories
* Tags
* Search articles
* Pagination
* Rich text editor
* Image upload
* Comments
* Responsive UI improvements

## Learning Outcomes

Through this project, I practiced:

* Building web servers using Go
* HTTP routing
* HTML template rendering
* Handling forms
* Cookie-based authentication
* CRUD operations
* Filesystem operations
* JSON encoding/decoding
* Organizing a Go project into packages
* Error handling

## Reference

Project idea from:

https://roadmap.sh/projects/personal-blog

## License

This project is created for learning purposes.
