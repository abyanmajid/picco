import { Box } from "@chakra-ui/react";

import CodeEditor from "@/components/editor/CodeEditor";
import { getLanguageVersions } from "@/actions/api";
import { SUPPORTED_LANGUAGES } from "@/lib/constants/languages";
import { redirect } from "next/navigation";

export default async function Page() {

  const languageVersions = await getLanguageVersions(SUPPORTED_LANGUAGES);
  if (languageVersions === null) {
    redirect("/error");
  }

  return (
    <Box minH="100vh" bg="#141212" color="#D1D5DB" px={6} py={8}>
      <CodeEditor languageVersions={languageVersions} />
    </Box>
  )
}