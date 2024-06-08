"use client";

import { useState, useRef } from "react";
import { Box, HStack } from "@chakra-ui/react";
import { Editor } from "@monaco-editor/react";
import * as monaco from "monaco-editor";

import LanguageSelector from "./LanguageSelector";
import Output from "./Output";
import { CODE_SNIPPETS } from "@/lib/constants/languages";
import { capitalize } from "@/lib/utils/string";

type Props = {
    languageVersions: { [key: string]: string | null };
}

export default function CodeEditor({ languageVersions }: Props) {
    const editorRef = useRef<monaco.editor.IStandaloneCodeEditor | null>(null);
    const [value, setValue] = useState("");
    const [language, setLanguage] = useState("python");

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
            <HStack spacing={4}>
                <Box w="50%">
                    <LanguageSelector language={capitalize(language)} onSelect={onSelect} languageVersions={languageVersions} />
                    <Editor
                        height="75vh"
                        theme="vs-dark"
                        language={language === "c++" ? "cpp" : language}
                        defaultValue={CODE_SNIPPETS[language as keyof typeof CODE_SNIPPETS]}
                        onMount={onMount}
                        value={value}
                        onChange={(newValue) => setValue(newValue || "")}
                    />
                </Box>
                <Output editorRef={editorRef} language={language} languageVersions={languageVersions} />
            </HStack>
        </Box>
    );
}