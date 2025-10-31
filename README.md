# 📝 To-Do CLI App (Golang)

A simple **Command Line Interface (CLI)** To-Do application built with [Golang](https://go.dev/).  
This project helps you manage tasks directly from the terminal — add, list, and mark tasks as done.  
Created as part of my Golang learning portfolio.

---

## 🚀 Features

- ✅ Add new tasks  
- 📋 List all tasks  
- 🏁 Mark tasks as completed  
- 💾 Persistent storage using JSON file  

---

## 🧩 Project Structure

todo-cli/
├── main.go # Entry point (CLI commands)
├── todo/
│ ├── task.go # Task struct definition
│ ├── storage.go # JSON file storage functions
│ └── manager.go # Core task management logic
└── tasks.json # Task data storage (auto-created)
