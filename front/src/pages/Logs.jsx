import { useState } from 'react'
import { fetchWithRefresh } from '../helpers/fetch'
import styles from './Logs.module.css'

const LOG_LEVELS = [
  { value: 'info', label: 'Info' },
  { value: 'warning', label: 'Warning' },
  { value: 'error', label: 'Error' },
  { value: 'debug', label: 'Debug' },
  { value: 'trace', label: 'Trace' },
]

const BROKERS = [
  { value: '', label: 'Default (Kafka)' },
  { value: 'kafka', label: 'Kafka' },
  { value: 'rabbit', label: 'RabbitMQ' },
]

export default function Logs() {
  const [formData, setFormData] = useState({
    level: 'info',
    message: '',
    broker: ''
  })
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)
  const [success, setSuccess] = useState(false)

  const handleSubmit = async (e) => {
    e.preventDefault()
    setLoading(true)
    setError(null)
    setSuccess(false)

    try {
      const body = {
        level: formData.level,
        message: formData.message,
      }

      // Only include broker if it's selected
      if (formData.broker) {
        body.broker = formData.broker
      }

      const response = await fetchWithRefresh('/api/v1/log', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
      })

      if (response.ok || response.status === 204) {
        setSuccess(true)
        setFormData({
          level: 'info',
          message: '',
          broker: ''
        })
        // Clear success message after 3 seconds
        setTimeout(() => setSuccess(false), 3000)
      } else {
        const errorText = await response.text()
        setError(errorText || 'Failed to send log')
      }
    } catch (err) {
      setError(err.message || 'An error occurred while sending the log')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <div>
          <h2 className={styles.title}>Logs</h2>
          <p className={styles.subtitle}>
            Send log messages to the system
          </p>
        </div>
      </div>

      {error && (
        <div className={styles.error}>
          {error}
          <button onClick={() => setError(null)}>×</button>
        </div>
      )}

      {success && (
        <div className={styles.success}>
          Log message sent successfully!
        </div>
      )}

      <div className={styles.formContainer}>
        <form onSubmit={handleSubmit} className={styles.form}>
          <div className={styles.formGroup}>
            <label htmlFor="level">Log Level</label>
            <select
              id="level"
              value={formData.level}
              onChange={(e) => setFormData({ ...formData, level: e.target.value })}
              required
              className={styles.select}
            >
              {LOG_LEVELS.map(level => (
                <option key={level.value} value={level.value}>
                  {level.label}
                </option>
              ))}
            </select>
          </div>

          <div className={styles.formGroup}>
            <label htmlFor="message">Message</label>
            <textarea
              id="message"
              value={formData.message}
              onChange={(e) => setFormData({ ...formData, message: e.target.value })}
              required
              className={styles.textarea}
              rows={6}
              placeholder="Enter your log message here..."
            />
          </div>

          <div className={styles.formGroup}>
            <label htmlFor="broker">Broker (Optional)</label>
            <select
              id="broker"
              value={formData.broker}
              onChange={(e) => setFormData({ ...formData, broker: e.target.value })}
              className={styles.select}
            >
              {BROKERS.map(broker => (
                <option key={broker.value} value={broker.value}>
                  {broker.label}
                </option>
              ))}
            </select>
            <p className={styles.helpText}>
              Select a message broker. If not specified, Kafka will be used by default.
            </p>
          </div>

          <div className={styles.formActions}>
            <button
              type="submit"
              disabled={loading}
              className={styles.submitButton}
            >
              {loading ? 'Sending...' : 'Send Log'}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}