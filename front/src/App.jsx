import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom"
import { useAuth } from "./contexts/auth"
import Login from "./components/login/Login"
import Header from "./components/header/Header"
import Navigation from "./components/navigation/Navigation"
import Footer from "./components/footer/Footer"
import Notification from "./components/notification/Notification"
import Dashboard from "./pages/Dashboard"
import Users from "./pages/Users"
import Files from "./pages/Files"
import Cache from "./pages/Cache"
import Logs from "./pages/Logs"
import Settings from "./pages/Settings"
import Help from "./pages/Help"
import styles from "./App.module.css"

export default function App() {
  const { user, register, login, logout, loading, error, clearError } = useAuth()

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

  if (user === null) {
    return (
      <>
        <Login register={register} login={login} />
        {error && <Notification message={error} onClose={clearError} />}
      </>
    )
  }

  return (
    <BrowserRouter>
      <div className={styles.appContainer}>
        <Header />
        
        <div className={styles.mainLayout}>
          <Navigation />
          
          <div className={styles.contentWrapper}>
            <main className={styles.contentArea}>
              <Routes>
                <Route path="/" element={<Dashboard />} />
                <Route path="/users" element={<Users />} />
                <Route path="/files" element={<Files />} />
                <Route path="/cache" element={<Cache />} />
                <Route path="/logs" element={<Logs />} />
                <Route path="/settings" element={<Settings />} />
                <Route path="/help" element={<Help />} />
                <Route path="*" element={<Navigate to="/" replace />} />
              </Routes>
            </main>
            
            <Footer />
          </div>
        </div>
      </div>
    </BrowserRouter>
  )
}
