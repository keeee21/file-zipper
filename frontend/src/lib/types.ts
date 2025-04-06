/*
 *
 *
 * Axios 関連
 */
export type AxiosException = {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  config: { [key: string]: any };

  request: XMLHttpRequest;

  response: AxiosExceptionResponse;

  isAxiosError: boolean;

  stack: string;

  message: string; // axios が生成したメッセージ
};

export type AxiosExceptionResponse = {
  status: number; // http error code

  statusText: string; // HTTP Status Code に対応するテキスト。なので、HTTP 400 response なら Bad Request が入ってる

  data: unknown; // 掘り出したい物。ここの type は ExceptionResponse に一致します

  headers: { [key: string]: string }; // HTTP Header

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  config: { [key: string]: any };
};
