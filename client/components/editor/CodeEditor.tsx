"use client";

import { useState, useRef } from "react";
import { Box, HStack, Flex, Spacer, useToast, Text } from "@chakra-ui/react";
import { Editor } from "@monaco-editor/react";
import * as monaco from "monaco-editor";

import LanguageSelector from "./LanguageSelector";
import Output from "./Output";
import IOSwitcher from "./IOSwitcher";
import { CODE_SNIPPETS } from "@/lib/constants/languages";
import { capitalize } from "@/lib/utils/string";
import { executeCode } from "@/actions/api";

type Props = {
    languageVersions: { [key: string]: string | null };
}

export default function CodeEditor({ languageVersions }: Props) {
    const toast = useToast();

    const editorRef = useRef<monaco.editor.IStandaloneCodeEditor | null>(null);
    const [value, setValue] = useState("");
    const [language, setLanguage] = useState("python");
    const [outputShown, setOutputShown] = useState(false);
    const [output, setOutput] = useState<string[] | null>(null);
    const [isLoading, setIsLoading] = useState(false);
    const [isError, setIsError] = useState(false);

    async function runCode() {
        if (editorRef.current) {
            const sourceCode = editorRef.current.getValue();
            try {
                setIsLoading(true);
                const { run: result } = await executeCode(language, languageVersions[language], sourceCode);
                setOutputShown(true);
                setOutput(result.output.split("\n"));
                result.stderr ? setIsError(true) : setIsError(false)
            } catch (error) {
                toast({
                    title: "Unexpected error occurred.",
                    description: "Please try again later.",
                    status: "error",
                    duration: 6000,
                })
            }
            finally {
                setIsLoading(false);
            }
        }
    }

    function onMount(editor: monaco.editor.IStandaloneCodeEditor) {
        editorRef.current = editor;
        editor.focus();
    }

    function onSelect(lang: string) {
        setLanguage(lang);
        setValue(CODE_SNIPPETS[lang as keyof typeof CODE_SNIPPETS]);
    }

    return (
        <Box>
            <Flex>
                <HStack flex="1" overflow="auto" spacing={2}>
                    <LanguageSelector language={capitalize(language)} onSelect={onSelect} languageVersions={languageVersions} />
                    <IOSwitcher outputShown={outputShown} setOutputShown={setOutputShown} />
                </HStack>
                <Spacer />
                <Output runCode={runCode} isLoading={isLoading} />
            </Flex>
            {
                outputShown ?
                    <Box
                        height="75vh"
                        p={2}
                        border="1px solid"
                        borderRadius={4}
                        color={
                            isError ? "red.400" : ""
                        }
                        borderColor={
                            isError ? "red.500" : "#333"
                        }
                    >
                        {
                            output ?
                                output.map(
                                    (line: string, i: number) => <Text key={i}>{line}</Text>
                                )
                                : 'Click "Run" to run your solution...'
                        }
                    </Box>
                    :
                    <Editor
                        height="75vh"
                        theme="vs-dark"
                        language={language === "c++" ? "cpp" : language}
                        defaultValue={CODE_SNIPPETS[language as keyof typeof CODE_SNIPPETS]}
                        onMount={onMount}
                        value={value}
                        onChange={(newValue) => setValue(newValue || "")}
                        options={{
                            minimap: {
                                enabled: false,
                            },
                        }}
                    />
            }

        </Box>
    );
}