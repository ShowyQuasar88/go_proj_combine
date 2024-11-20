import Navbar from "./components/layout/Navbar";
import Sidebar from "./components/layout/Sidebar";
import VideoFeed from "./components/layout/VideoFeed";

export default function Home() {
  return (
    <div className="flex h-[calc(100vh-64px)]">
      <aside className="w-[320px] border-r border-gray-800 overflow-y-auto">
        <Sidebar />
      </aside>

      <main className="flex-1 relative">
        <VideoFeed />
      </main>
    </div>
  )
} 