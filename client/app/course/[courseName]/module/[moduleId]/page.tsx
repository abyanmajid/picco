import { Box, SimpleGrid } from "@chakra-ui/react";

import CodeEditor from "@/components/editor/CodeEditor";
import { getLanguageVersions } from "@/actions/api";
import { SUPPORTED_LANGUAGES } from "@/lib/constants/languages";
import { redirect } from "next/navigation";
import rehypeHighlight from 'rehype-highlight';
import "./vs.css";

import ModuleContent from "@/content/data-structures-and-algorithms/module1.mdx"

const options = {
    mdxOptions: {
        remarkPlugins: [],
        rehypePlugins: [rehypeHighlight],
    }
}

export default async function LearnPage() {

    const languageVersions = await getLanguageVersions(SUPPORTED_LANGUAGES);
    if (languageVersions === null) {
        redirect("/error");
    }

    return (
        <SimpleGrid columns={2} spacing={0.25} overflow="hidden">
            <Box height="100vh" overflowY="auto" bg="#141212" color="#D1D5DB" px={6} py={6}>
                <ModuleContent />
            </Box>
            <Box height="100vh" bg="#171515" color="#D1D5DB" px={6} py={6}>
                <CodeEditor languageVersions={languageVersions} />
            </Box>
        </SimpleGrid>
    )
}