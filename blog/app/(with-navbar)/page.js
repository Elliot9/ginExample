'use client';

import { useEffect, useState, useCallback, useRef } from 'react';
import Image from "next/image";
import Link from "next/link";

export default function HomePage() {
  const [articles, setArticles] = useState([]);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [hasMore, setHasMore] = useState(true);
  const observer = useRef();
  const initialLoadDone = useRef(false);

  const lastArticleElementRef = useCallback(node => {
    if (loading) return;
    if (observer.current) observer.current.disconnect();
    observer.current = new IntersectionObserver(entries => {
      if (entries[0].isIntersecting && hasMore && !loading) {
        setPage(prevPage => prevPage + 1);
      }
    });
    if (node) observer.current.observe(node);
  }, [loading, hasMore]);

  const fetchArticles = useCallback(async (pageNum) => {
    setLoading(true);
    try {
      const res = await fetch(`http://localhost:8080/api/articles?page=${pageNum}`, { cache: 'no-store' });
      if (!res.ok) {
        throw new Error('Failed to fetch articles');
      }
      const data = await res.json();
      setArticles(prev => pageNum === 1 ? data.items : [...prev, ...data.items]);
      setHasMore(data.items.length > 0);
    } catch (error) {
      console.error('Error fetching articles:', error);
      setError(error.message);
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    if (!initialLoadDone.current || page > 1) {
      fetchArticles(page);
      if (page === 1) {
        initialLoadDone.current = true;
      }
    }
  }, [page, fetchArticles]);

  useEffect(() => {
    console.log('Articles updated:', articles.length);
  }, [articles]);

  if (error) return <div>錯誤: {error}</div>;

  return (
    <main className="container mx-auto px-4 py-8">
      <h1 className="text-2xl font-bold mb-4 text-black">News</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {articles.map((article, index) => (
          <div 
            key={article.id} 
            ref={index === articles.length - 1 ? lastArticleElementRef : null}
            className="bg-white rounded-lg shadow-md p-6"
          >
            <h2 className="text-xl text-black font-semibold mb-2">{article.title}</h2>
            <p className="text-gray-600 mb-4">{article.summary || '文章摘要...'}</p>
            <h5 className="text-gray-600 mb-4">{article.author || '未知'}</h5>
            <Link href={`/article/${article.id}`} className="text-blue-500 hover:underline">
              閱讀更多
            </Link>
          </div>
        ))}
      </div>
      {loading && <div className="text-center mt-4 text-red-800">加載中...</div>}
    </main>
  );
}