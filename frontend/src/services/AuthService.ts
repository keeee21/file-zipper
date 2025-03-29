export const LOCAL_STORAGE_ACCESS_TOKEN_KEY = 'accessToken';

class AuthService {
  getAccessToken(): string {
    return localStorage.getItem(LOCAL_STORAGE_ACCESS_TOKEN_KEY) || '';
  }
}

export default new AuthService();