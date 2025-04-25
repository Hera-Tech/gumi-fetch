"use server";
import { redirect } from "next/navigation";

export default async function search(formData: FormData) {
  redirect(`/results?s=${formData.get("search")}`);
}
