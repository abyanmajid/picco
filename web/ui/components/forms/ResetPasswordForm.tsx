import Container from "@/components/common/Container"
import PasswordInput from "@/components/auth/PasswordInput"
import { ResetPasswordButton } from "@/components/auth/AuthButtons"
import { Link } from "@nextui-org/link"

export default function ForgotPasswordForm() {
    return <>
        <Container className="w-full max-w-md grid gap-4 mb-2">
            <PasswordInput includeConfirmPassword={true} />
            <ResetPasswordButton />
        </Container>
    </>
}