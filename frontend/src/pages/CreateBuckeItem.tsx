import React, { useEffect } from 'react';
import { Link } from 'react-router-dom';
import { callCreateBucketItem } from '../api/bucket_item_api';

const CreateBucketItem = () => {
  const [name, setName] = React.useState<string>('');
  
  return (
    <>
      <h1>Create Bucket Item</h1>
      <form onSubmit={() => callCreateBucketItem(name)}>
        <div>
          <label>
            Name
            <input
              type='text'
              value={name}
              required
              onChange={(event) => setName(event.target.value)}
            />
          </label>
        </div>
        <div>
          <input type='submit' value='Create' />
        </div>
      </form>
      <button>
        <Link to='/bucket-items'>Back to Bucket Items</Link>
      </button>
    </>
  );
};

export default CreateBucketItem;
