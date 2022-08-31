import React, { useState, useEffect } from 'react';
import { BucketItem } from '../domain/entity/bucket_item';
import { useBucketItem } from '../api/bucket_item_api';

const BucketItems = () => {
  const { bucketItem, isLoading, isError } = useBucketItem();

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
          <tr>
            <td>{bucketItem?.id}</td>
            <td>{bucketItem?.name}</td>
          </tr>
        </tbody>
      </table>
    </>
  );
};

export default BucketItems;
