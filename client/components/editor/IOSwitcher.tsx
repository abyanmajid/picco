import { Box, Button } from "@chakra-ui/react";
import { LuTerminal } from "react-icons/lu";

type Props = {
    outputShown: boolean;
    setOutputShown: (shown: boolean) => void;
};

export default function IOSwitcher({ outputShown, setOutputShown }: Props) {
    function toggleOutput() {
        setOutputShown(!outputShown);
    }

    return (
        <Box>
            <Button
                colorScheme="gray"
                mb={4}
                onClick={toggleOutput}
                leftIcon={<LuTerminal />}
            >
                {outputShown ? "Code" : "Output"}
            </Button>
        </Box>
    );
}
