import { useState, useEffect } from 'react'
import { useParams, useNavigate, Link } from 'react-router-dom'
import styles from './UserEdit.module.css'
import { fetchWithRefresh } from '../helpers/fetch'
import { useAuth } from '../contexts/auth'

const AVAILABLE_PERMISSIONS = [
  { value: 'read:own', label: 'Read Own Data' },
  { value: 'write:own', label: 'Write Own Data' },
  { value: 'read:all', label: 'Read All Data' },
  { value: 'write:all', label: 'Write All Data' },
  { value: 'delete:all', label: 'Delete All Data' },
  { value: 'videos:read', label: 'Read Videos' },
  { value: 'videos:write', label: 'Write Videos' },
  { value: 'videos:delete', label: 'Delete Videos' },
  { value: 'videos:*', label: 'All Video Permissions' },
  { value: 'users:read', label: 'Read Users' },
  { value: 'users:write', label: 'Write Users' },
  { value: 'users:delete', label: 'Delete Users' },
  { value: 'users:*', label: 'All User Permissions' },
  { value: 'admin', label: 'Administrator (All Permissions)' },
]

export default function UserEdit() {
  const { id } = useParams()
  const navigate = useNavigate()
  const { user: currentUser } = useAuth()
  const [user, setUser] = useState(null)
  const [loading, setLoading] = useState(true)
  const [saving, setSaving] = useState(false)
  const [error, setError] = useState(null)
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    permissions: []
  })

  // Check if current user has permission to manage users
  const canManageUsers = currentUser?.permissions?.includes('users:write') ||
                         currentUser?.permissions?.includes('users:*') ||
                         currentUser?.permissions?.includes('admin')

  useEffect(() => {
    if (!canManageUsers) {
      navigate('/users')
      return
    }
    fetchUser()
  }, [id, canManageUsers, navigate])

  const fetchUser = async () => {
    try {
      setLoading(true)
      setError(null)
      
      // Получаем список пользователей и находим нужного
      const response = await fetchWithRefresh(`/api/v1/users?limit=1000&with_deleted=true`)
      if (response.ok) {
        const data = await response.json()
        const foundUser = data.users?.find(u => u.id === parseInt(id))
        
        if (foundUser) {
          setUser(foundUser)
          setFormData({
            name: foundUser.name || '',
            email: foundUser.email || '',
            permissions: foundUser.permissions || []
          })
        } else {
          setError('User not found')
        }
      } else {
        setError('Failed to fetch user')
      }
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    
    if (!formData.name.trim() || !formData.email.trim()) {
      setError('Name and email are required')
      return
    }

    try {
      setSaving(true)
      setError(null)

      const response = await fetchWithRefresh('/api/v1/users', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          id: parseInt(id),
          name: formData.name,
          email: formData.email,
          permissions: formData.permissions
        })
      })

      if (response.ok) {
        navigate('/users')
      } else {
        const data = await response.json()
        setError(data.message || 'Failed to update user')
      }
    } catch (err) {
      setError(err.message)
    } finally {
      setSaving(false)
    }
  }

  const togglePermission = (permission) => {
    const currentPermissions = formData.permissions || []
    const newPermissions = currentPermissions.includes(permission)
      ? currentPermissions.filter(p => p !== permission)
      : [...currentPermissions, permission]

    setFormData({ ...formData, permissions: newPermissions })
  }

  const handleDelete = async () => {
    if (!confirm('Are you sure you want to delete this user?')) return

    try {
      setSaving(true)
      setError(null)

      const response = await fetchWithRefresh(`/api/v1/users?id=${id}`, {
        method: 'DELETE'
      })

      if (response.ok) {
        navigate('/users')
      } else {
        const data = await response.json()
        setError(data.message || 'Failed to delete user')
      }
    } catch (err) {
      setError(err.message)
    } finally {
      setSaving(false)
    }
  }

  if (loading) {
    return (
      <div className={styles.container}>
        <div className={styles.loading}>Loading user...</div>
      </div>
    )
  }

  if (!user) {
    return (
      <div className={styles.container}>
        <div className={styles.error}>
          User not found
          <Link to="/users" className={styles.backLink}>Back to Users</Link>
        </div>
      </div>
    )
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <div>
          <Link to="/users" className={styles.backButton}>
            ← Back to Users
          </Link>
          <h2 className={styles.title}>Edit User</h2>
          <p className={styles.subtitle}>
            Edit user information and permissions
          </p>
        </div>
      </div>

      {error && (
        <div className={styles.error}>
          {error}
          <button onClick={() => setError(null)}>×</button>
        </div>
      )}

      <div className={styles.formContainer}>
        <form onSubmit={handleSubmit}>
          <div className={styles.formSection}>
            <h3 className={styles.sectionTitle}>Basic Information</h3>
            
            <div className={styles.formGroup}>
              <label htmlFor="name">Name</label>
              <input
                id="name"
                type="text"
                value={formData.name}
                onChange={(e) => setFormData({...formData, name: e.target.value})}
                required
                className={styles.input}
                placeholder="Enter user name"
              />
            </div>

            <div className={styles.formGroup}>
              <label htmlFor="email">Email</label>
              <input
                id="email"
                type="email"
                value={formData.email}
                onChange={(e) => setFormData({...formData, email: e.target.value})}
                required
                className={styles.input}
                placeholder="Enter user email"
              />
            </div>

            <div className={styles.formGroup}>
              <label>User ID</label>
              <div className={styles.idDisplay}>{user.id}</div>
            </div>

            {user.created_at && (
              <div className={styles.formGroup}>
                <label>Created At</label>
                <div className={styles.dateDisplay}>
                  {new Date(user.created_at).toLocaleString()}
                </div>
              </div>
            )}

            {user.updated_at && (
              <div className={styles.formGroup}>
                <label>Last Updated</label>
                <div className={styles.dateDisplay}>
                  {new Date(user.updated_at).toLocaleString()}
                </div>
              </div>
            )}

            {user.deleted_at && (
              <div className={styles.formGroup}>
                <label>Deleted At</label>
                <div className={styles.deletedDisplay}>
                  {new Date(user.deleted_at).toLocaleString()}
                </div>
              </div>
            )}
          </div>

          <div className={styles.formSection}>
            <h3 className={styles.sectionTitle}>Permissions</h3>
            <div className={styles.permissionsGrid}>
              {AVAILABLE_PERMISSIONS.map(perm => (
                <label key={perm.value} className={styles.permissionCheckbox}>
                  <input
                    type="checkbox"
                    checked={formData.permissions?.includes(perm.value)}
                    onChange={() => togglePermission(perm.value)}
                  />
                  <span>{perm.label}</span>
                </label>
              ))}
            </div>
          </div>

          <div className={styles.formActions}>
            <button
              type="button"
              onClick={() => navigate('/users')}
              className={styles.cancelButton}
              disabled={saving}
            >
              Cancel
            </button>
            <button
              type="button"
              onClick={handleDelete}
              className={styles.deleteButton}
              disabled={saving}
            >
              Delete User
            </button>
            <button
              type="submit"
              className={styles.submitButton}
              disabled={saving}
            >
              {saving ? 'Saving...' : 'Save Changes'}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}

