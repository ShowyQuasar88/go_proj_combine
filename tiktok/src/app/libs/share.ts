export interface ShareOption {
  id: string
  name: string
  icon: string
  color: string
  action: (url: string) => void | Promise<void>
}

export const copyToClipboard = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text)
    return true
  } catch (err) {
    console.error('复制失败:', err)
    return false
  }
}

export const openShareWindow = (url: string, title: string, w = 600, h = 400) => {
  const left = (window.screen.width / 2) - (w / 2)
  const top = (window.screen.height / 2) - (h / 2)
  return window.open(
    url,
    title,
    `toolbar=no, location=no, directories=no, status=no, menubar=no, scrollbars=no, resizable=no, copyhistory=no, width=${w}, height=${h}, top=${top}, left=${left}`
  )
}

export const shareOptions: ShareOption[] = [
  {
    id: 'wechat',
    name: '微信',
    icon: '💬',
    color: 'bg-green-500',
    action: (url) => {
      console.log('分享到微信:', url)
    }
  },
  {
    id: 'weibo',
    name: '微博',
    icon: '🔄',
    color: 'bg-red-500',
    action: (url) => {
      const shareUrl = `http://service.weibo.com/share/share.php?url=${encodeURIComponent(url)}`
      openShareWindow(shareUrl, '分享到微博')
    }
  },
  {
    id: 'qq',
    name: 'QQ',
    icon: '💭',
    color: 'bg-blue-500',
    action: (url) => {
      const shareUrl = `http://connect.qq.com/widget/shareqq/index.html?url=${encodeURIComponent(url)}`
      openShareWindow(shareUrl, '分享到QQ')
    }
  },
  {
    id: 'copy',
    name: '复制链接',
    icon: '📋',
    color: 'bg-gray-500',
    action: async (url) => {
      return await copyToClipboard(url)
    }
  }
]
