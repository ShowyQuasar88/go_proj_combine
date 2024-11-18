'use client'

import { useEffect, useState } from 'react'
import { X } from 'lucide-react'
import { ShareOption } from '@/app/libs/share'
import { copyToClipboard, openShareWindow } from '@/app/libs/share'

interface ShareModalProps {
  isOpen: boolean
  onClose: () => void
  videoUrl: string
  videoId: number
}

const shareOptions: ShareOption[] = [
  {
    id: 'wechat',
    name: '微信',
    icon: '💬',
    color: 'bg-green-500',
    action: (url) => {
      // 这里应该生成二维码
      console.log('分享到微信:', url)
    }
  },
  {
    id: 'weibo',
    name: '微博',
    icon: '🔄',
    color: 'bg-red-500',
    action: (url) => {
      const shareUrl = `http://service.weibo.com/share/share.php?url=${encodeURIComponent(url)}`
      openShareWindow(shareUrl, '分享到微博')
    }
  },
  {
    id: 'qq',
    name: 'QQ',
    icon: '💭',
    color: 'bg-blue-500',
    action: (url) => {
      const shareUrl = `http://connect.qq.com/widget/shareqq/index.html?url=${encodeURIComponent(url)}`
      openShareWindow(shareUrl, '分享到QQ')
    }
  },
  {
    id: 'copy',
    name: '复制链接',
    icon: '📋',
    color: 'bg-gray-500',
    action: async (url) => {
      const success = await copyToClipboard(url)
      if (success) {
        // 这里可以添加一个提示
        console.log('链接已复制')
      }
    }
  }
]

export default function ShareModal({ isOpen, onClose, videoUrl, videoId }: ShareModalProps) {
  const [showCopySuccess, setShowCopySuccess] = useState(false)

  useEffect(() => {
    if (showCopySuccess) {
      const timer = setTimeout(() => setShowCopySuccess(false), 2000)
      return () => clearTimeout(timer)
    }
  }, [showCopySuccess])

  if (!isOpen) return null

  const handleShare = async (option: ShareOption) => {
    const shareUrl = `${window.location.origin}/video/${videoId}`
    await option.action(shareUrl)
    
    if (option.id === 'copy') {
      setShowCopySuccess(true)
    } else {
      onClose()
    }
  }

  return (
    <div 
      className="fixed inset-0 z-50 flex items-center justify-center bg-black/60 animate-fade-in"
      onClick={(e) => {
        if (e.target === e.currentTarget) onClose()
      }}
    >
      <div className="bg-gray-900 rounded-lg w-full max-w-sm p-4 animate-slide-up">
        <div className="flex justify-between items-center mb-4">
          <h3 className="text-lg font-semibold">分享到</h3>
          <button 
            onClick={onClose}
            className="p-1 hover:bg-gray-800 rounded-full transition-colors"
          >
            <X className="w-6 h-6" />
          </button>
        </div>
        
        <div className="grid grid-cols-4 gap-4">
          {shareOptions.map((option) => (
            <button
              key={option.id}
              onClick={() => handleShare(option)}
              className={`
                flex flex-col items-center gap-2 p-2 rounded-lg
                transition-transform hover:scale-110
                ${option.id === 'copy' && showCopySuccess ? 'animate-bounce' : ''}
              `}
            >
              <span className={`
                w-12 h-12 flex items-center justify-center rounded-full
                ${option.color} text-2xl
              `}>
                {option.icon}
              </span>
              <span className="text-sm">{option.name}</span>
            </button>
          ))}
        </div>

        {showCopySuccess && (
          <div className="absolute bottom-4 left-1/2 transform -translate-x-1/2 bg-green-500 text-white px-4 py-2 rounded-full animate-fade-in">
            链接已复制
          </div>
        )}
      </div>
    </div>
  )
}
