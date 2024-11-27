'use client'

import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { useState, useEffect } from 'react'
import LoginModal from '../auth/LoginModal'
import { API_ROUTES } from '@/app/config/api'
import UserDropdown from './UserDropdown'
import { checkLoginStatus, getUsername } from '@/app/utils/auth'

export default function Navbar() {
  const router = useRouter()
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [username, setUsername] = useState('')
  const [isLoginModalOpen, setIsLoginModalOpen] = useState(false)

  useEffect(() => {
    const loggedIn = checkLoginStatus()
    setIsLoggedIn(loggedIn)
    if (loggedIn) {
      const username = getUsername()
      if (username) {
        setUsername(username)
      }
    }
  }, [])

  async function handleLogout() {
    try {
      // 1. 发送注销请求
      await fetch(API_ROUTES.AUTH.LOGOUT, { 
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include'  // 确保发送 cookie
      })
      
      // 2. 清除前端状态
      setIsLoggedIn(false)
      setUsername('')
      
      // 3. 重定向到首页
      router.push('/')
      
      // 4. 可选：清除其他前端状态（如果有的话）
      // 例如：清除本地存储的用户偏好设置等
      localStorage.removeItem('userPreferences')
      
    } catch (error) {
      console.error('登出失败:', error)
      // 5. 错误处理：显示错误提示
      alert('退出登录失败，请重试')
    }
  }

  return (
    <>
      <nav className="fixed top-0 z-50 w-full h-16 bg-black/80 backdrop-blur-sm border-b border-gray-800">
        <div className="h-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-full items-center">
            <div className="flex">
              <Link href="/" className="flex items-center">
                <span className="text-xl font-bold">TikTok Clone</span>
              </Link>
            </div>
            
            <div className="flex items-center">
              {isLoggedIn ? (
                <UserDropdown 
                  username={username} 
                  onLogout={handleLogout}
                />
              ) : (
                <button
                  onClick={() => setIsLoginModalOpen(true)}
                  className="px-4 py-2 rounded-md text-sm font-medium text-white bg-primary hover:bg-primary/90"
                >
                  登录
                </button>
              )}
            </div>
          </div>
        </div>
      </nav>

      <LoginModal 
        isOpen={isLoginModalOpen}
        onClose={() => setIsLoginModalOpen(false)}
        onSuccess={() => {
          setIsLoginModalOpen(false)
          const loggedIn = checkLoginStatus()
          setIsLoggedIn(loggedIn)
          if (loggedIn) {
            const username = getUsername()
            if (username) {
              setUsername(username)
            }
          }
        }}
      />
    </>
  )
} 