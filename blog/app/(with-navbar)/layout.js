import Navbar from '../navbar';

export default function WithNavbarLayout({ children }) {
  return (
    <div className="flex flex-col min-h-screen bg-gradient-to-br from-gray-100 to-gray-300">
      <Navbar />
      <div className="flex-grow">
        {children}
      </div>
    </div>
  )
}