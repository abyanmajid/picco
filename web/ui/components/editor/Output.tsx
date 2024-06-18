import { useState } from "react";
import { executeCode } from "@/actions/api";
import { Box, Button } from "@chakra-ui/react";
import * as monaco from "monaco-editor";
import { IoPlay } from "react-icons/io5";

type Props = {
    runCode: () => Promise<void>;
    isLoading: boolean;
}

export default function Output({ runCode, isLoading }: Props) {

    return (
        <Box>
            <Button
                isLoading={isLoading}
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