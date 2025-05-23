import axios from "axios";
import { notFound } from "next/navigation";
import Image from "next/image";

const API_BASE = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export default async function MovieDetails({
  params,
}: {
  params: { id: string };
}) {
  try {
    const { id } = await params;
    const res = await axios.get(`${API_BASE}/movies/${id}`);
    const movie = res.data;

    return (
      <div>
        <h1 className="text-3xl font-bold mb-2">{movie.title}</h1>
        <Image
          src={movie.Metadata.Poster || "/images/clapperboard.png"}
          alt={movie.Metadata.Title}
          width={512}
          height={512}
        />
        <p>
          <strong>Year:</strong> {movie.Metadata.Year}
        </p>
        <p>
          <strong>Genre:</strong> {movie.Metadata.Genre}
        </p>
        <p>
          <strong>Director:</strong> {movie.Metadata.Director}
        </p>
        <p>
          <strong>Plot:</strong> {movie.Metadata.Plot}
        </p>
        <p>
          <strong>IMDB Rating:</strong> {movie.Metadata.ImdbRating}
        </p>
      </div>
    );
  } catch (err) {
    console.error(err);
    return notFound();
  }
}
