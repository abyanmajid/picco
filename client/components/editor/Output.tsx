import { executeCode } from "@/actions/api";
import { Box, Button } from "@chakra-ui/react";
import * as monaco from "monaco-editor";
import { IoPlay } from "react-icons/io5";

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
        <Box>
            <Button
                variant="outline"
                colorScheme="green"
                mb={4}
                onClick={runCode}
                leftIcon={<IoPlay />}
            >
                Run
            </Button>
        </Box>
    )
}