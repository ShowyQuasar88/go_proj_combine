import Navbar from "./components/layout/Navbar";
import Sidebar from "./components/layout/Sidebar";
import VideoFeed from "./components/layout/VideoFeed";

export default function Home() {
  return (
    <div className="flex flex-col min-h-screen">
      {/* 顶部导航栏 */}
      <nav className="sticky top-0 z-50">
        <Navbar />
      </nav>

      <div className="flex flex-1">
        {/* 左侧边栏 */}
        <aside className="w-[348px] border-r border-gray-800 hidden lg:block p-4">
          <Sidebar />
        </aside>

        {/* 主内容区 - 视频Feed */}
        <main className="flex-1 h-[calc(100vh-60px)] overflow-y-auto">
          <div className="max-w-[600px] mx-auto">
            <VideoFeed />
          </div>
        </main>
      </div>
    </div>
  )
} 