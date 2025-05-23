"use client";
import { useEffect, useState } from "react";
import Link from "next/link";
import Image from "next/image";
import axios from "axios";
import { Movie } from "./movies/types";

const API_BASE = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export default function HomePage() {
  const [movies, setMovies] = useState<Movie[]>([]);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchMovies = async () => {
      setLoading(true);
      const res = await axios.get(`${API_BASE}/movies?page=${page}&limit=12`);
      console.log(res.data);
      setMovies(res.data);
      setLoading(false);
    };
    fetchMovies();
  }, [page]);

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">Movie Catalog</h1>
      {loading ? (
        <p>Loading...</p>
      ) : (
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
          {movies.map((movie) => (
            <Link href={`/movies/${movie.Name}`} key={movie.Name}>
              <div className="border rounded-xl shadow hover:shadow-lg p-2 cursor-pointer">
                <Image
                  src={movie.Metadata.Poster || "/images/clapperboard.png"}
                  alt={movie.Metadata.Title}
                  width={512}
                  height={512}
                  className="w-full h-48 object-cover mb-2 rounded"
                />
                <h2 className="text-md font-semibold truncate">
                  {movie.Metadata.Title}
                </h2>
              </div>
            </Link>
          ))}
        </div>
      )}
      <div className="mt-4 flex justify-center gap-2">
        <button
          onClick={() => setPage((p) => Math.max(1, p - 1))}
          className="px-4 py-2 border rounded"
        >
          Previous
        </button>
        <button
          onClick={() => setPage((p) => p + 1)}
          className="px-4 py-2 border rounded"
        >
          Next
        </button>
      </div>
    </div>
  );
}
