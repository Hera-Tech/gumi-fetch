import React from "react";

export default function ResultsSkeleton() {
  return (
    <ul className="flex flex-col h-full gap-4 py-4 animate-pulse">
      <li className=" bg-gray-900 p-2 rounded-lg flex w-full gap-4">
        <div className="w-[100px] h-[143px] bg-gray-700 rounded"></div>
        <div className="flex-1 rounded h-4 bg-gray-700"></div>
        <div className="self-center text-gray-700 bg-gray-700 w-20 px-4 py-1 rounded-lg shadow cursor-default select-none">
          Add
        </div>
      </li>
    </ul>
  );
}
