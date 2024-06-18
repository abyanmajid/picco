import withMDX from "@next/mdx";
import remarkPrism from "remark-prism";

const withMDXConfig = withMDX({
    extension: /\.mdx?$/,
    options: {
        remarkPlugins: [remarkPrism],
    },
});

/** @type {import('next').NextConfig} */
const nextConfig = {
    pageExtensions: ["ts", "tsx", "js", "jsx", "mdx"],
};

export default withMDXConfig(nextConfig);
