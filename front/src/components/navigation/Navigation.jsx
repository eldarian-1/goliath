import { NavLink } from "react-router-dom"
import styles from "./Navigation.module.css"

export default function Navigation() {
  const menuItems = [
    { id: 'dashboard', path: '/', icon: '📊', label: 'Dashboard' },
    { id: 'users', path: '/users', icon: '👥', label: 'Users' },
    { id: 'files', path: '/files', icon: '📁', label: 'Files' },
    { id: 'videos', path: '/videos', icon: '📹', label: 'Videos' },
    { id: 'cache', path: '/cache', icon: '💾', label: 'Cache' },
    { id: 'logs', path: '/logs', icon: '📝', label: 'Logs', badge: '12' },
  ]

  const settingsItems = [
    { id: 'settings', path: '/settings', icon: '⚙️', label: 'Settings' },
    { id: 'help', path: '/help', icon: '❓', label: 'Help' },
  ]

  return (
    <nav className={styles.navigation}>
      <div className={styles.navSection}>
        <div className={styles.navSectionTitle}>Main Menu</div>
        <ul className={styles.navMenu}>
          {menuItems.map(item => (
            <li key={item.id} className={styles.navItem}>
              <NavLink
                to={item.path}
                className={({ isActive }) => `${styles.navLink} ${isActive ? styles.active : ''}`}
                end={item.path === '/'}
              >
                <span className={styles.navIcon}>{item.icon}</span>
                <span className={styles.navText}>{item.label}</span>
                {item.badge && <span className={styles.navBadge}>{item.badge}</span>}
              </NavLink>
            </li>
          ))}
        </ul>
      </div>

      <div className={styles.navDivider}></div>

      <div className={styles.navSection}>
        <div className={styles.navSectionTitle}>System</div>
        <ul className={styles.navMenu}>
          {settingsItems.map(item => (
            <li key={item.id} className={styles.navItem}>
              <NavLink
                to={item.path}
                className={({ isActive }) => `${styles.navLink} ${isActive ? styles.active : ''}`}
              >
                <span className={styles.navIcon}>{item.icon}</span>
                <span className={styles.navText}>{item.label}</span>
              </NavLink>
            </li>
          ))}
        </ul>
      </div>

      <div className={styles.navFooter}>
        <div className={styles.navFooterItem}>
          <span className={styles.navIcon}>ℹ️</span>
          <span className={styles.navText}>Version 1.0.0</span>
        </div>
      </div>
    </nav>
  )
}
