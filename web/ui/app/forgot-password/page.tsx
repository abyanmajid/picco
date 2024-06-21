import Title from "@/components/common/Title"
import Subtitle from "@/components/common/Subtitle"
import Container from "@/components/common/Container"
import ForgotPasswordForm from "@/components/forms/ForgotPasswordForm";

export default function ForgotPasswordPage() {
    return (
        <Container className="flex flex-col justify-center items-center text-center space-y-4">
            <Container className="mb-4">
                <Title size="sm">Forgot Password</Title>
                <Subtitle>Send a password reset link to your email address</Subtitle>
            </Container>
            <ForgotPasswordForm />
        </Container>
    );
}