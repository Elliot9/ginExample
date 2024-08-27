"use client";

import { useState, useEffect } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';

const Navbar = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [username, setUsername] = useState('');
  const router = useRouter();

  useEffect(() => {
    // 檢查 localStorage 中是否存在 JWT token
    const token = localStorage.getItem('token');
    if (token) {
      // 這裡可以添加 token 驗證邏輯
      setIsLoggedIn(true);
      // 從 token 中解析用戶名,這裡僅為示例
      setUsername('User'); // 實際應用中,應該從 token 中解析用戶信息
    }
  }, []);

  const handleLogout = () => {
    localStorage.removeItem('token');
    setIsLoggedIn(false);
    setUsername('');
    router.push('/');
  };

  return (
    <nav className="bg-gray-100 shadow-md">
      <div className="container mx-auto px-4 py-3 flex justify-between items-center">
        <Link href="/" className="text-xl font-bold text-gray-800">
          Blogs
        </Link>
        <div className="space-x-4">
          {isLoggedIn ? (
            <>
              <span className="text-gray-600">Welcome, {username}</span>
              <button
                onClick={handleLogout}
                className="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded transition duration-300"
              >
                Sign out
              </button>
            </>
          ) : (
            <>
              <Link
                href="/signup"
                className="bg-gray-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded transition duration-300"
              >
                Sign up
              </Link>
              <Link
                href="/login"
                className="bg-white hover:bg-blue-600 hover:text-white text-gray-500 font-bold py-2 px-4 rounded border border-blue-500 transition duration-300"
              >
                Sign in
              </Link>
            </>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;