"use server";

export async function getShows() {
  const r = await fetch(process.env.BACKEND_URL + "/shows");
  const body = await r.json();
  const data = body.data;
  return data;
}
