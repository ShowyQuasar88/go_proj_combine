import Navbar from "./components/layout/Navbar";
import Sidebar from "./components/layout/Sidebar";
import VideoFeed from "./components/layout/VideoFeed";

export default function Home() {
  return (
    <div className="flex flex-col h-screen">
      <nav className="sticky top-0 z-50">
        <Navbar />
      </nav>

      <div className="flex flex-1 h-[calc(100vh-60px)]">
        <aside className="w-[348px] border-r border-gray-800 hidden lg:block">
          <Sidebar />
        </aside>

        <main className="flex-1 relative">
          <VideoFeed />
        </main>
      </div>
    </div>
  )
} 