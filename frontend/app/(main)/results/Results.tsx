import Button from "@/components/button";
import { registerShow } from "@/lib/mock";
import { search } from "@/lib/server";
import Image from "next/image";
import { redirect } from "next/navigation";
import React from "react";

export default async function Results({ query }: { query: string }) {
  const res = await search(query);
  async function addShow(formData: FormData) {
    "use server";
    registerShow(formData);
    redirect("/shows");
  }
  return (
    <ul className="flex flex-col h-full gap-4 py-4 overflow-y-scroll">
      {res.data.map((anime) => (
        <li
          key={anime.node.id}
          className=" bg-gray-900 p-2 rounded-lg flex w-full gap-4"
        >
          <div>
            <Image
              alt={anime.node.title}
              src={anime.node.main_picture.medium}
              width={100}
              height={143}
            />
          </div>
          <div>
            <div>{anime.node.title}</div>
            <div>{anime.node.media_type}</div>
          </div>
          <form className="ml-auto self-center" action={addShow}>
            <input
              name="main_picture"
              defaultValue={anime.node.main_picture.medium}
              readOnly
              hidden
            />
            <input
              name="title"
              defaultValue={anime.node.title}
              readOnly
              hidden
            />
            <input name="id" defaultValue={anime.node.id} readOnly hidden />
            <Button>Add</Button>
          </form>
        </li>
      ))}
    </ul>
  );
}
