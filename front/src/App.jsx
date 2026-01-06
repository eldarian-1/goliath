import { useState, useEffect } from "react"
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom"
import { useAuth } from "./contexts/auth"
import Login from "./components/login/Login"
import Header from "./components/header/Header"
import Navigation from "./components/navigation/Navigation"
import Footer from "./components/footer/Footer"
import Notification from "./components/notification/Notification"
import Dashboard from "./pages/Dashboard"
import Users from "./pages/Users"
import UserEdit from "./pages/UserEdit"
import Files from "./pages/Files"
import Videos from "./pages/Videos"
import Cache from "./pages/Cache"
import Logs from "./pages/Logs"
import Settings from "./pages/Settings"
import Help from "./pages/Help"
import About from "./pages/About"
import Documentation from "./pages/Documentation"
import Support from "./pages/Support"
import Privacy from "./pages/Privacy"
import styles from "./App.module.css"

export default function App() {
  const { user, register, login, logout, loading, error, clearError } = useAuth()
  const [showLoginModal, setShowLoginModal] = useState(false)

  // Close modal when user successfully logs in or registers
  useEffect(() => {
    if (user && showLoginModal) {
      setShowLoginModal(false)
    }
  }, [user, showLoginModal])

  if (loading) {
    return (
      <div className={styles.loadingContainer}>
        <div>
          <div className={styles.loadingSpinner}></div>
          <div className={styles.loadingText}>Loading...</div>
        </div>
      </div>
    )
  }

  return (
    <BrowserRouter>
      <div className={styles.appContainer}>
        <Header onLoginClick={() => setShowLoginModal(true)} />
        
        <div className={styles.mainLayout}>
          <Navigation />
          
          <div className={styles.contentWrapper}>
            <main className={styles.contentArea}>
              <Routes>
                <Route path="/" element={<Dashboard />} />
                <Route path="/users" element={<Users />} />
                <Route path="/users/:id/edit" element={<UserEdit />} />
                <Route path="/files" element={<Files />} />
                <Route path="/videos/*" element={<Videos />} />
                <Route path="/cache" element={<Cache />} />
                <Route path="/logs" element={<Logs />} />
                <Route path="/settings" element={<Settings />} />
                <Route path="/help" element={<Help />} />
                <Route path="/about" element={<About />} />
                <Route path="/documentation" element={<Documentation />} />
                <Route path="/support" element={<Support />} />
                <Route path="/privacy" element={<Privacy />} />
                <Route path="*" element={<Navigate to="/" replace />} />
              </Routes>
            </main>
            
            <Footer />
          </div>
        </div>

        {showLoginModal && (
          <Login
            register={register}
            login={login}
            onClose={() => setShowLoginModal(false)}
          />
        )}

        {error && <Notification message={error} onClose={clearError} />}
      </div>
    </BrowserRouter>
  )
}
