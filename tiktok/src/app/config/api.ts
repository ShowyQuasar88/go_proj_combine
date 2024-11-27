export const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:8000'

export const API_ROUTES = {
  AUTH: {
    LOGIN: `${API_BASE_URL}/v1/user/login`,
    LOGOUT: `${API_BASE_URL}/v1/user/logout`,
    CHECK: `${API_BASE_URL}/api/auth/check`,
  },
  VIDEO: {
    LIKE: (id: number) => `${API_BASE_URL}/api/videos/${id}/like`,
    COMMENT: (id: number) => `${API_BASE_URL}/api/videos/${id}/comment`,
    SHARE: (id: number) => `${API_BASE_URL}/api/videos/${id}/share`,
  }
}
