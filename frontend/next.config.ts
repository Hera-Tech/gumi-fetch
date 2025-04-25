import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  images: {
    remotePatterns: [new URL("https://cdn.myanimelist.net/images/**")],
  },
};

export default nextConfig;
