import AuthService from '@/services/AuthService';

const accessToken = AuthService.getAccessToken();

type UseGetRoomValidityResponse = {
  isValid: boolean;
};
/**
 * fileが有効期限切れかどうかを確認する
 * @param roomId
 * @return Promise<boolean>
 */
export const useGetRoomValidity = async (roomId: string): Promise<UseGetRoomValidityResponse> => {
  try {
    const res = await fetch(`/api/rooms/${roomId}/validity`, {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });

    const json = await res.json();
    return { isValid: json.data.isValid };
  } catch (error) {
    console.error('Error fetching file validity:', error);
    return { isValid: false };
  }
};
