import { Button } from "@nextui-org/button"
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
        <Button
            className="ml-1"
            variant="bordered"
            color="default"
            onClick={toggleOutput}
        >
            <LuTerminal /> {outputShown ? "Output" : "Code"}
        </Button>
    );
}
