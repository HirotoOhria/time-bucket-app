import axios from 'axios';

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
