"use client"

import React, { useState, useRef } from "react"
import { Editor } from "@monaco-editor/react";
import * as monaco from "monaco-editor";
import { CODE_SNIPPETS } from "@/utils/constants";
import Container from "../common/Container";
import { Button } from "@nextui-org/button";
import LanguageSelector from "./LanguageSelector";
import { capitalize } from "@/utils/helpers";
import IOSwitcher from "./IOSwitcher";
import Output from "./Output";

type Props = {
    languageVersions: { [key: string]: string | null };
}

export default function CodeEditor() {
    const editorRef = useRef<monaco.editor.IStandaloneCodeEditor | null>(null);
    const [value, setValue] = useState("");
    const [language, setLanguage] = useState("python");
    const [outputShown, setOutputShown] = useState(false);
    const [output, setOutput] = useState<string[] | null>(null);
    const [isLoading, setIsLoading] = useState(false);
    const [isError, setIsError] = useState(false);

    async function runCode() {
        // if (editorRef.current) {
        //     const sourceCode = editorRef.current.getValue();
        //     try {
        //         setIsLoading(true);
        //         const { run: result } = await executeCode(language, languageVersions[language], sourceCode);
        //         setOutputShown(true);
        //         setOutput(result.output.split("\n"));
        //         result.stderr ? setIsError(true) : setIsError(false)
        //     } catch (error) {
        //         // toast({
        //         //     title: "Unexpected error occurred.",
        //         //     description: "Please try again later.",
        //         //     status: "error",
        //         //     duration: 6000,
        //         // })
        //     }
        //     finally {
        //         setIsLoading(false);
        //     }
        // }
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
        <Container>
            <Container className="grid grid-cols-2 mb-4">
                <Container className="flex items-center mr">
                    <LanguageSelector />
                    <IOSwitcher outputShown={outputShown} setOutputShown={setOutputShown} />
                </Container>
                <Container className="justify-right text-right">
                    <Output runCode={runCode} isLoading={isLoading} />
                </Container>
            </Container>
            {
                outputShown ?
                    <div
                        style={{
                            height: "75vh",
                            padding: "16px",
                            border: "1px solid",
                            borderRadius: "4px",
                            color: isError ? "red" : "",
                            borderColor: isError ? "darkred" : "#333"
                        }}
                    >
                        {
                            output ?
                                output.map(
                                    (line: string, i: number) => <p key={i}>{line}</p>
                                )
                                : 'Click "Run" to run your solution...'
                        }
                    </div>
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
        </Container>
    );
}