'use client';

import { useEffect, useState } from 'react';
import { useParams } from 'next/navigation';
import Link from 'next/link';
import parse from 'html-react-parser';
export default function ArticlePage() {
  const [article, setArticle] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const { id } = useParams();

  useEffect(() => {
    async function fetchArticle() {
      try {
        const res = await fetch(`http://localhost:8080/api/articles/${id}`);
        if (!res.ok) {
          throw new Error('Failed to fetch article');
        }
        const data = await res.json();
        setArticle(data);
      } catch (error) {
        console.error('Error fetching article:', error);
        setError(error.message);
      } finally {
        setLoading(false);
      }
    }

    fetchArticle();
  }, [id]);

  if (loading) return <div className="text-center py-10 text-gray-500">加載中...</div>;
  if (error) return <div className="text-center py-10 text-red-500">錯誤: {error}</div>;
  if (!article) return <div className="text-center py-10">文章不存在</div>;

  return (
    <div className="max-w-4xl mx-auto px-4 py-8">
      <Link href="/" className="text-blue-500 hover:underline mb-4 inline-block">
        &larr; 返回文章列表
      </Link>
      <article className="bg-white shadow-lg rounded-lg overflow-hidden">
        <div className="p-6">
          <h1 className="text-3xl font-bold text-gray-800 mb-4">{article.title}</h1>
          <div className="flex items-center text-gray-600 text-sm mb-4">
            <span className="mr-4">作者: {article.author || '未知'}</span>
            <span>發布時間: {new Date(article.created_at).toLocaleDateString()}</span>
          </div>
          <div className="prose max-w-nonea text-black" style={{whiteSpace: "pre-wrap"}}>
            {article.content ? parse(article.content) : <p>此文章暫無內容</p>}
          </div>
        </div>
      </article>
    </div>
  );
}