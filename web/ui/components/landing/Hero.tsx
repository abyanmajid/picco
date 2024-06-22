import { GithubIcon } from "@/components/ui/icons";
import { Link } from "@nextui-org/link";
import { title, subtitle } from "@/components/primitives";
import { button as buttonStyles } from "@nextui-org/theme";
import { siteConfig } from "@/config/site";
import Title from "../common/Title";
import Subtitle from "../common/Subtitle";
import Container from "../common/Container";

export default function Hero() {
    return <>
        <Container className="inline-block max-w-2xl text-center justify-center">
            <Title>A&nbsp;</Title>
            <Title color="violet">blazingly fast&nbsp;</Title>
            <Title>way to&nbsp;</Title>
            <Title>learn programming&nbsp;</Title>
            <Title color="violet">for free.</Title>
            <Subtitle className="mt-4" >Avoid tutorial hell, and learn effectively by writing lots of code.</Subtitle>
        </Container >

        <Container className="flex gap-3">
            <Link
                className={buttonStyles({
                    color: "primary",
                    radius: "full",
                    variant: "shadow",
                })}
                href="/login"
            >
                Start Learning
            </Link>
            <Link
                isExternal
                className={buttonStyles({ variant: "bordered", radius: "full" })}
                href={siteConfig.links.github}
            >
                <GithubIcon size={20} />
                Source
            </Link>
        </Container>
    </>
}