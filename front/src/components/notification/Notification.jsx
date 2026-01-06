import { useEffect } from 'react';
import styles from './Notification.module.css';

export default function Notification({ message, onClose }) {
  useEffect(() => {
    const timer = setTimeout(() => {
      onClose();
    }, 5000);

    return () => clearTimeout(timer);
  }, [onClose]);

  return (
    <div className={styles.notificationOverlay}>
      <div className={styles.notificationContainer}>
        <div className={styles.notificationContent}>
          <div className={styles.notificationIcon}>⚠️</div>
          <div className={styles.notificationMessage}>{message}</div>
          <button className={styles.notificationClose} onClick={onClose} aria-label="Close notification">
            ✕
          </button>
        </div>
      </div>
    </div>
  );
}