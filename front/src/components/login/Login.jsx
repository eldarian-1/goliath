import { useState } from "react";
import styles from "./Login.module.css";

export default function Login({ register, login, onClose }) {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  const changeEmail = (event) => setEmail(event.target.value)
  const changePassword = (event) => setPassword(event.target.value)

  const handleRegister = () => register(email, password)
  const handleLogin = () => login(email, password)

  const handleBackdropClick = (e) => {
    if (e.target === e.currentTarget) {
      onClose()
    }
  }

  return (
    <div className={styles.modalOverlay} onClick={handleBackdropClick}>
      <div className={styles.loginCard}>
        <button className={styles.closeButton} onClick={onClose}>×</button>
        
        <div className={styles.loginHeader}>
          <h1 className={styles.loginTitle}>Eldarian Studio</h1>
          <p className={styles.loginSubtitle}>Welcome! Please login or register.</p>
        </div>
        
        <form className={styles.loginForm} onSubmit={(e) => e.preventDefault()}>
          <div className={styles.formGroup}>
            <label className={styles.formLabel} htmlFor="email">Email</label>
            <input
              id="email"
              type="email"
              className={styles.formInput}
              placeholder="Enter your email"
              onChange={changeEmail}
              value={email}
              required
            />
          </div>
          
          <div className={styles.formGroup}>
            <label className={styles.formLabel} htmlFor="password">Password</label>
            <input
              id="password"
              type="password"
              className={styles.formInput}
              placeholder="Enter your password"
              onChange={changePassword}
              value={password}
              required
            />
          </div>
          
          <div className={styles.buttonGroup}>
            <button className={`${styles.btn} ${styles.btnPrimary}`} onClick={handleLogin}>
              Login
            </button>
            <button className={`${styles.btn} ${styles.btnSecondary}`} onClick={handleRegister}>
              Register
            </button>
          </div>
        </form>
        
        <div className={styles.loginFooter}>
          <p>Secure authentication powered by Eldarian Studio</p>
        </div>
      </div>
    </div>
  )
}
