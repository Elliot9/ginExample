import { Inter } from "next/font/google";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Blogs",
  description: "Create Something Amazing",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en" className="h-full">
      <body className={`${inter.className} flex flex-col min-h-screen bg-gray-100`}>
        {children}
        <footer className="bg-gray-800 text-white text-center py-4 mt-auto">
          <p>&copy; 2024 Elliot. 保留所有權利。</p>
        </footer>
      </body>
    </html>
  )
}