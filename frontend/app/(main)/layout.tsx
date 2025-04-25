import Navbar from "@/components/navbar";
import Link from "next/link";
import React, { ComponentProps } from "react";

export default function layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="h-screen overflow-y-hidden lg:max-w-5xl mx-auto px-10 py-5 flex flex-col">
      <Navbar />
      <main className="bg-gray-800 h-full p-4">{children}</main>
    </div>
  );
}
