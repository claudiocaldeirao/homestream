import { useEffect, useState } from "react";
import Link from "next/link";
import Image from "next/image";
import axios from "axios";

const API_BASE = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

type OmdbMetadata = {
  Title: string;
  Year: string;
  Rated: string;
  Released: string;
  Runtime: string;
  Genre: string;
  Director: string;
  Plot: string;
  Language: string;
  Country: string;
  Awards: string;
  Poster: string;
  Metascore: string;
  ImdbRating: string;
  ImdbID: string;
  Response: string;
  Error: string;
};
type Movie = {
  Name: string;
  Path: string;
  Ext: string;
  Scanned: number;
  metadata: OmdbMetadata;
};

export default function Home() {
  const [movies, setMovies] = useState([]);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchMovies = async () => {
      setLoading(true);
      const res = await axios.get(`${API_BASE}/movies?page=${page}&limit=12`);
      setMovies(res.data);
      setLoading(false);
    };
    fetchMovies();
  }, [page]);

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">Movie Catalog</h1>
      {loading ? (
        <p>Loading...</p>
      ) : (
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
          {movies.map((movie: Movie) => (
            <Link
              href={`/movies/${movie.metadata.ImdbID}`}
              key={movie.metadata.ImdbID}
            >
              <div className="border rounded-xl shadow hover:shadow-lg p-2 cursor-pointer">
                <Image
                  src={movie.metadata.Poster}
                  alt={movie.metadata.Title}
                  className="w-full h-48 object-cover mb-2 rounded"
                />
                <h2 className="text-md font-semibold truncate">
                  {movie.metadata.Title}
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

// pages/movies/[id].js
export function MovieDetails({ movie }: { movie: Movie }) {
  return (
    <div className="p-4">
      <h1 className="text-3xl font-bold mb-2">{movie.metadata.Title}</h1>
      <Image
        src={movie.metadata.Poster}
        alt={movie.metadata.Title}
        className="w-64 mb-4"
      />
      <p>
        <strong>Year:</strong> {movie.metadata.Year}
      </p>
      <p>
        <strong>Genre:</strong> {movie.metadata.Genre}
      </p>
      <p>
        <strong>Director:</strong> {movie.metadata.Director}
      </p>
      <p>
        <strong>Plot:</strong> {movie.metadata.Plot}
      </p>
      <p>
        <strong>IMDB Rating:</strong> {movie.metadata.ImdbID}
      </p>
    </div>
  );
}

// pages/movies/[id].js - SSR
export async function getServerSideProps(context: { params: { id: string } }) {
  const { id } = context.params;
  const res = await axios.get(`${API_BASE}/movies/${id}`);
  return {
    props: {
      movie: res.data,
    },
  };
}
