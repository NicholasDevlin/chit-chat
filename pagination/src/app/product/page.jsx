"use client"
import { Button, Col, Row, Table } from "antd";
import React, { useState, useEffect } from "react";
import columns, { detailColumns } from "./config";
import { useRouter } from 'next/navigation';

export default function Product() {
  const route = useRouter();
  const [data, setData] = useState([]);
  const [detailData, setDetailData] = useState([]);
  const [loading, setLoading] = useState(false);
  const [product, setProduct] = useState({
    name: "",
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
    fetch(`http://localhost:80/product?${queryParams.toString()}`)
      .then((res) => res.json())
      .then((results) => {
        setData(results.data);
        if (results.data) {
          setDetailData(results.data[0].productDetail);
        }
        setLoading(false);
        setTableParams((prevData) => ({
          ...prevData,
          ...results.pagination,
        })
        );
      })
      .catch((error) => {
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
    setProduct((prevData) => ({ ...prevData, [name]: name == "age" ? parseInt(value) : value }));
  }

  const handleSubmit = async () => {
    product.pagination = tableParams;
    try {
      const response = await fetch("http://localhost:80/product", {
        method: "POST",
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(product),
      });

      const responseData = await response.json();

      if (responseData.success) {
        setProduct(tableParams);
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
    setProduct(record);
  }

  const getDetail = (record) => {
    let detail = data.find(x => x.uuid === record.uuid);
    setDetailData(detail.productDetail);
  }

  return (
    <>
      <Row className="p-8 h-screen">
        <Col span={24}>
          <Row className="mb-3">
            <Button type="primary" onClick={() => route.push("/product/form")}>Add Product</Button>
          </Row>
          <Table
            onRow={(record) => {
              return {
                onClick: () => getDetail(record)
              }
            }}
            columns={columns(handleUpdate)}
            expandable={{
              expandedRowRender: (record) => (
                <Table
                  columns={detailColumns}
                  dataSource={record.productDetail}
                />
              ),
              rowExpandable: (record) => record.productDetail.length != 0
            }}
            dataSource={data}
            pagination={tableParams}
            onChange={handleTableChange}
            loading={loading}
            rowKey={"uuid"}
          />
          <Table
            className="w-full"
            columns={detailColumns}
            dataSource={detailData}
            loading={loading}
            rowKey={"uuid"}
          />
        </Col>
      </Row>
    </>
  );
}