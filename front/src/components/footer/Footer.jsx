import styles from "./Footer.module.css"

export default function Footer() {
  const currentYear = new Date().getFullYear()

  return (
    <footer className={styles.footer}>
      <div className={styles.footerContent}>
        <div className={styles.footerLeft}>
          <div className={styles.footerBrand}>Goliath</div>
          <div className={styles.footerCopyright}>
            © {currentYear} Goliath. All rights reserved.
          </div>
        </div>
        
        <div className={styles.footerRight}>
          <ul className={styles.footerLinks}>
            <li>
              <a href="#" className={styles.footerLink}>About</a>
            </li>
            <li>
              <a href="#" className={styles.footerLink}>Documentation</a>
            </li>
            <li>
              <a href="#" className={styles.footerLink}>Support</a>
            </li>
            <li>
              <a href="#" className={styles.footerLink}>Privacy</a>
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