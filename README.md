# TSP Solver with Golang & React

This project solves the **Traveling Salesman Problem (TSP)** using Iterated Local Search (ILS) algorithm in a Go API.  
Solutions are continuously improved and streamed in real time via **Server-Sent Events (SSE)**.  
A **React** frontend with **Konva** dynamically visualizes the evolving routes in real time.

---

## Features

- TSP solved with metaheuristics (ILS)  
- Real-time result streaming via SSE  
- Interactive visualization with **React + Konva**  
- Fast backend in **Go**  

---

## Tech Stack

- **Backend**: Go (Golang), Event-Stream (SSE)  
- **Frontend**: React, Konva  

---

## Getting Started

### Prerequisites
- Go `1.20+`  
- Node.js `18+`  
- npm or yarn  

### Backend Setup
```bash
cd backend
go run main.go
```
The backend runs on: [http://localhost:8080](http://localhost:8080)

### Frontend Setup

You'll need a .env file to setup your API url, like this:

VITE_APT_URL=http://localhost:8080

```bash
cd Frontend
npm install
npm run dev
```
The frontend runs on: [http://localhost:3000](http://localhost:3000)

---

## Usage

1. Click on the plane to add cities (points). These will serve as the input for the ILS.
2. When you’re finished, press Solve.
3. The application will stream and visualize the TSP solution as it improves in real time.

---

## Roadmap

- Add more metaheuristic algorithms  
- Improve visualization options  
