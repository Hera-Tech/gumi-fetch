"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React, { ComponentProps } from "react";

function NavLink(props: ComponentProps<typeof Link>) {
  const pathname = usePathname();
  return (
    <Link
      {...props}
      className={`px-4 py-2 rounded-t-lg transition-colors ${
        pathname === props.href ? "bg-gray-800" : "bg-gray-900"
      }`}
    />
  );
}
export default function Navbar() {
  return (
    <nav className="flex gap-4">
      <NavLink href="/results">Results</NavLink>
      <NavLink href="/shows">Shows</NavLink>
    </nav>
  );
}
