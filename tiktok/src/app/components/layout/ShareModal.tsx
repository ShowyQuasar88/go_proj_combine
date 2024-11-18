'use client'

import { X } from 'lucide-react'

interface ShareModalProps {
  isOpen: boolean
  onClose: () => void
  videoUrl: string
}

const shareOptions = [
  { id: 'wechat', name: '微信', icon: '💬' },
  { id: 'weibo', name: '微博', icon: '🔄' },
  { id: 'qq', name: 'QQ', icon: '💭' },
  { id: 'copy', name: '复制链接', icon: '📋' }
]

export default function ShareModal({ isOpen, onClose, videoUrl }: ShareModalProps) {
  if (!isOpen) return null

  const handleShare = async (platform: string) => {
    switch (platform) {
      case 'copy':
        try {
          await navigator.clipboard.writeText(videoUrl)
          alert('链接已复制')
        } catch (err) {
          console.error('复制失败:', err)
        }
        break
      default:
        // 这里可以添加其他平台的分享逻辑
        console.log(`分享到 ${platform}:`, videoUrl)
    }
    onClose()
  }

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/60">
      <div className="bg-gray-900 rounded-lg w-full max-w-sm p-4">
        <div className="flex justify-between items-center mb-4">
          <h3 className="text-lg font-semibold">分享到</h3>
          <button onClick={onClose} className="p-1">
            <X className="w-6 h-6" />
          </button>
        </div>
        
        <div className="grid grid-cols-4 gap-4">
          {shareOptions.map((option) => (
            <button
              key={option.id}
              onClick={() => handleShare(option.id)}
              className="flex flex-col items-center gap-2 p-2 hover:bg-gray-800 rounded-lg transition-colors"
            >
              <span className="text-2xl">{option.icon}</span>
              <span className="text-sm">{option.name}</span>
            </button>
          ))}
        </div>
      </div>
    </div>
  )
}
