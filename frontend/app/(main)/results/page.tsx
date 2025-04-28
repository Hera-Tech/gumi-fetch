import React, { Suspense } from "react";
import SearchBar from "@/components/search-bar";
import ResultsSkeleton from "./ResultsSkeleton";
import Results from "./Results";
export default async function page({
  searchParams,
}: {
  searchParams: Promise<{ query: string }>;
}) {
  const searchString = (await searchParams).query;

  return (
    <div className="h-full py-4 flex-col flex gap-2">
      <SearchBar />
      <Suspense key={searchString} fallback={<ResultsSkeleton />}>
        <Results query={searchString ?? ""} />
      </Suspense>
    </div>
  );
}
