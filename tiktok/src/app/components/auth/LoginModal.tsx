import { X } from 'lucide-react'
import LoginForm from './LoginForm'

interface LoginModalProps {
  isOpen: boolean
  onClose: () => void
  onSuccess: () => void
}

export default function LoginModal({ isOpen, onClose, onSuccess }: LoginModalProps) {
  if (!isOpen) return null

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/80 backdrop-blur-sm">
      <div className="max-w-md w-full space-y-8 p-8 bg-gray-900 rounded-lg relative">
        <button
          onClick={onClose}
          className="absolute right-4 top-4 p-2 hover:bg-gray-800 rounded-full transition-colors"
        >
          <X className="w-6 h-6" />
        </button>

        <div className="text-center">
          <h2 className="text-3xl font-bold">登录 TikTok</h2>
          <p className="mt-2 text-gray-400">继续观看精彩视频</p>
        </div>
        <LoginForm onSuccess={onSuccess} />
      </div>
    </div>
  )
}
