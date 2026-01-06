import { useAuth } from "../../contexts/auth"
import styles from "./Header.module.css"

export default function Header({ onLoginClick }) {
  const { user, logout } = useAuth()

  const getInitials = (id) => {
    return `U${id.toString().slice(-2)}`
  }

  return (
    <header className={styles.header}>
      <div className={styles.headerLeft}>
        <h1 className={styles.headerLogo}>Goliath</h1>
        <span className={styles.headerSubtitle}>Management System</span>
      </div>
      
      <div className={styles.headerRight}>
        {user ? (
          <>
            <div className={styles.userInfo}>
              <div className={styles.userAvatar}>
                {getInitials(user.id)}
              </div>
              <div className={styles.userDetails}>
                <span className={styles.userName}>{user.name}</span>
                <span className={styles.userId}>ID: {user.id}</span>
              </div>
            </div>
            
            <button className={styles.logoutBtn} onClick={logout}>
              Logout
            </button>
          </>
        ) : (
          <button className={styles.loginBtn} onClick={onLoginClick}>
            Login / Register
          </button>
        )}
      </div>
    </header>
  )
}
