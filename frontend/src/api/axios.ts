import axios, { AxiosError } from 'axios';

export type BaseSuccessResponse = {
  isLoading: boolean;
  isError?: AxiosError<BackendErrorResponse>;
};

export type SuccessResponse = BaseSuccessResponse & {
  success: boolean | undefined;
};

export type BackendErrorResponse = {
  message: string;
};

export const baseAxios = axios.create({
  baseURL: 'http://localhost:8080', // バックエンドB のURL:port を指定する
  headers: {
    'Content-Type': 'application/json',
    'Access-Control-Allow-Origin': '*',
  },
  responseType: 'json',
});
