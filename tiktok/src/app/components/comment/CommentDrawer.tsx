import { X } from 'lucide-react'
import { useState } from 'react'

interface Comment {
  id: number
  author: string
  content: string
  timestamp: string
  likes: number
}

interface CommentDrawerProps {
  isOpen: boolean
  onClose: () => void
  videoId: number
}

export default function CommentDrawer({ isOpen, onClose, videoId }: CommentDrawerProps) {
  const [commentText, setCommentText] = useState('')
  // 模拟评论数据
  const comments: Comment[] = [
    {
      id: 1,
      author: "用户1",
      content: "真是太好看了！",
      timestamp: "2小时前",
      likes: 128
    },
    {
      id: 2,
      author: "用户2",
      content: "学到了学到了",
      timestamp: "1小时前",
      likes: 56
    }
  ]

  const handleSubmitComment = async () => {
    if (!commentText.trim()) return
    
    try {
      const res = await fetch('/api/auth/check')
      if (!res.ok) {
        // 如果未登录,打开登录模态框
        return
      }
      
      // TODO: 发送评论到后端
      console.log('发送评论:', commentText)
      setCommentText('')
    } catch (error) {
      console.error('发送评论失败:', error)
    }
  }

  if (!isOpen) return null

  return (
    <div 
      className={`
        fixed top-[63px] right-0 bottom-0 w-96
        bg-gray-900 border-l border-gray-800
        transform transition-transform duration-300 ease-in-out z-50
        ${isOpen ? 'translate-x-0' : 'translate-x-full'}
      `}
    >
      <div className="flex flex-col h-full">
        <div className="flex justify-between items-center p-4 border-b border-gray-800">
          <h3 className="text-lg font-semibold">评论 {comments.length}</h3>
          <button 
            onClick={onClose}
            className="p-2 hover:bg-gray-800 rounded-full"
          >
            <X className="w-6 h-6" />
          </button>
        </div>

        {/* 评论列表 */}
        <div className="flex-1 overflow-y-auto p-4">
          {comments.map(comment => (
            <div key={comment.id} className="mb-4 last:mb-0">
              <div className="flex items-start gap-3">
                <div className="flex-1">
                  <div className="flex items-center gap-2">
                    <span className="font-medium">@{comment.author}</span>
                    <span className="text-xs text-gray-400">{comment.timestamp}</span>
                  </div>
                  <p className="mt-1 text-sm">{comment.content}</p>
                </div>
              </div>
            </div>
          ))}
        </div>

        {/* 评论输入框 */}
        <div className="p-4 border-t border-gray-800">
          <div className="flex items-center gap-2">
            <input
              type="text"
              value={commentText}
              onChange={(e) => setCommentText(e.target.value)}
              placeholder="添加评论..."
              className="flex-1 bg-gray-800 rounded-full px-4 py-2 focus:outline-none focus:ring-2 focus:ring-primary"
            />
            <button 
              onClick={handleSubmitComment}
              disabled={!commentText.trim()}
              className="px-4 py-2 bg-primary rounded-full hover:bg-primary/90 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              发送
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}
