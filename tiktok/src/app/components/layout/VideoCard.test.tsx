import { expect, test, vi } from 'vitest'
import { render, screen, fireEvent } from '@testing-library/react'
import VideoCard from './VideoCard'

const mockVideo = {
  id: 1,
  author: '测试用户',
  description: '测试视频 #测试',
  videoUrl: 'https://example.com/test.mp4',
  likes: 100,
  comments: 50,
  shares: 20
}

test('VideoCard 基础渲染', () => {
  render(<VideoCard video={mockVideo} isActive={false} />)
  expect(screen.getByText('测试用户')).toBeDefined()
  expect(screen.getByText('测试视频 #测试')).toBeDefined()
})

test('双击点赞功能', async () => {
  const { container } = render(<VideoCard video={mockVideo} isActive={false} />)
  const videoContainer = container.firstChild as HTMLElement
  
  // 模拟双击
  fireEvent.click(videoContainer)
  await vi.advanceTimersByTimeAsync(100)
  fireEvent.click(videoContainer)
  
  // 检查点赞动画是否出现
  expect(screen.getByTestId('like-animation')).toBeDefined()
})

test('音量控制', () => {
  render(<VideoCard video={mockVideo} isActive={true} />)
  const volumeButton = screen.getByRole('button', { name: /volume/i })
  
  // 测试静音切换
  fireEvent.click(volumeButton)
  expect(screen.getByTestId('volume-slider')).toBeDefined()
}) 