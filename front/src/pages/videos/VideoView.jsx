import { useState, useEffect } from "react"
import { useParams, Link, useNavigate } from "react-router-dom"
import styles from "./VideoView.module.css"

export default function VideoView() {
  const { id } = useParams()
  const navigate = useNavigate()
  const [video, setVideo] = useState(null)
  const [loading, setLoading] = useState(true)
  const [isPlaying, setIsPlaying] = useState(false)

  // Mock data - replace with actual API call
  useEffect(() => {
    setTimeout(() => {
      const mockVideos = {
        "1": {
          id: "1",
          title: "Introduction to React Hooks",
          description: "Learn the basics of React Hooks including useState and useEffect. This comprehensive tutorial covers everything you need to know to get started with React Hooks and build modern React applications.",
          videoUrl: "https://www.w3schools.com/html/mov_bbb.mp4",
          thumbnail: "https://via.placeholder.com/1280x720/667eea/ffffff?text=React+Hooks",
          duration: "15:30",
          uploadDate: "2024-01-05",
          views: 1234,
          size: "45.2 MB",
          uploader: "John Doe",
          tags: ["React", "JavaScript", "Hooks", "Tutorial"]
        },
        "2": {
          id: "2",
          title: "Advanced JavaScript Patterns",
          description: "Deep dive into advanced JavaScript design patterns and best practices",
          videoUrl: "https://www.w3schools.com/html/mov_bbb.mp4",
          thumbnail: "https://via.placeholder.com/1280x720/764ba2/ffffff?text=JavaScript",
          duration: "22:15",
          uploadDate: "2024-01-04",
          views: 856,
          size: "67.8 MB",
          uploader: "Jane Smith",
          tags: ["JavaScript", "Patterns", "Advanced"]
        }
      }

      setVideo(mockVideos[id] || null)
      setLoading(false)
    }, 500)
  }, [id])

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
        <video
          controls
          poster={video.thumbnail}
          className={styles.video}
          onPlay={() => setIsPlaying(true)}
          onPause={() => setIsPlaying(false)}
        >
          <source src={video.videoUrl} type="video/mp4" />
          Your browser does not support the video tag.
        </video>
      </div>

      <div className={styles.videoDetails}>
        <div className={styles.videoHeader}>
          <div>
            <h1 className={styles.videoTitle}>{video.title}</h1>
            <div className={styles.videoStats}>
              <span className={styles.statItem}>
                👁️ {video.views.toLocaleString()} views
              </span>
              <span className={styles.statItem}>
                📅 {formatDate(video.uploadDate)}
              </span>
              <span className={styles.statItem}>
                ⏱️ {video.duration}
              </span>
              <span className={styles.statItem}>
                💾 {video.size}
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

        <div className={styles.videoMeta}>
          <div className={styles.uploaderInfo}>
            <div className={styles.uploaderAvatar}>
              {video.uploader.charAt(0)}
            </div>
            <div>
              <div className={styles.uploaderName}>{video.uploader}</div>
              <div className={styles.uploaderLabel}>Uploader</div>
            </div>
          </div>
        </div>

        <div className={styles.descriptionSection}>
          <h3 className={styles.sectionTitle}>Description</h3>
          <p className={styles.description}>{video.description}</p>
        </div>

        {video.tags && video.tags.length > 0 && (
          <div className={styles.tagsSection}>
            <h3 className={styles.sectionTitle}>Tags</h3>
            <div className={styles.tags}>
              {video.tags.map((tag, index) => (
                <span key={index} className={styles.tag}>
                  {tag}
                </span>
              ))}
            </div>
          </div>
        )}

        <div className={styles.actionsSection}>
          <h3 className={styles.sectionTitle}>Actions</h3>
          <div className={styles.actionsList}>
            <button className={styles.actionItem}>
              📥 Download
            </button>
            <button className={styles.actionItem}>
              🔗 Share
            </button>
            <button className={styles.actionItem}>
              ⭐ Add to Favorites
            </button>
            <button className={styles.actionItem}>
              📊 View Analytics
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}