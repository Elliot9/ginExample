"use client";

import { useAuth } from './auth';
import Link from 'next/link';
import { useRouter } from 'next/navigation';

const Navbar = () => {
  const router = useRouter();
  const { isLoggedIn, user, logout, loading } = useAuth();
  const handleLogout = () => {
    logout();
    router.push('/');
  };

  if (loading) {
    return <nav className="bg-gray-100 shadow-md">
      <div className="container mx-auto px-4 py-3 flex justify-between items-center">
        <div className="text-xl font-bold text-gray-800">Blogs</div>
        <div className="space-x-4">
          <span className="text-gray-600">載入中...</span>
        </div>
      </div>
    </nav>;
  }

  return (
    <nav className="bg-gray-100 shadow-md">
      <div className="container mx-auto px-4 py-3 flex justify-between items-center">
        <Link href="/" className="text-xl font-bold text-gray-800">
          Blogs
        </Link>
        <div className="space-x-4">
          {isLoggedIn ? (
            <>
              <span className="text-gray-600">{user.name}</span>
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