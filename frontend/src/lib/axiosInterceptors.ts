import $axios from '@/lib/axios';

import AuthService from '../services/AuthService';

const setupAxiosInterceptors = () => {
  // MEMO: 基本方針として、ここで影響を与えるのは token に関することのみに留める。
  //       store への保存、local strage への出し入れは、service へ一任する。

  // MEMO: request を行うとき、 accessToken がどこかに保存されていればログイン中と見なして、
  //       常に Token を付与したリクエストにする。
  // MEMO: もし、どうしても付けたくないリクエストが発生する場合、 router の meta data を利用して制御する
  // MEMO: 現状 API サーバーとの通信を前提に実装している。 $axios を別ドメインに対しても利用する場合、
  //       $axiosOther の要に別インスタンスでクライアントを用意すること。
  $axios.interceptors.request.use(
    (config) => {
      const accessToken = AuthService.getAccessToken();
      // MEMO: accessToken の出し入れは service に一任。

      if (config.headers && accessToken != '') {
        config.headers['Authorization'] = `Bearer ${accessToken}`;
      }
      return config;
    },
    (error) => {
      return Promise.reject(error);
    },
  );

  // MEMO: HTTP レスポンスによって、いくつか動作を変える。
  //   401 Unauthorized - 「tokenの期限切れ」なので、refresh token を試みる。
  //                       （ログインAPIの場合は認証失敗なので、そのAPI URLだけ除外する）
  $axios.interceptors.response.use(
    (response) => {
      return response;
    },
    async (error) => {
      // MEMO: 401 Unauthorized の場合、token の期限切れなので refresh token を試みる みたいなことを後々したい
      throw error;
    },
  );
};

export default setupAxiosInterceptors;
