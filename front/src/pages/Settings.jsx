export default function Settings() {
  return (
    <div>
      <h2 style={{ 
        fontSize: '1.875rem', 
        fontWeight: '700', 
        color: '#2d3748',
        marginBottom: '1rem'
      }}>
        Settings
      </h2>
      <p style={{ 
        color: '#718096', 
        fontSize: '1.125rem',
        lineHeight: '1.75',
        marginBottom: '2rem'
      }}>
        Configure system settings and preferences.
      </p>
      
      <div style={{
        background: 'white',
        borderRadius: '12px',
        padding: '1.5rem',
        boxShadow: '0 1px 3px rgba(0, 0, 0, 0.1)'
      }}>
        <p style={{ color: '#718096', textAlign: 'center', padding: '2rem' }}>
          Settings interface will be implemented here.
        </p>
      </div>
    </div>
  )
}