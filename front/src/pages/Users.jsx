import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import styles from './Users.module.css'
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

export default function Users() {
  const { user: currentUser } = useAuth()
  const [users, setUsers] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)
  const [showCreateModal, setShowCreateModal] = useState(false)
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    password: '',
    permissions: ['read:own', 'write:own']
  })

  // Check if current user has permission to manage users
  const canManageUsers = currentUser?.permissions?.includes('users:write') ||
                         currentUser?.permissions?.includes('users:*') ||
                         currentUser?.permissions?.includes('admin')

  useEffect(() => {
    fetchUsers()
  }, [])

  const fetchUsers = async () => {
    try {
      setLoading(true)
      const response = await fetchWithRefresh('/api/v1/users?limit=100')
      if (response.ok) {
        const data = await response.json()
        setUsers(data.users || [])
      } else {
        setError('Failed to fetch users')
      }
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const handleCreateUser = async (e) => {
    e.preventDefault()
    try {
      const response = await fetchWithRefresh('/api/v1/users', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          name: formData.name,
          email: formData.email,
          password: formData.password,
          permissions: formData.permissions
        })
      })

      if (response.ok) {
        setShowCreateModal(false)
        setFormData({
          name: '',
          email: '',
          password: '',
          permissions: ['read:own', 'write:own']
        })
        fetchUsers()
      } else {
        const data = await response.json()
        setError(data.message || 'Failed to create user')
      }
    } catch (err) {
      setError(err.message)
    }
  }


  const handleDeleteUser = async (userId) => {
    if (!confirm('Are you sure you want to delete this user?')) return

    try {
      const response = await fetchWithRefresh(`/api/v1/users?id=${userId}`, {
        method: 'DELETE'
      })

      if (response.ok) {
        fetchUsers()
      } else {
        const data = await response.json()
        setError(data.message || 'Failed to delete user')
      }
    } catch (err) {
      setError(err.message)
    }
  }

  const togglePermission = (permission) => {
    const currentPermissions = formData.permissions || []
    const newPermissions = currentPermissions.includes(permission)
      ? currentPermissions.filter(p => p !== permission)
      : [...currentPermissions, permission]

    setFormData({ ...formData, permissions: newPermissions })
  }

  if (loading) {
    return (
      <div className={styles.container}>
        <div className={styles.loading}>Loading users...</div>
      </div>
    )
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <div>
          <h2 className={styles.title}>Users Management</h2>
          <p className={styles.subtitle}>
            Manage user accounts and permissions
          </p>
        </div>
        {canManageUsers && (
          <button
            className={styles.createButton}
            onClick={() => setShowCreateModal(true)}
          >
            + Create User
          </button>
        )}
      </div>

      {error && (
        <div className={styles.error}>
          {error}
          <button onClick={() => setError(null)}>×</button>
        </div>
      )}

      <div className={styles.tableContainer}>
        <table className={styles.table}>
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Email</th>
              <th>Permissions</th>
              <th>Created</th>
              {canManageUsers && <th>Actions</th>}
            </tr>
          </thead>
          <tbody>
            {users.map(user => (
              <tr key={user.id}>
                <td>{user.id}</td>
                <td>{user.name}</td>
                <td>{user.email}</td>
                <td>
                  <div className={styles.permissions}>
                    {user.permissions?.map(perm => (
                      <span key={perm} className={styles.permissionBadge}>
                        {perm}
                      </span>
                    ))}
                  </div>
                </td>
                <td>{new Date(user.created_at).toLocaleDateString()}</td>
                {canManageUsers && (
                  <td>
                    <div className={styles.actions}>
                      <Link
                        to={`/users/${user.id}/edit`}
                        className={styles.editButton}
                      >
                        Edit
                      </Link>
                      <button
                        onClick={() => handleDeleteUser(user.id)}
                        className={styles.deleteButton}
                      >
                        Delete
                      </button>
                    </div>
                  </td>
                )}
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {showCreateModal && (
        <div className={styles.modal}>
          <div className={styles.modalContent}>
            <div className={styles.modalHeader}>
              <h3>Create New User</h3>
              <button onClick={() => setShowCreateModal(false)}>×</button>
            </div>
            <form onSubmit={handleCreateUser}>
              <div className={styles.formGroup}>
                <label>Name</label>
                <input
                  type="text"
                  value={formData.name}
                  onChange={(e) => setFormData({...formData, name: e.target.value})}
                  required
                  className={styles.input}
                />
              </div>
              <div className={styles.formGroup}>
                <label>Email</label>
                <input
                  type="email"
                  value={formData.email}
                  onChange={(e) => setFormData({...formData, email: e.target.value})}
                  required
                  className={styles.input}
                />
              </div>
              <div className={styles.formGroup}>
                <label>Password</label>
                <input
                  type="password"
                  value={formData.password}
                  onChange={(e) => setFormData({...formData, password: e.target.value})}
                  required
                  className={styles.input}
                />
              </div>
              <div className={styles.formGroup}>
                <label>Permissions</label>
                <div className={styles.permissionsGrid}>
                  {AVAILABLE_PERMISSIONS.map(perm => (
                    <label key={perm.value} className={styles.permissionCheckbox}>
                      <input
                        type="checkbox"
                        checked={formData.permissions.includes(perm.value)}
                        onChange={() => togglePermission(perm.value)}
                      />
                      {perm.label}
                    </label>
                  ))}
                </div>
              </div>
              <div className={styles.modalActions}>
                <button type="submit" className={styles.submitButton}>
                  Create User
                </button>
                <button
                  type="button"
                  onClick={() => setShowCreateModal(false)}
                  className={styles.cancelButton}
                >
                  Cancel
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  )
}