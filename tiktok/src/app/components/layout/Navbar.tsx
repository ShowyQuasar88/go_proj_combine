'use client'

import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { useState, useEffect } from 'react'
import LoginModal from '../auth/LoginModal'
import { API_ROUTES } from '@/app/config/api'

export default function Navbar() {
  const router = useRouter()
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [username, setUsername] = useState('')
  const [isLoginModalOpen, setIsLoginModalOpen] = useState(false)

  useEffect(() => {
    checkLoginStatus()
  }, [])

  async function checkLoginStatus() {
    try {
      const res = await fetch(API_ROUTES.AUTH.CHECK)
      if (res.ok) {
        const data = await res.json()
        setIsLoggedIn(true)
        setUsername(data.username)
      }
    } catch (error) {
      console.error('检查登录状态失败:', error)
    }
  }

  async function handleLogout() {
    try {
      await fetch(API_ROUTES.AUTH.LOGOUT, { method: 'POST' })
      setIsLoggedIn(false)
      setUsername('')
      router.push('/')
    } catch (error) {
      console.error('登出失败:', error)
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
            
            <div className="flex items-center space-x-4">
              {isLoggedIn ? (
                <>
                  <span className="text-gray-300">@{username}</span>
                  <button
                    onClick={handleLogout}
                    className="px-3 py-2 rounded-md text-sm font-medium text-gray-300 hover:text-white hover:bg-gray-900"
                  >
                    退出
                  </button>
                </>
              ) : (
                <button
                  onClick={() => setIsLoginModalOpen(true)}
                  className="px-3 py-2 rounded-md text-sm font-medium text-white bg-primary hover:bg-primary/90"
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
          checkLoginStatus()
        }}
      />
    </>
  )
} 