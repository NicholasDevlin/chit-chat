"use client"
import { useRouter } from 'next/navigation'
import { useState } from "react"
import PhoneInput from 'react-phone-number-input'
import 'react-phone-number-input/style.css'

export default function SignIn() {
  const [isSignIn, setIsSignIn] = useState(true);
  const [user, setUser] = useState({});
  const router = useRouter()

  const handleSubmit = async () => {
    try {
      const response = await fetch(`http://localhost:80/user/${isSignIn ? "login" : "register"}`, {
        method: "POST",
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(user),
      });

      const responseData = await response.json();

      if (responseData.success) {
        localStorage.setItem("authToken", responseData.data.token);
        router.push('/')
      } else {
        throw new Error(`${responseData.message}`);
      }
    } catch (error) {
      alert(error);
    }
  }

  return (
    <div className={`flex h-screen w-screen text-black bg-white absolute top-0 left-0 justify-center items-center`}>
      <div className="h-5/6 w-3/6 ">
        <div className="flex px-3 pt-3">
          <button onClick={() => setIsSignIn(true)} className={`${isSignIn ? 'bg-gray-300' : 'bg-gray-500'} rounded-t-lg p-2`}>
            Sign In
          </button>
          <button onClick={() => setIsSignIn(false)} className={`${!isSignIn ? 'bg-gray-300' : 'bg-gray-500'} rounded-t-lg p-2`}>
            Sign Up
          </button>
        </div>
        <div className="px-3 flex h-full w-full">
          {
            isSignIn ? <SignInPage onSubmit={handleSubmit} setUser={setUser} /> : <SignUpPage onSubmit={handleSubmit} setUser={setUser} />
          }
        </div>
      </div>
    </div>
  )
}

const SignInPage = ({ onSubmit, setUser }) => {
  const handleInputChange = (input) => {
    let name = input.target.name;
    let value = input.target.value;
    setUser((prevData) => ({ ...prevData, [name]: value }));
  }
  return (
    <div className="bg-gray-300 w-full h-5/6 text-lg flex justify-center flex-col">
      <div className="my-3 flex justify-center items-start w-full">
        <div className="min-w-32 me-3">Email </div>
        <input type="email" onChange={handleInputChange} className="p-1" placeholder="email" name="email" />
      </div>
      <div className="my-3 flex justify-center w-full">
        <div className="min-w-32 me-3">Password </div>
        <input type="password" onChange={handleInputChange} className="p-1" placeholder="password" name="password" />
      </div>
      <div className="flex justify-center">
        <button className="hover:bg-gray-400 w-2/6 p-2 border border-solid" onClick={onSubmit}>Sign in</button>
      </div>
    </div>
  )
}

const SignUpPage = ({ onSubmit, setUser }) => {
  const onChange = (phoneNumber) => {
    setUser((prevData) => ({ ...prevData, "phoneNumber": phoneNumber }));
  }
  const handleInputChange = (input) => {
    let name = input.target.name;
    let value = input.target.value;
    setUser((prevData) => ({ ...prevData, [name]: value }));
  }
  return (
    <div className="bg-gray-300 w-full h-5/6 text-lg flex justify-center flex-col">
      <div className="my-3 flex justify-center items-start w-full">
        <div className="min-w-32 me-3">Name </div>
        <input type="text" onChange={handleInputChange} className="p-1 w-64" placeholder="Name" name="name" />
      </div>
      <div className="my-3 flex justify-center items-start w-full">
        <div className="min-w-32 me-3">Phone Number </div>
        <div className="w-64">
          <PhoneInput
            international
            countryCallingCodeEditable={false}
            defaultCountry="ID"
            onChange={onChange}
          />
        </div>
      </div>
      <div className="my-3 flex justify-center items-start w-full">
        <div className="min-w-32 me-3">Email </div>
        <input type="email" onChange={handleInputChange} className="p-1 w-64" placeholder="Email" name="email" />
      </div>
      <div className="my-3 flex justify-center w-full">
        <div className="min-w-32 me-3">Password </div>
        <input type="password" onChange={handleInputChange} className="p-1 w-64" placeholder="Password" name="password" />
      </div>
      <div className="flex justify-center">
        <button className="hover:bg-gray-400 w-2/6 p-2 border border-solid" onClick={onSubmit}>Sign up</button>
      </div>
    </div>
  )
}