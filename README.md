
# Temporary Email Service API

A simple Temporary Email API service built in Go using the `chi` router.

---

## Features

- Create temporary email addresses
- Add and receive messages to temporary inboxes
- Automatic expiration and cleanup of emails
- Lightweight and fast API
- Built-in logger middleware

---

## Endpoints

| Method | Endpoint | Description |
|:------|:---------|:------------|
| POST | `/email` | Create a new temporary email |
| GET | `/email/{emailID}` | Get inbox (list of messages) for a temp email |
| POST | `/email/{emailID}/message` | Add a message to a temp email |

---

## Example Usage

### 1. Create a Temporary Email

```bash
curl -X POST http://localhost:8080/email
