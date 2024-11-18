export interface ShareOption {
  id: string
  name: string
  icon: string
  color: string
  action: (url: string) => void
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
