import { Box, Button, Menu, MenuButton, MenuItem, MenuList, Text } from "@chakra-ui/react";
import { capitalize } from "@/lib/utils/string";

const ACTIVE_COLOR = "blue.400";
const BACKGROUND_COLOR = "gray.900";

type Props = {
    language: string;
    onSelect: (lang: string) => void;
    languageVersions: { [key: string]: string | null };
};

export default function LanguageSelector({ language, onSelect, languageVersions }: Props) {

    return (
        <Box mb={4}>
            <Menu isLazy>
                <MenuButton
                    as={Button}
                    colorScheme="teal"
                >
                    {language}
                </MenuButton>
                <MenuList bg="#110c1b">
                    {Object.entries(languageVersions).map(([lang, version]) => (
                        <MenuItem
                            key={lang}
                            color={lang === language ? ACTIVE_COLOR : ""}
                            bg={lang === language ? BACKGROUND_COLOR : ""}
                            _hover={{
                                color: ACTIVE_COLOR,
                                bg: BACKGROUND_COLOR,
                            }}
                            onClick={() => onSelect(lang)}
                        >
                            {capitalize(lang)} <Text as="span" color="gray.600" fontSize="sm">&nbsp;({version})</Text>
                        </MenuItem>
                    ))}
                </MenuList>
            </Menu>
        </Box>
    );
}
