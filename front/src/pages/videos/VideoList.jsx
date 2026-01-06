import { useState, useEffect } from "react"
import { Link } from "react-router-dom"
import { fetchWithRefresh } from "../../helpers/fetch"
import styles from "./VideoList.module.css"

export default function VideoList() {
  const [videos, setVideos] = useState([])
  const [loading, setLoading] = useState(true)
  const [searchTerm, setSearchTerm] = useState("")
  const [sortBy, setSortBy] = useState("date")
  const [error, setError] = useState(null)

  useEffect(() => {
    fetchVideos()
  }, [])

  const fetchVideos = async () => {
    try {
      setLoading(true)
      setError(null)
      
      const response = await fetchWithRefresh("/api/v1/videos?limit=100")
      
      if (!response.ok) {
        throw new Error("Failed to fetch videos")
      }
      
      const data = await response.json()
      setVideos(data.videos || [])
    } catch (err) {
      setError(err.message)
      console.error("Error fetching videos:", err)
    } finally {
      setLoading(false)
    }
  }

  const filteredVideos = videos
    .filter(video =>
      video.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
      (video.description && video.description.toLowerCase().includes(searchTerm.toLowerCase()))
    )
    .sort((a, b) => {
      if (sortBy === "date") {
        return new Date(b.createdAt) - new Date(a.createdAt)
      } else if (sortBy === "title") {
        return a.title.localeCompare(b.title)
      }
      return 0
    })

  const formatDate = (dateString) => {
    const date = new Date(dateString)
    return date.toLocaleDateString("en-US", { 
      year: "numeric", 
      month: "short", 
      day: "numeric" 
    })
  }

  const formatFileSize = (bytes) => {
    if (bytes === 0) return "0 Bytes"
    const k = 1024
    const sizes = ["Bytes", "KB", "MB", "GB"]
    const i = Math.floor(Math.log(bytes) / Math.log(k))
    return Math.round(bytes / Math.pow(k, i) * 100) / 100 + " " + sizes[i]
  }

  const formatDuration = (seconds) => {
    if (!seconds) return "N/A"
    const mins = Math.floor(seconds / 60)
    const secs = seconds % 60
    return `${mins}:${secs.toString().padStart(2, '0')}`
  }

  if (loading) {
    return (
      <div className={styles.loadingContainer}>
        <div className={styles.spinner}></div>
        <p>Loading videos...</p>
      </div>
    )
  }

  if (error) {
    return (
      <div className={styles.errorContainer}>
        <div className={styles.errorIcon}>⚠️</div>
        <h3>Error loading videos</h3>
        <p>{error}</p>
        <button onClick={fetchVideos} className={styles.retryButton}>
          Try Again
        </button>
      </div>
    )
  }

  return (
    <div className={styles.listContainer}>
      <div className={styles.listHeader}>
        <h2 className={styles.sectionTitle}>Video Library</h2>
        <div className={styles.stats}>
          <span className={styles.statItem}>
            📹 {videos.length} videos
          </span>
        </div>
      </div>

      <div className={styles.controls}>
        <div className={styles.searchBox}>
          <span className={styles.searchIcon}>🔍</span>
          <input
            type="text"
            placeholder="Search videos..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className={styles.searchInput}
          />
        </div>

        <div className={styles.sortBox}>
          <label htmlFor="sort" className={styles.sortLabel}>Sort by:</label>
          <select
            id="sort"
            value={sortBy}
            onChange={(e) => setSortBy(e.target.value)}
            className={styles.sortSelect}
          >
            <option value="date">Upload Date</option>
            <option value="title">Title</option>
          </select>
        </div>
      </div>

      {filteredVideos.length === 0 ? (
        <div className={styles.emptyState}>
          <div className={styles.emptyIcon}>🎬</div>
          <h3>No videos found</h3>
          <p>Try adjusting your search or upload a new video</p>
        </div>
      ) : (
        <div className={styles.videoGrid}>
          {filteredVideos.map(video => (
            <Link
              key={video.id}
              to={`/videos/view/${video.id}`}
              className={styles.videoCard}
            >
              <div className={styles.thumbnailContainer}>
                <div className={styles.thumbnailPlaceholder}>
                  🎬
                </div>
                {video.duration && (
                  <div className={styles.duration}>{formatDuration(video.duration)}</div>
                )}
              </div>
              
              <div className={styles.videoInfo}>
                <h3 className={styles.videoTitle}>{video.title}</h3>
                <p className={styles.videoDescription}>{video.description}</p>
                
                <div className={styles.videoMeta}>
                  <span className={styles.metaItem}>
                    📅 {formatDate(video.createdAt)}
                  </span>
                  <span className={styles.metaItem}>
                    💾 {formatFileSize(video.fileSize)}
                  </span>
                </div>
              </div>
            </Link>
          ))}
        </div>
      )}
    </div>
  )
}