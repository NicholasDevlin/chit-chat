"use client"
import { useRouter } from 'next/router';
import { useEffect } from "react";

export default function Editor() {
  const { id } = router.query;
  useEffect(() => {
    debugger
    console.log(id);
  }, [id])
  return (
    <>hi</>
  );
}