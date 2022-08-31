import axios, { AxiosResponse } from 'axios';
import useSWR from 'swr';
import { BucketItem } from '../domain/entity/bucket_item';
import { baseAxios } from './axios';

type BucketItemResponse = {
  bucketItem?: BucketItem;
  isLoading: boolean;
  isError?: Error;
};

const fetcher = (url: string) => baseAxios.get(url).then((res) => res.data);

export const useBucketItem = (): BucketItemResponse => {
  const { data, error } = useSWR<AxiosResponse<BucketItem>, Error>(
    '/bucket_item/01GBMPC11X0G8BK9XB5BYVJ5EM',
    fetcher
  );

  return {
    bucketItem: data?.data,
    isLoading: !error && !data,
    isError: error,
  };
};
