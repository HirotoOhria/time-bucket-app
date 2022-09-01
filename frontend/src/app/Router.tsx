import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import BucketItems from 'pages/BucketItems';
import Top from '../pages/Top';
import CreateBucketItem from '../pages/CreateBuckeItem';

const Router = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Top />} />
        <Route path='bucket-items' element={<BucketItems />} />
        <Route path='bucket-item/create' element={<CreateBucketItem />} />
      </Routes>
    </BrowserRouter>
  );
};

export default Router;
