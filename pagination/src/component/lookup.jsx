import { Table } from "antd";
import { useEffect, useState } from "react"

export default function LookUp({ placeholder, dataSource, column, onSelect }) {
  const [open, setOpen] = useState(false);
  const [options, setOptions] = useState([] || dataSource);
  const [selected, setSelected] = useState({});
  useEffect(() => {
    setOptions(dataSource);
  }, [dataSource])

  const select = (record) => {
    onSelect(record);
    setSelected(record);
  }

  const handleBlur = () => {
    setTimeout(() => {
      setOpen(false);
    }, 100);
  };

  const handleOnChange = (e) => {
    setSelected((prevData) => {
      const newData = { ...prevData };
      newData.name = e.target.value;
      return newData;
    });
  }

  return (
    <>
      <div className="w-full max-h-96">
        <input type="text" style={{ color: '#000' }} onBlur={() => handleBlur()} value={selected.name} onFocus={() => setOpen(true)} onChange={handleOnChange} placeholder={placeholder} className="h-[30px] w-full mb-2" />
        {open ?
          <Table
            onRow={(record) => {
              return {
                onClick: () => select(record)
              }
            }}
            size="small"
            dataSource={options}
            columns={column}
            pagination={{
              position: [],
            }}
          /> : <></>}
      </div>
    </>
  )
}