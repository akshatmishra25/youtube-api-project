# YouTube API Project

## Overview

This project is an API to fetch the latest videos sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

## Requirements

- Go 1.23.1 or later
- MongoDB

## Setup

1. **Clone the repository**:
    ```sh
    git clone https://github.com/akshatmishra25/youtube-api-project.git
    cd youtube-api-project
    ```

2. **Create and configure the [.env](http://_vscodecontentref_/0) file**:
    Copy the [.env.sample](http://_vscodecontentref_/1) file to [.env](http://_vscodecontentref_/2) and fill in the required values.
    ```sh
    cp .env.sample .env
    ```

    Update the [.env](http://_vscodecontentref_/3) file with your MongoDB URI and YouTube API key:
    ```env
    PORT=8000
    MONGODB_URI=your_mongodb_uri
    DATABASE_NAME=VideoDB
    YOUTUBE_API_KEY=your_youtube_api_key
    SEARCH_QUERY=football
    ```

3. **Install dependencies**:
    ```sh
    go mod tidy
    ```

4. **Run the server**:
    ```sh
    go run main.go
    ```

## API Endpoints

### Get Videos

- **Endpoint**: `GET /videos`
- **Description**: Returns the stored video data in a paginated response sorted in descending order of published datetime.
- **Query Parameters**:
  - `page`: The page number (default is 1)
  - `limit`: The number of videos per page (default is 10)

#### Example Request

```sh
curl "http://localhost:8000/videos?page=1&limit=10"
```

#### Example Respnse

```sh
[
  {
    "id": "R3vhbJAgeVA",
    "title": "32nd Triveni Cup Football Tournament -2081 : New Road Team FC (NRT) VS Chame Yuwa Club, Manang🔥 🔥",
    "description": "LIVE! LIVE!! LIVE!!! WE ARE LIVE FROM BESISAHAR, LAMJUNG !! 32nd Triveni Cup Football Tournament -2081 : New Road ...",
    "publishedAt": "2024-12-29T11:18:52Z",
    "thumbnails": [
      "https://i.ytimg.com/vi/R3vhbJAgeVA/default.jpg"
    ]
  },
  {
    "id": "NaiIuZtYCC4",
    "title": "The People&#39;s Choice Award Football Boot Of The Year - The Football Boot Hour Podcast",
    "description": "As we come to the end of the year it's time to for the boot of the year awards, but it's one thing for us to tell you what we think, but, ...",
    "publishedAt": "2024-12-29T11:00:20Z",
    "thumbnails": [
      "https://i.ytimg.com/vi/NaiIuZtYCC4/default.jpg"
    ]
  },
  {
    "id": "wel2ag1x9cs",
    "title": "Canning Football ILVE 🛑 ফাইনাল ম্যাচ 💥🏆🔥 12/29/24",
    "description": "football India https://www.facebook.com/profile.php?id=61552141948192&mibextid=ZbWKwL India football live India football ...",
    "publishedAt": "2024-12-29T09:07:16Z",
    "thumbnails": [
      "https://i.ytimg.com/vi/wel2ag1x9cs/default.jpg"
    ]
  },
  {
    "id": "BcLTto7IyJE",
    "title": "32nd Triveni Cup Football Tournament -2081 : Triveni Yuwa Club VS Sundarbazar FC🔥 🔥",
    "description": "LIVE! LIVE!! LIVE!!! WE ARE LIVE FROM BESISAHAR, LAMJUNG !! 32nd Triveni Cup Football Tournament -2081 : Triveni Yuwa ...",
    "publishedAt": "2024-12-29T08:06:05Z",
    "thumbnails": [
      "https://i.ytimg.com/vi/BcLTto7IyJE/default.jpg"
    ]
  },
  {
    "id": "jlu0dMoqd4U",
    "title": "Live | The Football Show | Talk Show | Football | Football Analyst | T Sports",
    "description": "Live | The Football Show | Talk Show | Football | Football Analyst | T Sports Welcome to T Sports - Your ...",
    "publishedAt": "2024-12-29T06:53:33Z",
    "thumbnails": [
      "https://i.ytimg.com/vi/jlu0dMoqd4U/default.jpg"
    ]
  },
  {
    "id": "SqfElboSRuM",
    "title": "🔴 Việt Nam vs Singapore | เวียดนาม - สิงคโปร์ | Nghẹt Thở Cạnh Tranh Tấm Vé Đi Tiếp",
    "description": "Việt Nam vs Singapore | เวียดนาม - สิงคโปร์ | Nghẹt Thở Cạnh Tranh Tấm Vé Đi Tiếp Top Football TV - Cập nhật tin tức, Phân tích ...",
    "publishedAt": "2024-12-29T06:36:49Z",
    "thumbnails": [
      "https://i.ytimg.com/vi/SqfElboSRuM/default_live.jpg"
    ]
  },
  {
    "id": "vzOAFFdd_H4",
    "title": "Cardinals - Rams LIVE STREAM: Watch Now!",
    "description": "Cardinals vs Rams LIVE: How to Watch Online. Cardinals vs Rams LIVE STREAM: Don't Miss the Action! Catch all the action as ...",
    "publishedAt": "2024-12-29T06:15:10Z",
    "thumbnails": [
      "https://i.ytimg.com/vi/vzOAFFdd_H4/default.jpg"
    ]
  },
  {
    "id": "yyLEJywFT7A",
    "title": "Colorado Buffaloes&#39; Alamo Bowl postgame news conference",
    "description": "Head Coach Deion Sanders speaks after the Colorado Buffaloes lost to BYU 36-14 in the Alamo Bowl.",
    "publishedAt": "2024-12-29T04:36:20Z",
    "thumbnails": [
      "https://i.ytimg.com/vi/yyLEJywFT7A/default.jpg"
    ]
  },
  {
    "id": "Sp0iTOwJGqE",
    "title": "Alamo Bowl: BYU Cougars vs. Colorado Buffaloes | Full Game Highlights | ESPN CFB",
    "description": "Check out these highlights as the No. 17 BYU Cougars dominate the No. 23 Colorado Buffaloes, 36-14, in the 2024 Alamo Bowl.",
    "publishedAt": "2024-12-29T04:01:29Z",
    "thumbnails": [
      "https://i.ytimg.com/vi/Sp0iTOwJGqE/default.jpg"
    ]
  },
  {
    "id": "XZqZDMn4y6o",
    "title": "Experiment: Football VS Coca Cola Zero, Fanta, Mtn Dew, Powerade, Mirrinda and Mentos in the toilet",
    "description": "Experiment: Football VS Coca Cola Zero, Fanta, Mtn Dew, Powerade, Mirrinda and Mentos in the toilet.",
    "publishedAt": "2024-12-29T01:39:46Z",
    "thumbnails": [
      "https://i.ytimg.com/vi/XZqZDMn4y6o/default.jpg"
    ]
  }
]
```
