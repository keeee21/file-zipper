import AuthService from '@/services/AuthService';

const accessToken = AuthService.getAccessToken();
type UseGetFileResponse = {
  data: {
    fileNames: string[];
  };
};
/**
 * ファイル名を取得する
 * @param roomId
 * @returns Promise<UseGetFileResponse>
 */
export const useGetFileNames = async (roomId: string): Promise<UseGetFileResponse> => {
  try {
    const res = await fetch(`/api/files/${roomId}/name`, {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });

    const json = await res.json();
    console.log('File name:', json.data);
    return json;
  } catch (error) {
    console.error('Error getting file name:', error);
    return { data: { fileNames: [] } };
  }
};
