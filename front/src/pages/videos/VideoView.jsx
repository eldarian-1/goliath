import { useState, useEffect, useRef } from "react"
import { useParams, Link, useNavigate } from "react-router-dom"
import { fetchWithRefresh } from "../../helpers/fetch"
import styles from "./VideoView.module.css"

export default function VideoView() {
  const { id } = useParams()
  const navigate = useNavigate()
  const [video, setVideo] = useState(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)
  const [isPlaying, setIsPlaying] = useState(false)
  const videoRef = useRef(null)

  useEffect(() => {
    fetchVideoMetadata()
  }, [id])

  useEffect(() => {
    // Poll for progress updates if video is not fully processed
    if (!video || video.progress >= 100) {
      return
    }
    
    const intervalId = setInterval(() => {
      fetchVideoMetadata()
    }, 2000) // Poll every 2 seconds
    
    return () => {
      clearInterval(intervalId)
    }
  }, [video?.progress])

  const fetchVideoMetadata = async () => {
    try {
      setLoading(true)
      setError(null)
      
      const response = await fetchWithRefresh(`/api/v1/videos?limit=100`)
      
      if (!response.ok) {
        throw new Error("Failed to fetch video metadata")
      }
      
      const data = await response.json()
      const foundVideo = data.videos.find(v => v.id.toString() === id)
      
      if (!foundVideo) {
        setVideo(null)
      } else {
        setVideo(foundVideo)
      }
    } catch (err) {
      setError(err.message)
      console.error("Error fetching video:", err)
    } finally {
      setLoading(false)
    }
  }

  const handleDelete = () => {
    if (window.confirm("Are you sure you want to delete this video?")) {
      alert("Video deleted successfully!")
      navigate("/videos/list")
    }
  }

  const formatDate = (dateString) => {
    const date = new Date(dateString)
    return date.toLocaleDateString("en-US", {
      year: "numeric",
      month: "long",
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

  const getVideoUrl = () => {
    return `http://localhost:8080/api/v1/videos/${id}`
  }

  if (loading) {
    return (
      <div className={styles.loadingContainer}>
        <div className={styles.spinner}></div>
        <p>Loading video...</p>
      </div>
    )
  }

  if (!video) {
    return (
      <div className={styles.errorContainer}>
        <div className={styles.errorIcon}>❌</div>
        <h2>Video Not Found</h2>
        <p>The video you're looking for doesn't exist or has been removed.</p>
        <Link to="/videos/list" className={styles.backButton}>
          ← Back to Video List
        </Link>
      </div>
    )
  }

  return (
    <div className={styles.viewContainer}>
      <div className={styles.breadcrumb}>
        <Link to="/videos/list" className={styles.breadcrumbLink}>
          Videos
        </Link>
        <span className={styles.breadcrumbSeparator}>/</span>
        <span className={styles.breadcrumbCurrent}>{video.title}</span>
      </div>

      <div className={styles.videoPlayer}>
        {video.progress >= 100 ? (
          <video
            ref={videoRef}
            controls
            className={styles.video}
            onPlay={() => setIsPlaying(true)}
            onPause={() => setIsPlaying(false)}
          >
            <source src={getVideoUrl()} type={video.contentType} />
            Your browser does not support the video tag.
          </video>
        ) : (
          <div className={styles.processingMessage}>
            <div className={styles.processingIcon}>⏳</div>
            <h3>Video is being processed</h3>
            <p>Please wait while we convert your video. It will be available for playback shortly.</p>
            <div className={styles.progressContainer}>
              <div className={styles.progressBar}>
                <div
                  className={styles.progressFill}
                  style={{ width: `${video.progress}%` }}
                />
              </div>
              <div className={styles.progressText}>{video.progress}%</div>
            </div>
          </div>
        )}
      </div>

      <div className={styles.videoDetails}>
        <div className={styles.videoHeader}>
          <div>
            <h1 className={styles.videoTitle}>{video.title}</h1>
            <div className={styles.videoStats}>
              <span className={styles.statItem}>
                📅 {formatDate(video.createdAt)}
              </span>
              {video.duration && (
                <span className={styles.statItem}>
                  ⏱️ {formatDuration(video.duration)}
                </span>
              )}
              <span className={styles.statItem}>
                💾 {formatFileSize(video.fileSize)}
              </span>
            </div>
          </div>

          <div className={styles.actionButtons}>
            <button className={styles.editButton}>
              ✏️ Edit
            </button>
            <button className={styles.deleteButton} onClick={handleDelete}>
              🗑️ Delete
            </button>
          </div>
        </div>

        {video.description && (
          <div className={styles.descriptionSection}>
            <h3 className={styles.sectionTitle}>Description</h3>
            <p className={styles.description}>{video.description}</p>
          </div>
        )}

        <div className={styles.technicalInfo}>
          <h3 className={styles.sectionTitle}>Technical Information</h3>
          <div className={styles.infoGrid}>
            <div className={styles.infoItem}>
              <span className={styles.infoLabel}>File Name:</span>
              <span className={styles.infoValue}>{video.fileName}</span>
            </div>
            <div className={styles.infoItem}>
              <span className={styles.infoLabel}>Content Type:</span>
              <span className={styles.infoValue}>{video.contentType}</span>
            </div>
            <div className={styles.infoItem}>
              <span className={styles.infoLabel}>File Size:</span>
              <span className={styles.infoValue}>{formatFileSize(video.fileSize)}</span>
            </div>
            <div className={styles.infoItem}>
              <span className={styles.infoLabel}>Video ID:</span>
              <span className={styles.infoValue}>{video.id}</span>
            </div>
          </div>
        </div>

        <div className={styles.actionsSection}>
          <h3 className={styles.sectionTitle}>Actions</h3>
          <div className={styles.actionsList}>
            <a
              href={getVideoUrl()}
              download={video.fileName}
              className={styles.actionItem}
            >
              📥 Download
            </a>
            <button
              className={styles.actionItem}
              onClick={() => {
                navigator.clipboard.writeText(window.location.href)
                alert("Link copied to clipboard!")
              }}
            >
              🔗 Copy Link
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}