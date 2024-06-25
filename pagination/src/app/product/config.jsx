import {Button} from "antd";

const columns = (update) => {
  return ([
  {
    title: 'Name',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: 'Code',
    dataIndex: 'code',
    key: 'code',
  },
  {
    title: 'Action',
    dataIndex: 'action',
    key: 'action',
    render: (_, record) => {
      return (
        <>
          <Button onClick={() => update(record)}>
            Update
          </Button>
        </>
      );
    }
  }
])};

export const detailColumns =
  [
    {
      title: 'Size',
      dataIndex: 'size',
      key: 'size',
    },
    {
      title: 'Price',
      dataIndex: 'price',
      key: 'price',
    },
];

export default columns;