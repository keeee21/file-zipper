import axios, { AxiosError, type AxiosRequestConfig, type AxiosResponse } from 'axios';
import { setupCache, type CacheOptions as AxiosCacheOptions, buildMemoryStorage, buildKeyGenerator } from 'axios-cache-interceptor';
import { md5 } from 'js-md5';

const axiosCreateConfig: AxiosRequestConfig = {
  headers: {
    'Content-type': 'application/json',
    'Cache-Control': 'no-cache',
  },
};

if (import.meta.env.VITE_API_ENV === 'prod') {
  axiosCreateConfig.baseURL = 'https://api.file-zipper.com';
} else if (import.meta.env.VITE_API_ENV === 'stg') {
  axiosCreateConfig.baseURL = 'https://staging.api.file-zipper.com';
} else if (import.meta.env.VITE_API_ENV === 'local') {
  axiosCreateConfig.baseURL = import.meta.env.VITE_API_URL;
}

const options: AxiosCacheOptions = {
  ttl: 60 * 1000, // キャッシュの有効時間. msec
  storage: buildMemoryStorage(false, false, 100),
  generateKey: buildKeyGenerator((request) => {
    let t = '___';
    if (request.headers?.Authorization) {
      t = md5(request.headers.Authorization);
    }
    const custom = `${request.method}#${request.url}#${t}`;

    return {
      method: request.method,
      url: request.url,
      custom,
    };
  }),
};
const axInstance = axios.create(axiosCreateConfig);
const $axios = setupCache(axInstance, options);

$axios.interceptors.response.use(
  (response: AxiosResponse): AxiosResponse => {
    return response;
  },
  async (error: AxiosError): Promise<never> => {
    return Promise.reject(error);
  },
);

export function getCacheConfig(useCache: boolean): { cache: false } | undefined {
  return !useCache ? { cache: false } : undefined;
}

export default $axios;
