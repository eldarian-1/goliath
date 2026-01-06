import { useState, useEffect } from "react"
import { Link } from "react-router-dom"
import styles from "./VideoList.module.css"

export default function VideoList() {
  const [videos, setVideos] = useState([])
  const [loading, setLoading] = useState(true)
  const [searchTerm, setSearchTerm] = useState("")
  const [sortBy, setSortBy] = useState("date")

  // Mock data - replace with actual API call
  useEffect(() => {
    setTimeout(() => {
      setVideos([
        {
          id: "1",
          title: "Introduction to React Hooks",
          description: "Learn the basics of React Hooks including useState and useEffect",
          thumbnail: "https://via.placeholder.com/320x180/667eea/ffffff?text=React+Hooks",
          duration: "15:30",
          uploadDate: "2024-01-05",
          views: 1234,
          size: "45.2 MB"
        },
        {
          id: "2",
          title: "Advanced JavaScript Patterns",
          description: "Deep dive into advanced JavaScript design patterns and best practices",
          thumbnail: "https://via.placeholder.com/320x180/764ba2/ffffff?text=JavaScript",
          duration: "22:15",
          uploadDate: "2024-01-04",
          views: 856,
          size: "67.8 MB"
        },
        {
          id: "3",
          title: "CSS Grid Layout Tutorial",
          description: "Master CSS Grid with practical examples and real-world use cases",
          thumbnail: "https://via.placeholder.com/320x180/f093fb/ffffff?text=CSS+Grid",
          duration: "18:45",
          uploadDate: "2024-01-03",
          views: 2103,
          size: "52.1 MB"
        },
        {
          id: "4",
          title: "Node.js REST API Development",
          description: "Build a complete REST API using Node.js and Express",
          thumbnail: "https://via.placeholder.com/320x180/4facfe/ffffff?text=Node.js",
          duration: "28:20",
          uploadDate: "2024-01-02",
          views: 1567,
          size: "89.5 MB"
        },
        {
          id: "5",
          title: "Docker for Beginners",
          description: "Get started with Docker containerization and deployment",
          thumbnail: "https://via.placeholder.com/320x180/00f2fe/ffffff?text=Docker",
          duration: "20:10",
          uploadDate: "2024-01-01",
          views: 945,
          size: "61.3 MB"
        },
        {
          id: "6",
          title: "TypeScript Best Practices",
          description: "Learn TypeScript best practices and common patterns",
          thumbnail: "https://via.placeholder.com/320x180/43e97b/ffffff?text=TypeScript",
          duration: "25:40",
          uploadDate: "2023-12-31",
          views: 1789,
          size: "73.9 MB"
        }
      ])
      setLoading(false)
    }, 500)
  }, [])

  const filteredVideos = videos
    .filter(video =>
      video.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
      video.description.toLowerCase().includes(searchTerm.toLowerCase())
    )
    .sort((a, b) => {
      if (sortBy === "date") {
        return new Date(b.uploadDate) - new Date(a.uploadDate)
      } else if (sortBy === "views") {
        return b.views - a.views
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

  const formatViews = (views) => {
    if (views >= 1000) {
      return (views / 1000).toFixed(1) + "K"
    }
    return views.toString()
  }

  if (loading) {
    return (
      <div className={styles.loadingContainer}>
        <div className={styles.spinner}></div>
        <p>Loading videos...</p>
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
          <span className={styles.statItem}>
            👁️ {videos.reduce((sum, v) => sum + v.views, 0).toLocaleString()} total views
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
            <option value="views">Views</option>
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
                <img
                  src={video.thumbnail}
                  alt={video.title}
                  className={styles.thumbnail}
                />
                <div className={styles.duration}>{video.duration}</div>
              </div>
              
              <div className={styles.videoInfo}>
                <h3 className={styles.videoTitle}>{video.title}</h3>
                <p className={styles.videoDescription}>{video.description}</p>
                
                <div className={styles.videoMeta}>
                  <span className={styles.metaItem}>
                    👁️ {formatViews(video.views)} views
                  </span>
                  <span className={styles.metaItem}>
                    📅 {formatDate(video.uploadDate)}
                  </span>
                  <span className={styles.metaItem}>
                    💾 {video.size}
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