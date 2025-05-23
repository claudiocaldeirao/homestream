"use client";

import { useRouter } from "next/navigation";

export function WatchButton({ id }: { id: string }) {
  const router = useRouter();

  const handleClick = () => {
    router.push(`/movies/${id}/watch`);
  };

  return (
    <button
      onClick={handleClick}
      className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 transition"
    >
      Watch
    </button>
  );
}
