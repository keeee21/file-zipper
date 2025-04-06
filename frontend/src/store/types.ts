// stateの型
export type UserState = {
  user: User | null;
};

export type AuthState = {
  token: string;
};

// Google認証後、バックエンドでjwtを発行し、フロントエンドに返す型
export type GoogleAuthCallbackResponse = {
  message: string;
  user: GoogleAuthCallbackResponseUser;
};

export type GoogleAuthCallbackResponseUser = {
  email: string;
  name: string;
  googleSub: string;
  token: string;
};

// piniaで管理するuserの型
export type User = {
  id: number;
  name: string;
  email: string;
};

export type InfoUserResponse = {
  user: User | null;
};
