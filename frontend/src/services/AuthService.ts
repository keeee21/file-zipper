export const LOCAL_STORAGE_ACCESS_TOKEN_KEY = 'accessToken';

class AuthService {
  getAccessToken(): string {
    return localStorage.getItem(LOCAL_STORAGE_ACCESS_TOKEN_KEY) || '';
  }

  isLoggedIn(): boolean {
    const token = localStorage.getItem(LOCAL_STORAGE_ACCESS_TOKEN_KEY)
    if (!token) return false
  
    const payload = this.parseJwt(token)
    if (!payload || !payload.exp) return false
  
    const now = Math.floor(Date.now() / 1000)
    return payload.exp > now
  }

  parseJwt(token: string): any {
    try {
      return JSON.parse(atob(token.split('.')[1]))
    } catch (e) {
      return null
    }
  }
}

export default new AuthService();