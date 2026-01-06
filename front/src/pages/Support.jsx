import styles from "./Dashboard.jsx"

export default function Support() {
  return (
    <div>
      <h1>Support</h1>
      <div style={{ marginTop: '2rem' }}>
        <h2>Get Help</h2>
        <p style={{ marginTop: '1rem' }}>Need assistance? We're here to help you get the most out of Goliath.</p>

        <h2 style={{ marginTop: '2rem' }}>Contact Options</h2>
        <div style={{ marginTop: '1rem', lineHeight: '1.8' }}>
          <p><strong>Email Support:</strong> support@goliath.example.com</p>
          <p><strong>Response Time:</strong> Within 24 hours</p>
          <p><strong>Business Hours:</strong> Monday - Friday, 9:00 AM - 6:00 PM (UTC+3)</p>
        </div>

        <h2 style={{ marginTop: '2rem' }}>Community</h2>
        <ul style={{ marginTop: '1rem', lineHeight: '1.8' }}>
          <li>🐙 GitHub - Report issues and contribute</li>
          <li>💬 Discord - Join our community chat</li>
          <li>🐦 Twitter - Follow for updates</li>
        </ul>

        <h2 style={{ marginTop: '2rem' }}>Common Issues</h2>
        <div style={{ marginTop: '1rem' }}>
          <h3>Authentication Problems</h3>
          <p>If you're having trouble logging in, try clearing your browser cache or resetting your password.</p>
          
          <h3 style={{ marginTop: '1rem' }}>File Upload Issues</h3>
          <p>Ensure your files are within the size limit and in a supported format.</p>
          
          <h3 style={{ marginTop: '1rem' }}>Performance Issues</h3>
          <p>Check the system status page and ensure your network connection is stable.</p>
        </div>

        <h2 style={{ marginTop: '2rem' }}>System Status</h2>
        <p style={{ marginTop: '1rem' }}>All systems operational ✅</p>
      </div>
    </div>
  )
}