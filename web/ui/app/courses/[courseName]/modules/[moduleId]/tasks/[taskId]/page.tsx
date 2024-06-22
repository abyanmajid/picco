import React from "react";
import Container from "@/components/common/Container";
import Title from "@/components/common/Title"
import CodeEditor from "@/components/editor/CodeEditor";
import { redirect } from "next/navigation";
import { getLanguageVersions } from "@/actions/api";
import { SUPPORTED_LANGUAGES } from "@/utils/constants";
import { Breadcrumbs, BreadcrumbItem } from "@nextui-org/breadcrumbs";
import { Button } from "@nextui-org/button"
import TaskContent from "@/components/courses/TaskContent";

export default async function TaskPage() {
    const languageVersions = await getLanguageVersions(SUPPORTED_LANGUAGES);
    if (languageVersions === null) {
        redirect("/error");
    }
    return (
        <Container className="grid grid-cols-2 overflow-hidden gap-1 h-screen">
            <Container className="bg-neutral-900 p-8 max-h-screen overflow-y-auto">
                <TaskContent />
            </Container>
            <Container className="bg-neutral-900 p-8">
                <CodeEditor languageVersions={languageVersions} />
            </Container>
        </Container>
    );
}