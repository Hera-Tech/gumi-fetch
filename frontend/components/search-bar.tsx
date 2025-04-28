"use client";
import search from "@/app/action";
import React from "react";
import { FaSearch } from "react-icons/fa";
import Input from "./input";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { useDebouncedCallback } from "use-debounce";

export default function SearchBar() {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const { replace } = useRouter();

  const handleSearch = useDebouncedCallback((term: string) => {
    const params = new URLSearchParams(searchParams);
    if (term) {
      params.set("query", term);
    } else {
      params.delete("query");
    }
    replace(`${pathname}?${params.toString()}`);
  }, 300);
  return (
    <div className="w-2xl relative text-black">
      <Input
        name="search"
        onChange={(e) => {
          handleSearch(e.target.value);
        }}
        defaultValue={searchParams.get("query")?.toString()}
      />
      <div className="absolute right-2 top-1/2 -translate-y-1/2 ">
        <FaSearch />
      </div>
    </div>
  );
}
