import type { AxiosException } from '@/lib/types';
import type { InfoUserResponse } from '@/store/types';

import $axios, { getCacheConfig } from '../lib/axios';

import type { CacheAxiosResponse } from 'axios-cache-interceptor';

class UserService {
  async getInfo(useCache = false): Promise<InfoUserResponse> {
    let result: CacheAxiosResponse<InfoUserResponse>;
    try {
      result = await $axios.get('/api/user/info', getCacheConfig(useCache));
    } catch (ex) {
      const axiosException = ex as AxiosException;
      if (axiosException.isAxiosError) {
        throw new Error(`getInfo() - axios get error: ${axiosException.message}`);
      }
      throw new Error(`getInfo() - axios get error: ${ex}`);
    }
    return result.data;
  }
}

export default new UserService();
