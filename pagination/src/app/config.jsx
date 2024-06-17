import {Button} from "antd";

const columns = [
  {
    title: 'Name',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: 'Age',
    dataIndex: 'age',
    key: 'age',
  },
  {
    title: 'Address',
    dataIndex: 'address',
    key: 'address',
  },
  {
    title: 'Action',
    dataIndex: 'action',
    key: 'action',
    // render: (_, record) => {
    //   return (
    //     <>
    //       <Button>
    //         Update
    //       </Button>
    //     </>
    //   );
    // }
  }
];

export default columns;