import React from "react";
import { Dropdown, DropdownTrigger, DropdownMenu, DropdownItem, Button, Selection } from "@nextui-org/react";
import { capitalize } from "@/utils/helpers";

type Props = {
    languageVersions: { [key: string]: string | null };
};

export default function LanguageSelector() {
    const [selectedKeys, setSelectedKeys] = React.useState<Set<string>>(new Set(["python"]));

    const selectedValue = React.useMemo(
        () => Array.from(selectedKeys).join(", ").replaceAll("_", " "),
        [selectedKeys]
    );

    const handleSelectionChange = (keys: Selection) => {
        setSelectedKeys(new Set(keys as Set<string>));
    };

    return (
        <Dropdown>
            <DropdownTrigger>
                <Button className="capitalize">
                    {selectedValue}
                </Button>
            </DropdownTrigger>
            <DropdownMenu
                aria-label="Single selection example"
                variant="flat"
                disallowEmptySelection
                selectionMode="single"
                selectedKeys={selectedKeys}
                onSelectionChange={handleSelectionChange}
            >
                <DropdownItem key="python">Python</DropdownItem>
                <DropdownItem key="java">Java</DropdownItem>
                <DropdownItem key="cpp">C++</DropdownItem>
                <DropdownItem key="javascript">JavaScript</DropdownItem>
            </DropdownMenu>
        </Dropdown>
    );
}
