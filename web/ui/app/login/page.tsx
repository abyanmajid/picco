import Title from "@/components/common/Title"
import Subtitle from "@/components/common/Subtitle"
import Container from "@/components/common/Container"
import LoginForm from "@/components/forms/LoginForm";

export default function LoginPage() {
    return (
        <Container className="flex flex-col justify-center items-center text-center space-y-4">
            <Container className="mb-4">
                <Title size="sm">Welcome Back</Title>
                <Subtitle>Log in to your account to start learning</Subtitle>
            </Container>
            <LoginForm />
        </Container>
    );
}