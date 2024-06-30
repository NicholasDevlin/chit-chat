"use client"
import { Table, Row, Col, Button, Modal, Input } from "antd";
import columns from "@/app/customer/config"
import React, { useState, useEffect } from "react";
import NumberFormat from 'react-number-format';

export default function Customer() {
  const [data, setData] = useState([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [loading, setLoading] = useState(false);
  const [customer, setCustomer] = useState({
    name: "",
    address: "",
    age: 0
  });
  const [tableParams, setTableParams] = useState({
    current: 1,
    pageSize: 10,
    total: 0,
    sortName: true
  });

  const fetchData = () => {
    setLoading(true);
    const queryParams = new URLSearchParams(tableParams);
    fetch(`http://localhost:80/customer?${queryParams.toString()}`)
      .then((res) => res.json())
      .then((results) => {
        setData(results.data);
        setLoading(false);
        setTableParams((prevData) => ({
          ...prevData,
          ...results.pagination,
        })
        );
      })
      .catch((error) => {
        console.error("Error fetching data:", error);
        setLoading(false);
      });
  };

  useEffect(() => {
    fetchData();
  }, [tableParams.current]);

  const handleTableChange = (pagination, filters, sorter) => {
    setTableParams((prevData) => ({ ...prevData, ...pagination }));
  }

  const handleInputChange = (input) => {
    let name = input.target.name;
    let value = input.target.value;
    setCustomer((prevData) => ({ ...prevData, [name]: name == "name" ? value.toUpperCase() : value, [name]: name === "age" ? parseInt(value) : value }));
  }

  const handleSubmit = async () => {
    customer.pagination = tableParams;
    try {
      const response = await fetch("http://localhost:80/customer", {
        method: "POST",
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(customer),
      });

      const responseData = await response.json();

      if (responseData.success) {
        setCustomer(tableParams);
        setIsModalOpen(false);
        setTableParams((prevData) => ({ ...prevData, ...responseData.pagination }));
        fetchData();
      } else {
        throw new Error(`${responseData.message}`);
      }
    } catch (error) {
      alert(error);
    }
  }

  const handleUpdate = (record) => {
    setIsModalOpen(true);
    setCustomer(record);
  }

  return (
    <>
      <ModalEditor
        isOpen={isModalOpen}
        onCancel={() => setIsModalOpen(false)}
        customer={customer}
        handleInputChange={handleInputChange}
        onOk={handleSubmit} />
      <Row className="p-8 h-screen">
        <Col span={24}>
          <Row className="mb-3">
            <Button type="primary" onClick={() => setIsModalOpen(true)}>Add Customer</Button>
          </Row>
          <Table
            dataSource={data}
            columns={columns(handleUpdate)}
            pagination={tableParams}
            onChange={handleTableChange}
            loading={loading}
          />
        </Col>
      </Row>
    </>
  );
}

const ModalEditor = ({ isOpen, onOk, onCancel, customer, handleInputChange }) => {
  return (
    <Modal open={isOpen} onOk={onOk} onCancel={onCancel}>
      <Row className="my-5" gutter={10}>
        <h4>Name</h4>
        <Input value={customer.name} onChange={handleInputChange} name="name" placeholder="Name" />
      </Row>
      <Row className="my-5" gutter={10}>
        <h4>Address</h4>
        <Input value={customer.address} onChange={handleInputChange} name="address" placeholder="Address" />
      </Row>
      <Row className="my-5" gutter={10}>
        <h4>Age</h4>
        <Input value={customer.age} onChange={handleInputChange} name="age" placeholder="Age" type="number" />
      </Row>
      <Row className="my-5" gutter={10}>
        <h4>Test separator : </h4>
        <div>
          <NumberFormat className="border" thousandSeparator={true} />
        </div>
      </Row>
      <Row className="my-5" gutter={10}>
        <h4>Test format : </h4>
        <div>
          <NumberFormat className="border" format={"##/##/####"} mask={['D', 'D', 'M', 'M', 'Y', 'Y', 'Y', 'Y']} placeholder="DD/MM/YYYY" thousandSeparator={true} />
        </div>
      </Row>
    </Modal>);
}
