import { executeCode } from "@/actions/api";
import { Box, Text, Button } from "@chakra-ui/react";
import * as monaco from "monaco-editor";

type Props = {
    editorRef: React.RefObject<monaco.editor.IStandaloneCodeEditor>;
    language: string;
    languageVersions: { [key: string]: string | null };
}

export default function Output({ editorRef, language, languageVersions }: Props) {

    async function runCode() {
        if (editorRef.current) {
            const sourceCode = editorRef.current.getValue();
            try {
                const { run: result } = await executeCode(language, languageVersions[language], sourceCode);
                console.log(result);
            } catch (error) {

            }
        }
    }

    return (
        <Box w="50%">
            <Text mb={2} fontSize="lg">Output</Text>
            <Button
                variant="outline"
                colorScheme="green"
                mb={4}
                onClick={runCode}
            >
                Run Code
            </Button>
            <Box
                height="75vh"
                p={2}
                border="1px solid"
                borderRadius={4}
                borderColor="#333"
            >
                test
            </Box>
        </Box>
    )
}