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
  mockShows.splice(
    mockShows.findIndex((show) => show.id === id),
    1
  );
}

export async function registerShow(id: number) {
  const show = allShows.find((show) => show.id === id);
  if (show) mockShows.push(show);
}
