import { useState } from "react"
import styles from "./VideoUpload.module.css"

export default function VideoUpload() {
  const [selectedFile, setSelectedFile] = useState(null)
  const [title, setTitle] = useState("")
  const [description, setDescription] = useState("")
  const [uploading, setUploading] = useState(false)
  const [uploadProgress, setUploadProgress] = useState(0)
  const [dragActive, setDragActive] = useState(false)

  const handleDrag = (e) => {
    e.preventDefault()
    e.stopPropagation()
    if (e.type === "dragenter" || e.type === "dragover") {
      setDragActive(true)
    } else if (e.type === "dragleave") {
      setDragActive(false)
    }
  }

  const handleDrop = (e) => {
    e.preventDefault()
    e.stopPropagation()
    setDragActive(false)
    
    if (e.dataTransfer.files && e.dataTransfer.files[0]) {
      const file = e.dataTransfer.files[0]
      if (file.type.startsWith("video/")) {
        setSelectedFile(file)
      } else {
        alert("Please select a valid video file")
      }
    }
  }

  const handleFileChange = (e) => {
    if (e.target.files && e.target.files[0]) {
      const file = e.target.files[0]
      if (file.type.startsWith("video/")) {
        setSelectedFile(file)
      } else {
        alert("Please select a valid video file")
      }
    }
  }

  const handleUpload = async (e) => {
    e.preventDefault()
    
    if (!selectedFile || !title) {
      alert("Please select a file and enter a title")
      return
    }

    setUploading(true)
    
    // Simulate upload progress
    const interval = setInterval(() => {
      setUploadProgress(prev => {
        if (prev >= 100) {
          clearInterval(interval)
          return 100
        }
        return prev + 10
      })
    }, 300)

    // Simulate API call
    setTimeout(() => {
      clearInterval(interval)
      setUploadProgress(100)
      setTimeout(() => {
        alert("Video uploaded successfully!")
        setSelectedFile(null)
        setTitle("")
        setDescription("")
        setUploading(false)
        setUploadProgress(0)
      }, 500)
    }, 3000)
  }

  const formatFileSize = (bytes) => {
    if (bytes === 0) return "0 Bytes"
    const k = 1024
    const sizes = ["Bytes", "KB", "MB", "GB"]
    const i = Math.floor(Math.log(bytes) / Math.log(k))
    return Math.round(bytes / Math.pow(k, i) * 100) / 100 + " " + sizes[i]
  }

  return (
    <div className={styles.uploadContainer}>
      <h2 className={styles.sectionTitle}>Upload Video</h2>
      
      <form onSubmit={handleUpload} className={styles.uploadForm}>
        <div
          className={`${styles.dropZone} ${dragActive ? styles.dragActive : ""} ${selectedFile ? styles.hasFile : ""}`}
          onDragEnter={handleDrag}
          onDragLeave={handleDrag}
          onDragOver={handleDrag}
          onDrop={handleDrop}
        >
          {selectedFile ? (
            <div className={styles.fileInfo}>
              <div className={styles.fileIcon}>🎬</div>
              <div className={styles.fileName}>{selectedFile.name}</div>
              <div className={styles.fileSize}>{formatFileSize(selectedFile.size)}</div>
              <button
                type="button"
                onClick={() => setSelectedFile(null)}
                className={styles.removeButton}
              >
                Remove
              </button>
            </div>
          ) : (
            <>
              <div className={styles.uploadIcon}>📤</div>
              <p className={styles.dropText}>Drag and drop your video here</p>
              <p className={styles.dropSubtext}>or</p>
              <label className={styles.browseButton}>
                Browse Files
                <input
                  type="file"
                  accept="video/*"
                  onChange={handleFileChange}
                  className={styles.fileInput}
                />
              </label>
              <p className={styles.supportedFormats}>
                Supported formats: MP4, AVI, MOV, MKV
              </p>
            </>
          )}
        </div>

        <div className={styles.formGroup}>
          <label htmlFor="title" className={styles.label}>
            Title <span className={styles.required}>*</span>
          </label>
          <input
            type="text"
            id="title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className={styles.input}
            placeholder="Enter video title"
            required
          />
        </div>

        <div className={styles.formGroup}>
          <label htmlFor="description" className={styles.label}>
            Description
          </label>
          <textarea
            id="description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className={styles.textarea}
            placeholder="Enter video description (optional)"
            rows="4"
          />
        </div>

        {uploading && (
          <div className={styles.progressContainer}>
            <div className={styles.progressBar}>
              <div
                className={styles.progressFill}
                style={{ width: `${uploadProgress}%` }}
              />
            </div>
            <div className={styles.progressText}>{uploadProgress}%</div>
          </div>
        )}

        <button
          type="submit"
          className={styles.uploadButton}
          disabled={!selectedFile || !title || uploading}
        >
          {uploading ? "Uploading..." : "Upload Video"}
        </button>
      </form>
    </div>
  )
}