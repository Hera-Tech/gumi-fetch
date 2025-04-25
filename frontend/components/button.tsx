import React from "react";

export default function Button({
  className = "",
  children,
  ...props
}: React.ButtonHTMLAttributes<HTMLButtonElement>) {
  return (
    <button
      className={
        "bg-purple-800 px-4 py-1 rounded-lg shadow cursor-pointer " + className
      }
      {...props}
    >
      {children}
    </button>
  );
}
