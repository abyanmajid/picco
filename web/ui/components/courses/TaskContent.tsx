import { Button } from "@nextui-org/button"
import Container from "../common/Container"
import Title from "../common/Title"
import TaskContentHeader from "./TaskContentHeader"
import TaskInstructions from "@/content/introduction-to-programming/module-1/task-1.mdx"

export default function TaskContent() {
    return <>
        <TaskContentHeader />
        <Container className="prose dark:prose-invert">
            <TaskInstructions />
        </Container>
    </>
}