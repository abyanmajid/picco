import Title from "@/components/common/Title"
import Subtitle from "@/components/common/Subtitle"
import Container from "@/components/common/Container"
import ResetPasswordForm from "@/components/forms/ResetPasswordForm";

export default function ResetPasswordPage() {
    return (
        <Container className="flex flex-col justify-center items-center text-center space-y-4">
            <Container className="mb-4">
                <Title size="sm">Reset Your Password</Title>
                <Subtitle>Hi <span className="font-bold">user@example.com</span>, please your new password.</Subtitle>
                <Subtitle>This link expires in one hour.</Subtitle>
            </Container>
            <ResetPasswordForm />
        </Container>
    );
}