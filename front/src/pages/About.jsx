import styles from "./Dashboard.jsx"

export default function About() {
  return (
    <div>
      <h1>About Eldarian Studio</h1>
      <div style={{ marginTop: '2rem' }}>
        <p>Eldarian Studio is a comprehensive system management platform designed to help you manage users, files, cache, and logs efficiently.</p>
        
        <h2 style={{ marginTop: '2rem' }}>Features</h2>
        <ul style={{ marginTop: '1rem', lineHeight: '1.8' }}>
          <li>User Management - Create, update, and manage user accounts</li>
          <li>File Storage - Upload and manage files with S3 integration</li>
          <li>Cache Management - Monitor and control Redis cache</li>
          <li>Log Monitoring - Track system logs and events</li>
          <li>Real-time Metrics - Monitor system performance</li>
        </ul>

        <h2 style={{ marginTop: '2rem' }}>Technology Stack</h2>
        <ul style={{ marginTop: '1rem', lineHeight: '1.8' }}>
          <li>Frontend: React + Vite</li>
          <li>Backend: Go</li>
          <li>Database: PostgreSQL</li>
          <li>Cache: Redis</li>
          <li>Storage: S3</li>
          <li>Message Queue: Kafka/RabbitMQ</li>
          <li>Monitoring: Prometheus + Grafana</li>
        </ul>
      </div>
    </div>
  )
}