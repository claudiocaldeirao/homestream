export type OmdbMetadata = {
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

export type Movie = {
  _id: string;
  Name: string;
  Path: string;
  Ext: string;
  Scanned: number;
  Metadata: OmdbMetadata;
};
