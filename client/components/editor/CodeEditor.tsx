"use client";

import { useState, useRef } from "react";
import { Box, HStack } from "@chakra-ui/react";
import { Editor } from "@monaco-editor/react";
import * as monaco from "monaco-editor";

import LanguageSelector from "./LanguageSelector";
import { CODE_SNIPPETS } from "@/lib/constants/languages";

export default function CodeEditor() {
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
                    <LanguageSelector language={language} onSelect={onSelect} />
                    <Editor
                        height="75vh"
                        theme="vs-dark"
                        language={language}
                        defaultValue={CODE_SNIPPETS[language as keyof typeof CODE_SNIPPETS]}
                        onMount={onMount}
                        value={value}
                        onChange={(newValue) => setValue(newValue || "")}
                    />
                </Box>
            </HStack>
        </Box>
    );
}