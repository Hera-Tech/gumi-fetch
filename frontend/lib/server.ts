"use server";

export async function getShows() {
  const r = await fetch(process.env.BACKEND_URL + "/shows");
  const body = await r.json();
  const data = body.data;
  return data;
}

export async function search(searchString: string) {
  if (searchString === "") {
    return {
      data: [],
    };
  }
  const r = await fetch(
    process.env.BACKEND_URL +
      "/search?fields=media_type&limit=10&q=" +
      searchString
  );
  const body = (await r.json()) as {
    data: {
      node: {
        id: number;
        title: string;
        main_picture: {
          medium: string;
          large: string;
        };
        media_type: string;
      };
    }[];
    pagination: { next?: string; previous?: string };
  };
  const data = body.data.filter((anime) => anime.node.media_type === "tv"); // Filter only TV series
  return {
    data,
    pagination: {
      next: body.pagination.next,
      previous: body.pagination.previous,
    },
  };
}
