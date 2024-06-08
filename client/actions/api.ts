import axios from "axios";
import { version } from "os";

type LanguageEntry = {
    language: string;
    version: string;
    aliases: string[];
    runtime?: string;
};

const api = {
    piston: axios.create({
        baseURL: "https://emkc.org/api/v2/piston",
        headers: {
            'Content-Type': 'application/json'
        }
    })
}

export async function getLanguageVersions(languages: string[]): Promise<{ [key: string]: string | null }> {
    const response = await api.piston.get("/runtimes");
    console.log("API Response:", response.data);

    const versions: { [key: string]: string | null } = {};

    for (const language of languages) {
        const languageEntry = response.data.find((entry: LanguageEntry) => entry.language === language);
        if (!languageEntry) {
            console.log(`Language not found: ${language}`);
            versions[language] = null;
        } else {
            versions[language] = languageEntry.version;
        }
    }

    return versions;
}

export async function executeCode(language: string, version: string | null, sourceCode: string) {
    const response = await api.piston.post("/execute", {
        language: language,
        version: version,
        files: [
            {
                content: sourceCode,
            },
        ],
    });
    return response.data;
}