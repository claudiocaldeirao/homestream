# HomeStream

HomeStream is a self-hosted media streaming platform inspired by Netflix. It allows you to organize, enrich, and stream your local video library through a modern web interface.

## Features

- 🧠 **Automatic Metadata Fetching**: Scans local directories for video files and enriches them with metadata from an external API.
- 📦 **Backend in Go**: Handles scanning, metadata storage, and video streaming.
- 🗂️ **MongoDB Integration**: Stores metadata and media information.
- 🌐 **Next.js Frontend**: Provides a catalog-style interface with a built-in media player.
- 📺 **Media Streaming**: Streams videos directly from your local machine to your browser.

## Requirements

- Go (1.18+)
- MongoDB
- Node.js & npm (for the frontend)

## Environment Configuration

This project uses environment variables to configure paths and external integrations. You can define these variables directly in your terminal or by using a .env file.

#### 🔑 Expected Environment Variables

| Name              | Description                                      | Required | Default Value             |
| ----------------- | ------------------------------------------------ | -------- | ------------------------- |
| CATALOG_FOLDER    | Root directory where video files will be scanned | No       | /homestream_catalog       |
| MOVIES_COLLECTION | collection name for metadata                     | No       | movies                    |
| OMDB_API_KEY      | API key for OMDb                                 | Yes      |                           |
| MONGODB_URI       | MongoDB connection URI                           | No       | mongodb://localhost:27017 |
| MONGODB_DATABASE  | Name of the MongoDB database                     | No       | homestreamdb              |
| API_PORT          | Port which the api will be exposed               | No       | 8080                      |

## License

This project is licensed under the [MIT License](LISCENSE).
