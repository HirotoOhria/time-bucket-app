import axios, { AxiosError, AxiosResponse } from 'axios';
import useSWR from 'swr';
import { BucketItem } from '../domain/entity/bucket_item';
import { BackendErrorResponse, baseAxios } from './axios';

type BucketItemResponse = {
  bucketItem?: BucketItem;
  isLoading: boolean;
  isError?: AxiosError<BackendErrorResponse>;
};

const fetcher = (url: string) => baseAxios.get(url).then((res) => res.data);

export const useBucketItem = (): BucketItemResponse => {
  const { data, error } = useSWR<
    AxiosResponse<BucketItem>,
    AxiosError<BackendErrorResponse>
  >('/bucket_item/01GBMPC11X0G8BK9XB5BYVJ5EM', fetcher);

  return {
    bucketItem: data?.data,
    isLoading: !error && !data,
    isError: error,
  };
};
