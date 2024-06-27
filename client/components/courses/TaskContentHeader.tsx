import { Button } from "@nextui-org/button"
import Container from "../common/Container"

export default function TaskContentHeader() {
    return <Container className="mb-6 grid grid-cols-2">
        <Container>
            <Button className="mr-1">Modules</Button>
            <Button className="mr-1" variant="ghost">Previous Task</Button>
            <Button className="mr-1" variant="ghost">Next Task</Button>
        </Container>
        <Container className="justify-end text-right">
            <Button className="mr-1" variant="ghost" color="default" disabled>Mark as Complete</Button>
        </Container>
    </Container>
}