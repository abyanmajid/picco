import { Box, Button, Menu, MenuButton, MenuItem, MenuList, Text } from "@chakra-ui/react";
import { LANGUAGE_VERSIONS } from "@/lib/constants/languages";

const ACTIVE_COLOR = "blue.400";
const BACKGROUND_COLOR = "gray.900";

const languages = Object.entries(LANGUAGE_VERSIONS);

type Props = {
    language: string;
    onSelect: (lang: string) => void;
};

export default function LanguageSelector({ language, onSelect }: Props) {
    return (
        <Box ml={2} mb={4}>
            <Text mb={2} fontSize="lg">Language: </Text>
            <Menu isLazy>
                <MenuButton as={Button}>
                    {language}
                </MenuButton>
                <MenuList bg="#110c1b">
                    {languages.map(([lang, version]) => (
                        <MenuItem key={lang} color={
                            lang === language ? ACTIVE_COLOR : ""
                        }
                            bg={
                                lang === language ? BACKGROUND_COLOR : ""
                            }
                            _hover={{
                                color: ACTIVE_COLOR,
                                bg: BACKGROUND_COLOR,
                            }}
                            onClick={() => onSelect(lang)}>
                            {lang} <Text as="span" color="gray.600" fontSize="sm">&nbsp;({version})</Text>
                        </MenuItem>
                    ))}
                </MenuList>
            </Menu>
        </Box>
    );
}
