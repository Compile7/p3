import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import Auth from './Components/Auth'
import './index.css'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <Auth />
  </React.StrictMode>,
)
