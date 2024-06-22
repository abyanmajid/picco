import Title from "@/components/common/Title"
import Subtitle from "@/components/common/Subtitle"
import Container from "@/components/common/Container"
import LoginForm from "@/components/forms/LoginForm";
import { HeroHighlight } from "@/components/ui/hero-highlight";

export default function LoginPage() {
    return (
        <HeroHighlight>
            <Container className="container mx-auto max-w-full px-6 flex-grow">
                <Container className="flex flex-col justify-center items-center text-center space-y-4">
                    <Container className="mb-4">
                        <Title size="sm">Welcome Back</Title>
                        <Subtitle>Log in to your account to start learning</Subtitle>
                    </Container>
                    <LoginForm />
                </Container>
            </Container>
        </HeroHighlight>
    );
}