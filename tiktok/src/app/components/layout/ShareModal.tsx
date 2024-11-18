'use client'

import { useEffect, useState } from 'react'
import { X } from 'lucide-react'
import { QRCodeSVG } from 'qrcode.react'
import { ShareOption, shareOptions } from '@/app/libs/share'

interface ShareModalProps {
  isOpen: boolean
  onClose: () => void
  videoUrl: string
  videoId: number
}

export default function ShareModal({ isOpen, onClose, videoUrl, videoId }: ShareModalProps) {
  const [showCopySuccess, setShowCopySuccess] = useState(false)
  const [showQRCode, setShowQRCode] = useState(false)
  
  if (!isOpen) return null

  const handleShare = async (option: ShareOption) => {
    const shareUrl = `${window.location.origin}/video/${videoId}`
    
    if (option.id === 'wechat') {
      setShowQRCode(true)
      return
    }
    
    await option.action(shareUrl)
    if (option.id === 'copy') {
      setShowCopySuccess(true)
      setTimeout(() => setShowCopySuccess(false), 2000)
    } else {
      onClose()
    }
  }

  return (
    <div 
      className="fixed inset-0 z-50 flex items-center justify-center bg-black/60"
      onClick={(e) => {
        if (e.target === e.currentTarget) {
          setShowQRCode(false)
          onClose()
        }
      }}
    >
      <div className="bg-gray-900 rounded-lg w-full max-w-sm p-4">
        {!showQRCode ? (
          <>
            <div className="flex justify-between items-center mb-4">
              <h3 className="text-lg font-semibold">分享到</h3>
              <button 
                onClick={() => onClose()}
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
                  `}
                >
                  <span className={`
                    w-12 h-12 flex items-center justify-center rounded-full
                    ${option.color}
                  `}>
                    {option.icon}
                  </span>
                  <span className="text-sm">{option.name}</span>
                </button>
              ))}
            </div>
          </>
        ) : (
          <div className="text-center">
            <div className="flex justify-between items-center mb-4">
              <h3 className="text-lg font-semibold">微信扫码分享</h3>
              <button 
                onClick={() => setShowQRCode(false)}
                className="p-1 hover:bg-gray-800 rounded-full transition-colors"
              >
                <X className="w-6 h-6" />
              </button>
            </div>
            <div className="bg-white p-4 rounded-lg inline-block">
              <QRCodeSVG 
                value={`${window.location.origin}/video/${videoId}`}
                size={200}
                level="H"
              />
            </div>
            <p className="mt-4 text-sm text-gray-400">
              请使用微信扫描二维码进行分享
            </p>
          </div>
        )}

        {showCopySuccess && (
          <div className="absolute bottom-4 left-1/2 transform -translate-x-1/2 bg-green-500 text-white px-4 py-2 rounded-full">
            链接已复制
          </div>
        )}
      </div>
    </div>
  )
}
