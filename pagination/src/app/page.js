"use client"
import LookUp from "@/component/lookup";
import React, { useState, useEffect } from "react";

export default function Home() {
  const [data, setData] = useState({
    name: "",
    address: "",
    age: 0
  });
  const [options, setOptions] = useState([]);
  const [tableParams, setTableParams] = useState({
    current: 1,
    pageSize: 10,
    total: 0,
    sortName: true,
    name: ""
  });

  const fetchData = () => {
    const queryParams = new URLSearchParams(tableParams);
    fetch(`http://localhost:80/customer?${queryParams.toString()}`)
      .then((res) => res.json())
      .then((results) => {
        setOptions(results.data);
        setTableParams((prevData) => ({
          ...prevData,
          ...results.pagination,
        })
        );
      })
      .catch((error) => {
        console.error("Error fetching data:", error);
      });
  };

  useEffect(() => {
    fetchData();
  }, [tableParams])

  const columns =
    [
      {
        title: 'Name',
        dataIndex: 'name',
        key: 'name',
      },
      {
        title: 'Address',
        dataIndex: 'address',
        key: 'address',
      },
    ];

  const handleOnChange = (value) => {
    setTableParams({
      current: 1,
      pageSize: 10,
      total: 0,
      sortName: true,
      name: value,
    })
  }

  return (
    <>
      <div className="flex justify-between mb-3">
        <div>Name: {data.name}</div>
        <div>Address: {data.address}</div>
        <div>Age: {data.age}</div>
      </div>
      <div>
        <LookUp dataSource={options} onChange={handleOnChange} onSelect={setData} column={columns} />
      </div>
    </>
  );
}

