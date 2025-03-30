import AuthService from '@/services/AuthService'

const accessToken = AuthService.getAccessToken()
type UseGetSignedUrlResponse = {
  data: string[];
}

/**
 * signedUrlを取得する
 * @param roomId 
 * @param password 
 * @returns 
 */
export const useGetSignedUrl = async (
  roomId: string,
  password: string
): Promise<UseGetSignedUrlResponse> => {
  try {
    const res = await fetch(`/api/files/${roomId}/signed-url`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${accessToken}`,
      },
      body: JSON.stringify({ password }),
    });

    if (!res.ok) {
      const errorText = await res.text();
      throw new Error(`署名付きURL取得失敗: ${res.status} - ${errorText}`);
    }

    const json = await res.json();
    return { data: json.data };
  } catch (error) {
    console.error('Error getting signed URL:', error);
    return { data: [] };
  }
}