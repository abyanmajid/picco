import fs from "node:fs";
import nextMDX from "@next/mdx";
import rehypePrettyCode from "rehype-pretty-code";

/** @type {import('rehype-pretty-code').Options} */
const options = {
    theme: "dark-plus"
};

const withMDX = nextMDX({
    extension: /\.mdx?$/,
    options: {
        remarkPlugins: [],
        rehypePlugins: [[rehypePrettyCode, options]],
    },
});

/** @type {import('next').NextConfig} */
const nextConfig = { reactStrictMode: true };

export default withMDX(nextConfig);