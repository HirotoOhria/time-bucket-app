import React from 'react';
import { useBucketItems } from '../api/bucket_item_api';

const BucketItems = () => {
  const { bucketItems, isLoading, isError } = useBucketItems();

  if (isLoading) {
    return <div>Loading...</div>;
  }
  if (isError) {
    return <div>{isError.response?.data.message}</div>;
  }

  return (
    <>
      <h1>Bucket Items</h1>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>名前</th>
          </tr>
        </thead>
        <tbody>
          {bucketItems?.map((bucketItem, i) => (
            <tr key={i}>
              <td>{bucketItem?.id}</td>
              <td>{bucketItem?.name}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </>
  );
};

export default BucketItems;
