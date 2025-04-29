# 📨 Disposify - Temporary Email Service

Disposify is a sleek and modern full-stack application that allows users to create **temporary, disposable email addresses** to receive short-lived messages without revealing their personal inboxes. Ideal for protecting privacy, testing apps, or avoiding spam.

---

## 🌍 Live Project

- 🔗 **Frontend:** [https://disposify.netlify.app](https://disposify.netlify.app)
- 🔗 **Backend API:** Hosted via Render (public endpoint connected)

---

## ✨ Features

- 🔐 Generate unique, time-limited temporary email addresses
- 📥 Receive and display messages in a public inbox
- 🕐 Email auto-expiration (e.g., 10 minutes)
- ⚡ Fast and lightweight API using Go
- 💻 Interactive frontend UI built with vanilla HTML/CSS/JS
- 🌐 Fully deployed using Netlify and Render

---

## 📁 Folder Structure

temp-email-service/ ├── cmd/ │ └── server/ │ └── main.go ├── internal/ │ ├── handler/ │ │ └── handler.go │ └── service/ │ └── service.go ├── static/ │ ├── index.html │ ├── styles.css │ └── script.js ├── go.mod └── README.md



---

## 🚀 Getting Started

### Prerequisites

- Go 1.18+ installed
- Python 3 (to serve static files) or any static file server

---

### 1️⃣ Clone the Repo

```bash
git clone https://github.com/yourusername/temp-email-service.git
cd temp-email-service

--Run Backend on your local system
go run cmd/server/main.go

--Run Frontend
cd static
python3 -m http.server 5500
