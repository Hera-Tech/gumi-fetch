"use client";
import search from "@/app/action";
import React from "react";
import { FaSearch } from "react-icons/fa";
import Input from "./input";

export default function SearchBar({ defaultValue }: { defaultValue?: string }) {
  return (
    <form className="w-2xl relative text-black" action={search}>
      <Input name="search" defaultValue={defaultValue} />
      <div className="absolute right-2 top-1/2 -translate-y-1/2 ">
        <FaSearch />
      </div>
    </form>
  );
}
