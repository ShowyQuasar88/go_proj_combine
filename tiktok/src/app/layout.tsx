import { Metadata } from 'next'
import Navbar from '@/app/components/layout/Navbar'
import './globals.css'

export const metadata: Metadata = {
  title: 'TikTok Clone',
  description: 'A TikTok Clone built with Next.js',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="zh">
      <body className="bg-black text-white">
        <Navbar />
        <main className="pt-16">
          {children}
        </main>
      </body>
    </html>
  )
} 