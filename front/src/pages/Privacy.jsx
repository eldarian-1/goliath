import styles from "./Dashboard.jsx"

export default function Privacy() {
  return (
    <div>
      <h1>Privacy Policy</h1>
      <div style={{ marginTop: '2rem' }}>
        <p><em>Last updated: {new Date().toLocaleDateString()}</em></p>

        <h2 style={{ marginTop: '2rem' }}>Introduction</h2>
        <p style={{ marginTop: '1rem' }}>
          This Privacy Policy describes how Goliath collects, uses, and protects your personal information 
          when you use our platform.
        </p>

        <h2 style={{ marginTop: '2rem' }}>Information We Collect</h2>
        <ul style={{ marginTop: '1rem', lineHeight: '1.8' }}>
          <li><strong>Account Information:</strong> Username, email, and authentication credentials</li>
          <li><strong>Usage Data:</strong> System logs, access patterns, and feature usage</li>
          <li><strong>File Data:</strong> Files you upload and their metadata</li>
          <li><strong>Technical Data:</strong> IP addresses, browser type, and device information</li>
        </ul>

        <h2 style={{ marginTop: '2rem' }}>How We Use Your Information</h2>
        <ul style={{ marginTop: '1rem', lineHeight: '1.8' }}>
          <li>To provide and maintain our services</li>
          <li>To authenticate and authorize access</li>
          <li>To monitor system performance and security</li>
          <li>To improve our platform and user experience</li>
          <li>To comply with legal obligations</li>
        </ul>

        <h2 style={{ marginTop: '2rem' }}>Data Security</h2>
        <p style={{ marginTop: '1rem' }}>
          We implement industry-standard security measures to protect your data, including:
        </p>
        <ul style={{ marginTop: '1rem', lineHeight: '1.8' }}>
          <li>Encrypted data transmission (HTTPS/TLS)</li>
          <li>Secure authentication with JWT tokens</li>
          <li>Regular security audits and updates</li>
          <li>Access controls and monitoring</li>
        </ul>

        <h2 style={{ marginTop: '2rem' }}>Data Retention</h2>
        <p style={{ marginTop: '1rem' }}>
          We retain your data for as long as your account is active or as needed to provide services. 
          You can request deletion of your data at any time.
        </p>

        <h2 style={{ marginTop: '2rem' }}>Your Rights</h2>
        <ul style={{ marginTop: '1rem', lineHeight: '1.8' }}>
          <li>Access your personal data</li>
          <li>Correct inaccurate data</li>
          <li>Request deletion of your data</li>
          <li>Export your data</li>
          <li>Opt-out of certain data collection</li>
        </ul>

        <h2 style={{ marginTop: '2rem' }}>Contact Us</h2>
        <p style={{ marginTop: '1rem' }}>
          If you have questions about this Privacy Policy, please contact us at privacy@goliath.example.com
        </p>
      </div>
    </div>
  )
}