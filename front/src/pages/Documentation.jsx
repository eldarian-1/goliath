import styles from "./Dashboard.jsx"

export default function Documentation() {
  return (
    <div>
      <h1>Documentation</h1>
      <div style={{ marginTop: '2rem' }}>
        <h2>Getting Started</h2>
        <p style={{ marginTop: '1rem' }}>Welcome to the Goliath documentation. Here you'll find guides and references for using the platform.</p>

        <h2 style={{ marginTop: '2rem' }}>API Documentation</h2>
        <div style={{ marginTop: '1rem' }}>
          <h3>Authentication</h3>
          <p>Learn how to authenticate with the API using JWT tokens.</p>
          
          <h3 style={{ marginTop: '1rem' }}>Users API</h3>
          <p>Manage user accounts through the REST API.</p>
          
          <h3 style={{ marginTop: '1rem' }}>Files API</h3>
          <p>Upload, download, and manage files.</p>
          
          <h3 style={{ marginTop: '1rem' }}>Cache API</h3>
          <p>Interact with the Redis cache system.</p>
          
          <h3 style={{ marginTop: '1rem' }}>Logs API</h3>
          <p>Query and analyze system logs.</p>
        </div>

        <h2 style={{ marginTop: '2rem' }}>User Guides</h2>
        <ul style={{ marginTop: '1rem', lineHeight: '1.8' }}>
          <li>How to manage users</li>
          <li>Working with file storage</li>
          <li>Cache management best practices</li>
          <li>Monitoring and logging</li>
          <li>System configuration</li>
        </ul>

        <h2 style={{ marginTop: '2rem' }}>Developer Resources</h2>
        <p style={{ marginTop: '1rem' }}>Check the <code>/docs</code> directory in the repository for detailed technical documentation.</p>
      </div>
    </div>
  )
}