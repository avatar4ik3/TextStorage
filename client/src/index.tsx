import React from "react"
import ReactDOM from "react-dom/client"
import { TextApi } from "./apis/storageapi"
import App from "./App"

const root = ReactDOM.createRoot(document.getElementById("root") as HTMLElement)
root.render(<App api={new TextApi("http://localhost:5000")} />)
