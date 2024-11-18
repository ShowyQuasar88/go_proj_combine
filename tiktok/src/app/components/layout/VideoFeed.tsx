'use client'

import { useRef, useEffect, useState } from 'react'
import VideoCard from './VideoCard'

// 模拟视频数据生成器
const generateMockVideos = (start: number, count: number) => {
  return Array.from({ length: count }, (_, i) => ({
    id: start + i,
    authorId: start + i,
    author: `用户${start + i}`,
    description: `这是第${start + i}个视频 #抖音 #生活`,
    videoUrl: `https://example.com/video${start + i}.mp4`,
    likes: Math.floor(Math.random() * 10000),
    comments: Math.floor(Math.random() * 1000),
    shares: Math.floor(Math.random() * 500)
  }))
}

export default function VideoFeed() {
  const [videos, setVideos] = useState(generateMockVideos(1, 5))
  const [loading, setLoading] = useState(false)
  const [currentVideoIndex, setCurrentVideoIndex] = useState(0)
  const observerRef = useRef<IntersectionObserver>()
  const loadMoreRef = useRef<HTMLDivElement>(null)
  const videoRefs = useRef<(HTMLVideoElement | null)[]>([])

  // 设置无限滚动
  useEffect(() => {
    observerRef.current = new IntersectionObserver(
      (entries) => {
        if (entries[0].isIntersecting && !loading) {
          loadMoreVideos()
        }
      },
      { threshold: 0.1 }
    )

    if (loadMoreRef.current) {
      observerRef.current.observe(loadMoreRef.current)
    }

    return () => observerRef.current?.disconnect()
  }, [loading])

  // 设置视频播放观察器
  useEffect(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          const video = entry.target as HTMLVideoElement
          const index = videoRefs.current.indexOf(video)
          
          if (entry.isIntersecting) {
            setCurrentVideoIndex(index)
            video.play().catch(() => {
              // 处理自动播放失败的情况
              console.log('自动播放失败')
            })
          } else {
            video.pause()
          }
        })
      },
      { threshold: 0.6 } // 当视频60%可见时触发
    )

    videoRefs.current.forEach((video) => {
      if (video) observer.observe(video)
    })

    return () => observer.disconnect()
  }, [videos])

  useEffect(() => {
    const handleWheel = (e: WheelEvent) => {
      e.preventDefault()
      if (e.deltaY > 0 && currentVideoIndex < videos.length - 1) {
        setCurrentVideoIndex(prev => prev + 1)
      } else if (e.deltaY < 0 && currentVideoIndex > 0) {
        setCurrentVideoIndex(prev => prev - 1)
      }
    }

    const container = document.getElementById('video-feed-container')
    if (container) {
      container.addEventListener('wheel', handleWheel, { passive: false })
    }

    return () => {
      if (container) {
        container.removeEventListener('wheel', handleWheel)
      }
    }
  }, [currentVideoIndex, videos.length])

  const loadMoreVideos = async () => {
    try {
      setLoading(true)
      // 模拟加载延迟
      await new Promise(resolve => setTimeout(resolve, 1000))
      setVideos(prev => [...prev, ...generateMockVideos(prev.length + 1, 5)])
    } catch (error) {
      console.error('加载视频失败:', error)
      throw error // 这将触发error.tsx
    } finally {
      setLoading(false)
    }
  }

  return (
    <div 
      id="video-feed-container"
      className="h-full w-full relative overflow-hidden"
    >
      <div 
        className="h-full transition-transform duration-300"
        style={{
          transform: `translateY(-${currentVideoIndex * 100}%)`
        }}
      >
        {videos.map((video, index) => (
          <div key={video.id} className="h-full flex items-center justify-center relative">
            <div className="absolute right-4 h-full flex items-center">
              <div className="flex flex-col items-center space-y-6">
                {/* 互动按钮将在这里渲染 */}
              </div>
            </div>
            <VideoCard 
              video={video}
              ref={(el) => { videoRefs.current[index] = el }}
              isActive={currentVideoIndex === index}
            />
          </div>
        ))}
      </div>
    </div>
  )
}
