export default function Dashboard() {
  return (
    <div>
      <h2 style={{ 
        fontSize: '1.875rem', 
        fontWeight: '700', 
        color: '#2d3748',
        marginBottom: '1rem'
      }}>
        Dashboard
      </h2>
      <p style={{ 
        color: '#718096', 
        fontSize: '1.125rem',
        lineHeight: '1.75',
        marginBottom: '2rem'
      }}>
        Welcome to your dashboard. Here you can view system overview and statistics.
      </p>
      
      <div style={{
        display: 'grid',
        gridTemplateColumns: 'repeat(auto-fit, minmax(250px, 1fr))',
        gap: '1.5rem'
      }}>
        <div style={{
          padding: '1.5rem',
          background: 'linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%)',
          borderRadius: '12px',
          border: '2px solid rgba(102, 126, 234, 0.2)'
        }}>
          <div style={{ fontSize: '2rem', marginBottom: '0.5rem' }}>📊</div>
          <h3 style={{ fontSize: '1.125rem', fontWeight: '600', marginBottom: '0.5rem', color: '#2d3748' }}>
            Total Users
          </h3>
          <p style={{ fontSize: '2rem', fontWeight: '700', color: '#667eea' }}>
            1,234
          </p>
        </div>
        
        <div style={{
          padding: '1.5rem',
          background: 'linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%)',
          borderRadius: '12px',
          border: '2px solid rgba(102, 126, 234, 0.2)'
        }}>
          <div style={{ fontSize: '2rem', marginBottom: '0.5rem' }}>📁</div>
          <h3 style={{ fontSize: '1.125rem', fontWeight: '600', marginBottom: '0.5rem', color: '#2d3748' }}>
            Files Stored
          </h3>
          <p style={{ fontSize: '2rem', fontWeight: '700', color: '#667eea' }}>
            5,678
          </p>
        </div>
        
        <div style={{
          padding: '1.5rem',
          background: 'linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%)',
          borderRadius: '12px',
          border: '2px solid rgba(102, 126, 234, 0.2)'
        }}>
          <div style={{ fontSize: '2rem', marginBottom: '0.5rem' }}>💾</div>
          <h3 style={{ fontSize: '1.125rem', fontWeight: '600', marginBottom: '0.5rem', color: '#2d3748' }}>
            Cache Entries
          </h3>
          <p style={{ fontSize: '2rem', fontWeight: '700', color: '#667eea' }}>
            892
          </p>
        </div>
      </div>
    </div>
  )
}