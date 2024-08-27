import Image from "next/image";
import Link from "next/link";

export default function Home() {
  return (
    <main className="container mx-auto px-4 py-8">
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {/* 這裡可以添加文章列表 */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-xl font-semibold mb-2 text-black">文章標題XXXX</h2>  
          <p className="text-gray-600 mb-4">文章摘要...</p>
          <Link href="/article/1" className="text-blue-500 hover:underline">
            閱讀更多
          </Link>
        </div>
        {/* 可以重複上面的 div 來添加更多文章 */}
      </div>
    </main>
  )
}