"use client";

import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

const API_BASE = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

const VideoPlayerPage = () => {
  const { id } = useParams();
  const [videoUrl, setVideoUrl] = useState<string | null>(null);

  useEffect(() => {
    if (id) {
      setVideoUrl(`${API_BASE}/movies/${id}/watch`);
    }
  }, [id]);

  if (!videoUrl) return <div>Loading...</div>;

  return (
    <div className="flex flex-col items-center justify-center min-h-screen p-8 bg-black text-white">
      <h1 className="text-2xl font-bold mb-4">Now Playing</h1>
      <video
        src={videoUrl}
        controls
        className="w-full max-w-4xl h-auto rounded shadow-lg"
      >
        Your browser does not support the video tag.
      </video>
    </div>
  );
};

export default VideoPlayerPage;
