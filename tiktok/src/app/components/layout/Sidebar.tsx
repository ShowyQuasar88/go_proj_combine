'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'

export default function Sidebar() {
  const pathname = usePathname()
  
  return (
    <div className="flex flex-col h-full">
      {/* 主导航 */}
      <div className="flex flex-col space-y-2 mb-8">
        <Link 
          href="/"
          className={`flex items-center space-x-2 px-4 py-3 rounded-lg hover:bg-gray-900 ${
            pathname === '/' ? 'bg-gray-900' : ''
          }`}
        >
          <span className="text-xl">🏠</span>
          <span>推荐</span>
        </Link>
        <Link 
          href="/following"
          className={`flex items-center space-x-2 px-4 py-3 rounded-lg hover:bg-gray-900 ${
            pathname === '/following' ? 'bg-gray-900' : ''
          }`}
        >
          <span className="text-xl">👥</span>
          <span>关注</span>
        </Link>
      </div>

      {/* 建议关注的账号 */}
      <div className="border-t border-gray-800 pt-4">
        <h3 className="text-sm text-gray-400 px-4 mb-4">推荐关注</h3>
        {/* 这里后续添加推荐用户列表 */}
      </div>
    </div>
  )
}
