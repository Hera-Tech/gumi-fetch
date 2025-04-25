import React from "react";

export default function Input({
  className = "",
  ...props
}: React.InputHTMLAttributes<HTMLInputElement>) {
  return (
    <input
      className={"w-full bg-white text-black rounded-lg  px-4 py-1" + className}
      {...props}
    />
  );
}
