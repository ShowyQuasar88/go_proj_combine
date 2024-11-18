'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'

export default function Navbar() {
  const pathname = usePathname()
  
  return (
    <nav className="fixed top-0 z-50 w-full bg-black/80 backdrop-blur-sm border-b border-gray-800">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex">
            <Link href="/" className="flex items-center">
              <span className="text-xl font-bold">TikTok Clone</span>
            </Link>
          </div>
          
          <div className="flex items-center space-x-4">
            <Link 
              href="/following"
              className={`px-3 py-2 rounded-md text-sm font-medium ${
                pathname === '/following' 
                  ? 'text-white bg-gray-900' 
                  : 'text-gray-300 hover:text-white'
              }`}
            >
              Following
            </Link>
            <Link 
              href="/profile"
              className={`px-3 py-2 rounded-md text-sm font-medium ${
                pathname === '/profile' 
                  ? 'text-white bg-gray-900' 
                  : 'text-gray-300 hover:text-white'
              }`}
            >
              Profile
            </Link>
          </div>
        </div>
      </div>
    </nav>
  )
} 