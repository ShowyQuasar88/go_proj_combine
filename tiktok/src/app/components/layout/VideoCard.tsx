'use client'

import { forwardRef, useEffect, useImperativeHandle, useRef, useState } from 'react'
import { Heart, MessageCircle, Share2, Volume2, VolumeX } from 'lucide-react'
import ShareModal from './ShareModal'

interface VideoCardProps {
  video: {
    id: number
    author: string
    description: string
    videoUrl: string
    likes: number
    comments: number
    shares: number
  }
  isActive: boolean
}

const VideoCard = forwardRef<HTMLVideoElement, VideoCardProps>(({ video, isActive }, ref) => {
  const videoRef = useRef<HTMLVideoElement>(null)
  const [isPlaying, setIsPlaying] = useState(false)
  const [isLiked, setIsLiked] = useState(false)
  const [volume, setVolume] = useState(0.5)
  const [isMuted, setIsMuted] = useState(true)
  const [isVolumeSliderVisible, setIsVolumeSliderVisible] = useState(false)
  const [showLikeAnimation, setShowLikeAnimation] = useState(false)
  const [likePosition, setLikePosition] = useState({ x: 0, y: 0 })
  const lastTapTime = useRef(0)
  const likeCount = useRef(video.likes)
  const [isShareModalOpen, setIsShareModalOpen] = useState(false)

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

  // 处理双击事件
  const handleDoubleTap = (e: React.MouseEvent) => {
    const now = Date.now()
    const DOUBLE_TAP_DELAY = 300

    if (now - lastTapTime.current < DOUBLE_TAP_DELAY) {
      // 显示点赞动画
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
    }
    
    lastTapTime.current = now
  }

  return (
    <div 
      className="relative aspect-[9/16] max-w-[600px] mx-auto bg-gray-900 rounded-lg overflow-hidden group"
      onClick={handleDoubleTap}
      data-testid="video-container"
    >
      <video
        ref={videoRef}
        className="w-full h-full object-cover cursor-pointer"
        src={video.videoUrl}
        loop
        playsInline
        muted={!isActive}
        onClick={togglePlay}
        preload="metadata"
      />
      
      {/* 播放/暂停按钮 */}
      <div 
        className="absolute inset-0 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity"
        onClick={togglePlay}
      >
        <div className="p-4 rounded-full bg-black/50">
          {isPlaying ? (
            <span className="text-4xl">⏸️</span>
          ) : (
            <span className="text-4xl">▶️</span>
          )}
        </div>
      </div>

      {/* 视频信息 */}
      <div className="absolute bottom-0 left-0 right-0 p-4 bg-gradient-to-t from-black/60">
        <p className="font-bold text-lg">{video.author}</p>
        <p className="text-sm text-gray-200">{video.description}</p>
      </div>

      {/* 交互按钮 */}
      <div className="absolute right-4 bottom-20 flex flex-col items-center space-y-4">
        <button 
          className="p-3 rounded-full bg-gray-800/60 hover:bg-gray-700/60 transition-colors"
          onClick={(e) => {
            e.stopPropagation()
            setIsLiked(!isLiked)
            likeCount.current += isLiked ? -1 : 1
          }}
        >
          <Heart className={`w-6 h-6 ${isLiked ? 'fill-red-500 text-red-500' : ''}`} />
          <span className="text-xs">{likeCount.current}</span>
        </button>
        <button className="p-3 rounded-full bg-gray-800/60 hover:bg-gray-700/60 transition-colors">
          <MessageCircle className="w-6 h-6" />
          <span className="text-xs">{video.comments}</span>
        </button>
        <button 
          className="p-3 rounded-full bg-gray-800/60 hover:bg-gray-700/60 transition-colors"
          onClick={(e) => {
            e.stopPropagation()
            setIsShareModalOpen(true)
          }}
        >
          <Share2 className="w-6 h-6" />
          <span className="text-xs">{video.shares}</span>
        </button>
      </div>

      {/* 音量控制 */}
      <div 
        className="absolute right-4 bottom-4 flex flex-col items-center"
        data-testid="volume-control"
      >
        <button
          onClick={toggleMute}
          className="p-2 rounded-full bg-gray-800/60 hover:bg-gray-700/60 transition-colors"
        >
          {isMuted || volume === 0 ? (
            <VolumeX className="w-6 h-6" />
          ) : (
            <Volume2 className="w-6 h-6" />
          )}
        </button>
        
        {/* 音量滑块 */}
        <div 
          className={`absolute bottom-full mb-2 p-2 bg-gray-800/60 rounded-full transition-opacity duration-200 ${
            isVolumeSliderVisible ? 'opacity-100' : 'opacity-0 pointer-events-none'
          }`}
        >
          <input
            type="range"
            min="0"
            max="1"
            step="0.1"
            value={volume}
            onChange={handleVolumeChange}
            className="w-24 h-1 appearance-none bg-white/20 rounded-full cursor-pointer"
            style={{
              WebkitAppearance: 'none',
              transform: 'rotate(-90deg) translateX(-50%)',
              transformOrigin: 'left center'
            }}
          />
        </div>
      </div>

      {/* 点赞动画 */}
      {showLikeAnimation && (
        <div 
          className="absolute pointer-events-none"
          style={{
            left: likePosition.x - 50,
            top: likePosition.y - 50,
          }}
          data-testid="like-animation"
        >
          <Heart 
            className={`
              w-24 h-24 text-red-500 fill-red-500
              animate-like-burst
              transform -translate-x-1/2 -translate-y-1/2
            `}
          />
        </div>
      )}

      {/* 分享弹窗 */}
      <ShareModal 
        isOpen={isShareModalOpen}
        onClose={() => setIsShareModalOpen(false)}
        videoUrl={video.videoUrl}
      />
    </div>
  )
})

VideoCard.displayName = 'VideoCard'

export default VideoCard
