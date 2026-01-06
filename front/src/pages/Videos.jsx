import { useState } from "react"
import { Routes, Route, NavLink, Navigate } from "react-router-dom"
import VideoUpload from "./videos/VideoUpload"
import VideoList from "./videos/VideoList"
import VideoView from "./videos/VideoView"
import styles from "./Videos.module.css"

export default function Videos() {
  return (
    <div className={styles.videosContainer}>
      <div className={styles.videosHeader}>
        <h1 className={styles.title}>📹 Videos</h1>
        <p className={styles.subtitle}>Manage and view your video content</p>
      </div>

      <nav className={styles.tabsNav}>
        <NavLink
          to="/videos/list"
          className={({ isActive }) => `${styles.tab} ${isActive ? styles.activeTab : ''}`}
        >
          <span className={styles.tabIcon}>📋</span>
          <span>Video List</span>
        </NavLink>
        <NavLink
          to="/videos/upload"
          className={({ isActive }) => `${styles.tab} ${isActive ? styles.activeTab : ''}`}
        >
          <span className={styles.tabIcon}>⬆️</span>
          <span>Upload</span>
        </NavLink>
      </nav>

      <div className={styles.videosContent}>
        <Routes>
          <Route path="/" element={<Navigate to="/videos/list" replace />} />
          <Route path="/list" element={<VideoList />} />
          <Route path="/upload" element={<VideoUpload />} />
          <Route path="/view/:id" element={<VideoView />} />
        </Routes>
      </div>
    </div>
  )
}