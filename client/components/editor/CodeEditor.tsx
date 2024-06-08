"use client";

import { useState, useRef } from "react";
import { Box, HStack, SimpleGrid, Flex, Spacer } from "@chakra-ui/react";
import { Editor } from "@monaco-editor/react";
import * as monaco from "monaco-editor";

import LanguageSelector from "./LanguageSelector";
import Output from "./Output";
import IOSwitcher from "./IOSwitcher";
import { CODE_SNIPPETS } from "@/lib/constants/languages";
import { capitalize } from "@/lib/utils/string";

type Props = {
    languageVersions: { [key: string]: string | null };
}

export default function CodeEditor({ languageVersions }: Props) {
    const editorRef = useRef<monaco.editor.IStandaloneCodeEditor | null>(null);
    const [value, setValue] = useState("");
    const [language, setLanguage] = useState("python");
    const [outputShown, setOutputShown] = useState(false);

    function onMount(editor: monaco.editor.IStandaloneCodeEditor) {
        editorRef.current = editor;
        editor.focus();
    }

    function onSelect(lang: string) {
        setLanguage(lang);
        setValue(CODE_SNIPPETS[lang as keyof typeof CODE_SNIPPETS]);
    }

    return (
        <SimpleGrid columns={2} spacing={10} overflow="hidden">
            <Box>
                <Flex>
                    <HStack flex="1" overflow="auto" spacing={2}>
                        <LanguageSelector language={capitalize(language)} onSelect={onSelect} languageVersions={languageVersions} />
                        <IOSwitcher outputShown={outputShown} setOutputShown={setOutputShown} />
                    </HStack>
                    <Spacer />
                    <Output editorRef={editorRef} language={language} languageVersions={languageVersions} />
                </Flex>
                {
                    outputShown ?
                        <Box
                            height="85vh"
                            p={2}
                            border="1px solid"
                            borderRadius={4}
                            borderColor="#333"
                        >
                            test
                        </Box>
                        :
                        <Editor
                            height="85vh"
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
        </SimpleGrid>
    );
}