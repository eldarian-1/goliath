import { createRoot } from 'react-dom/client'
import App from './App.jsx'
import { AuthProvider } from './contexts/auth.jsx'

createRoot(document.getElementById('root')).render(
  <AuthProvider>
    <App />
  </AuthProvider>
)
