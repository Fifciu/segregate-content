import React from 'react'
import ReactDOM from 'react-dom/client'
import Home from './app/Home.tsx'
import NewProject from './app/NewProject.tsx'
import './index.css'
import { HashRouter, Route, Routes } from "react-router";

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <HashRouter>
      <Routes>
        <Route index element={<Home />} />
        <Route path="new-project" element={<NewProject />} />
        {/* <Route path="about" element={<About />} /> */}
      </Routes>
    </HashRouter>
  </React.StrictMode>,
)
