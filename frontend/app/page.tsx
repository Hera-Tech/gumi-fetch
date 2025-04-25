import Image from "next/image";
import { FaSearch } from "react-icons/fa";
import search from "./action";
import Link from "next/link";
import { redirect } from "next/navigation";
import Input from "@/components/input";
import SearchBar from "@/components/search-bar";

export default function Home() {
  return (
    <div className="h-screen flex flex-col items-center justify-center">
      <SearchBar />
      <Link href="/shows">Manage list</Link>
    </div>
  );
}
