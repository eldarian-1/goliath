import { Link } from "react-router-dom"
import styles from "./Footer.module.css"

export default function Footer() {
  const currentYear = new Date().getFullYear()

  return (
    <footer className={styles.footer}>
      <div className={styles.footerContent}>
        <div className={styles.footerLeft}>
          <div className={styles.footerBrand}>Eldarian Studio</div>
          <div className={styles.footerCopyright}>
            © {currentYear} Eldarian Studio. All rights reserved.
          </div>
        </div>
        
        <div className={styles.footerRight}>
          <ul className={styles.footerLinks}>
            <li>
              <Link to="/about" className={styles.footerLink}>About</Link>
            </li>
            <li>
              <Link to="/documentation" className={styles.footerLink}>Documentation</Link>
            </li>
            <li>
              <Link to="/support" className={styles.footerLink}>Support</Link>
            </li>
            <li>
              <Link to="/privacy" className={styles.footerLink}>Privacy</Link>
            </li>
          </ul>
          
          <div className={styles.footerSocial}>
            <div className={styles.socialIcon} title="GitHub">
              🐙
            </div>
            <div className={styles.socialIcon} title="Twitter">
              🐦
            </div>
            <div className={styles.socialIcon} title="Discord">
              💬
            </div>
          </div>
        </div>
      </div>
    </footer>
  )
}