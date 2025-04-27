"use server";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";

const mockResults = {
  data: [
    {
      node: {
        id: 459,
        title: "One Piece Movie 01",
        main_picture: {
          medium: "https://cdn.myanimelist.net/images/anime/1770/97704.jpg",
          large: "https://cdn.myanimelist.net/images/anime/1770/97704l.jpg",
        },
        media_type: "movie",
      },
    },
    {
      node: {
        id: 21,
        title: "One Piece",
        main_picture: {
          medium: "https://cdn.myanimelist.net/images/anime/1244/138851.jpg",
          large: "https://cdn.myanimelist.net/images/anime/1244/138851l.jpg",
        },
        media_type: "tv",
      },
    },
    {
      node: {
        id: 12859,
        title: "One Piece Film: Z",
        main_picture: {
          medium: "https://cdn.myanimelist.net/images/anime/6/44297.jpg",
          large: "https://cdn.myanimelist.net/images/anime/6/44297l.jpg",
        },
        media_type: "movie",
      },
    },
    {
      node: {
        id: 31772,
        title: "One Punch Man Specials",
        main_picture: {
          medium: "https://cdn.myanimelist.net/images/anime/1452/97840.webp",
          large: "https://cdn.myanimelist.net/images/anime/1452/97840l.webp",
        },
        media_type: "special",
      },
    },
  ],
  paging: {
    next: "https://api.myanimelist.net/v2/anime?offset=4&fields=media_type&q=one&limit=4",
  },
};

const allShows = [
  {
    id: 459,
    title: "One Piece",
    main_picture: {
      medium: "https://cdn.myanimelist.net/images/anime/1770/97704.jpg",
      large: "https://cdn.myanimelist.net/images/anime/1770/97704l.jpg",
    },
    media_type: "movie",
  },
  {
    id: 31772,
    title: "One Punch Man Specials",
    main_picture: {
      medium: "https://cdn.myanimelist.net/images/anime/1452/97840.webp",
      large: "https://cdn.myanimelist.net/images/anime/1452/97840l.webp",
    },
    media_type: "special",
  },
];

const mockShows = allShows;

export async function getResults(s: string) {
  return mockResults;
}

export async function getShows() {
  return { data: mockShows };
}

export async function unregisterShow(id: number) {
  const r = await fetch(process.env.BACKEND_URL + "/shows/" + id, {
    method: "DELETE",
  });
}

export async function registerShow(id: number) {
  const show = mockResults.data.find((show) => show.node.id === id);
  if (show) {
    const body = {
      id: show.node.id,
      source: "source",
      source_id: "source_id",
      title: show.node.title,
      main_picture: show.node.main_picture.medium,
    };
    const r = await fetch(process.env.BACKEND_URL + "/shows", {
      method: "POST",
      body: JSON.stringify(body),
    });
  }
}
