'use client'

import { forwardRef, useEffect, useImperativeHandle, useRef, useState } from 'react'
import { Heart, MessageCircle, Share2, Volume2, VolumeX } from 'lucide-react'
import ShareModal from './ShareModal'
import Link from 'next/link'
import LoginModal from '../auth/LoginModal'
import { API_ROUTES } from '@/app/config/api'
import CommentDrawer from '../comment/CommentDrawer'

interface VideoCardProps {
  video: {
    id: number
    author: string
    description: string
    videoUrl: string
    likes: number
    comments: number
    shares: number
    authorId: number
  }
  isActive: boolean
  isCommentOpen: boolean
  onCommentOpenChange: (open: boolean) => void
}

const VideoCard = forwardRef<HTMLVideoElement, VideoCardProps>(({ video, isActive, isCommentOpen, onCommentOpenChange }, ref) => {
  const videoRef = useRef<HTMLVideoElement>(null)
  const [isPlaying, setIsPlaying] = useState(false)
  const [isLiked, setIsLiked] = useState(false)
  const [volume, setVolume] = useState(0.5)
  const [isMuted, setIsMuted] = useState(true)
  const [isVolumeControlHovered, setIsVolumeControlHovered] = useState(false)
  const [showLikeAnimation, setShowLikeAnimation] = useState(false)
  const [likePosition, setLikePosition] = useState({ x: 0, y: 0 })
  const lastTapTime = useRef(0)
  const tapTimeout = useRef<NodeJS.Timeout>()
  const likeCount = useRef(video.likes)
  const [isShareModalOpen, setIsShareModalOpen] = useState(false)
  const shareCount = useRef(video.shares)
  const [isLoginModalOpen, setIsLoginModalOpen] = useState(false)

  // 暴露video元素给父组件
  useImperativeHandle(ref, () => videoRef.current as HTMLVideoElement)

  useEffect(() => {
    if (videoRef.current) {
      videoRef.current.volume = volume
      videoRef.current.muted = !isActive || isMuted
    }
  }, [volume, isActive, isMuted])

  const toggleMute = () => {
    setIsMuted(!isMuted)
  }

  const handleVolumeChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newVolume = parseFloat(e.target.value)
    setVolume(newVolume)
    setIsMuted(newVolume === 0)
  }

  useEffect(() => {
    if (isActive && videoRef.current) {
      videoRef.current.play().catch(console.error)
      setIsPlaying(true)
    } else if (!isActive && videoRef.current) {
      videoRef.current.pause()
      setIsPlaying(false)
    }
  }, [isActive])

  const togglePlay = () => {
    if (videoRef.current) {
      if (isPlaying) {
        videoRef.current.pause()
      } else {
        videoRef.current.play()
      }
      setIsPlaying(!isPlaying)
    }
  }

  const handleTap = (e: React.MouseEvent) => {
    const now = Date.now()
    const DOUBLE_TAP_DELAY = 300

    if (now - lastTapTime.current < DOUBLE_TAP_DELAY) {
      // 双击事件
      clearTimeout(tapTimeout.current)
      const rect = e.currentTarget.getBoundingClientRect()
      setLikePosition({
        x: e.clientX - rect.left,
        y: e.clientY - rect.top
      })
      
      if (!isLiked) {
        setIsLiked(true)
        likeCount.current += 1
      }
      
      setShowLikeAnimation(true)
      setTimeout(() => setShowLikeAnimation(false), 1000)
    } else {
      // 可能是单击，设置延时
      tapTimeout.current = setTimeout(() => {
        togglePlay()
      }, DOUBLE_TAP_DELAY)
    }
    
    lastTapTime.current = now
  }

  const handleShare = () => {
    setIsShareModalOpen(true)
  }

  const checkLogin = async () => {
    try {
      const res = await fetch(API_ROUTES.AUTH.CHECK)
      if (!res.ok) {
        setIsLoginModalOpen(true)
        return false
      }
      return true
    } catch (error) {
      console.error('检查登录状态失败:', error)
      return false
    }
  }

  return (
    <div className="relative w-full h-full">
      {/* 主视频区域容器 */}
      <div 
        className={`
          absolute inset-0 
          transition-all duration-300 ease-in-out
          ${isCommentOpen ? 'w-[calc(100%-384px)]' : 'w-full'}
        `}
      >
        {/* 视频元素 */}
        <video
          ref={videoRef}
          className="w-full h-full object-contain"
          src={video.videoUrl}
          loop
          playsInline
          muted={!isActive || isMuted}
          onClick={handleTap}
          preload="metadata"
        />

        {/* 用户信息 - 左下角 */}
        <div className="absolute bottom-8 left-4 z-10">
          <Link 
            href={`/user/${video.authorId}`}
            className="block text-lg font-semibold hover:text-primary transition-colors"
            onClick={(e) => e.stopPropagation()}
          >
            @{video.author}
          </Link>
          <p className="mt-2 text-sm text-gray-300 max-w-[80%]">
            {video.description}
          </p>
        </div>

        {/* 交互按钮 - 右侧 */}
        <div className="absolute right-4 top-1/2 -translate-y-1/2 z-10 flex flex-col items-center space-y-6">
          {/* 点赞按钮 */}
          <button 
            className="group flex flex-col items-center"
            onClick={async (e) => {
              e.stopPropagation()
              if (!(await checkLogin())) {
                return
              }
              setIsLiked(!isLiked)
              likeCount.current += isLiked ? -1 : 1
            }}
          >
            <div className="p-3 rounded-full bg-gray-800/60 group-hover:bg-gray-700/60">
              <Heart className={`w-6 h-6 ${isLiked ? 'fill-red-500 text-red-500' : ''}`} />
            </div>
            <span className="text-xs mt-1">{likeCount.current}</span>
          </button>

          {/* 评论按钮 */}
          <button 
            className="group flex flex-col items-center"
            onClick={(e) => {
              e.stopPropagation()
              onCommentOpenChange(true)
            }}
          >
            <div className="p-3 rounded-full bg-gray-800/60 group-hover:bg-gray-700/60">
              <MessageCircle className="w-6 h-6" />
            </div>
            <span className="text-xs mt-1">{video.comments}</span>
          </button>

          {/* 分享按钮 */}
          <button 
            className="group flex flex-col items-center"
            onClick={(e) => {
              e.stopPropagation()
              handleShare()
            }}
          >
            <div className="p-3 rounded-full bg-gray-800/60 group-hover:bg-gray-700/60">
              <Share2 className="w-6 h-6" />
            </div>
            <span className="text-xs mt-1">{shareCount.current}</span>
          </button>

          {/* 音量控制 */}
          <div 
            className="relative"
            onMouseEnter={() => setIsVolumeControlHovered(true)}
            onMouseLeave={() => {
              setTimeout(() => {
                setIsVolumeControlHovered(false)
              }, 300)
            }}
          >
            <button 
              className="flex flex-col items-center"
              onClick={(e) => {
                e.stopPropagation()
                toggleMute()
              }}
            >
              <div className="p-3 rounded-full bg-gray-800/60 hover:bg-gray-700/60">
                {isMuted ? <VolumeX className="w-6 h-6" /> : <Volume2 className="w-6 h-6" />}
              </div>
            </button>

            {/* 音量滑块 */}
            <div 
              className={`volume-slider absolute top-full left-1/2 -translate-x-1/2 mt-2 p-2 bg-gray-800/60 rounded-full transition-opacity duration-300 ease-in-out ${
                isVolumeControlHovered ? 'opacity-100' : 'opacity-0 pointer-events-none'
              }`}
            >
              <input
                type="range"
                min="0"
                max="1"
                step="0.1"
                value={volume}
                onChange={handleVolumeChange}
                onClick={(e) => e.stopPropagation()}
                className="h-24 w-2 appearance-none bg-white/20 rounded-full cursor-pointer"
                style={{
                  writingMode: 'vertical-lr',
                  WebkitAppearance: 'slider-vertical'
                }}
              />
            </div>
          </div>
        </div>
      </div>

      {/* 双击点赞动画 */}
      {showLikeAnimation && (
        <div 
          className="absolute z-20 pointer-events-none animate-scale-fade-out"
          style={{
            left: `${likePosition.x - 50}px`,
            top: `${likePosition.y - 50}px`
          }}
        >
          <Heart className="w-24 h-24 fill-red-500 text-red-500" />
        </div>
      )}

      {/* 分享态框 */}
      <ShareModal
        isOpen={isShareModalOpen}
        onClose={() => setIsShareModalOpen(false)}
        videoUrl={video.videoUrl}
        videoId={video.id}
        onShareSuccess={() => {
          shareCount.current += 1
        }}
      />

      <LoginModal 
        isOpen={isLoginModalOpen}
        onClose={() => setIsLoginModalOpen(false)}
        onSuccess={() => setIsLoginModalOpen(false)}
      />

      {/* 评论抽屉 - 固定在屏幕右侧 */}
      <CommentDrawer
        isOpen={isCommentOpen}
        onClose={() => onCommentOpenChange(false)}
        videoId={video.id}
      />
    </div>
  )
})

VideoCard.displayName = 'VideoCard'

export default VideoCard
