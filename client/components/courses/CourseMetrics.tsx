import Container from "../common/Container"
import { Avatar } from "@nextui-org/react"
import { Link } from "@nextui-org/link"

type Props = {
    creator: string,
    creatorProfileUrl: string,
}

export default function CourseMetrics({ creator, creatorProfileUrl }: Props) {
    return <>
        <Container className="flex items-center gap-4">
            <Avatar isBordered color="primary" src="/abyan-150x150.png" />
            <Container className="flex flex-col justify-start text-left">
                <Container className="text-neutral-300">
                    Created by <Link href={creatorProfileUrl} className="text-primary">{creator}</Link>
                </Container>
                <Container className="text-neutral-300">Last updated: 06/07/24</Container>
            </Container>
        </Container>
        <Container className="justify-end text-right">
            <Container className="text-neutral-300">26 modules</Container>
            <Container className="text-neutral-300">17 students</Container>
        </Container>
    </>
}