export const LOCAL_STORAGE_ACCESS_TOKEN_KEY = 'accessToken';

interface JwtPayload {
  exp?: number;
  [key: string]: unknown;
}

class AuthService {
  saveToken(token: string): void {
    localStorage.setItem(LOCAL_STORAGE_ACCESS_TOKEN_KEY, token);
  }

  getAccessToken(): string {
    return localStorage.getItem(LOCAL_STORAGE_ACCESS_TOKEN_KEY) || '';
  }

  isLoggedIn(): boolean {
    const token = localStorage.getItem(LOCAL_STORAGE_ACCESS_TOKEN_KEY);
    if (!token) return false;

    const payload = this.parseJwt(token);
    if (!payload || !payload.exp) return false;

    const now = Math.floor(Date.now() / 1000);
    return payload.exp > now;
  }

  parseJwt(token: string): JwtPayload | null {
    try {
      return JSON.parse(atob(token.split('.')[1]));
    } catch {
      return null;
    }
  }

  logout(): void {
    localStorage.removeItem(LOCAL_STORAGE_ACCESS_TOKEN_KEY);
  }
}

export default new AuthService();
