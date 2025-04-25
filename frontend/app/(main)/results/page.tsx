import React from "react";
import { getResults, registerShow } from "@/lib/mock";
import Input from "@/components/input";
import SearchBar from "@/components/search-bar";
import Image from "next/image";
import Button from "@/components/button";
import { revalidatePath } from "next/cache";
export default async function page({
  searchParams,
}: {
  searchParams: Promise<{ s: string }>;
}) {
  const searchString = (await searchParams).s;
  const res = await getResults(searchString);
  async function addShow(formData: FormData) {
    "use server";
    registerShow(parseInt(formData.get("id") as string));
    revalidatePath("/shows");
  }

  return (
    <div className="h-full">
      <SearchBar defaultValue={searchString} />
      <ul className="flex flex-col h-full gap-4 py-4">
        {res.data.map((anime) => (
          <li
            key={anime.node.id}
            className=" bg-gray-800 p-2 rounded-lg flex w-full gap-4"
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
              <input name="id" defaultValue={anime.node.id} readOnly hidden />
              <Button>Add</Button>
            </form>
          </li>
        ))}
      </ul>
    </div>
  );
}
