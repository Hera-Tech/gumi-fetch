import React from "react";
import { getResults, unregisterShow } from "@/lib/mock";
import Input from "@/components/input";
import SearchBar from "@/components/search-bar";
import Image from "next/image";
import Button from "@/components/button";
import { revalidatePath } from "next/cache";
import { IoIosRemoveCircle } from "react-icons/io";
import { getShows } from "@/lib/server";
import { Show } from "@/lib/types";
export default async function page({
  searchParams,
}: {
  searchParams: Promise<{ s: string }>;
}) {
  const res: Show[] = await getShows();
  async function action(formData: FormData) {
    "use server";
    unregisterShow(parseInt(formData.get("id") as string));
    revalidatePath("/shows");
  }

  return (
    <div className="h-full">
      <ul className="gap-4 grid-cols-6 grid">
        {res.map((anime) => (
          <li key={anime.mal_id} className="relative group overflow-hidden">
            <div className="w-full overflow-hidden rounded-lg">
              <Image
                alt={anime.title}
                src={anime.main_picture}
                width={100}
                height={143}
                className="w-full"
              />
            </div>
            <div className="absolute inset-0 z-10 bg-black/40 opacity-0 group-hover:opacity-100 transition-opacity ease-in-out duration-300"></div>
            <form
              className="absolute inset-0 z-20 opacity-0 group-hover:opacity-100 transition-opacity ease-in-out duration-300"
              action={action}
            >
              <input name="id" defaultValue={anime.mal_id} readOnly hidden />
              <button className="absolute inset-0 cursor-pointer flex items-center justify-center">
                <IoIosRemoveCircle size={50} className="text-red-500" />
              </button>
            </form>
            {/* <form className="ml-auto self-center" action={action}>
              <input name="id" defaultValue={anime.id} readOnly hidden />
              <Button>Remove</Button>
            </form> */}
          </li>
        ))}
      </ul>
    </div>
  );
}
